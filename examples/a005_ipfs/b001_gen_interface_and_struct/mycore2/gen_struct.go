package mycore2

import (
	"fmt"
	"os"
	"strings"

	"github.com/moonfdd/blockchain-demo/examples/a005_ipfs/b001_gen_interface_and_struct/mycore"
)

func BuildStruct(outPath string) {
	dir := `D:\gvm\.g\go\src\`
	packageNamePrefix := ``
	interfaceInfoList, structInfoList := mycore.LoadDir(dir, packageNamePrefix)
	if len(interfaceInfoList) == 0 {
		fmt.Println("no interface")
		return
	}
	if len(structInfoList) == 0 {
		fmt.Println("no struct")
		return
	}
	if true {
		dir2 := `E:\gosoft\code\kubo`
		packageNamePrefix2 := `github.com/ipfs/kubo`
		interfaceInfoList2, structInfoList2 := mycore.LoadDir(dir2, packageNamePrefix2)
		if len(interfaceInfoList2) == 0 {
			return
		}
		if len(structInfoList2) == 0 {
			return
		}
		interfaceInfoList = append(interfaceInfoList, interfaceInfoList2...)
		structInfoList = structInfoList2
	}
	fmt.Println(len(interfaceInfoList), len(structInfoList))

	imports := make(map[string]struct{})
	content := `package main
import (
{xxxstr}
)

func Test2() {
{yyystr}
}
`

	yyystr := ""
	count := 0
	// for i := 0; i < len(structInfoList); i++ {
	for i := 0; i < len(structInfoList); i++ {
		for j := 0; j < len(interfaceInfoList); j++ {
			count++
			if structInfoList[i].Package == "fmt" {
				imports["\"fmt\""] = struct{}{}
			} else {
				imports[structInfoList[i].GetAlias()+" \""+structInfoList[i].Package+"\""] = struct{}{}
			}
			if interfaceInfoList[j].Package == "fmt" {
				imports["\"fmt\""] = struct{}{}
			} else {
				imports[interfaceInfoList[j].GetAlias()+" \""+interfaceInfoList[j].Package+"\""] = struct{}{}
			}
			// imports[structInfoList[i].GetAlias()+" \""+structInfoList[i].Package+"\""] = struct{}{}
			// imports[interfaceInfoList[j].GetAlias()+" \""+interfaceInfoList[j].Package+"\""] = struct{}{}
			yyy := ""
			str := ""
			str = fmt.Sprintln("    if true {")
			yyy += str

			str = fmt.Sprintf("        type T1 = %s.%s\r\n", structInfoList[i].GetAlias(), structInfoList[i].Name)
			yyy += str

			str = fmt.Sprintf("        type T2 = %s.%s\r\n", interfaceInfoList[j].GetAlias(), interfaceInfoList[j].Name)
			yyy += str

			str = fmt.Sprintf("        var a interface{} = new (T1)\r\n")
			yyy += str

			str = fmt.Sprintf("         if _, ok := a.(T2); ok {\r\n")
			yyy += str

			str = fmt.Sprintf("            fmt.Println(\"%s.%s是%s.%s的子结构体\")\r\n", structInfoList[i].GetAlias(), structInfoList[i].Name, interfaceInfoList[j].GetAlias(), interfaceInfoList[j].Name)
			yyy += str

			str = fmt.Sprintln("        }")
			yyy += str

			str = fmt.Sprintln("    }")
			yyy += str
			yyystr += yyy
		}
	}
	fmt.Println("个数：", len(interfaceInfoList), len(structInfoList), count)
	// for _, interfaceInfo := range interfaceInfoList {
	// 	imports[interfaceInfo.GetAlias()+" \""+interfaceInfo.Package+"\""] = struct{}{}
	// 	yyy := ""
	// 	str := ""
	// 	str = fmt.Sprintln("    if true {")
	// 	yyy += str

	// 	str = fmt.Sprintf("        var a %s.%s\r\n", interfaceInfo.GetAlias(), interfaceInfo.Name)
	// 	yyy += str

	// 	str = fmt.Sprintf("        _ = a\r\n")
	// 	yyy += str

	// 	str = fmt.Sprintln("    }")
	// 	yyy += str
	// 	yyystr += yyy
	// }

	xxxstr := ""
	for k, _ := range imports {
		xxxstr += fmt.Sprintf("    %s\r\n", k)
	}
	// fmt.Println(xxxstr)
	// return
	content = strings.ReplaceAll(content, `{xxxstr}`, xxxstr)
	content = strings.ReplaceAll(content, `{yyystr}`, yyystr)
	// fmt.Println(content)
	f, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		panic(err)
	}

}
