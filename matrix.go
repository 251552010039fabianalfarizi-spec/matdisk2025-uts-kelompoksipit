package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateMatrix(n int, max int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = rand.Intn(max)
		}
	}
	return matrix
}

//SOAL 3

func multiplyMatrix(A, B [][]int, n int) [][]int {
	R := make([][]int, n)
	for i := 0; i < n; i++ {
		R[i] = make([]int, n)
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				R[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return R
}

func traceMatrix(R [][]int, n int) int {
	trace := 0
	for i := 0; i < n; i++ {
		trace += R[i][i]
	}
	return trace
}

//SOAL 4

func swapRows(M [][]int, r1 int, r2 int) {
	M[r1], M[r2] = M[r2], M[r1]
}

func findMax(M [][]int) (int, int, int) {
	max := M[0][0]
	row, col := 0, 0

	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[i]); j++ {
			if M[i][j] > max {
				max = M[i][j]
				row = i
				col = j
			}
		}
	}
	return max, row, col
}

func materi2() {
	rand.Seed(time.Now().UnixNano())

	N := 2

	// SOAL 3
	fmt.Println("Matrices generated (2x2)...")

	A := generateMatrix(N, 2)
	B := generateMatrix(N, 2)

	fmt.Println("Matrix A:", A)
	fmt.Println("Matrix B:", B)

	R := multiplyMatrix(A, B, N)
	trace := traceMatrix(R, N)

	fmt.Println("Hasil Matriks R:", R)
	fmt.Println("Trace:", trace)

	//SOAL 4
	M := generateMatrix(N, 5)

	fmt.Println("\nMatrix M (Generated):", M)
	fmt.Println("Menukar Baris 0 dan 1...")

	swapRows(M, 0, N-1)

	fmt.Println("Matrix M Terkini:", M)

	maxVal, row, col := findMax(M)
	fmt.Printf("Nilai Maksimum: %d ditemukan di Posisi (%d,%d)\n", maxVal, row, col)
}