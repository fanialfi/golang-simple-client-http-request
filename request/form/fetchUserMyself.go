package form

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/fanialfi/golang-simple-client-http-request/request"
)

func FetchUsersMyself(ID string) (request.Student, error) {
	var data request.Student
	client := &http.Client{}

	param := url.Values{}
	param.Set("id", ID)

	fmt.Println(request.BaseUrl + "/user?" + param.Encode())
	request, err := http.NewRequest("GET", request.BaseUrl+"/user?"+param.Encode(), nil)
	if err != nil {
		return data, fmt.Errorf("line 21 %s ", err.Error())
	}

	response, err := client.Do(request)
	if err != nil {
		return data, fmt.Errorf("line 26 %s ", err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		err = json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			return data, fmt.Errorf("line 33 %s ", err.Error())
		}
	} else {
		fmt.Println(response.Status)
		dataByte := make([]byte, response.ContentLength)
		n, err := response.Body.Read(dataByte)
		if err != nil {
			return data, fmt.Errorf("line 39 %s ", err.Error())
		}
		fmt.Println(string(dataByte[:n]))
	}

	return data, nil
}
