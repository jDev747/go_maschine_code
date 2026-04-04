package main

import "log"

var VARS map[any]any = make(map[any]any)

func CommandStrVar(instruction []byte) {
	if len(instruction) < 3 { 
		log.Fatal("PANIC: INVALID INSTRUCTION [VARSTR <MISSING> ...]  <GMC> MissingVarName")
	}
	 key := string(instruction[1:3])
	if len(instruction) == 3 {
		VARS[key] = nil
		return
	}
	VARS[key] = ParseString(instruction[3:])
}