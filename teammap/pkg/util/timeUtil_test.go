package util

import (
	"fmt"
	"testing"
	"time"
)

func TestGetMondayTimeStamp(t *testing.T) {
	result := GetMondayTimeStamp(time.Date(2021, 4, 15, 17, 14, 10, 0, time.UTC))
	fmt.Println(result)
	if result != time.Date(2021, 4, 12, 0, 0, 0, 0, time.UTC).Unix() {
		fmt.Println(result)
		t.Error("The result is not the same as the expected")
	}
}

func TestGetThis15MinTimeStamp(t *testing.T) {
	result := GetThis15MinTimeStamp(time.Date(2021, 4, 15, 17, 14, 10, 0, time.UTC))
	fmt.Println(result)
	if result != time.Date(2021, 4, 15, 17, 0, 0, 0, time.UTC).Unix() {
		fmt.Println(result, time.Date(2021, 4, 15, 17, 0, 0, 0, time.UTC).Unix())
		t.Error("The result is not the same as the expected")
	}
}
