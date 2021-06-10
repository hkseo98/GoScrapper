package main // main은 컴파일을 위해서 필요한 것일 뿐

import (
	"fmt"

	"github.io/hkseo98/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "first"
	dictionary.Add(word, "first word")
	err := dictionary.Delete(word)
	if err != nil {
		fmt.Println(err)
		return
	}
	val, err := dictionary.Search(word)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dictionary)
	fmt.Println(val)
}
