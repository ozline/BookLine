package middleware

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"
)

func GenerateTokenSHA256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GenerateMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func GenerateMD5Random() string {
	h := md5.New()
	h.Write([]byte(Int64ToBytes(GetTimestamp())))
	return hex.EncodeToString(h.Sum(nil))
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func GetTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}
