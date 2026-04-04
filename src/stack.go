package main

import (
	"fmt"
	"log"
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
	var stringtopush strings.Builder
	for _, byteitem := range instruction[2:] {
		convint := int(byteitem)
		if convint == 0xAC {
			break
		}
		if convint > len(DECODER)-1 {
			log.Fatal("PANIC: INVALID STRING [PUSHSTR STACK <INVALID STRING>] <GMC> InvalidChar: " + fmt.Sprint(int(byteitem)))
		}
		stringtopush.WriteString(string(DECODER[convint]))
	}
	if stack == 0 {
		STACK_ARG = append(STACK_ARG, stringtopush.String())
	} else {
		STACK_PERSONAL = append(STACK_PERSONAL, stringtopush.String())
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