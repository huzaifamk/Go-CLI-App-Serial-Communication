package helpers

import (
	"encoding/hex"
	"fmt"
)

func HexToBytes(s string) []byte {
	data, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("% x", data)
	return data
}