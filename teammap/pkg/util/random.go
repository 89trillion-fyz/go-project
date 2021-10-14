package util

import (
	"math/rand"
	"time"
)

// [0,n)
func RandomInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	if n <= 0 {
		return 0
	}
	return rand.Intn(n)
}

// [0,n)
func RandomInt32(n int32) int32 {
	rand.Seed(time.Now().UnixNano())
	if n <= 0 {
		return 0
	}
	return rand.Int31n(n)
}

// [0,n)
func RandomInt64(n int64) int64 {
	rand.Seed(time.Now().UnixNano())
	if n <= 0 {
		return 0
	}
	return rand.Int63n(n)
}

// 随机多次，随到的还能随到
// amount为抽的次数
// 返回amount个map的key
func GetResultsByWeight(weightMap map[uint32]int32, amount int32) []uint32 {
	result := make([]uint32, 0, amount)
	if amount <= 0 {
		return result
	}
	weightSum := int32(0)
	for _, w := range weightMap {
		if w > 0 {
			weightSum += w
		}
	}
	for index := amount; index > 0; index-- {
		tempSum := int32(0)
		rd := RandomInt32(weightSum) // [0,n)
		for i, w := range weightMap {
			if w <= 0 {
				continue
			}
			tempSum += w
			if rd < tempSum {
				result = append(result, i)
				// weightSum -= w
				break
			}
		}
	}
	return result
}

//[begin,end)
//范围内随机num个不同的结果
func GetRandomResultNotSame(begin, end int64, num int) []int64 {
	fixNum := end - begin
	if fixNum < int64(num) || num <= 0 || begin < 0 || end < 0 {
		return []int64{}
	}
	result := make([]int64, num)
	//范围过大时频繁创建/删除大array会有gc抖动问题，所以分数据量用两种方法
	if fixNum < 10000 {
		arr := make([]int64, fixNum)
		for i := int64(0); i < fixNum; i++ {
			arr[i] = i + begin
		}
		for j := 0; j < num; j++ {
			r := RandomInt(len(arr))
			result[j] = arr[r]
			DeleteInSlice(&arr, r)
		}
		return result
	}
	//会有数据分布不够离散的问题
	flag := map[int64]struct{}{}
	for i := 0; i < num; i++ {
		r := RandomInt64(fixNum)
		for {
			if _, ok := flag[r]; ok {
				r = (r + 1) % fixNum
				continue
			}
			result[i] = r + begin
			flag[r] = struct{}{}
			break
		}
	}
	return result
}

//[begin,end)
//范围内随机num个结果，采用相同种子
func GetRandomResult(begin, end int, num int) []int {
	if begin > end {
		return nil
	}

	itv := end - begin
	rand.Seed(time.Now().UnixNano())
	var result []int
	for i := 0; i < num; i++ {
		result = append(result, begin+rand.Intn(itv))
	}

	return result
}
