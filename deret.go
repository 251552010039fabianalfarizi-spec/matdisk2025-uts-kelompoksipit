package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func materi3() {
	rand.Seed(time.Now().UnixNano())

	//SOAL 5

	C1 := rand.Intn(5) + 1
	C2 := -1 * (rand.Intn(3) + 1)
	N := rand.Intn(12) + 1

	a0 := 1
	a1 := 1

	fmt.Printf("INPUT: C1=%d, C2=%d, N=%d\n", C1, C2, N)
	fmt.Println("Proses Perhitungan:")

	fmt.Printf("Suku 0: %d | ", a0)
	if N >= 1 {
		fmt.Printf("Suku 1: %d | ", a1)
	}

	prev2 := a0
	prev1 := a1
	current := 0

	for i := 2; i <= N; i++ {
		current = (C1 * prev1) + (C2 * prev2)
		fmt.Printf("Suku %d: %d", i, current)

		if i < N {
			fmt.Print(" | ")
		}

		prev2 = prev1
		prev1 = current
	}

	fmt.Printf("\nHASIL AKHIR. Suku ke-%d: %d\n\n", N, prev1)

	//SOAL 6

	a := float64(rand.Intn(7) + 1)
	r := rand.Float64() * 0.4
	N2 := rand.Intn(10) + 1

	SN := 0.0
	for i := 0; i < N2; i++ {
		SN += a * math.Pow(r, float64(i))
	}

	Sinf := a / (1 - r)
	percentage := (SN / Sinf) * 100

	fmt.Printf("Input Paket: a=%.0f, r=%.2f, N=%d\n", a, r, N2)
	fmt.Printf("Sum Berhingga S(%d): %.2f\n", N2, SN)
	fmt.Printf("Sum Tak Hingga S(inf): %.2f\n", Sinf)
	fmt.Printf("Persentase Kedekatan: %.2f%%\n", percentage)
}