package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"math"
	"net"
	"os"
	"strconv"
)

var ErrorNocommand = errors.New("no command")

var ErrorNotEnded = errors.New("")

// var ErrorNoData = errors.New("no return data")

func read(c net.Conn) error {
	// data := make([]byte, 4096)
	for {
		data := make([]byte, math.MaxInt32)
		n, err := c.Read(data) //server로부터 data 읽어오면
		// n, err := ioutil.ReadAll(c)
		if err != nil {
			if io.EOF == err {
				log.Println("연결 종료")
			}
			return err
		} else {
			// if bytes.Contains(data, []byte("EOF")) {
			// log.Println("find EOF!")
			// res := bytes.Trim(data, "EOF")
			// log.Printf("\n%v", string(res[:n])) //값 출력
			log.Printf("\n%v", string(bytes.Trim(data[:n], "EOF")))
			// log.Printf("")
			return nil
			// }

		}
	}
}
func sending(c net.Conn) error {
	var com string                   //command
	sc := bufio.NewScanner(os.Stdin) //init scanner

	sc.Scan() //stdinput으로 들어온 한 줄 그대로 scan
	if sc.Err() != nil {
		log.Println(sc.Err())
		return sc.Err()
	} else {
		com = sc.Text() //읽어온 데이터를 변수에 저장

		if com == "" {
			log.Println("insert command!")
			return ErrorNocommand
		} else {
			comlen := strconv.Itoa(len(com))
			_, err := c.Write([]byte(comlen + "\n" + com)) //server로 전송
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	var (
		network = "tcp"
		port    = ":3011"
	)

	conn, err := net.Dial(network, port) //client가 server와 연결할 객체 생성.

	if err != nil || conn == nil {
		log.Println(err)
		return
	}

	for {
		err := sending(conn)

		if err == ErrorNocommand {
			continue
		}

		if err != nil {
			log.Printf("%v", err)
			break
		}

		err = read(conn)
		if err == ErrorNotEnded {
			continue
		}

		// if err == ErrorNoData {
		// 	log.Println(err)
		// 	continue
		// }
		if err != nil {
			log.Println(err)
			break
		}
	}

	// } else {
	// 	log.Printf("connected to %s!\n", conn.LocalAddr().String())
	// 	for {
	// 		sending(conn)
	// 		read(conn)
	// 	}

	// }
	// // defer conn.Close()
}
