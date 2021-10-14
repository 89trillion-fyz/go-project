package util

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestIsMainBd(t *testing.T) {
	var bdId uint32

	bdId = 10203
	if result := IsMainBd(bdId); result != false {
		t.Error("The result is not the same as expected")
	}

	bdId = 10303
	if result := IsMainBd(bdId); result != false {
		t.Error("The result is not the same as expected")
	}

	bdId = 100103
	if result := IsMainBd(bdId); result != true {
		t.Error("The result is not the same as expected")
	}

	bdId = 200103
	if result := IsMainBd(bdId); result != true {
		t.Error("The result is not the same as expected")
	}

	bdId = 600103
	if result := IsMainBd(bdId); result != true {
		t.Error("The result is not the same as expected")
	}

	bdId = 500103
	if result := IsMainBd(bdId); result != false {
		t.Error("The result is not the same as expected")
	}
}

func TestIsBdValid(t *testing.T) {

	var bdId uint32

	bdId = 103
	if result := IsBuildIDValid(bdId); result != false {
		t.Error("The result is not the same as expected")
	}

	bdId = 100000
	if result := IsBuildIDValid(bdId); result != false {
		t.Error("The result is not the same as expected")
	}

	bdId = 9100003
	if result := IsBuildIDValid(bdId); result != true {
		t.Error("The result is not the same as expected")
	}

	bdId = 110100003
	if result := IsBuildIDValid(bdId); result != false {
		t.Error("The result is not the same as expected")
	}
}

func TestIsWallValid(t *testing.T) {

	var bdId uint32

	bdId = 103
	if result := IsWallValid(bdId); result != false {
		t.Error("The result is not the same as expected")
	}

	bdId = 1100003
	if result := IsWallValid(bdId); result != true {
		t.Error("The result is not the same as expected")
	}

	bdId = 1105003
	if result := IsWallValid(bdId); result != true {
		t.Error("The result is not the same as expected")
	}

	bdId = 11100003
	if result := IsWallValid(bdId); result != false {
		t.Error("The result is not the same as expected")
	}
}

func TestIsArmyValid(t *testing.T) {
	checkArmValid := IsArmyValid(uint32(103))
	if checkArmValid != false {
		t.Error("not equal to false as expected")
	}

	checkArmValid = IsArmyValid(uint32(102))
	if checkArmValid != false {
		t.Error("not equal to false as expected")
	}

	checkArmValid = IsArmyValid(uint32(220302))
	if checkArmValid != false {
		t.Error("not equal to false as expected")
	}

	checkArmValid = IsArmyValid(uint32(2200302))
	if checkArmValid != true {
		t.Error("not equal to true as expected")
	}
}

func TestTrim(t *testing.T) {

	str := "20402"

	checkTrim := Trim("\n\n     20402     \n\n")
	if str != checkTrim {
		t.Error("The result Trim function is not the same as expected")
	}

	checkTrim = Trim("\n\n     20402     \n\n   ")
	if str != checkTrim {
		t.Error("The result Trim function is not the same as expected")
	}

	checkTrim = Trim("\n\n     20402     \t\t   ")
	if str != checkTrim {
		t.Error("The result Trim function is not the same as expected")
	}
}

func TestMinUint32Of(t *testing.T) {
	if result := MinUint32Of(10, 5, 3); result != 3 {
		t.Error("The result is not the same as expected")
	}

	if result := MinUint32Of(100, 15, 20); result != 15 {
		t.Error("The result is not the same as expected")
	}
}

func TestGetStringInterface(t *testing.T) {

	var checkGetString interface{}

	switch checkGetString.(type) {
	case string:
		if result := GetStringInterface(checkGetString); result != checkGetString {
			t.Error("The result is not the same as expected")
		}

	default:
		if result := GetStringInterface(checkGetString); result != "" {
			t.Error("The result is not the same as expected")
		}

	}

}

func TestCalcRwdJoinId(t *testing.T) {
	rwdID := 100

	rwdType := 1000

	rwdResult := uint32(rwdID*100 + (rwdType % 100))

	if result := CalcRwdJoinId(rwdType, rwdID); result != rwdResult {

		t.Error("Hello")

	}
}

func TestGetUint32Interface(t *testing.T) {

	var checkGetUint32 interface{}

	switch checkGetUint32.(type) {

	case uint32:
		if result := GetUint32Interface(checkGetUint32); result != checkGetUint32 {
			t.Error("The result is not the same as expected")
		}

	default:
		if result := GetUint32Interface(checkGetUint32); result != 0 {
			t.Error("The result is not the same as expected")
		}

	}
}

func TestAnyToInt(t *testing.T) {

	var value interface{}

	value = "hello"
	if result := AnyToInt(value); result != 0 {
		t.Error("The result is not the same as expected")
	}

	value = "15"
	if result := AnyToInt(value); result != 15 {
		t.Error("The result is not the same as expected")
	}

	// 在进行string 转换 int的时候，如果string 超出最大值则会显示错误
	// value = "9223372036854775808"
	// if result := AnyToInt(value); result != 9223372036854775808 {
	//	t.Error("The result is not the same as expected")
	// }

	value = true
	if result := AnyToInt(value); result != 1 {
		t.Error("The result is not the same as expected")
	}

	value = false
	if result := AnyToInt(value); result != 0 {
		t.Error("The result is not the same as expected")
	}

	value = float32(30.5)
	if result := AnyToInt(value); result != 30 {
		t.Error("The result is not the same as expected")
	}

	// 这种情况下出错
	// value = uint64(18446744073709551615)
	// if result := AnyToInt(value); result != 18446744073709551615 {
	//	t.Error("The result is not the same as expected")
	// }

	// 在测试float64的时候数字较大会造成转换时的数字变化，造成不匹配
	value = float64(922337203685477580)
	if result := AnyToInt(value); result != 922337203685477632 {
		t.Error("The result is not the same as expected")
	}

}

func TestDeepCopy(t *testing.T) {

	// 测试map

	testValue := map[string]interface{}{
		"n1": "20402",
		"n2": "Oikura",
	}

	result := DeepCopy(testValue)

	if !reflect.DeepEqual(result, testValue) {

		t.Error("The result is not the same as the expected")

	}

	// 测试slice

	testValue1 := make([]interface{}, 2)

	testValue1[0] = "20402"
	testValue1[1] = "Oikura"

	result = DeepCopy(testValue1)

	if !reflect.DeepEqual(result, testValue1) {

		t.Error("The result is not the same as the expected")

	}

	var testValue2 interface{}

	if result := DeepCopy(testValue2); result != testValue2 {

		t.Error("The result is not the same as the expected")

	}

}

func TestInterfaceListToStringList(t *testing.T) {

	interfaceList := make([]interface{}, 2)

	// interfaceList[0] = 1
	// interfaceList[1] = 2

	result := InterfaceListToStringList(interfaceList)

	for i := range result {
		if reflect.TypeOf(result[i]) != reflect.TypeOf("") {
			t.Error("The result is not the same as the expected")
		}
	}
}

func TestGetCurDate(t *testing.T) {

	// 时间会随时改变
	time := time.Now()

	date := time.Format("20060102")

	if result := GetCurDate(); result != date {

		t.Error("The Date time is not correct")

	}

}

func TestItoa(t *testing.T) {

	var itoaValue interface{}

	itoaValue = int32(10)

	if result := Itoa(itoaValue); result != "10" {

		t.Error("The output is not the same as the expected")

	}

	itoaValue = uint32(10)

	if result := Itoa(itoaValue); result != "10" {

		t.Error("The output is not the same as the expected")

	}

	itoaValue = "10"

	if result := Itoa(itoaValue); result != "10" {

		t.Error("The output is not the same as the expected")

	}

	itoaValue = make([]interface{}, 2)

	if result := Itoa(itoaValue); result != "" {

		t.Error("The output is not the same as the expected")

	}

}

func TestAtoi(t *testing.T) {

	str := ""

	if result := Atoi(str); result != 0 {
		t.Error("The result is not the same as the expected")
	}

	str = "100"

	if result := Atoi(str); result != 100 {
		t.Error("The result is not the same as the expected")
	}

	str = "Oikura"

	if result := Atoi(str); result != 0 {
		t.Error("The result is not the same as the expected")
	}

}

func TestAtoInt64(t *testing.T) {

	str := ""

	if result := AtoInt64(str); result != int64(0) {
		t.Error("The result is not the same as the expected")
	}

	str = "100"

	if result := AtoInt64(str); result != int64(100) {
		t.Error("The result is not the same as the expected")
	}

	str = "Oikura"

	if result := AtoInt64(str); result != int64(0) {
		t.Error("The result is not the same as the expected")
	}
}

func TestAtoUint64(t *testing.T) {
	str := ""

	if result := AtoUint64(str); result != uint64(0) {
		t.Error("The result is not the same as the expected")
	}

	str = "100"

	if result := AtoUint64(str); result != uint64(100) {
		t.Error("The result is not the same as the expected")
	}

	str = "Oikura"

	if result := AtoUint64(str); result != uint64(0) {
		t.Error("The result is not the same as the expected")
	}
}

func TestBtoi(t *testing.T) {

	boolean := true

	if result := Btoi(boolean); result != 1 {
		t.Error("The result is not the same as the expected")
	}

	boolean = false

	if result := Btoi(boolean); result != 0 {
		t.Error("The result is not the same as the expected")
	}
}

func TestNtime(t *testing.T) {

	if uint32(time.Now().Unix()) != Ntime() {
		t.Error("The result is not the same as the expected")
	}

}

func TestChangeUintMapToInterface(t *testing.T) {

	newMap := map[string]uint32{}
	newChildMap := map[string]interface{}{}

	if result := ChangeUintMapToInterface(newMap); !reflect.DeepEqual(result, newChildMap) {

		t.Error("The result is not the same as the expected")

	}

	newMap["0"] = 1
	newMap["1"] = 2

	for k, v := range newMap {
		newChildMap[k] = v
	}

	if result := ChangeUintMapToInterface(newMap); !reflect.DeepEqual(result, newChildMap) {

		t.Error("The result is not the same as the expected")

	}

}

/*func TestCopyUintMap(t *testing.T) {

	newMap := map[uint32]uint32{}

	if result := CopyUintMap(newMap); !reflect.DeepEqual(result, newMap) {

		t.Error("The result is not the same as the expected")

	}

	newMap[0] = 1
	newMap[1] = 2

	if result := CopyUintMap(newMap); !reflect.DeepEqual(result, newMap) {

		t.Error("The result is not the same as the expected")

	}

}*/

func TestChangeMapTypetoint(t *testing.T) {

	newMap := map[string]string{}

	newChildMap := map[int32]uint32{}

	if result := ChangeMapTypetoint(newMap, true); !reflect.DeepEqual(result, newChildMap) {

		t.Error("The result is not the same as the expected")

	}
	if result := ChangeMapTypetoint(newMap, false); !reflect.DeepEqual(result, newChildMap) {

		t.Error("The result is not the same as the expected")

	}

	newMap["0"] = "0"
	newMap["1"] = "2"

	for k, v := range newMap {
		newChildMap[int32(Atoi(k))] = uint32(Atoi(v))
	}

	if result := ChangeMapTypetoint(newMap, false); !reflect.DeepEqual(result, newChildMap) {

		t.Error("The result is not the same as the expected")

	}
	if result := ChangeMapTypetoint(newMap, true); reflect.DeepEqual(result, newChildMap) {

		t.Error("The result is not the same as the expected")

	}

}

func TestChangeMapTypetoUint(t *testing.T) {

	newMap := map[string]string{}

	newChildMap := map[uint32]uint32{}

	if result := ChangeMapTypetoUint(newMap, false); !reflect.DeepEqual(result, newChildMap) {

		t.Error("The result is not the same as the expected")

	}

	newMap["0"] = "0"
	newMap["1"] = "2"

	for k, v := range newMap {
		newChildMap[uint32(Atoi(k))] = uint32(Atoi(v))
	}

	if result := ChangeMapTypetoUint(newMap, false); !reflect.DeepEqual(result, newChildMap) {

		t.Error("The result is not the same as the expected")

	}

}

func TestCopyintMap(t *testing.T) {

	newMap := map[int32]uint32{}

	if result := CopyintMap(newMap); !reflect.DeepEqual(result, newMap) {

		t.Error("The result is not the same as the expected")

	}

	newMap[0] = 1
	newMap[1] = 2

	if result := CopyintMap(newMap); !reflect.DeepEqual(result, newMap) {

		t.Error("The result is not the same as the expected")

	}

}

func TestGetBattlePassCurSeasonId(t *testing.T) {
	//
	//newTs := time.Now().Unix()
	//newNSeaon := int((newTs-1601942400)/(86400*15) + 1)
	//newOpen := 1601942400 + uint32((newNSeaon-1)*86400*15)
	//newClose := 1601942400 + uint32((newNSeaon)*86400*15) - 1
	//
	//if nSeason, open, close := GetBattlePassCurSeasonId(); nSeason != newNSeaon || open != newOpen || close != newClose {
	//	t.Error("The result is not the same as the expected")
	//}

}

func TestGetSpinFlopCurSeasonId(t *testing.T) {

	// newTs := time.Now().Unix()
	// newNSeaon := int((newTs-1606089600)/(86400*7) + 1)
	// newOpen := 1606089600 + uint32((newNSeaon-1)*86400*7)
	// newClose := 1606089600 + uint32((newNSeaon)*86400*7) - 1
	//
	// if nSeason, open, close := GetSpinFlopCurSeasonId(); nSeason != newNSeaon && open != newOpen && close != newClose {
	//
	//	t.Error("The result is not the same as the expected")
	//
	// }

}

func TestGetTeamWarSeasonId(t *testing.T) {

	// newSid := int((time.Now().Unix()-1611719400)/600 + 1)
	//
	// if result := GetTeamWarSeasonId(); result != newSid {
	//
	// 	t.Error("The result is not the same as the expected")
	//
	// }

}

func TestGetTeamWarStartAndEndTime(t *testing.T) {

	// newSid := 10
	//
	// newEnd := int64(1611719400 + newSid*600)
	//
	// if start, battleEnd, end := GetTeamWarStartAndEndTime(newSid); start != int64(1611719400) && battleEnd != newEnd && end != newEnd {
	//
	// 	t.Error("The result is not the same as the expected")
	//
	// }

}

func TestGetNextDayUnixTime(t *testing.T) {

	newTomorrowTs := time.Now().Unix() + 86400

	newNextTs := int(newTomorrowTs - newTomorrowTs%86400)

	if result := GetNextDayUnixTime(); result != newNextTs {

		t.Error("The result is not the same as the expected")

	}

}

func TestInArray(t *testing.T) {

	var newVal interface{}

	newVal = float32(2.0)

	var newArray interface{}

	newArray = 5

	if checkVal := InArray(newVal, newArray); checkVal != false {

		t.Error("The result is not the same as the expected")

	}

	newArray = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	if checkVal := InArray(newVal, newArray); checkVal != true {

		t.Error("The result is not the same as the expected")

	}

	newVal = float32(5.0)

	if checkVal := InArray(newVal, newArray); checkVal != false {

		t.Error("The result is not the same as the expected")

	}

	testValue1 := make([]interface{}, 2)

	testValue1 = append(testValue1, 10)
	testValue1 = append(testValue1, 20)

	newVal = 5

	if checkVal := InArray(newVal, testValue1); checkVal != false {

		t.Error("The result is not the same as the expected")

	}

	newVal = 10

	if checkVal := InArray(newVal, testValue1); checkVal != true {

		t.Error("The result is not the same as the expected")

	}

}

func TestHttpPost(t *testing.T) {

}

func TestGetCountryCodeByIp(t *testing.T) {

	ipStr := ""

	if result := GetCountryCodeByIp(ipStr); result != "AA" {

		t.Error("The result is not the same as the expected")

	}

	ipStr = "10.0.3.255"

	if result := GetCountryCodeByIp(ipStr); result != "AA" {

		t.Error("The result is not the same as the expected")

	}

}

func TestGetAppIdByPlatform(t *testing.T) {

	// setting.Setup()

	platform := ANDROID

	// 246
	if result := GetAppIdByPlatform(platform); result != 0 {

		t.Error("The result is not the same as the expected")

	}

	platform = IOS

	// 251
	if result := GetAppIdByPlatform(platform); result != 0 {

		t.Error("The result is not the same as the expected")

	}

}
func TestImplode(t *testing.T) {

	glue := "hello"

	pieces := []interface{}{1, 2, 3}

	testValue := "1hello2hello3"

	result := Implode(glue, pieces)
	if result != testValue {

		t.Error("The result is not the same as the expected")

	}

}

func TestIntranetIP(t *testing.T) {

	// var localIp []string
	// addrs, _ := net.InterfaceAddrs()
	// for _, address := range addrs {
	//	if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	//		if ipnet.IP.To4() != nil {
	//			localIp = []string{ipnet.IP.String()}
	//		}
	//	}
	// }
	//
	// if result, _ := IntranetIP(); result[0] != localIp[0] {
	//
	//	t.Error("The result is not the same as the expected")
	//
	// }
	//
	// localIp = []string{"10.0.2.14"}
	//
	// if result, _ := IntranetIP(); result[0] == localIp[0] {
	//
	//	t.Error("The result is not the same as the expected")
	//
	// }
}

func TestArraySum(t *testing.T) {

	array := []int{1, 2, 3, 4}

	sum := 10

	if result := ArraySum(array); result != sum {

		t.Error("The result is not the same as the expected")

	}

	sum = 11

	if result := ArraySum(array); result == sum {

		t.Error("The result is not the same as the expected")

	}

}

func TestRandomProbability(t *testing.T) {

	weight := make(map[string]int)

	weight["hello"] = 5
	weight["hi"] = 5
	weight["hi1"] = 5
	weight["hi2"] = 5
	array := []int{1, 2, 3, 4}
	total := 10
	count := 1

	checkTest := []int{}

	if result := RandomProbability(weight, array, total, count); reflect.TypeOf(checkTest) != reflect.TypeOf(result) {

		fmt.Println(result)
		t.Error("The result is not the same as the expected")

	}
}

func TestInArrayIndex(t *testing.T) {
	list1 := make([]int, 0)
	list2 := []int{1, 2, 2, 3, 4, 5, 9, 10, 11}
	list3 := []int{1, 2, 2, 3, 4, 5, 9, 10, 11}
	if result := InArrayIndex(1, list1); result != -1 {
		t.Error("The result is not the same as the expected")
	}
	if result := InArrayIndex(9, list2); result != 6 {
		t.Error("The result is not the same as the expected")
	}
	if result := InArrayIndex(29, list3); result != -1 {
		t.Error("The result is not the same as the expected")
	}
	if result := InArrayIndex(29, "aaa"); result != -1 {
		t.Error("The result is not the same as the expected")
	}
	if result := InArrayIndex("9", list3); result != -1 {
		t.Error("The result is not the same as the expected")
	}
	if result := InArrayIndex([]int{9}, list3); result != -1 {
		t.Error("The result is not the same as the expected")
	}
	if result := InArrayIndex(map[int]int{
		9: 9,
	}, list3); result != -1 {
		t.Error("The result is not the same as the expected")
	}
}
