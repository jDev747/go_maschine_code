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
func IntToBytes(int_ int) []byte {
	var bytes_ []byte
	str := strconv.FormatInt(int64(int_), 10)
	var split []string
	curstr := []byte("00")
	for i, char := range str {
		curstr[i%2] = byte(char)
		if i%2 == 0 {
			split = append(split, string(curstr))
		} //very wierd
	}
	for _, twodigits := range split {
		bytes_ = append(bytes_, byte(StrToInt(twodigits, 10)))
	}
	return bytes_
}