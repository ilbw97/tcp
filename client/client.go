package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

var ErrorNocommand = errors.New("no command")

var ErrorNotEnded = errors.New("")
var ErrorNotConverttoAtoi = errors.New("")

func read(c net.Conn) error {
	// var resultall []byte
	resultlen := 0

	for { //연결이 존재할 때
		data := make([]byte, 4096)
		n, err := c.Read(data) //server로부터 data 읽어오면
		if err != nil {
			if io.EOF == err {
				log.Println("연결 종료")
			}
			return err
		} else { //data를 읽어올 때 error 없을 경우
			if n > 0 {
				recvSize := data[:bytes.Index(data, []byte("\n"))] //data 첫부분에서 보낸 size check
				size, err := strconv.Atoi(string(recvSize))        //int로 converting
				log.Printf("receive size : %d\n", size)

				sizelen := len(strconv.Itoa(size))
				log.Printf("sizelen : %d\n", sizelen)

				totsize := size + sizelen
				log.Printf("totsize : %d\n", totsize)
				if err != nil {
					log.Println(err)

				} else {
					ns := bytes.Index(data, []byte("\n")) + 1
					log.Printf("ns : %d\n", ns)

					log.Printf("n - ns : %d\n", n-ns)
					// resultlen += (n - ns)
					log.Printf("resultlen : %v\n", resultlen)

					log.Printf("\n%v", string(data[ns:n]))

					if totsize > len(data) {
						a := 0

						for {
							q, err := c.Read(data)
							if err != nil {
								log.Printf("read q error : %v\n", err)
							}
							// 읽어온 만 큼 append

							// log.Printf("\n%v", string(data[:q]))

							// a += 1
							// log.Printf("%d번째 loop\n", a)
							resultlen += q
							if resultlen > totsize {
								// return nil
								log.Printf("resultlen : %v == totsize : %v\n", resultlen, totsize)
								break
							} else {
								log.Printf("\n%v", string(data[:q]))

								a += 1
								log.Printf("%d번째 loop\n", a)
							}

						}
					} else {
						return nil
					}
					return nil
				}
			}
		}
	}
}

// return nil
// }
// }

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
		if err == ErrorNotConverttoAtoi {
			continue
		}
		if err != nil {
			log.Println(err)
			break
		}
	}

}

// func sending(c net.Conn) error {
// 	var com string                   //command
// 	sc := bufio.NewScanner(os.Stdin) //init scanner

// 	sc.Scan() //stdinput으로 들어온 한 줄 그대로 scan
// 	if sc.Err() != nil {
// 		log.Println(sc.Err())
// 		return sc.Err()
// 	} else {
// 		com = sc.Text() //읽어온 데이터를 변수에 저장

// 		if com == "" {
// 			log.Println("insert command!")
// 			return ErrorNocommand
// 		} else {
// 			comlen := strconv.Itoa(len(com))
// 			_, err := c.Write([]byte(comlen + "\n" + com)) //server로 전송
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

// func main() {
// 	var (
// 		network = "tcp"
// 		port    = ":3011"
// 	)

// 	conn, err := net.Dial(network, port) //client가 server와 연결할 객체 생성.

// 	if err != nil || conn == nil {
// 		log.Println(err)
// 		return
// 	}

// 	for {
// 		err := sending(conn)

// 		if err == ErrorNocommand {
// 			continue
// 		}

// 		if err != nil {
// 			log.Printf("%v", err)
// 			break
// 		}

// 		err = read(conn)
// 		if err == ErrorNotEnded {
// 			continue
// 		}
// 		if err == ErrorNotConverttoAtoi {
// 			continue
// 		}
// 		if err != nil {
// 			log.Println(err)
// 			break
// 		}
// 	}

// }

// func sending(c net.Conn) error {
// 	var com string                   //command
// 	sc := bufio.NewScanner(os.Stdin) //init scanner

// 	sc.Scan() //stdinput으로 들어온 한 줄 그대로 scan
// 	if sc.Err() != nil {
// 		log.Println(sc.Err())
// 		return sc.Err()
// 	} else {
// 		com = sc.Text() //읽어온 데이터를 변수에 저장

// 		if com == "" {
// 			log.Println("insert command!")
// 			return ErrorNocommand
// 		} else {
// 			comlen := strconv.Itoa(len(com))
// 			_, err := c.Write([]byte(comlen + "\n" + com)) //server로 전송
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

// func main() {
// 	var (
// 		network = "tcp"
// 		port    = ":3011"
// 	)

// 	conn, err := net.Dial(network, port) //client가 server와 연결할 객체 생성.

// 	if err != nil || conn == nil {
// 		log.Println(err)
// 		return
// 	}

// 	for {
// 		err := sending(conn)

// 		if err == ErrorNocommand {
// 			continue
// 		}

// 		if err != nil {
// 			log.Printf("%v", err)
// 			break
// 		}

// 		err = read(conn)
// 		if err == ErrorNotEnded {
// 			continue
// 		}
// 		if err == ErrorNotConverttoAtoi {
// 			continue
// 		}
// 		if err != nil {
// 			log.Println(err)
// 			break
// 		}
// 	}

// }

// func sending(c net.Conn) error {
// 	var com string                   //command
// 	sc := bufio.NewScanner(os.Stdin) //init scanner

// 	sc.Scan() //stdinput으로 들어온 한 줄 그대로 scan
// 	if sc.Err() != nil {
// 		log.Println(sc.Err())
// 		return sc.Err()
// 	} else {
// 		com = sc.Text() //읽어온 데이터를 변수에 저장

// 		if com == "" {
// 			log.Println("insert command!")
// 			return ErrorNocommand
// 		} else {
// 			comlen := strconv.Itoa(len(com))
// 			_, err := c.Write([]byte(comlen + "\n" + com)) //server로 전송
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

// func main() {
// 	var (
// 		network = "tcp"
// 		port    = ":3011"
// 	)

// 	conn, err := net.Dial(network, port) //client가 server와 연결할 객체 생성.

// 	if err != nil || conn == nil {
// 		log.Println(err)
// 		return
// 	}

// 	for {
// 		err := sending(conn)

// 		if err == ErrorNocommand {
// 			continue
// 		}

// 		if err != nil {
// 			log.Printf("%v", err)
// 			break
// 		}

// 		err = read(conn)
// 		if err == ErrorNotEnded {
// 			continue
// 		}
// 		if err == ErrorNotConverttoAtoi {
// 			continue
// 		}
// 		if err != nil {
// 			log.Println(err)
// 			break
// 		}
// 	}

// }

// func sending(c net.Conn) error {
// 	var com string                   //command
// 	sc := bufio.NewScanner(os.Stdin) //init scanner

// 	sc.Scan() //stdinput으로 들어온 한 줄 그대로 scan
// 	if sc.Err() != nil {
// 		log.Println(sc.Err())
// 		return sc.Err()
// 	} else {
// 		com = sc.Text() //읽어온 데이터를 변수에 저장

// 		if com == "" {
// 			log.Println("insert command!")
// 			return ErrorNocommand
// 		} else {
// 			comlen := strconv.Itoa(len(com))
// 			_, err := c.Write([]byte(comlen + "\n" + com)) //server로 전송
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

// func main() {
// 	var (
// 		network = "tcp"
// 		port    = ":3011"
// 	)

// 	conn, err := net.Dial(network, port) //client가 server와 연결할 객체 생성.

// 	if err != nil || conn == nil {
// 		log.Println(err)
// 		return
// 	}

// 	for {
// 		err := sending(conn)

// 		if err == ErrorNocommand {
// 			continue
// 		}

// 		if err != nil {
// 			log.Printf("%v", err)
// 			break
// 		}

// 		err = read(conn)
// 		if err == ErrorNotEnded {
// 			continue
// 		}
// 		if err == ErrorNotConverttoAtoi {
// 			continue
// 		}
// 		if err != nil {
// 			log.Println(err)
// 			break
// 		}
// 	}

// }

// func sending(c net.Conn) error {
// 	var com string                   //command
// 	sc := bufio.NewScanner(os.Stdin) //init scanner

// 	sc.Scan() //stdinput으로 들어온 한 줄 그대로 scan
// 	if sc.Err() != nil {
// 		log.Println(sc.Err())
// 		return sc.Err()
// 	} else {
// 		com = sc.Text() //읽어온 데이터를 변수에 저장

// 		if com == "" {
// 			log.Println("insert command!")
// 			return ErrorNocommand
// 		} else {
// 			comlen := strconv.Itoa(len(com))
// 			_, err := c.Write([]byte(comlen + "\n" + com)) //server로 전송
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

// func main() {
// 	var (
// 		network = "tcp"
// 		port    = ":3011"
// 	)

// 	conn, err := net.Dial(network, port) //client가 server와 연결할 객체 생성.

// 	if err != nil || conn == nil {
// 		log.Println(err)
// 		return
// 	}

// 	for {
// 		err := sending(conn)

// 		if err == ErrorNocommand {
// 			continue
// 		}

// 		if err != nil {
// 			log.Printf("%v", err)
// 			break
// 		}

// 		err = read(conn)
// 		if err == ErrorNotEnded {
// 			continue
// 		}
// 		if err == ErrorNotConverttoAtoi {
// 			continue
// 		}
// 		if err != nil {
// 			log.Println(err)
// 			break
// 		}
// 	}

// }

// func sending(c net.Conn) error {
// 	var com string                   //command
// 	sc := bufio.NewScanner(os.Stdin) //init scanner

// 	sc.Scan() //stdinput으로 들어온 한 줄 그대로 scan
// 	if sc.Err() != nil {
// 		log.Println(sc.Err())
// 		return sc.Err()
// 	} else {
// 		com = sc.Text() //읽어온 데이터를 변수에 저장

// 		if com == "" {
// 			log.Println("insert command!")
// 			return ErrorNocommand
// 		} else {
// 			comlen := strconv.Itoa(len(com))
// 			_, err := c.Write([]byte(comlen + "\n" + com)) //server로 전송
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

// func main() {
// 	var (
// 		network = "tcp"
// 		port    = ":3011"
// 	)

// 	conn, err := net.Dial(network, port) //client가 server와 연결할 객체 생성.

// 	if err != nil || conn == nil {
// 		log.Println(err)
// 		return
// 	}

// 	for {
// 		err := sending(conn)

// 		if err == ErrorNocommand {
// 			continue
// 		}

// 		if err != nil {
// 			log.Printf("%v", err)
// 			break
// 		}

// 		err = read(conn)
// 		if err == ErrorNotEnded {
// 			continue
// 		}
// 		if err == ErrorNotConverttoAtoi {
// 			continue
// 		}
// 		if err != nil {
// 			log.Println(err)
// 			break
// 		}
// 	}

// }

// func sending(c net.Conn) error {
// 	var com string                   //command
// 	sc := bufio.NewScanner(os.Stdin) //init scanner

// 	sc.Scan() //stdinput으로 들어온 한 줄 그대로 scan
// 	if sc.Err() != nil {
// 		log.Println(sc.Err())
// 		return sc.Err()
// 	} else {
// 		com = sc.Text() //읽어온 데이터를 변수에 저장

// 		if com == "" {
// 			log.Println("insert command!")
// 			return ErrorNocommand
// 		} else {
// 			comlen := strconv.Itoa(len(com))
// 			_, err := c.Write([]byte(comlen + "\n" + com)) //server로 전송
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

// func main() {
// 	var (
// 		network = "tcp"
// 		port    = ":3011"
// 	)

// 	conn, err := net.Dial(network, port) //client가 server와 연결할 객체 생성.

// 	if err != nil || conn == nil {
// 		log.Println(err)
// 		return
// 	}

// 	for {
// 		err := sending(conn)

// 		if err == ErrorNocommand {
// 			continue
// 		}

// 		if err != nil {
// 			log.Printf("%v", err)
// 			break
// 		}

// 		err = read(conn)
// 		if err == ErrorNotEnded {
// 			continue
// 		}
// 		if err == ErrorNotConverttoAtoi {
// 			continue
// 		}
// 		if err != nil {
// 			log.Println(err)
// 			break
// 		}
// 	}

// }

// func sending(c net.Conn) error {
// 	var com string                   //command
// 	sc := bufio.NewScanner(os.Stdin) //init scanner

// 	sc.Scan() //stdinput으로 들어온 한 줄 그대로 scan
// 	if sc.Err() != nil {
// 		log.Println(sc.Err())
// 		return sc.Err()
// 	} else {
// 		com = sc.Text() //읽어온 데이터를 변수에 저장

// 		if com == "" {
// 			log.Println("insert command!")
// 			return ErrorNocommand
// 		} else {
// 			comlen := strconv.Itoa(len(com))
// 			_, err := c.Write([]byte(comlen + "\n" + com)) //server로 전송
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

// func main() {
// 	var (
// 		network = "tcp"
// 		port    = ":3011"
// 	)

// 	conn, err := net.Dial(network, port) //client가 server와 연결할 객체 생성.

// 	if err != nil || conn == nil {
// 		log.Println(err)
// 		return
// 	}

// 	for {
// 		err := sending(conn)

// 		if err == ErrorNocommand {
// 			continue
// 		}

// 		if err != nil {
// 			log.Printf("%v", err)
// 			break
// 		}

// 		err = read(conn)
// 		if err == ErrorNotEnded {
// 			continue
// 		}
// 		if err == ErrorNotConverttoAtoi {
// 			continue
// 		}
// 		if err != nil {
// 			log.Println(err)
// 			break
// 		}
// 	}

// }
