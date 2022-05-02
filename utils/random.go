package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	mathRand "math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// GenerateID GenerateID
func GenerateID() string {
	str, err := RandBytes(6)
	if err != nil {
		log.Print("GenerateID err:", err)
		return ""
	}
	return strings.ToUpper(hex.EncodeToString(str))
}

// GenerateToken File Token
func GenerateToken(val string) string {
	bytes, err := RandBytes(33)
	if err != nil {
		log.Println("Error:", err)
	}
	valBytes, _ := hex.DecodeString(val)
	return HmacSha256(valBytes, bytes)
}

// GenerateIv IV
func GenerateIv() []byte {
	result, _ := RandBytes(16)
	return result
}

// GenerateAad AAD
func GenerateAad() []byte {
	result, _ := RandBytes(16)
	return result
}

// GenerateSalt Salt
func GenerateSalt() []byte {
	result, _ := RandBytes(64)
	return result
}

// TimestampString timestamp
func TimestampString() string {
	now := time.Now().Unix()
	timestamp := strconv.FormatInt(now, 10)
	return timestamp
}

func GenerateUUID() string {
	return uuid.NewString()
}

// RandBytes RandBytes
func RandBytes(len int) ([]byte, error) {
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return []byte{}, err
	}

	return b, nil
}

// RandInt64 获得一个随机数
// min 最小值
// max 最大值
func RandInt64(min, max int64) int64 {
	randNum := int64(0)
	maxBigInt := big.NewInt(max)

	for randNum <= min {
		i, _ := rand.Int(rand.Reader, maxBigInt)
		randNum = i.Int64()
	}
	return randNum
}

// GenerateMsgID Create Message ID
func GenerateMsgID(address string, length ...int) string {
	baseID := fmt.Sprintf("%d", time.Now().UnixNano()/1000000)
	return baseID + fmt.Sprintf("%d", RandInt64(10000, 99999))

}

// StringToHex HexString to bytes
func StringToHex(str string) ([]byte, error) {
	dates := []byte(str)
	hexBytes, err := hex.DecodeString(hex.EncodeToString(dates))
	return hexBytes, err
}

func RandStr(strlen int) string {
	mathRand.Seed(time.Now().UnixNano())
	data := make([]byte, strlen)
	var num int
	for i := 0; i < strlen; i++ {
		num = mathRand.Intn(57) + 65
		for {
			if num > 90 && num < 97 {
				num = mathRand.Intn(57) + 65
			} else {
				break
			}
		}
		data[i] = byte(num)
	}
	return string(data)
}
