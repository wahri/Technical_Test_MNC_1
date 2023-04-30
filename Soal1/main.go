package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findMatchingStrings(n int, kata []string) interface{} {
	// konversi set string jadi lowercase
	for i := 0; i < n; i++ {
		kata[i] = kata[i][:len(kata[i])]
		kata[i] = strings.ToLower(kata[i])
	}

	// cek set string
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if kata[i] == kata[j] {
				return []int{i + 1, j + 1}
			}
		}
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// input jumlah string
	fmt.Print("Masukkan jumlah string: ")
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Input harus berupa angka")
		return
	}

	// input string
	var kata []string
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan string #%d: ", i+1)
		scanner.Scan()
		kata = append(kata, scanner.Text())
	}

	// Output
	result := findMatchingStrings(n, kata)
	fmt.Println("Output: ", result)
}
