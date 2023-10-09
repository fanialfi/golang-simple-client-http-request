package main

import (
	"flag"
	"fmt"

	"github.com/fanialfi/golang-simple-client-http-request/request/form"
)

// var wg = sync.WaitGroup{}

// func main() {
// 	totalRequest := 5

// 	for i := 0; i < totalRequest; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			users, err := request.FetchUsers()
// 			if err != nil {
// 				fmt.Println("Error!", err.Error())
// 				return
// 			}

// 			for _, elm := range users {
// 				fmt.Printf("ID : %s\t\tName : %s\t\tGrade : %d\n", elm.ID, elm.Name, elm.Grade)
// 			}
// 		}(i)
// 	}
// 	wg.Wait()

// 	fmt.Println()

// 	// fetchUserMyself()
// 	for i := 0; i < totalRequest; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			users, err := request.FetchUsersMyself()
// 			if err != nil {
// 				fmt.Println("Error!", err.Error())
// 				return
// 			}

// 			for _, elm := range users {
// 				fmt.Printf("ID : %s\t\tName : %s\t\tGrade : %d\n", elm.ID, elm.Name, elm.Grade)
// 			}
// 		}(i)
// 	}
// 	wg.Wait()
// }

// // implementasi form.FetchUser()
// func main() {
// 	id := "E001"
// 	totalRequest := 5

// 	for i := 0; i < totalRequest; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			user, err := form.FetchUser(id)
// 			if err != nil {
// 				fmt.Println("Error!", err.Error())
// 				return
// 			}
// 			fmt.Printf("ID: %s\tName: %s\tGrade: %d\n", user.ID, user.Name, user.Grade)
// 		}(i)
// 	}
// 	wg.Wait()
// }

// implementasi form.FetchUserMyself()
func main() {
	// id := "E002"
	var id = flag.String("id", "E001", "masukkan parameter id")
	flag.Parse()

	data, err := form.FetchUsersMyself(*id)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("ID: %s\tName: %s\tGrade: %d\n", data.ID, data.Name, data.Grade)
}
