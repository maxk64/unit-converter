package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var reader bufio.Reader = *bufio.NewReader(os.Stdin)

const (
	kB = 0
	MB = 1
	GB = 2
)

func selectUnit(header string) uint64 {
	fmt.Print(
		header + "\n" + `1. kB
2. MB
3. GB

4. KiB
5. MiB
6. GiB

7. Abort
> `)

	var choice uint64 = 0
	for {
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal("Error reading input!")
		}

		input = strings.TrimSpace(input)

		choice, err = strconv.ParseUint(input, 10, 3) // `bitSize` 3 here gives us the perfect range: 0 to 7

		if err != nil {
			fmt.Println("Please try again.")
		} else {
			break
		}
	}

	if choice == 7 {
		os.Exit(0)
	}

	return choice
}

var options = []string{"kB", "MB", "GB", "KiB", "MiB", "GiB"}

func stringifyChoice(choice uint64) string {
	return options[choice-1]
}

func main() {
	from := selectUnit("I want to convert from...")
	to := selectUnit("...to")
	fmt.Printf("Got it. How many %ss do you want to convert to %ss?\n> ", stringifyChoice(from), stringifyChoice(to))

	var toConvertInt uint64 = 0
	for {
		toConvertString, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal("Error reading input!")
		}

		toConvertString = strings.TrimSpace(toConvertString)

		toConvertInt, err = strconv.ParseUint(toConvertString, 10, 64)

		if err != nil {
			fmt.Println("Please try again.")
		} else {
			break
		}
	}

	toConvert := float64(toConvertInt)

	// 1. convert to kBs, the smallest unit here
	var kilobytes float64 = 0
	switch from {
	case 1:
		kilobytes = toConvert
	case 2:
		kilobytes = toConvert * 1000
	case 3:
		kilobytes = toConvert * math.Pow(1000, 3)
	case 4:
		kilobytes = toConvert
	case 5:
		kilobytes = toConvert * 1024
	case 6:
		kilobytes = toConvert * math.Pow(1024, 3)
	}
	fmt.Printf("Resulting kBs: %f\n", kilobytes)

	// 2. convert the kBs to the respective unit
	var result float64 = 0
	switch to {
	case 1:
		result = kilobytes
	case 2:
		result = kilobytes / 1000
	case 3:
		result = kilobytes / 1000 / 1000
	case 4:
		result = (kilobytes + 24)
	case 5:
		result = (kilobytes + 24) / 1024
	case 6:
		result = (kilobytes + 24) / 1024 / 1024
	}
	fmt.Printf("%f %ss are %f %ss.\n", toConvert, stringifyChoice(from), result, stringifyChoice(to))
}
