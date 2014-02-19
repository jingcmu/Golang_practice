//实现了一个简单的server-client程序
package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

func server() {
	//listen on a port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		//accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//handle the connection
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	//receive the message
	var msg string
	defer c.Close()
	for i := 0; i < 100; i++ {
		err := gob.NewDecoder(c).Decode(&msg)
		time.Sleep(time.Second * 1)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Received ", msg, i)
		}
	}

}

func client() {
	//connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	defer c.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	//send the message
	msg := "Hello World"
	for i := 0; i < 100; i++ {
		fmt.Println("Sending", msg, i)
		err = gob.NewEncoder(c).Encode(msg)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
