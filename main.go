package main

import (
	"fmt"
	"sync"

	"github.com/fanialfi/golang-simple-client-http-request/request"
)

var wg = sync.WaitGroup{}

func main() {
	totalRequest := 5

	for i := 0; i < totalRequest; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			users, err := request.FetchUsers()
			if err != nil {
				fmt.Println("Error!", err.Error())
				return
			}

			for _, elm := range users {
				fmt.Printf("ID : %s\t\tName : %s\t\tGrade : %d\n", elm.ID, elm.Name, elm.Grade)
			}
		}(i)
	}
	wg.Wait()

	fmt.Println()

	// fetchUserMyself()
	for i := 0; i < totalRequest; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			users, err := request.FetchUsersMyself()
			if err != nil {
				fmt.Println("Error!", err.Error())
				return
			}

			for _, elm := range users {
				fmt.Printf("ID : %s\t\tName : %s\t\tGrade : %d\n", elm.ID, elm.Name, elm.Grade)
			}
		}(i)
	}
	wg.Wait()
}
