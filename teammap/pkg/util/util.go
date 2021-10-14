package util

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"hash/crc32"
	"log"
	"math"
	"math/rand"
	"net"
	"net/http"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"teammap/pkg/errorlog"
	"teammap/pkg/myerr"
	"teammap/pkg/setting"
	"time"
	"unsafe"

	"github.com/oschwald/geoip2-golang"
)

func GetArmyLv(id uint32) (uint32, uint32) {
	id /= 100
	lv := id % 100
	name := id / 100
	return name, lv
}

func GetArmyCID(id uint32) uint32 {
	id /= 100
	return id
}

func GetBdTypeLv(bdType uint32, bdLv uint32) uint32 {
	return bdType*100 + bdLv
}

// 捕获panic，打印stack信息
func RecoverPanic() {
	if e := recover(); e != nil {
		errorlog.Panic("panic at")
	}
}

// Trim 移除两端空格、\t和换行符
func Trim(str string) string {
	return strings.Trim(str, " \n\t")
}

func Decimal(value float64) float64 {
	val, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return val
}

func GetStringInterface(i interface{}) string {
	s, ok := i.(string)
	if !ok {
		return ""
	}
	return s
}

func GetUint32Interface(i interface{}) uint32 {
	u, ok := i.(uint32)
	if !ok {
		return 0
	}
	return u
}

// 转Int
func AnyToInt(value interface{}) int {
	if value == nil {
		return 0
	}
	switch val := value.(type) {
	case int:
		return int(val)
	case int8:
		return int(val)
	case int16:
		return int(val)
	case int32:
		return int(val)
	case int64:
		return int(val)
	case uint:
		return int(val)
	case uint8:
		return int(val)
	case uint16:
		return int(val)
	case uint32:
		return int(val)
	case uint64:
		return int(val)
	case *string:
		v, err := strconv.Atoi(*val)
		if err != nil {
			return 0
		}
		return v
	case string:
		v, err := strconv.Atoi(val)
		if err != nil {
			return 0
		}
		return v
	case float32:
		return int(val)
	case float64:
		return int(val)
	case bool:
		if val {
			return 1
		} else {
			return 0
		}
	case json.Number:
		v, _ := val.Int64()
		return int(v)
	}

	return 0
}

func AnyToFloat64(value interface{}) float64 {
	// string
	if value == nil {
		return 0
	}
	switch val := value.(type) {
	case int:
		return float64(val)
	case int8:
		return float64(val)
	case int16:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case float32:
		return float64(val)
	case float64:
		return val
	case uint:
		return float64(val)
	case uint8:
		return float64(val)
	case uint16:
		return float64(val)
	case uint32:
		return float64(val)
	case uint64:
		return float64(val)

	case *string:
		float, err := strconv.ParseFloat(*val, 64)
		if err != nil {
			return float64(0)
		}
		return float
	case string:
		float, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return float64(0)
		}
		return float
	case bool:
		if val {
			return float64(1)
		} else {
			return float64(0)
		}
	case json.Number:
		v, _ := val.Float64()
		return v
	}
	return float64(0)
}

// 多个输入中找最小uint32
func MinUint32Of(vars ...uint32) uint32 {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

// 根据奖励/cost类型和id拼接出 联合id
func CalcRwdJoinId(rwdType int, rwdId int) uint32 {
	joinId := rwdId*100 + (rwdType % 100)
	return uint32(joinId)
}

// 拆分物品id和类型
func SplitRwdJoinId(joinId int) (int, int) {
	return joinId % 100, joinId / 100
}

// stick army/hero name and level
func StickArmyOrHeroNameAndLv(name uint32, level uint32) int {
	return int(name*100 + level%100)
}

func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}

		return newSlice
	}

	return value
}

func InterfaceListToStringList(data []interface{}) []string {
	result := make([]string, len(data))
	for i, info := range data {
		if info == nil {
			result[i] = ""
			continue
		}
		result[i] = info.(string)
	}
	return result
}

// 获取当前日期 YYYYMMDD 格式
func GetCurDate() string {
	t := time.Now()
	str := t.Format("20060102")
	return str
}

func Max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func Itoa(a interface{}) string {
	switch at := a.(type) {
	case int, int8, int16, int64, int32:
		return strconv.FormatInt(reflect.ValueOf(a).Int(), 10)
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatInt(int64(reflect.ValueOf(a).Uint()), 10)
	case float32, float64:
		return strconv.FormatFloat(reflect.ValueOf(a).Float(), 'f', 0, 64)
	case string:
		return at
	}
	return ""
}

func Atoi(a string) int {
	if a == Strnull {
		return 0
	}
	r, e := strconv.Atoi(a)
	if e == nil {
		return r
	}
	return 0
}

func AtoInt64(a string) int64 {
	r, e := strconv.ParseInt(a, 10, 64)
	if e == nil {
		return r
	}
	return 0
}

func AtoInt32(a string) int32 {
	r, e := strconv.ParseInt(a, 10, 32)
	if e == nil {
		return int32(r)
	}
	return 0
}

func AtoUint64(a string) uint64 {
	r, e := strconv.ParseUint(a, 10, 64)
	if e == nil {
		return r
	}
	return 0
}

func AtoUint32(a string) uint32 {
	r, e := strconv.ParseUint(a, 10, 32)
	if e == nil {
		return uint32(r)
	}
	return 0
}

func Btoi(a bool) int {
	if a {
		return 1
	}
	return 0
}

func Ntime() uint32 {
	return uint32(time.Now().Unix())
}

func ChangeUintMapToInterface(father map[string]uint32) map[string]interface{} {
	childMap := map[string]interface{}{}
	if len(father) == 0 {
		return childMap
	}
	for k, v := range father {
		childMap[k] = v
	}
	return childMap
}

func ChangeMapTypetoint(father map[string]string, delZeroFlag bool) map[int32]uint32 {
	childMap := map[int32]uint32{}
	if len(father) == 0 {
		return childMap
	}
	for k, v := range father {
		if delZeroFlag && uint32(Atoi(v)) == 0 {
			continue
		}
		childMap[int32(Atoi(k))] = uint32(Atoi(v))
	}
	return childMap
}

func ChangeInterfaceMapTypeToInt(father map[string]interface{}, delZeroFlag bool) map[int32]uint32 {
	childMap := map[int32]uint32{}
	if len(father) == 0 {
		return childMap
	}
	for k, v := range father {
		if delZeroFlag && uint32(AnyToInt(v)) == 0 {
			continue
		}
		childMap[int32(Atoi(k))] = uint32(AnyToInt(v))
	}
	return childMap
}

func ChangeMapTypetoUint(father map[string]string, dealZeroFlag bool) map[uint32]uint32 {
	childMap := map[uint32]uint32{}
	if len(father) == 0 {
		return childMap
	}
	for k, v := range father {
		realV := uint32(Atoi(v))
		if dealZeroFlag && realV == 0 {
			continue
		}
		childMap[uint32(Atoi(k))] = realV
	}
	return childMap

}

func CopyUintMap(father map[uint32]uint32, dealZeroFlag bool) map[uint32]uint32 {
	childMap := map[uint32]uint32{}
	if len(father) == 0 {
		return childMap
	}
	for k, v := range father {
		if dealZeroFlag && v == 0 {
			continue
		}
		childMap[k] = v
	}

	return childMap
}

func CopyintMap(father map[int32]uint32) map[int32]uint32 {
	childMap := map[int32]uint32{}
	if len(father) == 0 {
		return childMap
	}
	for k, v := range father {
		childMap[k] = v
	}

	return childMap
}

// 当前转盘和翻牌的期数
// 奇数是转盘，偶数是翻牌
func GetSpinFlopCurSeasonId() (int, uint32, uint32) {
	ts := time.Now().Unix()
	startTs := 1606089600 - 86400*7
	nSeason := (ts-int64(startTs))/(86400*7) + 1
	return int(nSeason), uint32(startTs) + uint32((nSeason-1)*86400*7), uint32(startTs) + uint32((nSeason)*86400*7) - 1
	// return 9, 1601561033, 1711561033
}

func GetCurrDayStartTime() int {
	tomorrowTs := time.Now().Unix()
	nextTs := tomorrowTs - tomorrowTs%86400

	return int(nextTs)
}

func GetNextDayUnixTime() int {
	tomorrowTs := time.Now().Unix() + 86400
	nextTs := tomorrowTs - tomorrowTs%86400

	return int(nextTs)
}

// 判断元素是否在数组中
func InArray(val interface{}, array interface{}) bool {
	k := reflect.TypeOf(array).Kind()
	if k != reflect.Slice && k != reflect.Array {
		return false
	}
	s := reflect.ValueOf(array)
	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == val {
			return true
		}
	}

	return false
}

// 通用http post方法
func HttpPost(url string, body []byte, params map[string]string, headers map[string]string) (*http.Response, error) {
	// add post body
	var req *http.Request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
		// return nil, errors.New("new request is fail: %v \n")
	}
	// add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	// add headers
	for key, val := range headers {
		req.Header.Add(key, val)
	}
	// http client
	client := &http.Client{}
	log.Printf("Go %s URL : %s \n", http.MethodPost, req.URL.String())
	return client.Do(req)
}

func GetCountryCodeByIp(ipStr string) string {
	if len(ipStr) == 0 {
		return "AA"
	}

	db, err := geoip2.Open("/usr/share/GeoIP/GeoIP2-Country.mmdb")
	if err != nil {
		return "AA"
	}
	defer db.Close()
	ip := net.ParseIP(ipStr)
	country, err := db.Country(ip)
	if err != nil {
		return "AA"
	}

	return country.Country.IsoCode
}

func GetAppIdByPlatform(platform string) int {
	if platform == ANDROID {
		return setting.AndroidSetting.AppId
	} else {
		return setting.IOSSetting.AppId
	}
}

func Implode(glue string, pieces []interface{}) string {
	data := make([]string, len(pieces))
	for i, s := range pieces {
		data[i] = GetAnything(s)
	}
	return strings.Join(data, glue)
}

func GetAnything(i interface{}) string {
	switch i.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool, string:
		return fmt.Sprint(i)
	default:
		t := reflect.TypeOf(i)
		v := reflect.ValueOf(i)
		switch t.Kind() {
		case reflect.Struct:
			b := bytes.Buffer{}
			b.WriteString("{ ")
			for ii := 0; ii < v.NumField(); ii++ {
				b.WriteString(GetAnything((&v).Field(ii).Interface()))
				b.WriteString(" ")
			}
			b.WriteString("}")
			return b.String()
		case reflect.Ptr:
			return fmt.Sprint(GetAnything((&v).Elem().Interface()))
		case reflect.Slice:
			b := bytes.Buffer{}
			b.WriteString("[")
			for ii := 0; ii < v.Len(); ii++ {
				b.WriteString(GetAnything(v.Index(ii).Interface()))
				b.WriteString(",")
			}
			b.WriteString("]")
			return b.String()
		case reflect.Map:
			b := bytes.Buffer{}
			b.WriteString("{")
			for _, key := range v.MapKeys() {
				b.WriteString(fmt.Sprint(key))
				b.WriteString(":")
				b.WriteString(GetAnything(v.MapIndex(key).Interface()))
				b.WriteString(",")
			}
			b.WriteString("}")
			return b.String()
		}
	}
	return ""
}

// / 获取本机内网ip
func IntranetIP() (ips []string, err error) {
	ips = make([]string, 0)

	ifaces, e := net.Interfaces()
	if e != nil {
		return ips, e
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}

		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}

		// ignore docker and warden bridge
		if strings.HasPrefix(iface.Name, "docker") || strings.HasPrefix(iface.Name, "w-") {
			continue
		}

		addrs, e := iface.Addrs()
		if e != nil {
			return ips, e
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}

			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}

			ipStr := ip.String()
			if isIntranet(ipStr) {
				ips = append(ips, ipStr)
			}
		}
	}

	return ips, nil
}

func isIntranet(ipStr string) bool {
	if strings.HasPrefix(ipStr, "10.") || strings.HasPrefix(ipStr, "192.168.") {
		return true
	}

	if strings.HasPrefix(ipStr, "172.") {
		// 172.16.0.0-172.31.255.255
		arr := strings.Split(ipStr, ".")
		if len(arr) != 4 {
			return false
		}

		second, err := strconv.ParseInt(arr[1], 10, 64)
		if err != nil {
			return false
		}

		if second >= 16 && second <= 31 {
			return true
		}
	}

	return false
}

func ArraySum(weight []int) int {
	total := 0
	for _, v := range weight {
		total += v
	}
	return total
}

func RandomProbability(weight map[string]int, keys []int, total, count int) []int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := []int{}
	for count > 0 {
		randV := random.Intn(total)
		for _, key := range keys {
			if randV < weight[Itoa(key)] {
				result = append(result, key)
				break
			}
		}
		count--
	}

	return result
}

func RandomResultByProbability(weight []int, count int) []int {
	total := 0
	for _, wei := range weight {
		total += wei
	}
	result := []int{}
	for count > 0 {
		count--
		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		randVal := random.Intn(total)

		tmp := 0
		index := -1
		for i, val := range weight {
			tmp += val
			if randVal < tmp {
				index = i
				break
			}
		}
		if index == -1 {
			continue
		}
		result = append(result, index)
	}
	return result
}

func RandomResultFromSlice(items []string, count int) []string {
	if count > len(items) {
		return items
	}

	idxLst := rand.Perm(len(items))
	var results []string
	for idx := 0; idx < count; idx++ {
		results = append(results, items[idxLst[idx]])
	}

	return results
}

/**
 * @Description  洗牌
 * @Param
 * @return
 **/
func Shuffle(vals []uint32) []uint32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ret := make([]uint32, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}

func Curl(url string, method string, params *map[string]string, body *[]byte, headers *map[string]string) (*http.Response, error) {
	var (
		req *http.Request
		err error
	)

	// new request
	switch method {
	case http.MethodGet:
		req, err = http.NewRequest(http.MethodGet, url, nil)
	case http.MethodPost:
		req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(*body))
	default:
		return nil, errors.New("invalid http method")
	}
	if err != nil {
		return nil, err
	}

	// add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range *params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}

	// add headers
	if *headers != nil {
		for key, val := range *headers {
			req.Header.Set(key, val)
		}
	}

	// http client
	client := &http.Client{}

	return client.Do(req)
}

//删除slice里元素，但不保证原顺序
func DeleteInSlice(s *[]int64, index int) {
	l := len(*s)
	if l > index && index >= 0 {
		(*s)[index] = (*s)[l-1]
		(*s)[l-1] = 0
		((*reflect.SliceHeader)(unsafe.Pointer(s))).Len--
	}
}

// 通常先扣物品再加奖励，需要把扣物品后的balance和加奖励后的balance整合。
// costBalance里有的货币/道具种类如果rwdBalance里也有会直接覆盖，对应花钻石又抽到了钻石的情况
func AddCostBalanceToNew(costBalance map[uint32]uint64, rwdBalance map[uint32]uint64) {
	for id, amount := range costBalance {
		if _, ok := rwdBalance[id]; !ok {
			rwdBalance[id] = amount
		}
	}
}

func CsvReader(path, msg string) [][]string {
	cast, e := os.Open(path)
	if e != nil {
		panic(msg)
	}
	r, e := csv.NewReader(cast).ReadAll()
	if e != nil {
		panic(msg)
	}
	return r
}

//严禁线上使用
/*func PrintTimeCost() func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		fmt.Printf("time cost = %v\n", tc)
	}
}
*/

//ArrayCompare 比较数组是否相等
func ArrayCompare(a, b []uint32) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	//重新排序，防止因为顺序不一样而返回false
	a1 := make([]int, len(a))
	b1 := make([]int, len(b))
	for i := range a {
		a1[i] = int(a[i])
		b1[i] = int(b[i])
	}
	sort.Ints(a1)
	sort.Ints(b1)
	for i, v := range a1 {
		if v != b1[i] {
			return false
		}
	}
	return true
}

//ArrayContain 判断该数组是否包含该元素
func ArrayContain(array []uint32, num uint32) bool {
	for _, value := range array {
		if value == num {
			return true
		}
	}
	return false
}

/**
 * @Description: IP 转为 地址值
 * @Author: SkyDo
 * @param ip
 * @return uint
 * @return error
 */
func Ip2Long(ip string) (uint, error) {
	b := net.ParseIP(ip).To4()
	if b == nil {
		return 0, errors.New("invalid ip")
	}

	return uint(b[3]) | uint(b[2])<<8 | uint(b[1])<<16 | uint(b[0])<<24, nil
}

/**
 * @Description: 检测IP是否给定的区间内
 * @Author: SkyDo
 * @param ip
 * @param ipList
 * @return bool
 */
func CheckIp(ip string, ipList []string) bool {
	ipAllow := false
	for _, ipAddr := range ipList {
		authIp := strings.Split(ipAddr, "/")
		if len(authIp) >= 2 {
			//网段形式
			ipVal, _ := Ip2Long(ip)
			ipAddrVal, _ := Ip2Long(authIp[0])
			mark := Atoi(authIp[1])
			ipVal = ipVal >> (32 - mark)
			ipAddrVal = ipAddrVal >> (32 - mark)
			if ipVal == ipAddrVal {
				ipAllow = true
			}
		} else if ipAddr == ip {
			//确定IP，且命中
			ipAllow = true
		}

		if ipAllow {
			break
		}
	}

	return ipAllow
}

//公会任务目标任务id分割  轮回，任务具体的小id, 种类
func GetTeamTargetTaskIdInfo(taskId int) (int, int, int, *myerr.MyErr) {
	//taskId 编号是五位数字起步
	if taskId < 10000 {
		return 0, 0, 0, myerr.INVALID_PARAMS
	}
	return taskId / 10000, (taskId % 10000) / 100, taskId % 100, myerr.SUCCESS
}

//获取元素在该数组的索引
func InArrayIndex(val interface{}, array interface{}) int {
	k := reflect.TypeOf(array).Kind()
	if k != reflect.Slice && k != reflect.Array {
		return -1
	}
	s := reflect.ValueOf(array)
	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == val {
			return i
		}
	}
	return -1
}

func ArrayChunk(s interface{}, size int) []interface{} {
	if size < 1 {
		return nil
	}

	k := reflect.TypeOf(s).Kind()
	if k != reflect.Slice {
		return nil
	}

	v := reflect.ValueOf(s)
	length := v.Len()
	if length <= 0 {
		return nil
	}

	chunks := int(math.Ceil(float64(length) / float64(size)))
	var n []interface{}
	for i, end := 0, 0; chunks > 0; chunks-- {
		start := i * size
		end = (i + 1) * size
		if end > length {
			end = length
		}

		n = append(n, v.Slice(start, end))
		i++
	}
	return n
}

func GetUserGroup(userId string, groupSize int) int {
	crcVal := crc32.ChecksumIEEE([]byte(userId))
	groupIdx := crcVal % uint32(groupSize)

	return int(groupIdx)
}

//根据权重，随机生成新的随机数序列，按照新权重降序排序
func RandomWeight(weightMap map[uint32]int32) []uint32 {
	var members []uint32
	type kv struct {
		Key   uint32
		Value float64
	}
	var NewWeight []kv
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for cid, weight := range weightMap {
		rWeight := math.Pow(r.Float64(), 1/float64(weight))
		NewWeight = append(NewWeight, kv{cid, rWeight})
	}
	sort.Slice(NewWeight, func(i, j int) bool {
		return NewWeight[i].Value > NewWeight[j].Value // 降序
	})
	for _, kv := range NewWeight {
		members = append(members, kv.Key)
	}
	return members
}

//并发map
type ConcurrentMap struct {
	items        map[string]interface{}
	sync.RWMutex // Read Write mutex, guards access to internal map.
}

func (c *ConcurrentMap) Init() {
	c.items = make(map[string]interface{})
}

func (c *ConcurrentMap) Get(key string) interface{} {
	defer c.RUnlock()
	c.RLock()
	if v, ok := c.items[key]; ok {
		return v
	}
	return nil
}

func (c *ConcurrentMap) Set(k string, v interface{}) {
	defer c.Unlock()
	c.Lock()
	c.items[k] = v
}

//根据客户端版本分段获取 最大cvc
func GetMaxCvc(cvc int, cvcList []int) int {
	for _, v := range cvcList {
		if cvc >= v {
			return v
		}
	}
	return cvc
}
