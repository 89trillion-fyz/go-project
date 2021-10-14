package util

import (
	"testing"
)

func TestRandomInt32(t *testing.T) {
	var rand int32
	for i := 0; i < 100; i++ {
		rand = RandomInt32(10)
		if rand < 0 || rand > 9 {
			t.Error("invalid random result ", rand)
		}
	}

	rand = RandomInt32(-1)
	if rand != 0 {
		t.Error("not equal 0 as expect")
	}
}

//func TestGetResultsByWeight(t *testing.T) {
//	weightMap := map[uint32]int32{}
//	weightMap[101] = 10
//	weightMap[201] = 20
//	weightMap[301] = 30
//
//	count := int32(10000)
//	lst := GetResultsByWeight(weightMap, count)
//	if len(lst) != int(count) {
//		t.Error("length not as expect")
//	}
//
//	resultMap := map[uint32]int{}
//	for _, v := range lst {
//		if v != 101 && v != 201 && v != 301 {
//			t.Error("random not expect value")
//		}
//
//		resultMap[v]++
//	}
//
//	r21 := float32(resultMap[201]) / float32(resultMap[101])
//	if r21 < 1.8 || r21 > 2.2 {
//		t.Errorf("result not as expect weight, expect 2:1, actual %d : %d = %f", resultMap[201], resultMap[101], r21)
//	}
//
//	r31 := float32(resultMap[301]) / float32(resultMap[101])
//	if r31 < 2.8 || r31 > 3.2 {
//		t.Errorf("result not as expect weight, expect 3:1, actual %d : %d = %f", resultMap[301], resultMap[101], r31)
//	}
//}

func TestRandomResultFromSliceUint32(t *testing.T) {
	for i := 0; i < 10; i++ {
		cids := []uint32{180, 181, 182}
		cids = Shuffle(cids)
	}

}
