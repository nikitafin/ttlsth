package mul

import (
	"testing"
)

func TestMul(t *testing.T) {
	expectedStr := "aaaaaa"
	resultStr := Mul("a", 6)
	if expectedStr != resultStr {
		t.Errorf("%s != %s", expectedStr, resultStr)
	}

	expectedNum := 15
	resultNum := Mul(5, 3)
	if expectedNum != resultNum {
		t.Errorf("%s != %s", expectedStr, resultStr)
	}

	var expectedUn interface{}
	resultUn := Mul(5.3, 3)
	if expectedUn != resultUn {
		t.Errorf("%s != %s", expectedStr, resultStr)
	}
}
