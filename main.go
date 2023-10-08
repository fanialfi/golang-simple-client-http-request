package main

import (
	"fmt"
)

func main() {
	users, err := fetchUsers()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	for _, elm := range users {
		fmt.Printf("ID : %s\t\tName : %s\t\tGrade : %d\n", elm.ID, elm.Name, elm.Grade)
	}

	fmt.Println()

	// fetchUserMyself()
	users, err = fetchUsersMyself()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	for _, elm := range users {
		fmt.Printf("ID : %s\t\tName : %s\t\tGrade : %d\n", elm.ID, elm.Name, elm.Grade)
	}
}
