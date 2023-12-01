package comm

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"lottery-system/conf"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// 当前时间戳
func NowUnix() int {
	return int(time.Now().In(conf.SysTimeLocation).Unix())
}

// 将时间戳转为 yyyy-mm-dd H:i:s 格式

func FormatUnixTime(t int64) string {
	if t > 0 {
		return time.Unix(t, 0).Format(conf.SysTimeform)
	} else {
		return time.Now().Format(conf.SysTimeform)
	}
}

// 将时间戳转为 yyyy-mm-dd 格式

func FormatUnixTimeShort(t int64) string {
	if t > 0 {
		return time.Unix(t, 0).Format(conf.SysTimeformShort)
	} else {
		return time.Now().Format(conf.SysTimeformShort)
	}
}

// 将字符串转成时间
func ParseTime(str string) (time.Time, error) {
	return time.ParseInLocation(conf.SysTimeform, str, conf.SysTimeLocation)
}

// 得到一个随机数
func Random(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if max < 1 {
		return r.Int()
	} else {
		return r.Intn(max)
	}

}

// 对一个字符串进行加密
func encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))

	return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext is too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil

}

// 在预定义字符前添加 \
// ' " \
// http://www.ruanyifeng.com/blog/2007/10/ascii_unicode_and_utf-8.html

func AddSlashes(str string) string {
	temRune := []rune{}
	strRune := []rune(str)
	for _, ch := range strRune {
		switch ch {
		case []rune{'\\'}[0], []rune{'"'}[0], []rune{'\''}[0]:
			temRune = append(temRune, []rune{'\\'}[0])
			temRune = append(temRune, ch)
		default:
			temRune = append(temRune, ch)
		}
	}
	return string(temRune)
}

// 删除 \

func StripsSlashs(str string) string {
	dstRune := []rune{}
	strRune := []rune(str)

	for i := 0; i < len(strRune); i++ {
		if strRune[i] == []rune{'\\'}[0] {
			i++
		}

		if i == len(strRune) {
			return string(dstRune)
		}

		dstRune = append(dstRune, strRune[i])
	}
	return string(dstRune)
}

// 将字符串的ip 地址转换为数字

func Ip4toInt(ip string) int64 {
	bits := strings.Split(ip, ".")
	var sum int64 = 0
	if len(bits) == 4 {
		b0, _ := strconv.Atoi(bits[0])
		b1, _ := strconv.Atoi(bits[1])
		b2, _ := strconv.Atoi(bits[2])
		b3, _ := strconv.Atoi(bits[3])
		sum += int64(b0) << 24
		sum += int64(b1) << 16
		sum += int64(b2) << 8
		sum += int64(b3)
	}
	return sum
}

// 得到当前时间到下一天零点的延时

func NextDayDuration() time.Duration {
	year, month, day := time.Now().Add(time.Hour * 24).Date()

	next := time.Date(year, month, day, 0, 0, 0, 0, conf.SysTimeLocation)

	return time.Until(next)
}

// 判断我们系统中的字节序类型
/**
在给定的 Go 代码中，var num uint16 = 1 定义了一个 uint16 类型的变量 num，并初始化为 1。uint16 表示一个无符号的 16 位整数，它的取值范围是 0 到 65535。

在函数中，通过使用 binary.LittleEndian.PutUint16(buf, num) 将 num 编码为字节序列，这里使用了小端序（Little Endian）。在小端序中，低位字节（最低有效字节）在前，高位字节在后。

因此，如果 buf 的第一个字节的最右侧比特是 1，则表示系统是小端序。这是因为在二进制表示中，1 的最右侧比特是 1。所以，这个函数的目的是检查系统的字节序，如果返回 true，则表示系统是小端序。
**/
func isLittleEndian() bool {
	var num uint16 = 1
	buf := make([]byte, 2)

	// 使用 binary.LittleEndian 将 uint16 编码为字节序列
	binary.LittleEndian.PutUint16(buf, num)
	// 如果小端序的话，buf 的第一个字节的最右侧比特应该是 1
	return buf[0]&1 == 1
}

func isBigEndian() bool {
	var num uint16 = 1
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, num)
	return buf[0]&1 == 0
}
