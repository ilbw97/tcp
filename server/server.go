package main

import (
	"bytes"
	"io"
	"log"
	"math"
	"net"
	"os/exec"
	"strconv"
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
	log.Printf("serving %s\n", conn.RemoteAddr().String())
	recvData := make([]byte, math.MaxInt32) //값을 읽어와 저장할 버퍼 생성s
	defer conn.Close()
	for {
		n, err := conn.Read(recvData) //client가 값을 줄 때까지 blocking 되어 대기하다가 값을 주면 읽어들인다.
		log.Printf("n : %d\n", n)

		if nil != err { //입력이 종료되면 종료
			if io.EOF == err {
				log.Printf("connection is closed from client : %v\n", conn.RemoteAddr().String())
				return
			}
			log.Printf("Failed to receive data : %v\n", err)
		}
		// if bytes.Contains(recvData, []byte("\n")) {
		recvSize := recvData[:bytes.Index(recvData, []byte("\n"))]

		// excommand := recvData[bytes.Index(recvData, []byte("\n")):n]

		size, err := strconv.Atoi(string(recvSize))
		if err != nil {
			return
		}
		log.Printf("recieve command size : %d\n", size)
		// recvCommand := make([]byte, size+1)

		r, err := conn.Read(recvData[:size])
		if err != nil {
			return
		}
		log.Printf("r : %d\n", r)
		excommand := recvData[:r]
		if r > 0 {
			result := execute(string(excommand))
			result = append(result, "EOF"...)
			_, err := conn.Write(result)
			if err != nil {
				log.Printf("write err : %v\n", err)
				return
			}
		}

		// // recvCommand := make([]byte, size)
		// if n > 0 { // 받아온 길이만큼 슬라이스를 잘라서 출력
		// 	// excommand := recvCommand[:n]

		// }
		// }
	}
}
func execute(command string) []byte {
	cmd := exec.Command("bash", "-c", command)
	// cmd := exec.Command("/bin/bash", "-c", command)
	log.Printf("Execute Command : %v\n", cmd)

	cmdres, err := cmd.Output()
	cmdreslen := []byte(strconv.Itoa(len(cmdres)) + "\n")
	if err != nil {
		log.Println(err)
		return append(cmdres, []byte("Command error : "+err.Error())...)
	}
	if string(cmdres) == "" {
		return append(cmdres, ([]byte("No output data"))...)
	}

	log.Printf("stdout: %v bytes\n%s", string(bytes.Trim(cmdreslen, "\n")), string(cmdres))
	// return append(cmdres, "EOF"...)
	return append(cmdreslen, cmdres...)
}
