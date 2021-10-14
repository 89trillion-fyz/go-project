package util

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	tk := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEyMzQ1NiwiZXhwIjoxNjM1NTYyNzEzLCJpc3MiOiJnaW4tYmxvZyJ9.85lKsws_twdw9i1Wcr1lkjSevYdsyfggTZHXcxwgB0A"
	gen, _ := GenerateToken("123456")
	fmt.Println(gen)
	if gen == tk || gen[:40] != tk[:40] {
		t.Error("generate token is not same as expect")
	}
	fmt.Println(gen)
}

func TestParseToken(t *testing.T) {
	tk := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEyMzQ1NiwiZXhwIjoxNjM1NTYyNzEzLCJpc3MiOiJnaW4tYmxvZyJ9.85lKsws_twdw9i1Wcr1lkjSevYdsyfggTZHXcxwgB0A"
	claim, _ := ParseToken(tk)
	if claim == nil {
		t.Error("parse token fail")
	} else {
		if claim.UserId != 123456 {
			t.Error("parse token userid not as expected")
		}
	}
}
