package common

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"math/big"
	"path/filepath"
)

//GenerateRandomBytes is to generate randorm bytes
func GenerateRandomBytes(bytes int) []byte {
	b := make([]byte, bytes)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil
	}
	return b
}

// CopyBytes Returns an exact copy of the provided bytes
func CopyBytes(b []byte) (copiedBytes []byte) {
	copiedBytes = make([]byte, len(b))
	copy(copiedBytes, b)

	return
}

//AbsolutePath get absolute path
func AbsolutePath(Datadir string, filename string) string {
	if filepath.IsAbs(filename) {
		return filename
	}
	return filepath.Join(Datadir, filename)
}

//BigD bytes to big
func BigD(data []byte) *big.Int { return BytesToBig(data) }

//BytesToBig bytes to big
func BytesToBig(data []byte) *big.Int {
	n := new(big.Int)
	n.SetBytes(data)

	return n
}

//BytesToHexString convert bytes to hex string
func BytesToHexString(bytes []byte) string {

	return hex.EncodeToString(bytes)
}

func BytesToInt64(data []byte) int64 {

	bytesBuffer := bytes.NewBuffer(data)

	var value int64
	binary.Read(bytesBuffer, binary.BigEndian, &value)
	return value

}
func Int64ToBytes(data int64) []byte {

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, data)
	return bytesBuffer.Bytes()

}
