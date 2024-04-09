package pointers_and_memo

import (
	"fmt"
	"strconv"
)

// ===================== Structs and methods for  ====================

type User struct {
	userName, email string
	age             int
}

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil}
}

func (ll *LinkedList) append(data int) {
	newNode := &Node{data: data, next: ll.head}
	ll.head = newNode
}

func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d", current.data)
		current = current.next
	}
	fmt.Println("")
}

func (u *User) Create(a, b string, c int) {
	cs := strconv.Itoa(c)
	u.userName = a
	u.email = b
	u.age = c

	s := make(map[string]string)
	s["User Name: "] = a
	s["Email: "] = b
	s["Age: "] = cs

	for key, value := range s {
		fmt.Printf("%s: %s\n", key, value)
	}
}

//=========================================================================

func Pointers() {

	fmt.Println("-------------------- Pointers and Memory Managemnt -----------------------")

	// - Create a program that demonstrates the use of pointers to swap the values of two variables: ✅
	fmt.Println(" - Create a program that demonstrates the use of pointers to swap the values of two variables.")
	// Vars
	x := "Before "
	y := "Switch"
	z := "After "
	var val1, val2 int
	val1 = 5
	val2 = 2024
	fmt.Println(x+y, "\nvalue 1: ", val1, "\nvalue 2: ", val2)

	// Function
	swap := func(a, b *string, v1, v2 *int) {
		temp := *a
		*a = *b
		*b = temp

		temp2 := *v1
		*v1 = *v2
		*v2 = temp2
	}

	// Calling the function
	swap(&x, &z, &val1, &val2)

	// Printing
	fmt.Println(x+y, "\nvalue 1: ", val1, "\nvalue 2: ", val2)

	// - Write a function that accepts a pointer to an integer and increments its value : ✅
	fmt.Println("- Write a function that accepts a pointer to an integer and increments its value.")
	// variable
	x2 := 0
	// Function
	incrementPointer := func(a *int) {
		*a++
	}
	// Calling function
	incrementPointer(&x2)
	// Printing
	fmt.Println(x2)

	//  - Implement a linked list data structure using pointers to manage memory dynamically ✅
	fmt.Println(" - Implement a linked list data structure using pointers to manage memory dynamically.")
	ll := NewLinkedList()
	ll.append(11)
	ll.append(12)
	ll.append(13)
	ll.append(33)
	ll.append(32)
	ll.append(31)
	fmt.Println("Linked List:")
	ll.Display()
}

func DynamicAllocation() {

	// - Memory Allocation: Using the `new` and `make` functions for dynamic memory allocation: ❌

	fmt.Println(" - Memory Allocation: Using the `new` and `make` functions for dynamic memory allocation")

	user := User{
		userName: "User Name",
		email:    "default@email.com",
		age:      24,
	}
	fmt.Println(user)
	user.Create("Djahmet", "djahmet@mail.com", 33)
	fmt.Println(user)
	u := new(User)
	u.Create("BOLO", "pinkoko@gmail.com", 28)

}
