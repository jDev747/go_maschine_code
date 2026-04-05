package main

import (
	"fmt"
	"log"
)
var OPTION_AUTOCLEAR_ARG = false
var STACK_ARG []any
var STACK_RETURN []any
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
		STACK_RETURN = append(STACK_RETURN, str)
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
			STACK_RETURN = make([]any, 0, 5)
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
	i := ParseInt(instruction[2:])
	if instruction[1] == 0 {
		STACK_ARG = append(STACK_ARG, i)
	} else {
		STACK_RETURN = append(STACK_RETURN, i)
	}
}
func CommandPushVar(instruction []byte) {
	if len(instruction) < 2 {
		log.Fatal("PANIC: INVALID INSTRUCTION [PUSHVAR <MISSING> ...] <GMC> MissingStack")
	}
	if len(instruction) < 4 {
		log.Fatal("PANIC: INVALID INSTRUCTION [PUSHVAR STACK <MISSING>] <GMC> MissingVar")
	}
	stack := int(instruction[1])
	key := string(instruction[2:])
	value, ok := VARS[key]
	if !ok || value == nil{
		log.Fatal("PANIC: INVALID INSTRUCTION [PUSHVAR STACK <VAR-DOES-NOT-EXIST>]")
	}
	if stack == 0 {
		STACK_ARG = append(STACK_ARG, value)
	} else {
		STACK_RETURN = append(STACK_RETURN, value)
	}
}