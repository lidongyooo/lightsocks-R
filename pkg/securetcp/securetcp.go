package securetcp

import (
	"github.com/lidongyooo/lightsocks-R/pkg/cipher"
	"io"
	"log"
	"net"
)

// 加密传输的 TCP Socket
type SecureTCPConn struct {
	io.ReadWriteCloser
	Cipher *cipher.Cipher
}

func ListenSecureTCP(laddr *net.TCPAddr, cipher *cipher.Cipher, handleConn func(localConn *SecureTCPConn), didListen func(listenAddr net.Addr)) error {
	listener, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		return err
	}

	defer listener.Close()

	if didListen != nil {
		didListen(listener.Addr())
	}

	for {
		localConn, err := listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}

		// localConn被关闭时直接清除所有数据 不管没有发送的数据
		localConn.SetLinger(0)
		go handleConn(&SecureTCPConn{
			ReadWriteCloser: localConn,
			Cipher: cipher,
		})
	}
}
