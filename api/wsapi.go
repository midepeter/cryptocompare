package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"midepeter/devtest/config"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func Open() {
	fmt.Println("Websocket connection called")
	var h = http.Header{}
	api := config.GetConfig()
	url := "wss://streamer.cryptocompare.com/v2/price"
	h.Add("Authorization", fmt.Sprintln("Apikey "+api.Key.Apikey))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c, _, err := websocket.DefaultDialer.Dial(url, h)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			fmt.Fprintf(os.Stdout, "%s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	f, err := os.Open("sub.json")
	if err != nil {
		fmt.Println(err)
	}
	fbyte, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, fbyte)
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

//price handler over websockets
func PricehandlerWs(w http.ResponseWriter, r *http.Request) {
	Open()
}
