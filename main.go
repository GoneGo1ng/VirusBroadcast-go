package main

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
	"virusbroadcast/constants"
	"virusbroadcast/panel"
	"virusbroadcast/person"
)

func Refresh(ws *websocket.Conn) {
	pp := person.GetInstance()

	// 设置初始感染人员
	// rand.Seed(time.Now().UnixNano())
	for i := 0; i < constants.OriginalCount; i++ {
		p := pp.Persons[rand.Intn(len(pp.Persons))]
		for {
			if !p.IsInfected() {
				break
			}
			p = pp.Persons[rand.Intn(len(pp.Persons))]
		}
		p.BeInfected()
	}

	// 定时发送所绘制都病毒传播图片
	ticker := time.NewTicker(500 * time.Millisecond)
	for _ = range ticker.C {
		b := panel.Paint()

		base64Str := "data:image/png;base64," + base64.StdEncoding.EncodeToString(b.Bytes())

		if err := websocket.Message.Send(ws, base64Str); err != nil {
			fmt.Println("send failed:", err)
			break
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func main() {
	// 接受websocket的路由地址
	http.Handle("/websocket", websocket.Handler(Refresh))
	// html页面
	http.HandleFunc("/", index)
	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
