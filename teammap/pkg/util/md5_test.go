package util

import "testing"

func TestEncodeMD5(t *testing.T) {

	testVar := "468b12b7e3a4b6e84e08327a56e9e156"

	if result := EncodeMD5("Oikura"); result != testVar {

		t.Error("EncodeMD5 result is not the same as expected")

	}

}
