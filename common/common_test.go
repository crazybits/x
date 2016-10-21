package common

import (
	"fmt"
	"testing"
)

func TestRandomBytes32(t *testing.T) {

	b := GenerateRandomBytes(32)
	if b == nil {

		t.Fail()

	}

	fmt.Println("rand bytes =", b)

}

func TestRandomBytes64(t *testing.T) {

	b := GenerateRandomBytes(64)
	if b == nil {

		t.Fail()

	}
	fmt.Println("rand bytes =", b)

}

func TestRandomBytes128(t *testing.T) {

	b := GenerateRandomBytes(128)
	if b == nil {

		t.Fail()

	}
	fmt.Println("rand bytes =", b)

}
