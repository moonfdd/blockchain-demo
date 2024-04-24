package mycore

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
)

// 接口信息
type InterfaceInfo struct {
	// 接口名
	Name string
	// 包路径
	Package string
	// 别名
	Alias string
}

func (i *InterfaceInfo) GetAlias() string {
	if i.Package == "fmt" {
		return "fmt"
	}
	hash := md5.Sum([]byte(i.Package))
	return i.Alias + "_" + hex.EncodeToString(hash[:])
}

// 结构体信息
type StructInfo struct {
	// 结构体名
	Name string
	// 包路径
	Package string
	// 别名
	Alias string
}

func (i *StructInfo) GetAlias() string {
	if i.Package == "fmt" {
		return "fmt"
	}
	hash := md5.Sum([]byte(i.Package))
	return i.Alias + "_" + hex.EncodeToString(hash[:])
}

// 同一个目录下的go文件列表，只要大写的接口和结构体，并且都不是泛型
func LoadPathList(pathList []string) (alias string, interfaceNameList []string, structNameList []string) {
	for _, path := range pathList {
		p, a, b := LoadPath(path)
		alias = p
		if len(a) > 0 {
			interfaceNameList = append(interfaceNameList, a...)
		}
		if len(b) > 0 {
			structNameList = append(structNameList, b...)
		}
	}
	return
}

// 单个go文件，只要大写的接口和结构体，并且都不是泛型
func LoadPath(path string) (packageName string, interfaceNameList []string, structNameList []string) {
	content, err := os.ReadFile(path)
	// fmt.Println(path)
	if err != nil {
		fmt.Println("read file error", err)
		return
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", content, parser.ParseComments)
	if err != nil {
		fmt.Println("parse file error", err, path)
		return
	}
	// fmt.Println(f.Name) //package名称
	if f.Name.Name == "main" {
		return
	}
	packageName = f.Name.Name

	for i := 0; i < len(f.Decls); i++ {
		// fmt.Printf("%T\r\n", f.Decls[i])
		// continue
		decl, ok := f.Decls[i].(*ast.GenDecl)
		if !ok {
			// fmt.Println("not a type decl 1", f.Decls[i])
			continue
		}
		// fmt.Println("type decl 1", f.Decls[i])
		e, _ := json.MarshalIndent(decl, "", "  ")
		_ = e

		if decl.Tok != token.TYPE {
			// fmt.Println("not a type decl 2", string(e))
			continue
		}
		if len(decl.Specs) != 1 {
			// fmt.Println("not a type decl 3", string(e))
			continue
		}

		spec, ok2 := decl.Specs[0].(*ast.TypeSpec)
		if !ok2 {
			// fmt.Println("not a type decl 4", string(e))
			continue
		}
		interfaceType, ok3 := spec.Type.(*ast.InterfaceType)
		if !ok3 {
			// fmt.Println("not a type decl 5", spec.Name.Name)
			// continue
		} else {
			// fmt.Println("interfaceType", interfaceType)
			_ = interfaceType
			// 如果接口是泛型，过滤
			if spec.TypeParams != nil {
				continue
			}

			//接口小写，过滤
			if len(spec.Name.Name) == 0 {
				continue
			}
			bytesname := []byte(spec.Name.Name)
			if bytesname[0] < 'A' || bytesname[0] > 'Z' {
				continue
			}

			interfaceNameList = append(interfaceNameList, spec.Name.Name)
		}

		structType, ok4 := spec.Type.(*ast.StructType)
		if !ok4 {
			// fmt.Println("not a type decl 6", decl)
			// continue
		} else {
			_ = structType
			// fmt.Println(spec.Name.Name, "结构体", structType)
			// for i := 0; i < len(decl.Specs); i++ {
			// 	fmt.Println(spec.Name.Name, "结构体")
			// 	// fmt.Println("spec.Doc = ", spec.Doc)
			// 	// fmt.Println("spec.Name = ", spec.Name)
			// 	fmt.Println("spec.TypeParams.Closing = ", spec.TypeParams.Closing)
			// 	fmt.Println("spec.TypeParams.List = ", spec.TypeParams.List)
			// 	fmt.Println("spec.TypeParams.Opening = ", spec.TypeParams.Opening)
			// 	// fmt.Println("spec.Assign = ", spec.Assign)
			// 	// fmt.Println("spec.Type = ", spec.Type)
			// 	// fmt.Println("spec.Comment = ", spec.Comment)
			// 	fmt.Println("---------")
			// 	// Doc        *CommentGroup // associated documentation; or nil
			// 	// Name       *Ident        // type name
			// 	// TypeParams *FieldList    // type parameters; or nil
			// 	// Assign     token.Pos     // position of '=', if any
			// 	// Type       Expr          // *Ident, *ParenExpr, *SelectorExpr, *StarExpr, or any of the *XxxTypes
			// 	// Comment    *CommentGroup // line comments; or nil
			// }
			// 如果结构体是泛型，过滤
			if spec.TypeParams != nil {
				continue
			}

			// 结构体小写，过滤
			if len(spec.Name.Name) == 0 {
				continue
			}
			bytesname := []byte(spec.Name.Name)
			if bytesname[0] < 'A' || bytesname[0] > 'Z' {
				continue
			}

			structNameList = append(structNameList, spec.Name.Name)
		}
	}
	return
}

// 检查结构体是否是泛型
func isGeneric(t types.Type) bool {
	// 这里可以通过检查是否有类型参数来判断是否是泛型
	// 但是由于Go不支持真正的泛型，我们只能通过接口来模拟，
	// 这里我们简单检查是否实现了任何接口。
	if _, ok := t.Underlying().(*types.Interface); ok {
		return true
	}
	return false
}
