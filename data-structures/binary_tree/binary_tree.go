// all elements in a binary tree are same
package main

import "fmt"

type tree struct {
	left  *tree
	value int
	right *tree
}

func main() {
	var t *tree
	t=insert(1000,t)
	t=insert(10,t)
	t=insert(2,t)
	t=insert(90,t)
	t=insert(-2,t)
	
	traverse(t)

}

func traverse(t *tree) {
	// fmt.Println("h")
	if t == nil {
		return
	}
	traverse(t.left)
	fmt.Println(t.value)
	traverse(t.right)
}

func insert(x int , t *tree) *tree{
	if t == nil{
		// fmt.Println(x)
		return &tree{nil,x,nil}
		
	}
	if t.value == x{
		return t
	}
	if x<t.value{
		t.left = insert(x,t.left)
		return t
	}
	t.right = insert(x,t.right)
	return t
}
