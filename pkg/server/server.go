package server

import (
	"github.com/lidongyooo/lightsocks-R/pkg/cipher"
	"github.com/lidongyooo/lightsocks-R/pkg/password"
	"io"
	"log"
	"net"
)

type Server struct {
	Cipher *cipher.Cipher
	ListenAddr *net.TCPAddr
}

// 加密传输的 TCP Socket
type SecureTCPConn struct {
	io.ReadWriteCloser
	Cipher *cipher.Cipher
}

func NewLsServer(pw, listenAddr string) (*Server, error) {
	bsPassword, err := password.ParsePassword(pw)
	if err != nil {
		return nil, err
	}

	structListenAddr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		return nil, err
	}

	return &Server{
		Cipher: cipher.NewCipher(bsPassword),
		ListenAddr: structListenAddr,
	}, nil
}

func (server *Server) Listen(didListen func(listenAddr net.Addr)) error {
	return ListenSecureTCP(server.ListenAddr, server.Cipher, server.handleConn, didListen)
}

func (server *Server) handleConn (localConn *SecureTCPConn)  {
	log.Println(localConn)
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