package main

import (
	"log"
	"strconv"
)

func IntToBin(int_ int) string {
	return strconv.FormatInt(int64(int_), 2)
}
func BinToInt(bin string) int64 {
	val, err := strconv.ParseInt(bin, 2, 0)
	if err != nil {
		log.Fatal(err)
	}
	return val
}
func StrToInt(str string, base int) int64 {
	i, err := strconv.ParseInt(str, base, 64)
		if err != nil {
			log.Fatal(err)
		}
	return i
}
func IntToBytes(n int) []byte {
    if n == 0 {
        return []byte{0}
    }
    var bytes []byte
    for n > 0 {
        bytes = append([]byte{byte(n & 0xFF)}, bytes...)
        n >>= 8
    }
    return bytes
} // this is a standart algorithm for doing stuff like this apparently