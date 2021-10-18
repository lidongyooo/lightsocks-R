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

func (secureSocket *SecureTCPConn) DecodeRead(bs []byte) (n int, err error) {
	n, err = secureSocket.Read(bs)
	if err != nil {
		return
	}

	secureSocket.Cipher.Decode(bs)
	return
}

func (secureSocket *SecureTCPConn) EncodeWrite(bs []byte) (n int, err error) {
	secureSocket.Cipher.Encode(bs)
	return secureSocket.Write(bs)
}

func DialTCPSecure(raddr *net.TCPAddr, c *cipher.Cipher) (*SecureTCPConn, error) {
	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		return nil, err
	}

	return &SecureTCPConn{
		ReadWriteCloser: conn,
		Cipher: c,
	}, nil
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
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}

		// conn被关闭时直接清除所有数据 不管没有发送的数据
		conn.SetLinger(0)
		go handleConn(&SecureTCPConn{
			ReadWriteCloser: conn,
			Cipher: cipher,
		})
	}
}
