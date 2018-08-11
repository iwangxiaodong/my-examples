package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
        gen()
	fmt.Println("gen!")
}

func gen() {
	var data = []byte("package main\r\nfunc RootDefaultPage() string {\r\n	return \"demo-index\"\r\n}")

	err2 := ioutil.WriteFile("/go/src/func/generated.go", data, 0666)
	if err2 != nil {
		fmt.Printf("%!!(MISSING)!(MISSING)s(MISSING)", err2)
	}
}
