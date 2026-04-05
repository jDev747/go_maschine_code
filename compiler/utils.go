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
// the old one was VERY bad
func IntToBytes(n int) []byte {
    if n == 0 {
        return []byte{0}
    }
    s := strconv.Itoa(n)
    if len(s)%2 == 1 {
        s = "0" + s
    }
    result := make([]byte, len(s)/2)
    for i := 0; i < len(s); i += 2 {
        pair, _ := strconv.Atoi(s[i : i+2])
        result[i/2] = byte(pair)
    }
    return result
}