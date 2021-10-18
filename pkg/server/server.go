package server

import (
	"github.com/lidongyooo/lightsocks-R/pkg/cipher"
	"github.com/lidongyooo/lightsocks-R/pkg/password"
	"github.com/lidongyooo/lightsocks-R/pkg/securetcp"
	"net"
)

type Server struct {
	Cipher *cipher.Cipher
	ListenAddr *net.TCPAddr
}

func New(pw, listenAddr string) (*Server, error) {
	bsPassword, err := password.ParsePassword(pw)
	if err != nil {
		return nil, err
	}

	structListenAddr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		return nil, err
	}

	return &Server{
		Cipher: cipher.New(bsPassword),
		ListenAddr: structListenAddr,
	}, nil
}

func (server *Server) Listen(didListen func(listenAddr net.Addr)) error {
	return securetcp.ListenSecureTCP(server.ListenAddr, server.Cipher, server.handleConn, didListen)
}

func (server *Server) handleConn (localConn *securetcp.SecureTCPConn)  {
	defer localConn.Close()

	//buffer := make([]byte, 256)
	//_, err := localConn.
}