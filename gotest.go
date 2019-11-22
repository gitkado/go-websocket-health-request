package main

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

// websocketリクエスト情報ストラクト
type Req struct {
	Jsonrpc string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  map[string]string `json:"params"`
}

var addr = flag.String("addr", "127.0.0.1", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	// 現在時刻
	now_time := time.Now()
	// 無限ループ終了時刻
	ed_time := now_time.Add(3 * time.Minute)
	// websocketリクエスト情報
	req := Req{
		Jsonrpc: "2.0",
		Method:  "health",
		Params: map[string]string{
			"date": now_time.String(),
		},
	}

	for {
		// struct->json
		jsonBytes, _ := json.Marshal(req)
		jsonStr := string(jsonBytes)
		log.Println("req:", jsonStr)

		// {"jsonrpc": "2.0", "method": "health", "params": {"date": "2019-11-22 19:23:30.757800829 +0900 JST"}}
		if err := c.WriteMessage(websocket.TextMessage, []byte(jsonStr)); err != nil {
			log.Println("write health:", err)
			return
		}

		// 無限ループ終了時刻 < 現在時刻 の場合、ループ終了
		if ed_time.Before(now_time) {
			// close
			if err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
				log.Println("write close:", err)
			}
			return
		}

		// 0.5秒休憩
		time.Sleep(time.Millisecond * 500)
		// 現在時刻を更新
		now_time = time.Now()
		req.Params["date"] = now_time.String()
	}
}
