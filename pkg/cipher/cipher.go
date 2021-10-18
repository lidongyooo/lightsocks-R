package cipher

import (
	"github.com/lidongyooo/lightsocks-R/pkg/password"
)

type Cipher struct {
	// 编码用的密码
	encodePassword *password.Password
	// 解码用的密码
	decodePassword *password.Password
}

func (cipher *Cipher) Decode(bs []byte) {
	for i, v := range bs {
		bs[i] = cipher.decodePassword[v]
	}
}

func (cipher *Cipher) Encode (bs []byte) {
	for i, v := range bs {
		bs[i] = cipher.encodePassword[v]
	}
}

// 新建一个编码解码器
func New(encodePassword *password.Password) *Cipher {
	decodePassword := &password.Password{}
	for i, v := range encodePassword {
		decodePassword[v] = byte(i)
	}

	return &Cipher{
		encodePassword: encodePassword,
		decodePassword: decodePassword,
	}
}
