package local

import (
	"github.com/lidongyooo/lightsocks-R/pkg/cipher"
	"github.com/lidongyooo/lightsocks-R/pkg/password"
	"github.com/lidongyooo/lightsocks-R/pkg/securetcp"
	"log"
	"net"
)

type Local struct {
	Cipher *cipher.Cipher
	LocalAddr *net.TCPAddr
	RemoteAddr *net.TCPAddr
}

func New(pw, laddr, raddr string) (*Local, error) {
	bsPw, err := password.ParsePassword(pw)
	if err != nil {
		return nil, err
	}

	structLAddr, err := net.ResolveTCPAddr("tcp", laddr)
	if err != nil {
		return nil, err
	}

	structRAddr, err := net.ResolveTCPAddr("tcp", raddr)
	if err != nil {
		return nil, err
	}

	return &Local{
		Cipher: cipher.New(bsPw),
		LocalAddr: structLAddr,
		RemoteAddr: structRAddr,
	}, nil
}

func (local *Local) Listen(didListen func(lAddr net.Addr)) error {
	return securetcp.ListenSecureTCP(local.LocalAddr, local.Cipher, local.handleConn, didListen)
}

func (local *Local) handleConn(userConn *securetcp.SecureTCPConn) {
	defer userConn.Close()

	proxyServer, err := securetcp.DialTCPSecure(local.RemoteAddr, local.Cipher)
	if err != nil {
		log.Println(err)
		return
	}

	defer proxyServer.Close()
}