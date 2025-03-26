package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type SortStats struct {
	comparisons int
	swaps       int
}

func createMatrix(M, N int) [][]int {
	matrix := make([][]int, M)
	rand.Seed(time.Now().UnixNano())
	for i := range matrix {
		matrix[i] = make([]int, N)
		for j := range matrix[i] {
			matrix[i][j] = rand.Intn(100)
		}
	}
	return matrix
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%d\t", val)
		}
		fmt.Println()
	}
	fmt.Println()
}

func getValidatedSize(prompt string) int {
	var size int
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&size)
		if err != nil || size <= 0 {
			fmt.Println("Error: Size must be a positive integer.")
			// Clear input buffer
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		return size
	}
}

func bubbleSort(arr []int, stats *SortStats) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			stats.comparisons++
			if math.Abs(float64(arr[j])) > math.Abs(float64(arr[j+1])) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				stats.swaps++
			}
		}
	}
}

func selectionSort(arr []int, stats *SortStats) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			stats.comparisons++
			if math.Abs(float64(arr[j])) < math.Abs(float64(arr[minIdx])) {
				minIdx = j
			}
		}
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
			stats.swaps++
		}
	}
}

func insertionSort(arr []int, stats *SortStats) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 {
			stats.comparisons++
			if math.Abs(float64(arr[j])) > math.Abs(float64(key)) {
				arr[j+1] = arr[j]
				j--
				stats.swaps++
			} else {
				break
			}
		}
		arr[j+1] = key
	}
}

func shellSort(arr []int, stats *SortStats) {
	n := len(arr)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i
			for j >= gap {
				stats.comparisons++
				if math.Abs(float64(arr[j-gap])) > math.Abs(float64(temp)) {
					arr[j] = arr[j-gap]
					j -= gap
					stats.swaps++
				} else {
					break
				}
			}
			arr[j] = temp
		}
	}
}

func quickSort(arr []int, left, right int, stats *SortStats) {
	if left < right {
		pivot := arr[right]
		i := left - 1

		for j := left; j < right; j++ {
			stats.comparisons++
			if math.Abs(float64(arr[j])) <= math.Abs(float64(pivot)) {
				i++
				if i != j {
					arr[i], arr[j] = arr[j], arr[i]
					stats.swaps++
				}
			}
		}

		if i+1 != right {
			arr[i+1], arr[right] = arr[right], arr[i+1]
			stats.swaps++
		}

		pivotIdx := i + 1

		quickSort(arr, left, pivotIdx-1, stats)
		quickSort(arr, pivotIdx+1, right, stats)
	}
}

func printComparisonTable(methods []struct {
	name   string
	sortFn func([]int, *SortStats)
}, matrix [][]int) {
	// Print table header
	fmt.Println("\n+----------------+--------------+-------+")
	fmt.Println("|    Method      | Comparisons | Swaps |")
	fmt.Println("+----------------+--------------+-------+")

	// Calculate and print each method's stats
	for _, method := range methods {
		totalStats := &SortStats{}
		M, N := len(matrix), len(matrix[0])

		for col := 0; col < N; col++ {
			column := make([]int, M)
			for row := 0; row < M; row++ {
				column[row] = matrix[row][col]
			}
			stats := &SortStats{}
			method.sortFn(column, stats)
			totalStats.comparisons += stats.comparisons
			totalStats.swaps += stats.swaps
		}

		// Format with fixed width for alignment
		fmt.Printf("| %-14s | %12d | %6d |\n",
			method.name,
			totalStats.comparisons,
			totalStats.swaps)
	}

	// Print table footer
	fmt.Println("+----------------+--------------+-------+")
}

func main() {
	M := getValidatedSize("Enter number of rows M: ")
	N := getValidatedSize("Enter number of columns N: ")

	matrix := createMatrix(M, N)

	methods := []struct {
		name   string
		sortFn func([]int, *SortStats)
	}{
		{"Bubble Sort", bubbleSort},
		{"Selection Sort", selectionSort},
		{"Insertion Sort", insertionSort},
		{"Shell Sort", shellSort},
		{"Quick Sort", func(arr []int, stats *SortStats) {
			quickSort(arr, 0, len(arr)-1, stats)
		}},
	}

	fmt.Println("\nOriginal matrix:")
	printMatrix(matrix)

	// Sort and print each method's result
	for _, method := range methods {
		sortedMatrix := make([][]int, M)
		for i := range sortedMatrix {
			sortedMatrix[i] = make([]int, N)
			copy(sortedMatrix[i], matrix[i])
		}

		for col := 0; col < N; col++ {
			column := make([]int, M)
			for row := 0; row < M; row++ {
				column[row] = sortedMatrix[row][col]
			}
			stats := &SortStats{}
			method.sortFn(column, stats)
			for row := 0; row < M; row++ {
				sortedMatrix[row][col] = column[row]
			}
		}

		fmt.Printf("Matrix sorted by %s:\n", method.name)
		printMatrix(sortedMatrix)
	}

	// Print formatted comparison table
	printComparisonTable(methods, matrix)
}
