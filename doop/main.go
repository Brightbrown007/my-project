package main

import (
	"os"
)

func printStr(s string) {
	os.Stdout.Write([]byte(s))
}

func printNbr(n int64) {
	if n == 0 {
		printStr("0")
		return
	}
	if n < 0 {
		printStr("-")
		n = -n
	}
	var digits [20]byte
	i := 0
	for n > 0 {
		digits[i] = byte(n%10 + '0')
		n /= 10
		i++
	}
	for j := i - 1; j >= 0; j-- {
		os.Stdout.Write([]byte{digits[j]})
	}
}

func atoi(s string) (int64, bool) {
	var res int64
	sign := int64(1)
	i := 0
	if len(s) == 0 {
		return 0, false
	}
	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		res = res*10 + int64(s[i]-'0')
	}
	return res * sign, true
}

func absUint64(x int64) uint64 {
	if x < 0 {
		return uint64(-x)
	}
	return uint64(x)
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, ok1 := atoi(os.Args[1])
	op := os.Args[2]
	b, ok2 := atoi(os.Args[3])

	if !ok1 || !ok2 {
		return
	}

	const maxInt64 uint64 = 9223372036854775807

	switch op {
	case "+":
		res := a + b
		if (a > 0 && b > 0 && res < 0) || (a < 0 && b < 0 && res > 0) {
			return
		}
		printNbr(res)

	case "-":
		res := a - b
		if (a > 0 && b < 0 && res < 0) || (a < 0 && b > 0 && res > 0) {
			return
		}
		printNbr(res)

	case "*":
		if a == 0 || b == 0 {
			printStr("0\n")
			return
		}

		absA := absUint64(a)
		absB := absUint64(b)

		allowed := maxInt64
		if (a < 0) != (b < 0) {
			allowed++
		}

		if absA > allowed/absB {
			return
		}

		printNbr(a * b)

	case "/":
		if b == 0 {
			printStr("No division by 0\n")
			return
		}
		printNbr(a / b)

	case "%":
		if b == 0 {
			printStr("No modulo by 0\n")
			return
		}
		printNbr(a % b)

	default:
		return
	}

	printStr("\n")
}
