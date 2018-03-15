package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

func start() {
	listenAds := net.JoinHostPort(cfg.ListenAddress, strconv.Itoa(int(cfg.ListenPort)))

	listener, err := net.Listen("tcp", listenAds)
	if err != nil {
		log.Fatalf("listen on address '%v' err, %v\n", listenAds, err)
	}
	defer listener.Close()

	log.Printf("listen on address '%v'\n", listenAds)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept err, %v\n", err)
			continue
		}

		go handleConnection(&conn)
	}

	log.Println("server stop")
}

func handleConnection(conn *net.Conn) {
	tgtConn, err := net.DialTimeout("tcp", targetAds, time.Second*dialTimeoutSeconds)
	if err != nil {
		(*conn).Close()
		log.Printf("connect to target '%v' fail: %v\n", targetAds, err)
		return
	}

	pipe(conn, &tgtConn)
}

func pipe(connA, connB *net.Conn) {
	go pipeTo(connA, connB)
	pipeTo(connB, connA)
}

func pipeTo(src, dst *net.Conn) {
	buf := make([]byte, bufferSize)
	s := *src
	d := *dst

	for {
		// set timeout
		s.SetDeadline(time.Now().Add(time.Minute * SwitchTimeoutMinutes))

		n, err := s.Read(buf)
		if err != nil {
			break
		}

		data := buf[:n]

		_, err = d.Write(data)
		if err != nil {
			break
		}
	}

	(*src).Close()
	(*dst).Close()
}
