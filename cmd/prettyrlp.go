package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	prettyrlp "prettyrlp/pkg"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: prettyrlp {hex data}")
		os.Exit(1)
	}

	hexString := args[0]

	// Convert hex string to []byte, e.g. "1234" -> {0x12, 0x34}
	data, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatal("error decoding hex string: ", err)
	}

	result, bytesProcessed, err := prettyrlp.Parse(data, 0)
	if err != nil {
		log.Fatalf("error parsing RLP data: %v (bytesProcessed=%d)", err, bytesProcessed)
	}

	fmt.Println(result)
}
