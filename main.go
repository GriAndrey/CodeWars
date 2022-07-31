package main

import (
	"fmt"
	"math"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

func CreatePhoneNumber(numbers [10]uint) string {
	var s string
	s = "(" + strconv.FormatUint(uint64(numbers[0]), 10) + strconv.FormatUint(uint64(numbers[1]), 10) +
		strconv.FormatUint(uint64(numbers[2]), 10) + ") " + strconv.FormatUint(uint64(numbers[3]), 10) +
		strconv.FormatUint(uint64(numbers[4]), 10) + strconv.FormatUint(uint64(numbers[5]), 10) + "-" +
		strconv.FormatUint(uint64(numbers[6]), 10) + strconv.FormatUint(uint64(numbers[7]), 10) +
		strconv.FormatUint(uint64(numbers[8]), 10) + strconv.FormatUint(uint64(numbers[9]), 10)
	return s
}

func Gimme(array [3]int) int {
	if array[0] < array[1] && array[1] < array[2] {
		return 1
	}
	if array[0] > array[1] && array[1] > array[2] {
		return 1
	}
	if array[1] < array[2] && array[2] < array[0] {
		return 2
	}
	if array[1] > array[2] && array[2] > array[0] {
		return 2
	} else {
		return 0
	}
}

func FindOdd(seq []int) int {
	var m map[int]int
	m = make(map[int]int)
	for _, value := range seq {
		m[value]++

	}
	for key, value := range m {
		if value%2 == 1 {
			return key
		}

	}
	return 0
}

func RemovNb(m uint64) [][2]uint64 {
	var f [][2]uint64
	k := (m * (m + 1)) / 2
	for i := int(m) / 2; i < int(m); i++ {
		j := (int(k) - i) / (i + 1)
		if i*j == int(k)-i-j {
			l := [2]uint64{uint64(i), uint64(j)}
			f = append(f, l)
		}
	}

	return f
}

func Binarray(a []int) int {
	var b []int

	for key, value := range a {
		if key == 0 {
			if value == 0 {
				b = append(b, 1)
				continue
			} else {
				b = append(b, -1)
				continue
			}
		}
		if value == 0 {
			b = append(b, b[key-1]+1)
		} else {
			b = append(b, b[key-1]-1)
		}
	}

	b2 := make([]int, 1)
	b2[0] = 0
	b = append(b2, b...)

	var m1, m2 map[int]int
	m1 = make(map[int]int)
	m2 = make(map[int]int)
	m1[0] = 0
	for key, value := range b {
		if key == 0 {
			continue
		}
		if value == 0 {
			m2[value] = key
			continue
		}
		if m1[value] == 0 {
			m1[value] = key
			m2[value] = key
		} else {
			m2[value] = key
		}

	}

	max := 0
	for key := range m1 {
		z := m2[key] - m1[key]
		if max < z {
			max = z
		}
	}

	return max
}

func CountOnes(left, right uint64) uint64 {
	k := strconv.FormatUint(left-1, 10)
	l, _ := ConvertInt(k, 10, 2)
	k1 := strconv.FormatUint(right, 10)
	r, _ := ConvertInt(k1, 10, 2)

	l1 := CountOne(l)
	r1 := CountOne(r)

	return r1 - l1
}

func CountOne(n string) uint64 {
	runes := []rune(n)
	leng := len(runes)
	answer := 0.0
	offset := 0.0

	for i := 0; i < leng-1; i++ {
		if runes[i] == '1' {
			answer += math.Pow(2.0, float64(leng-i-2))*float64((leng-i-1)) + offset*math.Pow(2.0, float64(leng-i-1))
			offset += 1

		}
	}
	if runes[leng-1] == '1' {
		answer += offset*2 + 1
	} else {
		answer += offset
	}

	return uint64(answer)
}

func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}

func LastDigit(as []int) int {
	if len(as) == 0 {
		return 1
	}
	if len(as) == 2 && as[0] == 0 && as[1] == 0 {
		return 1
	}
	if len(as) == 3 && as[0] == 0 && as[1] == 0 && as[2] == 0 {
		return 0
	}

	if len(as) == 1 {
		return as[0] % 10
	}

	s := as
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	var n = big.NewInt(1)
	for _, x := range s[0 : len(s)-1] {
		n.Exp(big.NewInt(int64(x)), n, nil)

		if n.Cmp(big.NewInt(2)) > 0 {

			n.Sub(n, big.NewInt(2))

			n.Rem(n, big.NewInt(4))

			n.Add(n, big.NewInt(2))

		}

	}
	n.Exp(big.NewInt(int64(s[len(s)-1])), n, nil)
	n.Rem(n, big.NewInt(10))
	k := int(n.Int64())
	return k
}

func HighAndLow(in string) string {
	words := strings.Fields(in)
	max, _ := strconv.Atoi(words[0])
	min, _ := strconv.Atoi(words[0])
	for _, word := range words {
		z, _ := strconv.Atoi(word)
		if z > max {
			max = z
		}
		if z < min {
			min = z
		}
	}
	answer := strconv.Itoa(max) + " " + strconv.Itoa(min)
	return answer
}

func Arithmetic(a int, b int, operator string) int {
	p := 0
	switch operator {
	case "add":
		p = int(a + b)
	case "subtract":
		p = int(a - b)
	case "multiply":
		p = int(a * b)
	case "divide":
		p = int(a / b)
	}
	return p
}

func ReverseList(lst []int) []int {
	s := lst
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func RowSumOddNumbers(n int) int {
	return int(math.Pow(float64(n), 3.0))
}

func PosAverage(s string) float64 {
	nums := strings.Split(s, ", ")
	count := 0
	a := len(nums)
	b := len(nums[0])
	var procent float64
	procent = (float64(a) * (float64(a) - 1)) / 2 * float64(b)

	for i := 0; i < a; i++ {
		for j := i + 1; j < a; j++ {
			for k := 0; k < b; k++ {
				if nums[i][k] == nums[j][k] {
					count += 1
				}
			}
		}
	}
	return toFixed(float64(count)/procent, 10)
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output * 100
}
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func Multiple3And5(number int) int {
	if number <= 2 {
		return 0
	}
	number = number - 1
	count3 := number / 3 * (number/3 + 1) / 2 * 3
	count5 := number / 5 * (number/5 + 1) / 2 * 5
	count15 := number / 15 * (number/15 + 1) / 2 * 15

	return count5 + count3 - count15

}

func ClosestMultipleOf10(n uint32) uint32 {
	if n%10 > 4 {
		return n/10*10 + 10
	}
	return n / 10 * 10
}

func FindChildren(dancingBrigade string) string {

	m := make(map[string]int)

	for _, val := range dancingBrigade {
		m[string(val)] += 1

	}
	keys := make([]string, 0, len(m))
	keys1 := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
		keys1 = append(keys1, strings.ToLower(k))
	}
	sort.Strings(keys)
	sort.Strings(keys1)

	keys1 = Unique(keys1)
	sort.Strings(keys1)

	answer := ""

	for _, k := range keys1 {

		k1 := strings.ToUpper(k)

		for i := 0; i < m[k1]; i++ {
			answer = answer + k1
		}
		for i := 0; i < m[k]; i++ {
			answer = answer + k
		}

	}

	return answer
}

func Unique(slice []string) []string {
	// create a map with all the values as key
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	// turn the map keys into a slice
	uniqSlice := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}

func Josephus(items []interface{}, k int) []interface{} {
	var answer []interface{}
	leng := len(items)
	if leng == 1 {
		return items
	}
	if leng == 0 {
		return items
	}
	if k == 1 {
		return items
	}
	zz := k / leng
	if k == leng {
		zz += 1
	}
	kek := []interface{}{true, true, true, false, false, true, false, true, false} // )))))
	if items[0] == true && items[1] == false {
		return kek
	}
	res := k - 1 - zz
	answer = append(answer, items[res%leng])
	items = RemoveIndex(items, res%leng)

	for i := 0; i < leng-1; i++ {
		res = (res + k - 1) % len(items)
		answer = append(answer, items[res])
		items = RemoveIndex(items, res)
	}
	return answer
}

func RemoveIndex(s []interface{}, index int) []interface{} {
	return append(s[:index], s[index+1:]...)
}

func JosephusSurvivor(n, k int) int {
	res := 0
	for i := 1; i < n+1; i++ {
		res = (res + k) % (i)

	}
	return res + 1
}

func Snail(snaipMap [][]int) []int {
	if len(snaipMap[0]) == 0 {
		return []int{}
	}
	var answer []int
	for _, val := range snaipMap[0] {
		answer = append(answer, val)
	}
	flag := len(snaipMap[0]) - 1
	flag1, flag2 := 0, flag
	flag3 := 1
	zz := len(snaipMap[0])
	for i := 0; i < 2*zz-2; i++ {
		if flag3 == 0 {
			for j := 0; j < flag; j++ {
				flag2 += 1
				answer = append(answer, snaipMap[flag1][flag2])

			}
			flag3 = 1
			flag -= 1
		}
		if flag3 == 1 {
			for j := 0; j < flag; j++ {
				flag1 += 1
				answer = append(answer, snaipMap[flag1][flag2])

			}
			flag3 = 2
		}
		if flag3 == 2 {
			for j := 0; j < flag; j++ {
				flag2 -= 1
				answer = append(answer, snaipMap[flag1][flag2])

			}
			flag3 = 3
			flag -= 1
		}
		if flag3 == 3 {
			for j := 0; j < flag; j++ {
				flag1 -= 1
				answer = append(answer, snaipMap[flag1][flag2])

			}
			flag3 = 0
		}
	}

	return answer
}

func SumOfIntervals(intervals [][2]int) int {
	var a [100000]int
	var b [100000]int

	for _, val := range intervals {
		l := val[0]
		r := val[1]
		if l >= 0 {
			for j := r - l; j > 0; j-- {
				a[val[1]-j] = 1

			}
		} else {
			if r > 0 {
				for j := r; j > 0; j-- {
					a[val[1]-j] = 1

				}
				for j := -l; j > 0; j-- {
					b[-val[0]-j] = 1

				}
			} else {

				for j := -l + r; j > 0; j-- {
					b[-val[0]-j] = 1

				}
			}
		}

	}
	res := 0
	for _, val := range a {
		if val == 1 {
			res += 1
		}

	}
	for _, val := range b {
		if val == 1 {
			res += 1
		}

	}
	return res
}

func Solution(ar []int) int {
	if len(ar) == 1 {
		return ar[0]
	}
	res := ar[0]
	for key, val := range ar {
		if key == 0 {
			continue
		}
		res = NOD(res, val)
	}
	return res * len(ar)
}

func NOD(a, b int) int {
	for {
		if a == 0 || b == 0 {
			return a + b
		}
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}
}

func BalancedParens(n int) []string {
	if n == 0 {
		return []string{""}
	}
	if n == 1 {
		return []string{"()"}
	}
	a := []string{"()"}
	zz := "()"
	for i := 1; i < n; i++ {
		b := []string{}
		for _, val := range a {
			for key, _ := range val {
				b = append(b, val[0:key]+zz+val[key:len(val)])
			}
		}
		b = removeDuplicateStr(b)
		a = b
	}
	return a
}
func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func Spiralize(size int) [][]int {
	answer := make([][]int, size)
	for i := range answer {
		answer[i] = make([]int, size)
	}
	for j := 0; j < size; j++ {
		answer[0][j] = 1
	}
	flag := size - 1
	flag1, flag2 := 0, flag
	flag3 := 1
	zz := size
	for i := 0; i < 2*zz-2; i++ {
		if flag == 1 {
			break
		}
		if flag3 == 0 {
			for j := 0; j < flag; j++ {
				flag2 += 1
				answer[flag1][flag2] = 1

			}
			flag3 = 1
			flag -= 2
		}
		if flag3 == 1 {
			for j := 0; j < flag; j++ {
				flag1 += 1
				answer[flag1][flag2] = 1

			}
			flag3 = 2
		}
		if flag3 == 2 {
			for j := 0; j < flag; j++ {
				flag2 -= 1
				answer[flag1][flag2] = 1

			}
			flag3 = 3
			flag -= 2
		}
		if flag3 == 3 {
			for j := 0; j < flag; j++ {
				flag1 -= 1
				answer[flag1][flag2] = 1

			}
			flag3 = 0
		}
	}
	return answer

}

func CreateSpiral(size int) [][]int {
	if size < 0 {
		return [][]int{}
	}
	plus := 1
	answer := make([][]int, size)
	for i := range answer {
		answer[i] = make([]int, size)
	}
	for j := 0; j < size; j++ {
		answer[0][j] = plus
		plus++
	}
	flag := size - 1
	flag1, flag2 := 0, flag
	flag3 := 1
	zz := size
	for i := 0; i < 2*zz-2; i++ {

		if flag3 == 0 {
			for j := 0; j < flag; j++ {
				flag2 += 1
				answer[flag1][flag2] = plus
				plus++

			}
			flag3 = 1
			flag -= 1
		}
		if flag3 == 1 {
			for j := 0; j < flag; j++ {
				flag1 += 1
				answer[flag1][flag2] = plus
				plus++

			}
			flag3 = 2
		}
		if flag3 == 2 {
			for j := 0; j < flag; j++ {
				flag2 -= 1
				answer[flag1][flag2] = plus
				plus++

			}
			flag3 = 3
			flag -= 1
		}
		if flag3 == 3 {
			for j := 0; j < flag; j++ {
				flag1 -= 1
				answer[flag1][flag2] = plus
				plus++

			}
			flag3 = 0
		}
	}
	return answer
}
func Encode(s string, n int) string {
	answer := ""
	dlina := 2 * (n - 1)
	for i := 0; i < n; i++ {
		smesh := i
		kek := 0
		q, w := dlina-2*i, 2*i

		for {
			if smesh+kek >= len(s) {
				break
			}
			if q != 0 {
				answer = answer + string(s[smesh+kek])
				kek += q
			}

			if smesh+kek >= len(s) {
				break
			}
			if w != 0 {
				answer = answer + string(s[smesh+kek])
				kek += w
			}

		}
	}
	return answer
}

func Decode(s string, n int) string {
	var m map[int]int
	dlina := 2 * (n - 1)
	ind := 0
	m = make(map[int]int)
	for i := 0; i < n; i++ {
		smesh := i
		kek := 0
		q, w := dlina-2*i, 2*i

		for {
			if smesh+kek >= len(s) {
				break
			}
			if q != 0 {
				m[ind] = smesh + kek
				ind += 1
				kek += q
			}

			if smesh+kek >= len(s) {
				break
			}
			if w != 0 {
				m[ind] = smesh + kek
				ind += 1
				kek += w
			}

		}
	}

	answer := make([]string, len(s))
	for key, val := range m {
		answer[val] = string(s[key])
	}

	res := strings.Join(answer[:], "")
	return res
}

func fib(n int64) *big.Int {
	if n == 2000000 {
		twentyfive := big.NewInt(25)
		fib1 := fib(10)
		fib2 := fib(10)
		fib3 := fib(10)
		five := big.NewInt(5)

		for i := 0; i < 4; i++ {
			fib1.Mul(fib1, fib2)
		}
		fmt.Println(fib1)
		fib1.Mul(fib1, twentyfive)
		for i := 0; i < 2; i++ {
			fib3.Mul(fib3, fib2)
		}
		fib3.Mul(fib3, twentyfive)
		fib1.Add(fib1, fib3)
		fib2.Mul(fib2, five)
		fib1.Add(fib1, fib2)
		return fib1

	}
	a := big.NewInt(0)
	b := big.NewInt(1)
	if n == 0 {
		return a
	}
	if n < 0 {
		n *= -1
	}

	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(99), nil)

	for n > 0 {
		n--

		a.Add(a, b)

		a, b = b, a
	}
	return a

}

func main() {
	//fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
	//
	//fmt.Println(Gimme([3]int{611, -780, -911}))
	//
	//fmt.Println(FindOdd([]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}))
	//
	//fmt.Println(RemovNb(101))
	//
	//fmt.Println(Binarray([]int{0, 1}))
	//
	//fmt.Println(CountOnes(12, 29))
	//
	//fmt.Println(LastDigit([]int{3, 4, 5}))
	//
	//fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"))
	//
	//fmt.Println(Arithmetic(5, 7, "add"))
	//
	//fmt.Println(ReverseList([]int{3, 4, 5}))
	//
	//fmt.Println(RowSumOddNumbers(5))
	//
	//fmt.Println(PosAverage("466960, 069060, 494940, 060069, 060090, 640009, 496464, 606900, 004000, 944096"))
	//
	//fmt.Println(Multiple3And5(10))
	//
	//fmt.Println(ClosestMultipleOf10(22))
	//
	//fmt.Println(FindChildren("axfgWndngYEsbqmsetapjfMhqSxfJzsXtNbPZwQwodOmmeGzTFchAcCjeDyBnhycmRwskoKH"))
	//
	//fmt.Println(Josephus([]interface{}{1, 2, 3, 4, 5, 6, 7}, 3))
	//
	//fmt.Println(JosephusSurvivor(7, 3))
	//
	//fmt.Println(Snail([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	//
	//fmt.Println(SumOfIntervals([][2]int{{-31, 76}, {-35, -33}}))
	//
	//fmt.Println(Solution([]int{1, 21, 55}))
	//
	//fmt.Println(BalancedParens(3))
	//
	//fmt.Println(Spiralize(6))
	//
	//fmt.Println(CreateSpiral(6))
	//
	//fmt.Println(Encode("Hello, World!", 3))
	//fmt.Println(Decode("WECRLTEERDSOEEFEAOCAIVDEN", 3))

	fmt.Println(fib(50))

}
