package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func readData(filePath string) ([]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		data = append(data, value)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func calculateStats(data []float64) (int, int, int, int) {
	var sum float64
	for _, value := range data {
		sum += value
	}
	average := int(math.Round(sum / float64(len(data))))

	sortedData := make([]float64, len(data))
	copy(sortedData, data)
	sort.Float64s(sortedData)
	medianIndex := len(sortedData) / 2
	median := int(math.Round((sortedData[medianIndex-1] + sortedData[medianIndex]) / 2))

	var variance float64
	for _, value := range data {
		variance += math.Pow(value-float64(average), 2)
	}
	variance = variance / float64(len(data))
	varianceInt := int(math.Round(variance))

	stdDev := int(math.Round(math.Sqrt(variance)))

	return average, median, varianceInt, stdDev
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . data.txt")
		os.Exit(1)
	}

	filePath := os.Args[1]
	data, err := readData(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	avg, median, variance, stdDev := calculateStats(data)

	fmt.Println("Average:", avg)
	fmt.Println("Median:", median)
	fmt.Println("Variance:", variance)
	fmt.Println("Standard Deviation:", stdDev)
}
