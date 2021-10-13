package password

import (
	"encoding/base64"
	"math/rand"
	"time"
)

const PASSWORDLENGTH = 256

type password [PASSWORDLENGTH]byte

func init() {
	// 更新随机种子，防止生成一样的随机密码
	rand.Seed(time.Now().Unix())
}

// 采用base64编码把密码转换为字符串
func (p *password) String() string {
	return base64.StdEncoding.EncodeToString(p[:])
}

// 产生 256个byte随机组合的 密码，最后会使用base64编码为字符串存储在配置文件中
// 不能出现任何一个重复的byte位，必须又 0-255 组成，并且都需要包含
func RandPassword() string {
	// 随机生成一个由  0~255 组成的 byte 数组
	intArr := rand.Perm(PASSWORDLENGTH)

	password := &password{}
	for i, v := range intArr {
		password[i] = byte(v)
		if i == v {
			// 确保不会出现如何一个byte位出现重复
			return RandPassword()
		}
	}

	return password.String()
}