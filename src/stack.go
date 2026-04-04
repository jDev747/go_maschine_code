package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)
var OPTION_AUTOCLEAR_ARG = false
var STACK_ARG []any
var STACK_PERSONAL []any
func ReadStackArg(index int) any {
	if len(STACK_ARG) < index+1 {
		log.Fatal("PANIC: INVALID STACK [STACK ARG] <GMC> MissingParamInStack")
	}
	return STACK_ARG[index]
}
func CommandPushStr(instruction []byte) {
	if len(instruction) < 2 {
		log.Fatal("PANIC: INVALID INSTRUCTION [PUSHSTR <MISSING> ...] <GMC> MissingStack")
	}
	if len(instruction) < 3 {
		log.Fatal("PANIC: INVALID INSTRUCTION [PUSHSTR STACK <MISSING>] <GMC> MissingString")
	}
	stackany := instruction[1]

	stack := int(stackany)
	if stack > 1 {
		log.Fatal("PANIC: INVALID INSTRUCTION [PUSHSTR <TO_BIG> ...] <GMC> InvalidStack: " + fmt.Sprint(stack))
	}
	str := ParseString(instruction[2:])
	if stack == 0 {
		STACK_ARG = append(STACK_ARG, str)
	} else {
		STACK_PERSONAL = append(STACK_PERSONAL, str)
	}
}
func CommandClearStack(instruction []byte) {
	if len(instruction) < 2 {
			log.Fatal("PANIC: INVALID INSTRUCTION [CLEARSTACK <MISSING>] <GMC> MissingStack")
		}
		stack := instruction[1]
		if stack > 1 {
			log.Fatal("PANIC: INVALID INSTRUCTION [PUSHSTR <TO_BIG> ...] <GMC> InvalidStack: " + fmt.Sprint(stack))
		}
		if stack == 0 {
			STACK_ARG = make([]any, 0, 5)
		} else {
			STACK_PERSONAL = make([]any, 0, 5)
		}
}
func CommandPushInt(instruction []byte) {
	if len(instruction) < 2 {
		log.Fatal("PANIC: INVALID INSTRUCTION [PUSHINT <MISSING> ...] <GMC> MissingStack")
	}
	if len(instruction) < 3 {
		log.Fatal("PANIC: INVALID INSTRUCTION [PUSHINT STACK <MISSING>] <GMC> MissingInt")
	}
	stack := int(instruction[1])
	if stack > 1 {
		log.Fatal("PANIC: INVALID INSTRUCTION [PUSHINT <TO_BIG> ...] <GMC> InvalidStack: " + fmt.Sprint(stack))
	}
	var stringtopush strings.Builder
	for _, byteitem := range instruction[2:] {
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
if instruction[1] == 0 {
	STACK_ARG = append(STACK_ARG, i)
} else {
	STACK_PERSONAL = append(STACK_PERSONAL, i)
}
}