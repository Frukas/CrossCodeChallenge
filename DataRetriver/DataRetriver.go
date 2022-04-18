package dataretriver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type NumberSet struct {
	Numbers []float32 `json:"numbers"`
}

var byteChannel = make(chan []byte)

func PageDataRetrive() ([]float32, int) {

	i := 1
	speed := 1
	//speedLImiter sets the maximum of offset HTTP calls.
	speedLimiter := 30
	var DataSet []float32

	for {
		select {
		case value := <-byteChannel:
			if string(value) == `{"numbers":[]}` {
				return DataSet, speed
			} else {
				DataSet = append(DataSet, resultoTofloatSlice(value)...)
			}
			speed++
		default:
			if i < speed+speedLimiter {
				go getPageData(i)
				i++
			}
		}
	}
}

func getPageData(in int) {

	pagesSTR := fmt.Sprintf("http://challenge.dienekes.com.br/api/numbers?page=%d", in)

	resp, err := http.Get(pagesSTR)
	if err != nil {
		// handle error
		fmt.Println("Error when getting the page index: ", in, err)
		go getPageData(in)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error reading the data from site: ", err)
	}

	byteChannel <- body

}

func resultoTofloatSlice(res []byte) []float32 {
	var numberRs NumberSet

	json.Unmarshal(res, &numberRs)

	return numberRs.Numbers

}
