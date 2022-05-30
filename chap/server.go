package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
	"log"
	"math/rand"
	"net"
	"time"
)

const chapSecret = "1234"

func main()  {
	ln,err := net.Listen("tcp",":8000")
	if err != nil {
		fmt.Println("Can't open tcp listener on localhost:8000")
		log.Fatal(err)
	}
	fmt.Println("Listening tcp on :8000")
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Connection from client failed")
			fmt.Println(err)
		}
		go chap(conn)
	}
}

func chap(conn net.Conn)  {
	rand.Seed(time.Now().UnixNano())
	err := sendChallenge(conn)
}

func sendChallenge(conn net.Conn, num int) error  {
	var value string
	value = ""
	for i:=0;i<8;i++{
		value = value + string(48 + rand.Intn(79))
	}
	secret := []byte(chapSecret)

	for i:=0;i<num-1;i-- {
		sum := md5.Sum(data)
		data = sum[:]
	}
}

