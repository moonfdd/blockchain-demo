package mycore2

import (
	"fmt"
	"os"
	"strings"

	"github.com/moonfdd/blockchain-demo/examples/a004_eth/b014_gen_interface_and_struct/mycore"
)

func BuildInterface(outPath string) {
	dir := `D:\gvm\.g\go\src\`
	packageNamePrefix := ``
	interfaceInfoList, _ := mycore.LoadDir(dir, packageNamePrefix)
	if len(interfaceInfoList) == 0 {
		return
	}

	imports := make(map[string]struct{})
	content := `package main
import (

{xxxstr}
)

func Test() {
{yyystr}
}
`
	yyystr := ""
	for i := 0; i < len(interfaceInfoList); i++ {
		for j := 0; j < len(interfaceInfoList); j++ {
			if i == j {
				continue
			}
			if interfaceInfoList[i].Package == "fmt" {
				imports["\"fmt\""] = struct{}{}
			} else {
				imports[interfaceInfoList[i].GetAlias()+" \""+interfaceInfoList[i].Package+"\""] = struct{}{}
			}
			yyy := ""
			str := ""
			str = fmt.Sprintln("    if true {")
			yyy += str

			str = fmt.Sprintf("        type T1 = %s.%s\r\n", interfaceInfoList[i].GetAlias(), interfaceInfoList[i].Name)
			yyy += str

			str = fmt.Sprintf("        type T2 = %s.%s\r\n", interfaceInfoList[j].GetAlias(), interfaceInfoList[j].Name)
			yyy += str

			str = fmt.Sprintf("        var a T1 = struct{ T1 }{}\r\n")
			yyy += str

			// str = fmt.Sprintf("        var b T2 = struct{ T2 }{}\r\n")
			// yyy += str

			str = fmt.Sprintf("         if _, ok := a.(T2); ok {\r\n")
			yyy += str

			str = fmt.Sprintf("            fmt.Println(\"%s.%s是%s.%s的子接口\")\r\n", interfaceInfoList[i].GetAlias(), interfaceInfoList[i].Name, interfaceInfoList[j].GetAlias(), interfaceInfoList[j].Name)
			yyy += str

			str = fmt.Sprintln("        }")
			yyy += str

			// str = fmt.Sprintf("         if _, ok := b.(T1); ok {\r\n")
			// yyy += str

			// str = fmt.Sprintf("            fmt.Println(\"%s.%s是%s.%s的子接口\")\r\n", interfaceInfoList[j].GetAlias(), interfaceInfoList[j].Name, interfaceInfoList[i].GetAlias(), interfaceInfoList[i].Name)
			// yyy += str

			// str = fmt.Sprintln("        }")
			// yyy += str

			str = fmt.Sprintln("    }")
			yyy += str
			yyystr += yyy
		}
	}
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
