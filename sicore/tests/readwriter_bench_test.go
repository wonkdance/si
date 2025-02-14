package sicore_test

import (
	"net"
	"testing"
	"time"

	"github.com/wonksing/si/v2/sicore"
	"github.com/wonksing/si/v2/siutils"
)

// func requestTcp(b *testing.B) {

// 	conn, err := net.DialTimeout("tcp", "127.0.0.1:10000", 6*time.Second)
// 	siutils.NilFailB(b, err)
// 	defer conn.Close()

// 	err = conn.SetWriteDeadline(time.Now().Add(6 * time.Second))
// 	siutils.NilFailB(b, err)
// 	err = conn.SetReadDeadline(time.Now().Add(12 * time.Second))
// 	siutils.NilFailB(b, err)

// 	err = conn.(*net.TCPConn).SetWriteBuffer(4096)
// 	siutils.NilFailB(b, err)
// 	err = conn.(*net.TCPConn).SetReadBuffer(4096)
// 	siutils.NilFailB(b, err)

// 	s := sicore.NewReadWriterWithValidator(conn, conn, tcpValidator())
// 	received, err := s.WriteAndRead(createDataToSend())
// 	siutils.NilFailB(b, err)

// 	l, err := strconv.Atoi(string(received[:7]))
// 	siutils.NilFailB(b, err)
// 	assert.Equal(b, l, len(received))
// }

// func BenchmarkReadWriter_Tcp_WriteAndRead(b *testing.B) {
// 	if onlinetest != "1" {
// 		b.Skip("skipping online tests")
// 	}
// 	for i := 0; i < b.N; i++ {
// 		requestTcp(b)
// 	}
// }

// func requestTcpWithConn(b *testing.B, conn net.Conn) {

// 	s := sicore.NewReadWriterWithValidator(conn, conn, tcpValidator())
// 	received, err := s.WriteAndRead(createSmallDataToSend())
// 	siutils.NilFailB(b, err)

// 	l, err := strconv.Atoi(string(received[:7]))
// 	siutils.NilFailB(b, err)
// 	assert.Equal(b, l, len(received))
// }
// func BenchmarkReadWriter_Tcp_WriteAndReadReuseConn(b *testing.B) {
// 	if onlinetest != "1" {
// 		b.Skip("skipping online tests")
// 	}
// 	conn, err := net.DialTimeout("tcp", "127.0.0.1:10000", 6*time.Second)
// 	siutils.NilFailB(b, err)
// 	defer conn.Close()

// 	// err = conn.SetWriteDeadline(time.Now().Add(6 * time.Second))
// 	// siutils.NilFailB(b, err)
// 	// err = conn.SetReadDeadline(time.Now().Add(12 * time.Second))
// 	// siutils.NilFailB(b, err)

// 	err = conn.(*net.TCPConn).SetWriteBuffer(4096)
// 	siutils.NilFailB(b, err)
// 	err = conn.(*net.TCPConn).SetReadBuffer(4096)
// 	siutils.NilFailB(b, err)

// 	for i := 0; i < b.N; i++ {
// 		requestTcpWithConn(b, conn)
// 	}
// }

// func requestTcpWithConn2(b *testing.B, s *sicore.ReadWriter, conn net.Conn) {

// 	s.WriteAndRead(createSmallDataToSend())
// }

// func tcpValidatorDummy() sicore.ReadValidator {
// 	return sicore.ValidateFunc(func(b []byte, errIn error) (bool, error) {
// 		return true, nil
// 	})
// }

// func BenchmarkReadWriter_Tcp_WriteAndReadReuseConn2(b *testing.B) {
// 	if onlinetest != "1" {
// 		b.Skip("skipping online tests")
// 	}
// 	conn, err := net.DialTimeout("tcp", "127.0.0.1:10000", 6*time.Second)
// 	siutils.NilFailB(b, err)
// 	defer conn.Close()

// 	err = conn.SetWriteDeadline(time.Now().Add(6 * time.Second))
// 	siutils.NilFailB(b, err)
// 	err = conn.SetReadDeadline(time.Now().Add(12 * time.Second))
// 	siutils.NilFailB(b, err)

// 	err = conn.(*net.TCPConn).SetWriteBuffer(4096)
// 	siutils.NilFailB(b, err)
// 	err = conn.(*net.TCPConn).SetReadBuffer(4096)
// 	siutils.NilFailB(b, err)

// 	s := sicore.NewReadWriterSizeWithValidator(conn, conn, 4096, tcpValidator())
// 	for i := 0; i < b.N; i++ {
// 		requestTcpWithConn2(b, s, conn)
// 	}
// }

func Benchmark_Tcp_Basic(b *testing.B) {
	if onlinetest {
		b.Skip("skipping online tests")
	}
	conn, err := net.DialTimeout("tcp", "127.0.0.1:10000", 6*time.Second)
	siutils.AssertNilFailB(b, err)
	defer conn.Close()

	// err = conn.SetWriteDeadline(time.Now().Add(6 * time.Second))
	// siutils.NilFailB(b, err)
	// err = conn.SetReadDeadline(time.Now().Add(12 * time.Second))
	// siutils.NilFailB(b, err)

	err = conn.(*net.TCPConn).SetWriteBuffer(4096)
	siutils.AssertNilFailB(b, err)
	err = conn.(*net.TCPConn).SetReadBuffer(4096)
	siutils.AssertNilFailB(b, err)

	for i := 0; i < b.N; i++ {
		buf := make([]byte, 1024)
		conn.Write(createSmallDataToSend())
		conn.Read(buf)
	}
}

func Benchmark_Tcp_Request(b *testing.B) {
	if onlinetest {
		b.Skip("skipping online tests")
	}
	conn, err := net.DialTimeout("tcp", "127.0.0.1:10000", 6*time.Second)
	siutils.AssertNilFailB(b, err)
	defer conn.Close()

	err = conn.SetWriteDeadline(time.Now().Add(6 * time.Second))
	siutils.AssertNilFailB(b, err)
	err = conn.SetReadDeadline(time.Now().Add(12 * time.Second))
	siutils.AssertNilFailB(b, err)

	err = conn.(*net.TCPConn).SetWriteBuffer(4096)
	siutils.AssertNilFailB(b, err)
	err = conn.(*net.TCPConn).SetReadBuffer(4096)
	siutils.AssertNilFailB(b, err)

	for i := 0; i < b.N; i++ {
		// rw := sicore.GetReadWriterWithOptions(conn, []sicore.ReaderOption{SetTcpEofChecker()}, conn, nil)
		rw := sicore.GetReadWriterWithReadWriter(conn)
		rw.Reader.ApplyOptions(SetTcpEofChecker())
		rw.Request(createSmallDataToSend())
		sicore.PutReadWriter(rw)

	}
}
