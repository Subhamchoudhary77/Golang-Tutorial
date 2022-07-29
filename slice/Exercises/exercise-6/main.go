// Problem Statement- 1: Consider the following slice declaration: friends := []string{"Marry", "John", "Paul", "Diana"}

// Using copy() function create a copy of the slice. Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.

// Problem Statement 2: Consider the following slice declaration: friends := []string{"Marry", "John", "Paul", "Diana"}

// Using append() function create a copy of the slice. Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.

package main

import "fmt"

func main() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	myFriends := make([]string, len(friends))

	copy(myFriends, friends)

	myFriends[0] = "Subham"

	fmt.Println(friends, myFriends)

	trueFriends := []string{}
	trueFriends = append(trueFriends, friends...)
	trueFriends[0] = "Sergio"
	fmt.Println(friends, trueFriends)
}
