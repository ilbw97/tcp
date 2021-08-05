package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":3011") //client가 server와 연결 시도
	if nil != err {
		log.Println(err)
	}

	go func() {
		data := make([]byte, 4096)

		for {
			n, err := conn.Read(data) //server로부터 data 읽어오면
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("Server send : " + string(data[:n])) //값 출력
			time.Sleep(time.Duration(3) * time.Second)
		}
	}()
	for {
		var s string
		fmt.Scanln(&s)        //사용자가 입력하면
		conn.Write([]byte(s)) //server로 전송
		time.Sleep(time.Duration(3) * time.Second)
	}
}
