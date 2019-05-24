package util

import (
	"testing"
)

func TestEncode(t *testing.T) {
	tcs := []int{64, 758, 2987, 459667}
	exp := []string{"10", "Bs", "kh", "1mEJ"}
	for i, tc := range tcs {
		r := Encode(tc)
		if r != exp[i] {
			t.Errorf("The %dth case is incorrect, expect %v, but %v.", i, exp[i], r)
		}
	}
}


func TestTo64Base(t *testing.T) {
	exp := []byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '-', '.',
	}
	for i := 0; i < 64; i++ {
		r := to64Base(i)
		if r != exp[i] {
			t.Errorf("The %dth case is wrong, expect %v, but %v", i, exp[i], r)
		}
	}
} 
