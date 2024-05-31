package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findUpdate(x int, arr []int, secondChance []bool, frames int) bool {
	for i := 0; i < frames; i++ {
		if arr[i] == x {
			secondChance[i] = true
			return true
		}
	}
	return false
}

func replaceUpdate(x int, arr []int, secondChance []bool, frames int, pointer int) int {
	for {
		if !secondChance[pointer] {
			arr[pointer] = x
			return (pointer + 1) % frames
		}
		secondChance[pointer] = false
		pointer = (pointer + 1) % frames
	}
}

func printFaults(referenceString string, frames int) {
	pointer, pf := 0, 0
	arr := make([]int, frames)
	for i := range arr {
		arr[i] = -1
	}
	secondChance := make([]bool, frames)

	str := strings.Split(referenceString, " ")
	for _, s := range str {
		x, _ := strconv.Atoi(s)
		if !findUpdate(x, arr, secondChance, frames) {
			pointer = replaceUpdate(x, arr, secondChance, frames, pointer)
			pf++
		}
	}

	fmt.Println("Número total de falhas de página:", pf)
	fmt.Println("Lista final:")
	for i := 0; i < frames; i++ {
		fmt.Print(arr[i], " -> ")
	}
	fmt.Println("(head)")
}

func main() {
	var referenceString string
	var frames int
	var resp string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Digite uma string para referência. Ex: 0 1 2 3 4...")
		referenceString, _ = reader.ReadString('\n')
		referenceString = strings.TrimSpace(referenceString)
		fmt.Println("Digite o número de frames.")
		fmt.Scanf("%d", &frames)
		printFaults(referenceString, frames)
		fmt.Println("Deseja continuar? (S/N)")
		resp, _ = reader.ReadString('\n')
		resp = strings.TrimSpace(resp)
		if resp != "S" {
			break
		}
	}
}
