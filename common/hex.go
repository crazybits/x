package common

import "encoding/hex"

//BytesToHexSting bytes to hex string
func BytesToHexSting(bytes []byte) string {

	return hex.EncodeToString(bytes)

}

//HexStringToBytes hex string to bytes
func HexStringToBytes(hexstr string) []byte {

	bytes, err := hex.DecodeString(hexstr)
	if err != nil {

		return nil

	}
	return bytes

}
