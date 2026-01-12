package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ada(a []int, x int) bool {
	for _, v := range a {
		if v == x {
			return true
		}
	}
	return false
}

func gen(n, max int) []int {
	s := []int{}
	for len(s) < n {
		x := rand.Intn(max)
		if !ada(s, x) {
			s = append(s, x)
		}
	}
	return s
}

func materi1() {
	rand.Seed(time.Now().UnixNano())

	// SOAL 1
	N := 128
	A, B, C := gen(5, N), gen(5, N), gen(3, N)

	R := []int{}
	for _, a := range A {
		if !ada(B, a) && !ada(C, a) {
			R = append(R, a)
		}
	}
	for _, b := range B {
		if !ada(A, b) && !ada(C, b) {
			R = append(R, b)
		}
	}
	for _, a := range A {
		if ada(C, a) && !ada(R, a) {
			R = append(R, a)
		}
	}

	fmt.Println("Soal 1 R:", R)

	f1 := []int{}
	for _, x := range R {
		if x%2 == 0 && x < N/4 {
			f1 = append(f1, x)
		}
	}
	fmt.Println("Filter:", f1)

	//  SOAL 2
	M, K := 8, 22
	S := gen(M, K)

	fmt.Println("\nSoal 2 S:", S)
	jml := 0
	for i := 0; i < len(S); i++ {
		for j := i + 1; j < len(S); j++ {
			if S[i]+S[j] > K {
				fmt.Println("{", S[i], ",", S[j], "}")
				jml++
			}
		}
	}
	fmt.Println("Total:", jml)
}