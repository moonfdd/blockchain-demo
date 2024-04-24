package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"net/mail"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// os.Stderr:标准错误输出
var errerLog = log.New(os.Stderr, "ERROR:", log.Lshortfile)
var serialNumberLimit = new(big.Int).Lsh(big.NewInt(1), 128)

func main() {
	//certshop.exe ca
	var command string
	//用户没有输入任何命令
	if len(os.Args) < 1 {
		command = ""
	} else {
		command = os.Args[1]
	}
	switch command {
	//证书颁发机构
	case "ca":
		createCA(os.Args[2:], "ca", "/CN=certshop-ca/O=xiongdilian/OU=qukuailian", 3650)
	//中级证书颁发机构
	case "ica":
		createCA(os.Args[2:], "ca/ica", "/CN=certshop-ca1/O=xiongdilian1/OU=qukuailian1", 365*5)
	//服务器证书
	case "server":
		createCertificate(os.Args[2:], "ca/server", "/CN=server",
			"127.0.0.1,192.168.12.16,764217451@qq.com",
			365*5, x509.KeyUsageDigitalSignature|x509.KeyUsageDataEncipherment, []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth})
		//客户端证书
	case "client":
		createCertificate(os.Args[2:], "ca/client", "/CN=client",
			"",
			365*5, x509.KeyUsageDigitalSignature|x509.KeyUsageDataEncipherment, []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth})
	default:
		errerLog.Println("Usage: ca | ica | server | client")
	}

}

/*
思路：
1.如果创建的是CA，则需要保存其自签证书和私钥，并将其证书拷贝到ca.pem中
2.如果创建的是中级CA，则需要将其父级政CA拷贝到自己证书的后后面，且将父级CA中ca.pem拷贝到自己的文件夹中
*/
func createCA(args []string, path, defaultDn string, defaultValidity int) {
	//参数一：名称   参数二：错误处理策略
	fs := flag.NewFlagSet("ca", flag.PanicOnError)
	dn := fs.String("dn", defaultDn, "证书主题")
	maxPathLength := fs.Int("maxPathLength", 5, "可以颁发的证书数量")
	validity := fs.Int("validity", defaultValidity, "证书有效期")
	overwrite := fs.Bool("overwrite", false, "是否覆盖原有文件")
	//解析参数
	err := fs.Parse(args)
	if err != nil {
		errerLog.Fatalf("将参数解析到命令失败:%s", err)
	}
	//解析之后剩余的参数大于1
	if len(fs.Args()) > 1 {
		errerLog.Fatalf("参数非法:%s", strings.Join(fs.Args(), ","))
	} else if len(fs.Args()) == 1 { //解析之后剩余的参数等于1
		path = fs.Arg(0)
	}
	if !*overwrite {
		checkExisting(path)
	}
	//   C:/ca/ica/aa.txt===>C:/ca/ica
	//    ca==> .
	ca := filepath.Dir(path)
	var caCert *x509.Certificate
	var caKey *ecdsa.PrivateKey
	if ca != "." {
		//从文件中读取父级CA的证书
		caCert = parseCert(ca)
		//判断是否是CA
		if !caCert.IsCA {
			errerLog.Fatalf("%s 不是一个有效的证书颁发机构!", ca)
		} else if !(caCert.MaxPathLen > 0) {
			errerLog.Fatalf("证书颁发机构 %s 无法颁发证书!", ca)
		}
		*maxPathLength = caCert.MaxPathLen - 1
		caKey = parseKey(ca)
	}
	//生成私钥
	key, derKey, err := generatePrivateKey()
	if err != nil {
		errerLog.Fatalf("私钥生成失败:%s", err)
	}

	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		errerLog.Fatalf("证书编号生成失败:%s", err)
	}
	notBofor := time.Now().UTC()
	notAfter := notBofor.AddDate(0, 0, *validity)

	template := x509.Certificate{
		//序列号
		SerialNumber: serialNumber,
		//主题
		Subject: *parseDn(caCert, *dn),
		//生效时间
		NotBefore: notBofor,
		//失效时间
		NotAfter:              notAfter,
		BasicConstraintsValid: true,
		//是否是CA
		IsCA:           true,
		MaxPathLen:     *maxPathLength,
		MaxPathLenZero: *maxPathLength == 0,
		//证书的用途：用于数字签名和证书签发
		KeyUsage: x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}
	//判断是否获取到父级证书
	if caCert == nil {
		caCert = &template
		caKey = key
	}
	//创建证书
	derCert, err := x509.CreateCertificate(rand.Reader, &template, caCert, &key.PublicKey, caKey)
	if err != nil {
		errerLog.Fatalf("CA证书创建失败: %s", err)
	}
	//保存证书
	saveCert(path, derCert)
	//保存私钥
	saveKey(path, derKey)

	if caCert != &template {
		copyFile(filepath.Join(filepath.Dir(path), "ca.pem"), filepath.Join(path, "ca.pem"))
	} else {
		copyFile(filepath.Join(path, path+".crt"), filepath.Join(path, "ca.pem"))
	}
}

// 拷贝文件
func copyFile(source string, dest string) {
	//打开原文件
	sourceFile, err := os.Open(source)
	if err != nil {
		errerLog.Fatalf("原文件 %s 打开失败:%s", source, err)
	}
	defer sourceFile.Close()
	//打开目标文件
	destFile, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		errerLog.Fatalf("目标文件 %s 打开失败:%s", source, err)
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, sourceFile); err != nil {
		errerLog.Fatalf("拷贝文件 %s 失败:%s", source, err)
	}

}

// 将私钥保存文件
func saveKey(directory string, derKey []byte) {
	fileName := filepath.Join(directory, filepath.Base(directory)+".key")
	keyFiel, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)

	if err != nil {
		errerLog.Fatalf("私钥文件 %s 打开失败:%s", fileName, err)
	}
	defer keyFiel.Close()

	block := &pem.Block{Type: "EC PRIVATE KEY", Bytes: derKey}
	if err := pem.Encode(keyFiel, block); err != nil {
		errerLog.Fatalf("证书 %s 序列化失败:%s", fileName, err)
	}
}

// 保存证书
// 如果保存的是中级CA证书，需要将其父级证书拷贝到中级证书后面
func saveCert(directory string, derCert []byte) {
	//创建文件夹
	createDirectory(directory)
	//ca   ca/ca.crt
	fileName := filepath.Join(directory, filepath.Base(directory)+".crt")
	//打开文件
	certFiel, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		errerLog.Fatalf("文件 %s 打开失败:%s", fileName, err)
	}
	defer certFiel.Close()

	block := &pem.Block{Type: "CERTIFICATE", Bytes: derCert}
	if err := pem.Encode(certFiel, block); err != nil {
		errerLog.Fatalf("证书 %s 序列化失败:%s", fileName, err)
	}
	//如果是中级CA
	if filepath.Dir(directory) != "." {
		//   ca/ica  ca/ca.crt
		newpath := filepath.Join(filepath.Dir(directory), filepath.Base(filepath.Dir(directory)))
		//打开父级证书
		caFile, err := os.Open(newpath + ".crt")
		if err != nil {
			errerLog.Fatalf("证书打开失败:%s", err)
		}
		defer caFile.Close()

		_, err = io.Copy(certFiel, caFile)
		if err != nil {
			errerLog.Fatalf("证书拷贝失败:%s", err)
		}
		err = certFiel.Sync()
		if err != nil {
			errerLog.Fatalf("刷新失败:%s", err)
		}
	}
}

// 创建文件夹
func createDirectory(directory string) {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		if err := os.MkdirAll(directory, 0777); err != nil {
			errerLog.Fatalf("文件夹 ./%s 创建失败: %s", directory, err)
		}
	}
}

// 解析用户输入的证书主题
func parseDn(ca *x509.Certificate, dn string) *pkix.Name {
	var caName pkix.Name
	if ca != nil {
		caName = ca.Subject
	} else {
		caName = pkix.Name{}
	}

	newName := &pkix.Name{}
	//  -dn="/CN=Chain/O=xiongdilian/OU=qukuailian"   CN=Chain/O=xiongdilian/OU=qukuailian
	for _, element := range strings.Split(strings.Trim(dn, "/"), "/") {
		value := strings.Split(element, "=")
		if len(value) != 2 {
			errerLog.Fatalf("参数 %s 非法", element)
		}
		switch strings.ToUpper(value[0]) {
		//名称
		case "CN":
			newName.CommonName = value[1]
		//国家名称
		case "C":
			if value[1] == "" {
				caName.Country = []string{}
			} else {
				newName.Country = append(newName.Country, value[1])
			}
		//地点
		case "L":
			if value[1] == "" {
				caName.Locality = []string{}
			} else {
				newName.Locality = append(newName.Locality, value[1])
			}
			//州或省
		case "ST":
			if value[1] == "" {
				caName.Province = []string{}
			} else {
				newName.Province = append(newName.Province, value[1])
			}
			//组织
		case "O":
			if value[1] == "" {
				caName.Organization = []string{}
			} else {
				newName.Organization = append(newName.Organization, value[1])
			}
			//部门
		case "OU":
			if value[1] == "" {
				caName.OrganizationalUnit = []string{}
			} else {
				newName.OrganizationalUnit = append(newName.OrganizationalUnit, value[1])
			}
		default:
			errerLog.Fatalf("参数 %s 非法", element)
		}
	}
	if ca != nil {
		newName.Country = append(caName.Country, newName.Country...)
		newName.Locality = append(caName.Locality, newName.Locality...)
		newName.Province = append(caName.Province, newName.Province...)
		newName.Organization = append(caName.Organization, newName.Organization...)
		newName.OrganizationalUnit = append(caName.OrganizationalUnit, newName.OrganizationalUnit...)
	}
	return newName
}

// 生成私钥
// 返回值一：私钥
// 返回值二：序列化后的私钥，用于吸入文件
func generatePrivateKey() (*ecdsa.PrivateKey, []byte, error) {
	//生成秘钥对
	key, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	//将私钥系列化
	derKey, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return nil, nil, err
	}
	return key, derKey, nil
}

// 从文件中读取私钥
func parseKey(path string) *ecdsa.PrivateKey {
	newpath := filepath.Join(path, filepath.Base(path)+".key")
	der, err := ioutil.ReadFile(newpath)
	if err != nil {
		errerLog.Fatalf("私钥文件 %s 读取失败:%s", newpath, err)
	}
	//反序列化
	block, _ := pem.Decode(der)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		errerLog.Fatalf("私钥文件 %s 编码失败:%s", newpath, err)
	}

	//从块中解析私钥文件
	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		errerLog.Fatalf("证书文件 %s 解析失败: %s", newpath, err)
	}
	return key
}

// 从文件中读取证书
func parseCert(path string) *x509.Certificate {
	//    ca    ca/ca.crt
	newpath := filepath.Join(path, filepath.Base(path)+".crt")
	der, err := ioutil.ReadFile(newpath)
	if err != nil {
		errerLog.Fatalf("证书文件 %s 读取失败:%s", newpath, err)
	}
	//反序列化
	block, _ := pem.Decode(der)
	if block == nil || block.Type != "CERTIFICATE" {
		errerLog.Fatalf("证书文件 %s 编码失败:%s", newpath, err)
	}

	//从块中解析证书
	crt, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		errerLog.Fatalf("证书文件 %s 解析失败: %s", newpath, err)
	}
	return crt
}

// 检查指定路径下是否存在证书文件，私钥，父级证书文件
func checkExisting(path string) {
	//ca   ca/ca.crt
	fullPath := filepath.Join(path, filepath.Base(path)) //   /ca/ca
	//判断证书是否存在
	if _, err := os.Stat(fullPath + ".crt"); err == nil {
		errerLog.Fatalf("文件:%s 已经存在.", "./"+fullPath+".crt")
	}
	//判断私钥是否存在
	if _, err := os.Stat(fullPath + ".key"); err == nil {
		errerLog.Fatalf("文件:%s 已经存在.", "./"+fullPath+".key")
	}
	//判断父级证书是否存在
	if _, err := os.Stat(fullPath + "ca.pem"); err == nil {
		errerLog.Fatalf("文件:%s 已经存在.", "./"+fullPath+"ca.pem")
	}
}

/*
创建证书
1.用户输入的参数
2.保存路径
3.证书主题
4.主题备用名称
5.证书有效时间
6.证书用途
7.额外用途
*/
func createCertificate(args []string, path, defaultDn, defaultSan string, defaultValidity int, keyUsage x509.KeyUsage, extKeyUsage []x509.ExtKeyUsage) {
	//参数一：名称   参数二：错误处理策略
	fs := flag.NewFlagSet("ca", flag.PanicOnError)
	dn := fs.String("dn", defaultDn, "证书主题")
	san := fs.String("san", defaultSan, "备用主题")
	validity := fs.Int("validity", defaultValidity, "证书有效期")
	overwrite := fs.Bool("overwrite", false, "是否覆盖原有文件")
	//解析参数
	err := fs.Parse(args)
	if err != nil {
		errerLog.Fatalf("将参数解析到命令失败:%s", err)
	}
	//解析之后剩余的参数大于1
	if len(fs.Args()) > 1 {
		errerLog.Fatalf("参数非法:%s", strings.Join(fs.Args(), ","))
	} else if len(fs.Args()) == 1 { //解析之后剩余的参数等于1
		path = fs.Arg(0)
	}
	if !*overwrite {
		checkExisting(path)
	}
	//   C:/ca/ica/aa.txt===>C:/ca/ica
	//    ca==> .
	ca := filepath.Dir(path)
	//读取父级证书
	caCert := parseCert(ca)
	//父级证书不是有效的证书颁发机构
	if !caCert.IsCA {
		errerLog.Fatalf("%s 不是证书颁发机构", filepath.Dir(path))
	}
	//获取父级私钥
	caKey := parseKey(ca)

	//生成私钥
	key, derKey, err := generatePrivateKey()
	if err != nil {
		errerLog.Fatalf("私钥生成失败:%s", err)
	}

	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		errerLog.Fatalf("证书编号生成失败:%s", err)
	}
	notBofor := time.Now().UTC()
	notAfter := notBofor.AddDate(0, 0, *validity)

	template := x509.Certificate{
		//序列号
		SerialNumber: serialNumber,
		//主题
		Subject: *parseDn(caCert, *dn),
		//生效时间
		NotBefore: notBofor,
		//失效时间
		NotAfter:              notAfter,
		BasicConstraintsValid: true,
		//是否是CA
		IsCA: true,
		//证书的用途：用于数字签名和证书签发
		KeyUsage:       keyUsage,
		ExtKeyUsage:    extKeyUsage,
		EmailAddresses: []string{},
		IPAddresses:    []net.IP{},
		DNSNames:       []string{},
	}
	//解析ip，邮箱，DNS
	parseSubjectNames(*san, &template)
	derCert, err := x509.CreateCertificate(rand.Reader, &template, caCert, &key.PublicKey, caKey)
	if err != nil {
		errerLog.Fatalf("证书文件 %s 创建失败:%s", path, err)
	}
	saveCert(path, derCert)
	saveKey(path, derKey)
	//将父级的ca.pem拷贝到自己的文件夹下
	copyFile(filepath.Join(filepath.Dir(path), "ca.pem"), filepath.Join(path, "ca.pem"))

}

// "127.0.0.1,764217451@qq.com"
func parseSubjectNames(san string, template *x509.Certificate) {
	if san != "" {
		for _, h := range strings.Split(san, ",") {
			if ip := net.ParseIP(h); ip != nil {
				template.IPAddresses = append(template.IPAddresses, ip)
			} else if email := parseEmailAddress(h); email != nil {
				template.EmailAddresses = append(template.EmailAddresses, email.Address)
			} else {
				template.DNSNames = append(template.DNSNames, h)
			}
		}
	}
}

func parseEmailAddress(address string) (email *mail.Address) {
	var err error
	email, err = mail.ParseAddress(address)
	if err == nil && email != nil {
		return email
	}
	return nil
}
