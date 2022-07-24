package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))

	fmt.Println(Gimme([3]int{611, -780, -911}))

	fmt.Println(FindOdd([]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}))

	fmt.Println(RemovNb(101))

	fmt.Println(Binarray([]int{0, 1}))

	fmt.Println(CountOnes(12, 29))

	fmt.Println(LastDigit([]int{3, 4, 5}))

	fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"))
}
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
