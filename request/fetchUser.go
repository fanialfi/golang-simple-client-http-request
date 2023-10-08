package request

import (
	"encoding/json"
	"net/http"
)

var baseUrl string = "http://localhost:8080"

type student struct {
	ID    string
	Name  string
	Grade int
}

// digunakan untuk melakukan request ke http://localhost:8080/users
// yang telah dibuat di https://github.com/fanialfi/golang-web-service-api-server.git
// lalu menerima response dari request tersebut lalu menampilkannya
func FetchUsers() ([]student, error) {
	var err error

	// &http.Client{}
	// digunakan untuk membuat instance (object) dari tipe data `http.Client`
	// diperlukan untuk eksekusi request
	var client = &http.Client{}

	var data []student

	// http.NewRequest(method, url string, body io.Reader)(*Request, error)
	// digunakan untuk membuat request baru menggunakan context.Background
	request, err := http.NewRequest("GET", baseUrl+"/users", nil)
	if err != nil {
		return nil, err
	}

	// *Client.Do(req *Request) (*Response, error)
	// digunakan untuk mengirimkan http request dan mengembalikan http response
	// mengembalikan instance bertipe http.Response,
	// yang didalamnya terdapat informasi yang dikembalikan dari server
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// data response perlu di close setelah tidak dipakai
	defer response.Body.Close()

	// data http.Response.Body merepresentasikan response body
	// response body ini bertipe io.ReadCloser, data-nya bisa di read dan di tutup
	//
	// *json.Decoder digunakan untuk mengkonversi data response body menjadi bentuk json
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
