package helpers

import (
	"encoding/hex"
	"fmt"
)

// This is the helper fuction, due to some technical issues, it has not been called from here but it's content has been exported in the open.go file.
func HexToBytes(s string) []byte {
	data, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("% x", data)
	return data
}
