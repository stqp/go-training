package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	lc "github.com/stqp/go-training/ch02/ex02/lengthconv"
	tc "github.com/stqp/go-training/ch02/ex02/tempconv"
	wc "github.com/stqp/go-training/ch02/ex02/weightconv"
)

func usageAndExit(exitCode int) {
	fmt.Println("Please specify a number.")
	os.Exit(exitCode)
}

func main() {

	if len(os.Args) > 2 {
		usageAndExit(1)
	}

	var input string
	if len(os.Args) == 2 {
		input = os.Args[1]
	} else {
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		input = stdin.Text()
	}

	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println(err)
		usageAndExit(1)
	}

	fmt.Printf("\nLength:\n")
	fmt.Printf(" %s = %s , %s = %s\n\n", lc.Feet(num), lc.FToM(lc.Feet(num)), lc.Meter(num), lc.MToF(lc.Meter(num)))

	fmt.Printf("Temperature:\n")
	fmt.Printf(" %s = %s , %s = %s\n\n", tc.Celsius(num), tc.CToF(tc.Celsius(num)), tc.Fahrenheit(num), tc.FToC(tc.Fahrenheit(num)))

	fmt.Printf("Weight:\n")
	fmt.Printf(" %s = %s , %s = %s\n\n", wc.Pound(num), wc.PToK(wc.Pound(num)), wc.Kilogramme(num), wc.KToP(wc.Kilogramme(num)))

}
