package main

import "fmt"
import "../ch3"
import "../ch4"

func main() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July",
		8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)
	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}
	endlessSummer := summer[:7]
	fmt.Println(endlessSummer)
	ptr := [32]byte{}
	ch3.One(ptr[:])
	fmt.Println(ptr)
	fmt.Println("--example1--")
	a := [...]int{0, 1, 2, 3, 4, 5}
	ch4.Reverse(a[:2])
	ch4.Reverse(a[2:])
	ch4.Reverse(a[:])
	fmt.Println(a)
	fmt.Println("--example2--")
	var x, y []int
	for i := 0; i < 10; i++ {
		y = ch4.AppendInt(x, i)
		fmt.Printf("%d ycap=%d xcap=%d %v\n", i, cap(y), cap(x), y)
		x = y
	}
	fmt.Println("--example3--")
	data := []string{"one", "", "three"}
	fmt.Println(ch4.Nonempty(data))
	fmt.Println(data)

}
