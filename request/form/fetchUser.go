package form

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/fanialfi/golang-simple-client-http-request/request"
)

func FetchUser(ID string) (request.Student, error) {
	var err error
	var client = &http.Client{}
	var data request.Student

	// url.Values{} akan menghasilkan object yang digunakan sebagai form data request
	var param = url.Values{}

	// Set() akan mensetting parameter pada url (example => www.example.com?id=ID)
	// saat ini bentuk param masih dalam map[string][]string
	// maka perlu di encode dulu menggunakan param.Encode()
	param.Set("id", ID)

	// pada statement ini sebelum data param di encode menjadi bentuk *byte.Buffer
	// param di encode dulu dengan cara panggil function Encode() (param.Encode())
	// karena function byte.NewBufferString() menerima argument dalam bentuk string
	var payload = bytes.NewBufferString(param.Encode())

	// digunakan membuat request baru menggunakan context.Background
	// data payload dapat digunakan sebagai argument sebagai parameter ketiga
	// karena *bytes.Buffer adalah salah satu implementasi dari io.Reader
	// salah satu bukti implementasi *bytes.Buffer merupakan implementasi dari interface io.Reader adalah
	// *bytes.Buffer memiliki method Read(p []byte)(n int, err error) yang diperlukan oleh interface io.Reader
	//
	// karena disini menggunakan request method POST pastika juga web server yang menangani endpoint yang terkait juga mendukung / bisa menangani method POST
	request, err := http.NewRequest("GET", request.BaseUrl+"/user", payload)
	if err != nil {
		fmt.Println("38", err.Error())
		return data, err
	}

	// karena data yang akan dikirim di encode, maka pada request header perlu di set Content-Type nya yang sesuai
	// disini Content-Type di set sebagai "application/x-www-form-urlencoded"
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	fmt.Println("request url is", request.URL)

	// client.Do() digunakan untuk mengirimkan http.Request dan mengembailkan http.Response
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("49", err.Error())
		return data, err
	}

	// tipe data response.Body adalah interface io.ReadCloser
	defer response.Body.Close()

	// jek terlebih dahulu response code nya apakah 200 OK
	if response.StatusCode == http.StatusOK {
		// karena data dari response.Body adalah []byte json, maka perlu di decoding ke tipe data aslinya
		// dimana tipe data response nya adalah struct Student
		// dan variabel data adalah instance dari struct Student
		err = json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			return data, fmt.Errorf("line 67 %s ", err.Error())
		}
	} else {
		// jika response code selain 200 ok, maka data response di baca secara manual
		dataByte := make([]byte, response.ContentLength)
		n, err := response.Body.Read(dataByte)
		if err != nil {
			fmt.Printf("line 74 %s", err.Error())
		}
		fmt.Println(string(dataByte[:n]))
	}

	return data, nil
}
