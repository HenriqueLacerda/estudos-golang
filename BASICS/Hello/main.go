package main

// func main() {

// 	fmt.Println("Hello ")
// 	print("Hello ")

// 	fmt.Println("go" + "lang")

//     fmt.Println("1+1 =", 1+1)
//     fmt.Println("7.0/3.0 =", 7.0/3.0)

//     fmt.Println(true && false)
//     fmt.Println(true || false)
//     fmt.Println(!true)

// }
import "fmt"

func main() {
	// array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	array := []int{3}
	target := 10
	var buffer []int = []int{}

	for i := 0; i < len(array)-1; i++ {
		for x := i + 1; x < len(array)-1; x++ {
			if array[i]+array[x] == target {
				buffer = append(buffer, array[i], array[x])
			}
		}
	}
	// return buffer
	fmt.Printf("len=%d cap=%d %v\n", len(buffer), cap(buffer), buffer)
}
