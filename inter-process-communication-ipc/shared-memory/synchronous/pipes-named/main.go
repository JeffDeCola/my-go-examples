package main

func foo(p *int) int {
	*p = 3
	return *p
}

func bar(n int) int {
	return n + foo(&n)
}

func main() {
	println(bar(2))
}
