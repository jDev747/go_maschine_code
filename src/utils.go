package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var DECODER string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!?.()[]{}_&%$'\"/\\@<>|+-*~#= \n\t" // add nmore charaters //it is now in utils.go
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
func ParseString(bytearr []byte) string{
	var stringtopush strings.Builder
	for _, byteitem := range bytearr {
		convint := int(byteitem)
		if convint == 0xAC { //STREND / INTEND
			break
		}
		if convint > len(DECODER)-1 {
			log.Fatal("PANIC: INVALID STRING <GMC> InvalidChar: " + fmt.Sprint(int(byteitem)))
		}
		stringtopush.WriteString(string(DECODER[convint]))
	}
	return stringtopush.String()
}
func ParseInt(bytearr []byte) int{					//this is wierd
var stringtopush strings.Builder
	for _, byteitem := range bytearr {
		if byteitem == byte(0xAC) { //STREND / INTEND
			break
		} 
		bin := fmt.Sprintf("%08s", IntToBin(int(byteitem)))
		tensplace := BinToInt(bin[:4])
		fmt.Fprint(&stringtopush, tensplace)
		onesplace := BinToInt(bin[4:])
		fmt.Fprint(&stringtopush, onesplace)
		if tensplace > 9 {
			log.Fatal("PANIC: INVALID INT [PUSHINT <INVALID INT>] <GMC> InvalidInt: " + fmt.Sprint(tensplace))
		}
		if onesplace > 9 {
			log.Fatal("PANIC: INVALID INT [PUSHINT <INVALID INT>] <GMC> InvalidInt: " + fmt.Sprint(tensplace))
		}
	}
	i, err := strconv.Atoi(stringtopush.String())
	if err != nil {
		log.Fatal(err)
	}
	return i 
}