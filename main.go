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

// 파라미터 여러개 받으려면 ... <- 사용한다.
func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	for index, number := range numbers {
		fmt.Println(index, number)
	}

	for _, number := range numbers {
		total += number
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return total
}

func canIDrink(age int) bool {
	if age < 18 {
		return false
	}
	return true
}

// if 절에 변수를 선언할 수 있다. 이는 if-else에 해당하는 영역에서만 사용 가능하다.
func canIDrink2(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func canIDrink3(age int) bool {
	switch koreanAge := age + 2; {
	case koreanAge < 18:
		return false
	case koreanAge == 18:
		return true
	case koreanAge > 18:
		return true
	}
	return false
}

// struct
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	// 초기화 선언 방법 두가지
	name := "seunghyun"
	result := superAdd(1, 2, 3, 4, 5)
	// var name string = "seunghyun"
	something.SayHi()

	totalLength, upperName := lenAndUpper(name)
	totalLength2, upperName2 := lenAndUpper2(name)
	totalLength3, upperName3 := lenAndUpper3(name)

	fmt.Println(totalLength, upperName)
	fmt.Println(totalLength2, upperName2)
	fmt.Println(totalLength3, upperName3)

	fmt.Println(multiply(2, 5))
	fmt.Println(result)

	fmt.Println(canIDrink(18))
	fmt.Println(canIDrink2(18))
	fmt.Println(canIDrink3(18))

	repeatMe("seung", "hyun", "learns", "go", "lang")

	something.SayBye()

	a := 2
	b := a
	fmt.Println(a, b)
	fmt.Println(&a, &b) // 메모리 주소가 다르다는 것을 알 수 있다.

	c := 2
	d := c
	c = 10
	fmt.Println(c, d)
	fmt.Println(&c, &d)

	e := 2
	f := &e
	fmt.Println(&e, f) // 둘다 메모리 주소가 출력된다.
	fmt.Println(e, *f) // *을 사용해야 값이 보인다.

	g := 2
	h := &g
	*h = 10 // *을 사용하면 포인터로 값을 바꿀 수 있다.
	fmt.Println(g)

	// Arrays
	names := [5]string{"seung", "hyun", "mun"}
	names[3] = "learns"
	names[4] = "go lang"
	fmt.Println(names)

	// Slices
	names2 := []string{"seung", "hyun", "mun"}
	names2 = append(names2, "leanrs", "go lang")
	fmt.Println(names2)

	// Maps
	maps := map[string]string{"name": "seunghyun"}
	fmt.Println(maps)

	for key, value := range maps {
		fmt.Println(key, value)
	}

	// using struct
	favFood := []string{"hi", "yo"}
	seunghyun := person{name: "seunghyun", age: 31, favFood: favFood}
	fmt.Println(seunghyun)
}
