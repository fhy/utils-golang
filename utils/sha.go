package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"time"
)

// Sha256Digest Sha256Digest
func Sha256Digest(plaintext string) (result string) {
	hash := sha256.New()
	hash.Write([]byte(plaintext))
	md := hash.Sum(nil)
	result = hex.EncodeToString(md)
	return
}

// HmacSha512 Hmac Sha512
func HmacSha512(plaintex string) string {
	mac := hmac.New(sha512.New, []byte("49f41477fa1bfc3b4792d5233b6a659f4b"))
	mac.Write([]byte(plaintex))
	md := mac.Sum(nil)
	return hex.EncodeToString(md)
}

// Sha1Digest Sha1Digest
func Sha1Digest(plaintext string) (result string) {
	hash := sha1.New()
	hash.Write([]byte(plaintext))
	md := hash.Sum(nil)
	result = hex.EncodeToString(md)
	return
}

func Sha256Hex(plaintext string) (result string) {
	result = hex.EncodeToString(Sha256HexBytes(plaintext))
	return
}

func Sha256(plaintext []byte) []byte {
	hash := sha256.New()
	hash.Write(plaintext)
	md := hash.Sum(nil)
	hash2 := sha256.New()
	hash2.Write(md)
	res := hash2.Sum(nil)
	return res
}

func Sha256HexBytes(plaintext string) []byte {
	hash := sha256.New()
	hash.Write([]byte(plaintext))
	md := hash.Sum(nil)
	hash2 := sha256.New()
	hash2.Write(md)
	res := hash2.Sum(nil)
	return res
}

func Sha1Hex(plaintext string) (result string) {
	hash := sha1.New()
	hash.Write([]byte(plaintext))
	md := hash.Sum(nil)
	hash2 := sha1.New()
	hash2.Write(md)
	res := hash2.Sum(nil)
	result = hex.EncodeToString(res)
	return
}

// HmacSha1 Hmac Sha1
func HmacSha1(plainText string) string {
	randBytes, _ := RandBytes(128)
	h := hmac.New(sha1.New, randBytes)
	h.Write([]byte(plainText))
	return hex.EncodeToString(h.Sum(nil))
}

// HmacSha256 Hmac Sha256
func HmacSha256(pubKey, plaintex []byte) string {
	mac := hmac.New(sha256.New, pubKey)
	mac.Write(plaintex)
	md := mac.Sum(nil)
	return hex.EncodeToString(md)
}

// HmacSha1WithKey Hmac sha1 with key
func HmacSha1WithKey(plainText []byte, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write(plainText)
	return hex.EncodeToString(h.Sum(nil))
}

// HmacSha512Rand hmac sha512
func HmacSha512Rand(category string) string {
	secString := category + Int64ToString(time.Now().UnixNano())
	randBytes, _ := RandBytes(64)
	mac := hmac.New(sha512.New, randBytes)
	mac.Write([]byte(secString))
	md := mac.Sum(nil)
	return hex.EncodeToString(md)
}

// HmacSha1WithKey Hmac sha1 with key
func HmacSha1WithKeyByte(plainText []byte, key string) []byte {
	h := hmac.New(sha1.New, []byte(key))
	h.Write(plainText)
	return h.Sum(nil)
}
