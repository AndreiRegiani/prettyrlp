package prettyrlp

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

func checkResult(t *testing.T, hexString string, expected string) {
	data, _ := hex.DecodeString(hexString)
	expectedBytes := len(data)
	result, bytesProcessed, err := Parse(data, 0)

	if err != nil {
		t.Errorf("parsing error: %v", err)
	}

	if result != expected {
		t.Errorf("actual=%v, expected=%v", result, expected)
	}

	if bytesProcessed != expectedBytes {
		t.Errorf("actual=%v, expected=%v", bytesProcessed, expectedBytes)
	}
}

func TestEmptyString(t *testing.T) {
	checkResult(t, "80", "String \n")
}

func TestString(t *testing.T) {
	checkResult(t, "8e416e647265692052656769616e69", "String Andrei Regiani\n")
}

func TestStringLong(t *testing.T) {
	dataSize := 512
	hexString := fmt.Sprintf("b90200%s", strings.Repeat("41", dataSize))
	expected := fmt.Sprintf("StringLong %s\n", strings.Repeat("A", dataSize))
	checkResult(t, hexString, expected)
}

func TestList(t *testing.T) {
	checkResult(t, "cc86616e6472656984616c6578", `List {
    String andrei
    String alex
}
`)
}
