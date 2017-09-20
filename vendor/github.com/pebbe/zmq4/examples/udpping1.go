//
//  UDP ping command
//  Model 1, does UDP work inline
//

//  this doesn't use ZeroMQ at all

package main

import (
	"fmt"
	"log"
	"syscall"
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
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err != nil {
		log.Fatalln(err)
	}

	//  Ask operating system to let us do broadcasts from socket
	if err := syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_BROADCAST, 1); err != nil {
		log.Fatalln(err)
	}

	//  Bind UDP socket to local port so we can receive pings
	if err := syscall.Bind(fd, &syscall.SockaddrInet4{Port: PING_PORT_NUMBER, Addr: [4]byte{0, 0, 0, 0}}); err != nil {
		log.Fatalln(err)
	}

	buffer := make([]byte, PING_MSG_SIZE)

	//  We use syscall.Select to wait for activity on the UDP socket.
	//  We send a beacon once a second, and we collect and report
	//  beacons that come in from other nodes:

	rfds := &syscall.FdSet{}
	timeout := &syscall.Timeval{}

	//  Send first ping right away
	ping_at := time.Now()

	bcast := &syscall.SockaddrInet4{Port: PING_PORT_NUMBER, Addr: [4]byte{255, 255, 255, 255}}
	for {
		dur := int64(ping_at.Sub(time.Now()) / time.Microsecond)
		if dur < 0 {
			dur = 0
		}
		timeout.Sec, timeout.Usec = dur/1000000, dur%1000000
		FD_ZERO(rfds)
		FD_SET(rfds, fd)
		_, err := syscall.Select(fd+1, rfds, nil, nil, timeout)
		if err != nil {
			log.Fatalln(err)
		}

		//  Someone answered our ping
		if FD_ISSET(rfds, fd) {
			_, addr, err := syscall.Recvfrom(fd, buffer, 0)
			if err != nil {
				log.Fatalln(err)
			}
			a := addr.(*syscall.SockaddrInet4)
			fmt.Printf("Found peer %v.%v.%v.%v:%v\n", a.Addr[0], a.Addr[1], a.Addr[2], a.Addr[3], a.Port)
		}
		if time.Now().After(ping_at) {
			//  Broadcast our beacon
			fmt.Println("Pinging peers...")
			buffer[0] = '!'
			if err := syscall.Sendto(fd, buffer, 0, bcast); err != nil {
				log.Fatalln(err)
			}
			ping_at = time.Now().Add(PING_INTERVAL)
		}
	}

}

func FD_SET(p *syscall.FdSet, i int) {
	p.Bits[i/64] |= 1 << uint(i) % 64
}

func FD_ISSET(p *syscall.FdSet, i int) bool {
	return (p.Bits[i/64] & (1 << uint(i) % 64)) != 0
}

func FD_ZERO(p *syscall.FdSet) {
	for i := range p.Bits {
		p.Bits[i] = 0
	}
}
