package main

import (
	"log"
	"os"
	"strings"
)

var DECODER string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!?.()[]{}_&%$'\"/\\@<>|+-*~#= \n\t"
var encoderMap map[rune]byte

func init() {
	encoderMap = make(map[rune]byte)
	for i, char := range DECODER {
		encoderMap[char] = byte(i)
	}
}

func Raise(msg string, i string) {
	log.Fatalf("Compilation error: %s \nInstruction: %s", msg, i)
}
func Encode(msg string) []byte {
	list := make([]byte, 0, len(msg))
	for _, char := range msg {
		if val, ok := encoderMap[char]; ok {
			list = append(list, val)
		} else {
			Raise("Unknown character in string", string(char))
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
		i := StrToInt(split[1], 16)
		compiled = append(compiled, byte(0xAA), byte(i))
	case "pushstr":
		if len(split) < 2 {
			Raise("Missing <stack>", instruction)
		}
		if len(split) < 3 {
			Raise("Missing <string>", instruction)
		}

		stack := StrToInt(split[1], 16)
		compiled = append(compiled, byte(0xAB), byte(stack))
		compiled = append(compiled, Encode(strings.Join(split[2:], " "))...)
		compiled = append(compiled, byte(0xAC))
	case "call":
		if len(split) < 2 {
			Raise("Missing <function>", instruction)
		}
		i := StrToInt(split[1], 16)
		compiled = append(compiled, byte(0xAD), byte(i))
	case "pushint":
		if len(split) < 2 {
			Raise("Missing <stack>", instruction)
		}
		if len(split) < 3 {
			Raise("Missing <int>", instruction)
		}
		stack := StrToInt(split[1], 16)
		compiled = append(compiled, byte(0xB0), byte(stack))
		compiled = append(compiled, IntToBytes(int(StrToInt(split[2], 10)))...)
	case "clearstack":
		if len(split) < 2 {
			Raise("Missing <stack>", instruction)
		}
		stack := StrToInt(split[1], 16)
		compiled = append(compiled, byte(0xAF), byte(stack))
	case "varstr":
		if len(split) < 2 {
			Raise("Missing <varname>", instruction)
		}
		if split[1][0:2] == "A8" || split[1][2:4] == "A8" {
			Raise("Invalid: <varname> CANNOT contain byte A8", instruction)
		}
		compiled = append(compiled, byte(0xB1))
		compiled = append(compiled, byte(StrToInt(split[1][0:2], 16)))
		compiled = append(compiled, byte(StrToInt(split[1][2:4], 16)))
		if len(split) > 2 {
			compiled = append(compiled, Encode(strings.Join(split[2:], " "))...)
		}
	case "pushvar":
		if len(split) < 2 {
			Raise("Missing <stack>", instruction)
		}
		if len(split) < 3 {
			Raise("Missing <varname>", instruction)
		}
		compiled = append(compiled, byte(0xB2))
		compiled = append(compiled, byte(StrToInt(split[1], 16)))
		compiled = append(compiled, byte(StrToInt(split[2][0:2], 16)))
		compiled = append(compiled, byte(StrToInt(split[2][2:4], 16)))
	case "varint":
		if len(split) < 2 {
			Raise("Missing <varname>", instruction)
		}
		if split[1][0:2] == "A8" || split[1][2:4] == "A8" {
			Raise("Invalid: <varname> CANNOT contain byte A8", instruction)
		}
		compiled = append(compiled, byte(0xB3))
		compiled = append(compiled, byte(StrToInt(split[1][0:2], 16)))
		compiled = append(compiled, byte(StrToInt(split[1][2:4], 16)))
		if len(split) > 2 {
			compiled = append(compiled, IntToBytes(int(StrToInt(split[2], 10)))...)
		}
	case "readstack":
		if len(split) < 2 {
			Raise("Missing <stack>", instruction)
		}
		if len(split) < 3 {
			Raise("Missing <varname>", instruction)
		}
		if split[2][0:2] == "A8" || split[2][2:4] == "A8" {
			Raise("Invalid: <varname> CANNOT contain byte A8", instruction)
		}
		compiled = append(compiled, byte(0xB4), byte(StrToInt(split[1], 16)))
		compiled = append(compiled, byte(StrToInt(split[2][0:2], 16)))
		compiled = append(compiled, byte(StrToInt(split[2][2:4], 16)))
	case "//":
		//ignore: this will be a comment
	}
	return compiled
}
func main() {
	var complete []byte
	for _, i := range GetInstructions("../test/smth.gac") {
		instruction := CompileInstruction(i)
		complete = append(complete, instruction...)
		complete = append(complete, byte(0xA8))
	}
	os.WriteFile("../test/smth.gmc", complete, 0777)
}
