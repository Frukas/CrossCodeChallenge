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

var st = make(chan []byte)

func PageDataRetrive() ([]float32, int) {

	i := 1
	speed := 1
	speedLimiter := 30
	//sts := []string{""}
	var sts []float32

	for {
		select {
		case v := <-st:
			if string(v) == `{"numbers":[]}` {
				return sts, speed
			} else {
				sts = append(sts, resultoTofloatSlice(v)...)
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
		fmt.Println("Err when get the page index: ", in, err)
		go getPageData(in)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Erro no corpo no indice: ", err)
	}

	st <- body

	defer resp.Body.Close()

}

func resultoTofloatSlice(res []byte) []float32 {
	var numberRs NumberSet

	json.Unmarshal(res, &numberRs)

	return numberRs.Numbers

}
