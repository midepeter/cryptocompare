package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"midepeter/devtest/config"
	"midepeter/devtest/db"
)

var fsyms = []string{"BTC", "XRP", "ETH", "BCH", "EOS", "LTC", "XMR", "DASH"}
var tsyms = []string{"USD", "EUR", "GBP", "JPY", "RUR"}

func Pricehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(fsyms[0])
	apistr := config.GetConfig()
	var data db.Result
	r.Header.Add("Authorization", fmt.Sprintln("Apikey "+apistr.Key.Apikey))
	str := fmt.Sprintf("https://min-api.cryptocompare.com/data/pricemultifull?fsyms=%v&tsyms=%v", fsyms[0], tsyms[0])
	resp, err := http.Get(str)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("unable to unmarshal data")
		panic(err)
	}
	fmt.Fprintln(w, "", data)
}
