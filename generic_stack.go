package main

import "fmt"

type stack[T any] struct {
	data []T
}

func (t *stack[T]) push(v ...T) {

	t.data = append(t.data, v...)

}

func (t *stack[T]) pop() {
	if len(t.data) != 0 {
		ind := len(t.data) - 1
		fmt.Printf("poping... %v\n", t.data[ind])
		t.data = t.data[:ind]
	}

	fmt.Println("stack is empty ")

}

func (t *stack[T]) empty() {

	if len(t.data) > 0 {
		t.data = []T{}
	}

}

func main() {

	// create int type stack
	s1 := stack[int]{}
	s1.push(3, 6) // adding element
	fmt.Printf(" stack after element added\n %v\n", s1.data)
	s1.pop()
	fmt.Printf(" stack after one element poped \n %v\n", s1.data)
	s1.empty()
	fmt.Printf(" stack get cleaned \n %v\n", s1.data)

	// creating string type stack
	s2 := stack[string]{}
	s2.push("boru", "adil") // adding element
	fmt.Printf(" stack after element added\n %v\n", s2.data)
	s1.pop()
	fmt.Printf(" stack after one element poped \n %v\n", s2.data)
	s2.empty()
	fmt.Printf(" stack get cleaned \n %v\n", s2.data)

	// creating string type stack
	s3 := stack[float64]{}
	s3.push(64.45, 100.45) // adding element
	fmt.Printf(" stack after element added\n %v\n", s3.data)
	s1.pop()
	fmt.Printf(" stack after one element poped \n %v\n", s3.data)
	s3.empty()
	fmt.Printf(" stack get cleaned \n %v\n", s3.data)

}
