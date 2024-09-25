package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (us UserSlice) Len() int {
	return len(us)
}

func (us UserSlice) Less(i, j int) bool {
	return us[i].Age < us[j].Age
}

func (us UserSlice) Swap(i, j int) {
	// temp := us[i]
	// us[i] = us[j]
	// us[j] = temp

	// Cara yang lebih gampang
	us[i], us[j] = us[j], us[i]
}

func main() {
	users := []User{
		{"Eko", 30},
		{"Budi", 35},
		{"Joko", 40},
		{"Adit", 28},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
