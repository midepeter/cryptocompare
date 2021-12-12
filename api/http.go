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

//price handler over http
func Pricehandler(w http.ResponseWriter, r *http.Request) {
	apistr := config.GetConfig()
	var data db.Result
	r.Header.Add("Authorization", fmt.Sprintln("Apikey "+apistr.Key.Apikey))
	resp, err := http.Get("https://min-api.cryptocompare.com/data/pricemultifull?+fsyms=" + fsyms[0] + "&tsyms=" + tsyms[0])
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
	fmt.Print(data)
}
