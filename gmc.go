package main //TODO: Find a good name!
import (
	"bytes"
	// "fmt"
	"log" 
	"os"
)
func safeIndex(slice []any, index int) [2]any{
	if index < 0 {
		log.Fatal("PANIC: SAFEINDEX PARAM < 0 <GMC>")
	}
	if len(slice) < index {
		return [2]any{0, false}
	} else {
		return [2]any{slice[index], true}
	}
} 
func GetInstructions(path string) [][]byte {
	bytesoriginal, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	bytesnew := bytes.Split(bytesoriginal, []byte{0xA8})
	cleaned := make([][]byte, 0, len(bytesnew))
	for _, bytearr := range bytesnew {
		
		if len(bytearr) > 0 {
			cleaned = append(cleaned, bytearr)
		}
	}
	return cleaned
}
func ReadInstruction(intruction []byte) {
	// TODO: make this
}

func main() {
	//tests go here
}
