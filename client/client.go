package main

import (
	"bufio"
	"log"
	"math"
	"net"
	"os"
	"time"
)

func main() {
	var (
		network = "tcp"
		port    = ":3011"
	)
	conn, err := net.Dial(network, port) //client가 server와 연결할 객체 생성.
	if nil != err {
		log.Println(err)
	}

	go func() {
		data := make([]byte, math.MaxInt32)
		// data := make([]byte, 4096)
		for {
			n, err := conn.Read(data) //server로부터 data 읽어오면
			if err != nil {
				log.Println(err)
				return
			}
			log.Printf("result : \n%s\n", string(data[:n])) //값 출력
			time.Sleep(time.Duration(3) * time.Second)
		}
	}()
	for {
		var com string                   //command
		sc := bufio.NewScanner(os.Stdin) //init scanner

		sc.Scan() //stdinput으로 들어온 한 줄 그대로 scan

		com = sc.Text() //읽어온 데이터를 변수에 저장

		log.Printf("before converting : %s to Server\n", com)
		log.Printf("after converting : %v to Server\n", []byte(com))
		conn.Write([]byte(com)) //server로 전송

		time.Sleep(time.Duration(1) * time.Second)
	}
}
