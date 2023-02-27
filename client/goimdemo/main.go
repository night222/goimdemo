package main

import (
	"fmt"
	"net/url"
	"time"
	goimmessage "websocket_client/common/goim_message"

	"github.com/gorilla/websocket"
)

var addr = "127.0.0.1:8080"

func main() {
	u := url.URL{
		Scheme: "ws",
		Host:   addr,
		Path:   "/sendweb",
	}
	token := ""
	wmsg := &goimmessage.Message{
		FormId:   1,
		TargetId: 2,
		Type:     goimmessage.PrivateChat,
		Media:    goimmessage.Text,
		Content:  "你好呀!",
	}
	dialer := websocket.DefaultDialer
	dialer.Subprotocols = append(dialer.Subprotocols, token)
	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println("dialer.Dial", err)
		return
	}
	fmt.Println("建立连接成功")
	defer conn.Close()
	err = conn.WriteJSON(wmsg)
	if err != nil {
		fmt.Println("write:", err)
		return
	}
	rmsg := &goimmessage.Message{}
	for {
		err = conn.ReadJSON(rmsg)
		if err != nil {
			fmt.Println("read:", err)
			break
		}
		fmt.Printf("接受到%d的信息:%s", rmsg.FormId, rmsg.Content)
		err = conn.WriteJSON(wmsg)
		if err != nil {
			fmt.Println("write:", err)
			break
		}
		time.Sleep(time.Second)
	}
}
