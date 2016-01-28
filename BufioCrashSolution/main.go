package main

import (
	"bufio"
	"fmt"
	"net"

	"time"
)

func main() {

	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	go func(listner net.Listener) {
		for {
			conn, err := listner.Accept()
			if err == nil && conn != nil {
				brdr := bufio.NewReaderSize(conn, 16)
				go readln(brdr)
				//comment this go routine the again build and run it now program will not crash
				go readln(brdr)
			} else {
				fmt.Println(err.Error())
				time.Sleep(10 * time.Second)
			}
		}
	}(listner)

	time.Sleep(10000 * time.Second)
}
func readln(brdr *bufio.Reader) error {
	line := make([]byte, 200)

	for {
		bite, err := brdr.ReadByte()
		if err == nil {
			line = append(line, bite)
			fmt.Println("go routine 1 : %s", string(line))
			time.Sleep(1 * time.Second)
		} else {
			fmt.Println("go routine 1 %s", err.Error())
			return err
		}
	}
}
