package main

import (
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":3011") //socket 열어준다
	if nil != err {
		log.Println(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept() //socker에 연결한다
		if nil != err {
			log.Println(err)
			continue
		}
		defer conn.Close()
		go ConnHandler(conn) //연결을 parameter로 넘겨주고 ConnHandler go routine 실행
	}
}

func ConnHandler(conn net.Conn) {
	recvBuf := make([]byte, 4096)
	for {
		n, err := conn.Read(recvBuf) //client가 값을 줄 때까지 blocking 되어 대기하다가 값을 주면 읽어들인다.
		if nil != err {
			if io.EOF == err {
				log.Println(err)
				return
			}
			log.Println(err)
		}
		if 0 < n {
			data := recvBuf[:n]
			log.Println(string(data))
			_, err := conn.Write(data[:n])
			if err != nil {
				log.Println(err)
				return
			}
		}
	}

}
