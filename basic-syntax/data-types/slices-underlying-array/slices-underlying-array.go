package main

import "fmt"

func main() {

	fmt.Println("Lets create an array x of length 10 and fill that array with numbers 20-29")
	fmt.Println("var x [10]int")
	var x [10]int
	for i := 0; i < len(x); i++ {
		x[i] = i + 20
	}
	fmt.Printf("    Array x: %T len %2v Cap %2v %v\n", x, len(x), cap(x), x)

	fmt.Println("Lets create a slice and fill that slice with numbers 20-29")
	fmt.Println("var y []int")
	var y []int
	for i := 0; i < 10; i++ {
		y = append(y, i+20)
	}
	fmt.Printf("    Slice y:   %T len %2v Cap %2v %v\n", y, len(y), cap(y), y)

	fmt.Println("Lets make a slice with length 10, capacity 10")
	fmt.Println("z := make([]int, 10)")
	z := make([]int, 10)
	for i := 0; i < 10; i++ {
		z[i] = i + 20
	}
	fmt.Printf("    Slice z:   %T len %2v Cap %2v %v\n", z, len(z), cap(z), z)

	fmt.Println("Lets make a slice with length 10, capacity 18")
	fmt.Println("n := make([]int, 10, 18)")
	n := make([]int, 10, 18)
	for i := 0; i < 10; i++ {
		n[i] = i + 20
	}
	fmt.Printf("    Slice n:   %T len %2v Cap %2v %v\n", n, len(n), cap(n), n)

	fmt.Println("Now lets demonstrate appending to a slice of length 0, capacity 5")
	fmt.Println("s := make([]int, 0, 5)")
	s := make([]int, 0, 5)
	for i := 0; i < 20; i++ {
		s = append(s, i+20)
		fmt.Printf("    Slice s:   %T len %2v Cap %2v %v\n", s, len(s), cap(s), s)
	}

	fmt.Println("Now lets demonstrate appending to a slice of length 0, capacity 7")
	fmt.Println("t := make([]int, 0, 7)")
	t := make([]int, 0, 7)
	for i := 0; i < 20; i++ {
		t = append(t, i+20)
		fmt.Printf("    Slice t:   %T len %2v Cap %2v %v\n", t, len(t), cap(t), t)
	}

}
