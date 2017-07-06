package main

import (
	"io"
	"net"
	"testing"
)

const (
	NR_BYTES_TRANSMIT = 1024 * 512
)

func panicIfError(err error) {
	if err != nil && err != io.EOF {
		panic(err)
	}
}

//func BenchmarkBuffSize_512(b *testing.B) {
//	benchmarkBuffSize(b, 512)
//}
//
//func BenchmarkBuffSize_1024(b *testing.B) {
//	benchmarkBuffSize(b, 1024)
//}
//
//func BenchmarkBuffSize_2048(b *testing.B) {
//	benchmarkBuffSize(b, 2048)
//}
//
//func BenchmarkBuffSize_4096(b *testing.B) {
//	benchmarkBuffSize(b, 4096)
//}

func BenchmarkBuffSize(b *testing.B) {
	benchmarkBuffSize(b, 16000)
}

var (
	listener net.Listener
)

func init() {
	var err error
	listener, err = net.Listen("tcp", "127.0.0.1:12000")
	panicIfError(err)
}

func benchmarkBuffSize(b *testing.B, bufsize int) {

	go func() {
		conn, err := net.Dial("tcp", "localhost:12000")
		//fmt.Fprintln(os.Stderr, "connected", conn, err)
		panicIfError(err)
		//conn.Write([]byte{'a'})
		recvbuf := make([]byte, 8192*2*2)
		recvBytes := 0
		for {
			n, err := conn.Read(recvbuf)
			panicIfError(err)
			recvBytes += n
			//if recvBytes >= NR_BYTES_TRANSMIT {
			//	break
			//}
		}
	}()

	conn, err := listener.Accept()
	//fmt.Fprintln(os.Stderr, "accepted", conn, err)
	panicIfError(err)
	sndbuf := make([]byte, bufsize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sendBytes := 0

		for {
			n, err := conn.Write(sndbuf)
			panicIfError(err)
			sendBytes += n
			if sendBytes >= NR_BYTES_TRANSMIT {
				break
			}
		}

	}

}
