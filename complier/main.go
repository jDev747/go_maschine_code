package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var DECODER string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!?.()[]{}_&%$'\"/\\@<>|+-*~#= \n\t"

func Raise(msg string, i string) {
	log.Fatalf("Compilation error: %s \nInstruction: %s", msg, i)
}
func Encode(msg string) []byte {
	list := make([]byte, 0, len(msg))
	for _, char := range msg {
		for i, thing := range DECODER {
			if thing == char {
				list = append(list, byte(i))
			}
		}
	}
	return list
}
func GetInstructions(path string) []string {
	bytesoriginal, err := os.ReadFile(path)
	content := string(bytesoriginal)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(content, "\n")
}
func CompileInstruction(instruction string) []byte {
	splitold := strings.Split(instruction, " ")
	split := make([]string, 0, len(splitold))
	for _, thing := range splitold {
		split = append(split, strings.TrimSpace(thing))
	}
	compiled := make([]byte, 0, 10)
	itype := split[0]
	switch itype {
	case "exit":
		compiled = append(compiled, byte(0xA9))
	case "option":
		if len(split) < 2 {
			Raise("Missing <option>", instruction)
		}
		i, err := strconv.ParseInt(split[1], 16, 0)
		if err != nil {
			log.Fatal(err)
		}
		compiled = append(compiled, byte(0xAA), byte(i))
	case "pushstr":
		if len(split) < 2 {
			Raise("Missing <stack>", instruction)
		}
		if len(split) < 3 {
			Raise("Missing <string>", instruction)
		}

		i, err := strconv.ParseInt(split[1], 16, 0)
		if err != nil {
			log.Fatal(err)
		}
		compiled = append(compiled, byte(0xAB), byte(i))
		compiled = append(compiled, Encode(strings.Join(split[2:], " "))...)
		compiled = append(compiled, byte(0xAC))
	case "call":
		if len(split) < 2 {
			Raise("Missing <function>", instruction)
		}
		i, err := strconv.ParseInt(split[1], 16, 0)
		if err != nil {
			log.Fatal(err)
		}
		compiled = append(compiled, byte(0xAD), byte(i))
	}
	return compiled
}
func main() {
	complete := make([]byte, 0, 30)
	for _, i := range GetInstructions("smth.gmc") {
		instruction := CompileInstruction(i)
		complete = append(complete, instruction...)
		complete = append(complete, byte(0xA8))
	}
	for _, item := range complete {
		fmt.Println(item)
	}
}
