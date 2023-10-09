package request

import (
	"encoding/json"
	"net/http"
)

// function ini kegunaannya sama seperti fetchUsers()
func FetchUsersMyself() ([]Student, error) {
	url := "http://localhost:8080/users"
	var data []Student

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
