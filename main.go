package main // main은 컴파일을 위해서 필요한 것일 뿐

import (
	"fmt"
)

func main() {
	names := []string{}
	names = append(names, "dcdc", "dcdv")
	fmt.Println(names)
}
