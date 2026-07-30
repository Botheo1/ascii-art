package ascii

import (
	"fmt"
	"os"
)

// Open the right text file
func OpenFile(fileName string) []byte {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error at opening file", err)
	}
	defer file.Close()
	bytes, err2 := os.ReadFile(fileName)
	if err2 != nil {
		fmt.Println("Error at reading file", err)
	}
	return bytes
}
