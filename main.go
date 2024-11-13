package main

import (
	"fmt"
	"math"
	"sync"
)

type Function func()

func main() {

	functions := map[string]Function{
		"squareMaster": squareExample,
		"matrix":       matrix,
	}

	fmt.Println("Opened function:")
	for name := range functions {
		fmt.Printf("- %s\n", name)
	}

	fmt.Print("\ninput command: ")
	var input string
	fmt.Scanln(&input)

	if function, exists := functions[input]; exists {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			function()
		}()
		wg.Wait()
	} else {
		fmt.Println("Функция не найдена.")
	}

}
func matrix() {
	fmt.Println("input matrix 3x3:")
	matrix := [3][3]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf(" input element[%d][%d] ", i, j)
			fmt.Scanln(&matrix[i][j])
		}
	}
	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}
	fmt.Println("you wan`t find determinant on matrix?")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "y":
		determinate := matrix[0][0]*(matrix[1][1]*matrix[2][2]-matrix[1][2]*matrix[2][1]) - matrix[0][1]*(matrix[1][0]*matrix[2][2]-matrix[1][2]*matrix[2][0]) + matrix[0][2]*(matrix[1][0]*matrix[2][1]-matrix[1][1]*matrix[2][0])
		fmt.Printf("determinate: %d", determinate)
	case "n":
		break
	}
}
func squareExample() {
	fmt.Println("input args:")
	//args := os.Args[1:]
	//a, _ := strconv.ParseFloat(args[0], 64)
	//b, _ := strconv.ParseFloat(args[1], 64)
	//c, _ := strconv.ParseFloat(args[2], 64)
	var a, b, c float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	D := math.Pow(b, 2) - 4*(a*c)
	if D > 0 {
		x1 := (-b - math.Sqrt(D)) / 2 * a
		x2 := (-b + math.Sqrt(D)) / 2 * a
		fmt.Println("x1:", x1, "x2:", x2)
	} else if D == 0 {
		x1 := -b / (2 * a)
		fmt.Println("x1:", x1, "D = 0")
	} else if D < 0 {
		D = math.Sqrt(-1 * D)
		x1 := complex(-b, -math.Sqrt(D))
		x2 := complex(-b, math.Sqrt(D))
		fmt.Printf("%v/2*a \n", x1)
		fmt.Printf("%v/2*a \n", x2)
	}
}
