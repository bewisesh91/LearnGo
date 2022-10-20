package main

import (
	"fmt"
	"strings"

	"github.com/bewisesh91/learngo/something"
)

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func lenAndUpper2(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func lenAndUpper3(name string) (length int, uppercase string) {
	defer fmt.Println("function done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	name := "seunghyun"
	// var name string = "seunghyun"
	something.SayHi()

	totalLength, upperName := lenAndUpper(name)
	totalLength2, upperName2 := lenAndUpper2(name)
	totalLength3, upperName3 := lenAndUpper3(name)

	fmt.Println(totalLength, upperName)
	fmt.Println(totalLength2, upperName2)
	fmt.Println(totalLength3, upperName3)

	fmt.Println(multiply(2, 5))

	repeatMe("seung", "hyun", "learns", "go", "lang")

	something.SayBye()
}
