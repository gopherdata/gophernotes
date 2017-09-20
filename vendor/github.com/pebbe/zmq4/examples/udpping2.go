//
//  UDP ping command
//  Model 2, uses the GO net library
//

//  this doesn't use ZeroMQ at all

package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

const (
	PING_PORT_NUMBER = 9999
	PING_MSG_SIZE    = 1
	PING_INTERVAL    = 1000 * time.Millisecond //  Once per second
)

func main() {

	log.SetFlags(log.Lshortfile)

	//  Create UDP socket
	bcast := &net.UDPAddr{Port: PING_PORT_NUMBER, IP: net.IPv4bcast}
	conn, err := net.ListenUDP("udp", bcast)
	if err != nil {
		log.Fatalln(err)
	}

	buffer := make([]byte, PING_MSG_SIZE)

	//  We send a beacon once a second, and we collect and report
	//  beacons that come in from other nodes:

	//  Send first ping right away
	ping_at := time.Now()

	for {
		if err := conn.SetReadDeadline(ping_at); err != nil {
			log.Fatalln(err)
		}

		if _, addr, err := conn.ReadFrom(buffer); err == nil {
			//  Someone answered our ping
			fmt.Println("Found peer", addr)
		}

		if time.Now().After(ping_at) {
			//  Broadcast our beacon
			fmt.Println("Pinging peers...")
			buffer[0] = '!'
			if _, err := conn.WriteTo(buffer, bcast); err != nil {
				log.Fatalln(err)
			}
			ping_at = time.Now().Add(PING_INTERVAL)
		}
	}
}
