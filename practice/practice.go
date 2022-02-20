package main

// import "fmt"

// func main() {
// 	var var1 int = 10
// 	// & を用いることで変数のアドレスの取得が可能です。
// 	// ポインタ型変数名の前に * をつけることで変数の中身へのアクセスが可能です。
// 	var var2 *int = &var1
// 	fmt.Println(var1) //=> 10
// 	fmt.Println(&var1) //=> 0xc00010c008
// 	fmt.Println(var2) //=> 0xc00010c008
//     fmt.Println(*var2) //=> 10
// }


// import (
// 	"fmt"
// )

// type Person struct {
// 	firstName string 
// 	age int
// }

// func main() {
// 	// フィールドにセットする
// 	var mike Person
// 	mike.firstName = "Mike"
// 	mike.age = 20
// 	fmt.Println(mike.firstName, mike.age)

// 	bob := Person{"Bob", 30}
// 	fmt.Println(bob.firstName, bob.age)

// 	sam := Person{age: 15, firstName: "Sam"}
// 	fmt.Println(sam.firstName, sam.age)
	
// } 