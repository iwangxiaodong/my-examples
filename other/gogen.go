package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("gen....")

	if len(os.Args) > 1 {
		filePath := os.Args[1]
		fmt.Println(filePath)
		genDefaultPageSourceFile(filePath)
	}
}

func genDefaultPageSourceFile(sourceFilePath string) {

	//os.RemoveAll(sourceFilePath + "/generated.go")
	fulllFilePath := filepath.Join("/go/src/func/", sourceFilePath)
	b, _ := PathExists(fulllFilePath)
	if !b {
		fmt.Println("文件不存在" + fulllFilePath)
		return
	}

	//	判断编程语言!!!	后续通过func.yaml的runtime来判断
	b1, _ := PathExists("/go/src/func/func.go")
	if !b1 {
		fmt.Println("只支持go编程语言自动生成默认首页！")
		return
	}

	f, err := ioutil.ReadFile(fulllFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ee := hex.EncodeToString(f)
	//fmt.Println(ee)

	sourceCode := strings.Replace(rootDefaultPageTemplate, "@@@", ee, -1)
	//fmt.Println(sourceCode)

	var data = []byte(sourceCode)

	err2 := ioutil.WriteFile("/go/src/func/generated.go", data, 0666)
	if err2 != nil {
		fmt.Printf("%s", err2)
	}
}

const rootDefaultPageTemplate = `// This file is automatically generated.
// Modifications to this file will be overwritten.
package main

import (
	"encoding/hex"
)

// RootDefaultPage - Root default page
func RootDefaultPage() string {
	ret, _ := hex.DecodeString("@@@")
	return string(ret)
}`

// PathExists -
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
