package mycore

import (
	"container/list"
	"os"
	"path/filepath"
	"strings"
)

func LoadDir(dir string, packagePrefix string) (interfaceInfoList []*InterfaceInfo, structInfoList []*StructInfo) {
	queue := list.New()
	queue.PushBack(dir)
	for queue.Len() > 0 {
		firstDir := queue.Front().Value.(string)
		queue.Remove(queue.Front())
		//获取当前目录下的文件或目录名(包含路径)
		filepathNames, _ := filepath.Glob(filepath.Join(firstDir, "*"))
		pathList := make([]string, 0)
		for i := range filepathNames {
			// filepathNames[i] = strings.ReplaceAll(filepathNames[i], "\\", "/")
			fileInfo, _ := os.Stat(filepathNames[i])
			if fileInfo.IsDir() {
				if fileInfo.Name() == "internal" {
					continue
				}
				if fileInfo.Name() == "testdata" {
					continue
				}
				if fileInfo.Name() == "vender" {
					continue
				}
				if fileInfo.Name() == "vendor" {
					continue
				}
				// if fileInfo.Name() == "arena" {
				// 	continue
				// }
				// if fileInfo.Name() == "syscall" {
				// 	continue
				// }
				// if fileInfo.Name() == "syslog" {
				// 	continue
				// }
				// if fileInfo.Name() == "cgo" {
				// 	continue
				// }
				queue.PushBack(filepathNames[i])
				// fmt.Println(filepathNames[i])
			} else {
				//文件集合
				if strings.HasSuffix(filepathNames[i], "_test.go") {
					continue
				}
				if strings.HasSuffix(filepathNames[i], "_unix.go") {
					continue
				}
				if strings.HasSuffix(filepathNames[i], "_nofuse.go") {
					continue
				}
				if strings.HasSuffix(filepathNames[i], "_darwin.go") {
					continue
				}
				if strings.HasSuffix(filepathNames[i], "_notsupp.go") {
					continue
				}
				if !strings.HasSuffix(filepathNames[i], ".go") {
					continue
				}
				pathList = append(pathList, filepathNames[i])
			}
		}
		if len(pathList) == 0 {
			continue
		}
		var alias string
		var interfaceNameList []string
		var structNameList []string
		alias, interfaceNameList, structNameList = LoadPathList(pathList)
		if alias == "" || alias == "main" {
			continue
		}
		if len(interfaceNameList) == 0 && len(structNameList) == 0 {
			continue
		}
		temp := strings.ReplaceAll(firstDir, dir, packagePrefix)
		temp = strings.ReplaceAll(temp, "\\", "/")
		// fmt.Println(firstDir)
		// fmt.Println(dir)
		// fmt.Println(packagePrefix)
		// fmt.Println(temp)
		// fmt.Println("-----")
		if len(interfaceNameList) > 0 {
			for _, interfaceName := range interfaceNameList {
				if interfaceName == "Ordered" && temp == "cmp" {
					continue
				}
				if interfaceName == "RoutingMessage" && temp == "syscall" {
					continue
				}
				interfaceInfoList = append(interfaceInfoList, &InterfaceInfo{
					Alias:   alias,
					Name:    interfaceName,
					Package: temp,
				})
			}
		}
		if len(structNameList) > 0 {
			for _, structName := range structNameList {
				structInfoList = append(structInfoList, &StructInfo{
					Alias:   alias,
					Name:    structName,
					Package: temp,
				})
			}
		}

	}
	// fmt.Println(len(interfaceInfoList), len(structInfoList))
	// os.Exit(0)
	return
}
