// my-go-examples pointers2.go

package main

import "fmt"

func byval(q *int) {
	fmt.Printf("   byval -- q type=%T: &q=%p q=&i=%p  *q=i=%v\n", q, &q, q, *q)
	*q = 33
	fmt.Printf("   byval -- q type=%T: &q=%p q=&i=%p  *q=i=%v\n", q, &q, q, *q)
	q = nil
}

func main() {

	i := int(42)

	//The original int type
	fmt.Printf("\n1. main  -- i type= %T: &i=%p i=%v\n", i, &i, i)

	// Make a pointer to i
	p := &i
	fmt.Printf("2. main  -- p type=%T: &p=%p p=&i=%p  *p=i=%v\n", p, &p, p, *p)
	fmt.Println("")

	// Pass address i to function by "value" (but we say reference)
	byval(&i)
	fmt.Println("")

	fmt.Printf("3. main  -- i type= %T: &i=%p i=%v\n", i, &i, i)
	fmt.Printf("4. main  -- p type=%T: &p=%p p=&i=%p  *p=i=%v\n", p, &p, p, *p)
	fmt.Println("")

	// Reset i
	i = 42
	// Pass ptr p to function by "value" (but we say reference)
	byval(p)
	fmt.Println("")

	fmt.Printf("5. main  -- i type= %T: &i=%p i=%v\n", i, &i, i)
	fmt.Printf("6. main  -- p type=%T: &p=%p p=&i=%p  *p=i=%v\n\n", p, &p, p, *p)
}
