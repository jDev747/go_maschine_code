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
