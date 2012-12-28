package main

import (
	"bytes"
	"container/list"
	"fmt"
	"math"
)

// problem 1
func sumOfMultiplesOf3and5(max int) int {
	sum := 0
	for i := 0; i < max; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

// generative fibonnaci sequence using closure
func fibonacci() func() int {
	a, b := 1, 2
	return func() int {
		fib := a
		a, b = b, a+b
		return fib
	}
}

func problem2() {
	fib := fibonacci()
	val := fib()
	sum := 0
	for val <= 4000000 {
		if val%2 == 0 {
			sum += val
			fmt.Println(val)
		}
		val = fib()
	}
	fmt.Println(val)
}

func PrintList(l *list.List) {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for e := l.Front(); e != nil; e = e.Next() {
		buffer.WriteString(" " + fmt.Sprintf("%s", e.Value) + " ")
	}
	buffer.WriteString("]")
	fmt.Println(buffer.String())
}

func PrimeSieve(maximum int64) *list.List {
	s := make([]bool, maximum)
	for i := range s {
		s[i] = true
	}
	sqrt := int64(math.Sqrt(float64(maximum)))
	for i := int64(2); i < sqrt; i++ {
		if s[i] {
			for j := i * i; j < maximum; j += i {
				s[j] = false
			}
		}
	}
	primes := list.New()
	for i := int64(2); i < maximum; i++ {
		if s[i] {
			primes.PushBack(i)
		}
	}
	return primes
}

// use trial division
func Factors(n int64) *list.List {
	factors := list.New()
	if n == 1 {
		factors.PushBack(n)
		return factors
	}
	primes := PrimeSieve(int64(math.Sqrt(float64(n))) + 1)
	for p := primes.Front(); p != nil; p = p.Next() {
		if p.Value.(int64)*p.Value.(int64) > n {
			break
		}
		for n%p.Value.(int64) == 0 {
			factors.PushBack(p.Value.(int64))
			n = n / p.Value.(int64)
		}
	}
	if n > 1 {
		factors.PushBack(n)
	}
	return factors
}

func Prime(p int64) bool {
	if p == 2 {
		return true
	}
	if p%2 == 0 {
		return false
	}
	for i := int64(3); i <= (int64(math.Sqrt(float64(p)))); i += 2 {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func NextPrime() func() int64 {
	var prime int64 = 2
	return func() int64 {
		p := prime
		n := prime + 1
		for ; !Prime(n); n++ {
		}
		prime = n
		return p
	}
}

func problem3() {
	factors := Factors(600851475143)
	fmt.Println(factors.Back().Value.(int64))
}

func problem7() {
	prime := NextPrime()
	var n int64
	for i := 1; i <= 10001; i++ {
		n = prime()
	}
	fmt.Println(n)
}

func main() {
	problem7()
}
