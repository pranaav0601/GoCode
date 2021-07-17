package main

import "fmt"

func main() {
	var array1 = []string{"s1", "s2", "s3"}
	var array2 []string = reversear(array1)
	fmt.Print(array2)
}
func reversear(ar []string) []string {
	var length int = len(ar) - 1
	var rev = make([]string, length+1)
	for i := len(ar) - 1; i >= 0; i-- {
		fmt.Println(ar[i])
		rev[len(ar)-1-i] = ar[i]
	}
	return rev
}
