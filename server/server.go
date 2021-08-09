package main

import (
	"io"
	"log"
	"math"
	"net"
	"os/exec"
)

func main() {
	var (
		network = "tcp"
		port    = ":3011"
	)
	server, err := net.Listen(network, port) //socket 열어준다..

	if nil != err {
		log.Printf("Failed to Listen : %v\n", err)
	}
	defer server.Close()

	for {
		conn, err := server.Accept() //연결 기다림.. 계속 block해서 기다리다가, 연결이 들어왔을 경우 값을 return
		if nil != err {
			log.Printf("Accept Error : %v\n", err)
			continue
		}
		go ConnHandler(conn) //연결을 parameter로 넘겨주고 ConnHandler go routine 실행
	}
}

func ConnHandler(conn net.Conn) {
	recvCommand := make([]byte, math.MaxInt32) //값을 읽어와 저장할 버퍼 생성
	// recvBuf := make([]byte, 4096) //값을 읽어와 저장할 버퍼 생성
	log.Printf("serving %s\n", conn.RemoteAddr().String())
	defer conn.Close()
	for {
		n, err := conn.Read(recvCommand) //client가 값을 줄 때까지 blocking 되어 대기하다가 값을 주면 읽어들인다.
		if nil != err {                  //입력이 종료되면 종료
			if io.EOF == err {
				log.Printf("connection is closed from client : %v\n", conn.RemoteAddr().String())
				return
			}
			log.Printf("Failed to receive data : %v\n", err)
		}

		if n > 0 { // 받아온 길이만큼 슬라이스를 잘라서 출력
			excommand := recvCommand[:n]
			result := execute(string(excommand))
			// result = append(result, "EOF"...)
			_, err := conn.Write(result)
			if err != nil {
				log.Printf("write err : %v\n", err)
				return
			}
		}
	}
}
func execute(command string) []byte {
	cmd := exec.Command("bash", "-c", command)
	// cmd := exec.Command("/bin/bash", "-c", command)
	log.Printf("Execute Command : %v\n", cmd)

	cmdres, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return []byte("Command error : " + err.Error())
	}
	if string(cmdres) == "" {
		return ([]byte("No output data"))
	}

	log.Printf("stdout: %v bytes\n%s", len(cmdres), string(cmdres))
	return append(cmdres, "EOF"...)
}
