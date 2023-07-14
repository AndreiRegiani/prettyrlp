package prettyrlp

import (
	"encoding/binary"
	"fmt"
	"strings"
)

// RLP prefixes (byte header)
const (
	String     = 0x80
	StringLong = 0xb7
	List       = 0xc0
	ListLong   = 0xf7
)

// Parse RLP data and returns a formatted string representation of it.
func Parse(data []byte, indentationLevel int) (result string, bytesProcessed int, err error) {
	if len(data) == 0 {
		return // Empty string
	}

	indentationStr := strings.Repeat("    ", indentationLevel)
	prefix := data[0]

	switch {
	// String (max 55 bytes)
	case prefix >= String && prefix < StringLong:
		length := prefix - String
		value := data[1 : length+1]
		result = fmt.Sprintf("%sString %s\n", indentationStr, value)
		bytesProcessed = len(value) + 1
		return

	// StringLong
	case prefix >= StringLong && prefix < List:
		sizePrefix := int(prefix - StringLong)
		length := 0

		if sizePrefix == 2 { // Most common case, max 64 KB
			length = int(binary.BigEndian.Uint16(data[1 : sizePrefix+1]))
		} else if sizePrefix == 4 { // max 4 GB
			length = int(binary.BigEndian.Uint32(data[1 : sizePrefix+1]))
		} else if sizePrefix == 8 { // max 18 exabytes
			length = int(binary.BigEndian.Uint64(data[1 : sizePrefix+1]))
		} else {
			err = fmt.Errorf("invalid size of length field: bytes=%d", sizePrefix)
			return
		}

		value := data[1+sizePrefix : 1+sizePrefix+length]
		result = fmt.Sprintf("%sStringLong %s\n", indentationStr, value)
		bytesProcessed = length + sizePrefix + 1
		return

	// List (max 55 bytes)
	case prefix >= List && prefix < ListLong:
		length := prefix - List
		value := data[1 : length+1]
		bytesProcessed++
		bytesRemaining := len(value)
		result = fmt.Sprintf("%sList {\n", indentationStr)

		for bytesRemaining > 0 {
			nestedResult, processed, err2 := Parse(value, indentationLevel+1)
			if err != nil {
				err = err2
				return
			}

			result += nestedResult
			bytesRemaining -= processed
			bytesProcessed += processed
			value = value[processed:]

			if processed == 0 {
				break
			}
		}

		result += fmt.Sprintf("%s}\n", indentationStr)
		return

	// ListLong
	case prefix >= ListLong:
		err = fmt.Errorf("ListLong is not supported")
		return
	}

	return
}
