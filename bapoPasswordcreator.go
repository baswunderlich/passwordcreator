package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!?\\\"+-:;.[]()'="

func main() {
	version := "1.0.0"
	fmt.Println("BapoPasswordgenerator started " + version)
	fmt.Println("Bitte nur folgende Zeichen für die Erstellung verwenden: 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!?\\\"+-:;.[]()'=")
	fmt.Print("Bitte Laenge des Passworts eingeben (default 25):")
	stringOfLength := readNextString()
	var lengthOfPassword int = 0
	var err error = nil
	lengthOfPassword, err = strconv.Atoi(stringOfLength)
	if err != nil {
		lengthOfPassword = 25
	}
	fmt.Print("Bitte Seed eingeben:")
	seed := readNextString()
	fmt.Print("Bitte Plattform eingeben:")
	plattform := readNextString()
	fmt.Println("Your password:", GeneratePassword(lengthOfPassword, seed, plattform))
}

func GeneratePassword(lengthOfPassword int, seed string, plattform string) string {
	result := ""
	var seedNum int = 0
	seedNumTable := make([]int, 2)
	for i := 0; i < len(seed); i++ {
		seedNum = seedNum ^ int(seed[i])
		seedNumTable = append(seedNumTable, int(seedNum))
	}
	for i, v := range seedNumTable {
		seedNumTable[i] = v ^ seedNum
	}

	plattformTable := make([]byte, 2)
	for i := 0; i < len(plattform); i++ {
		plattformTable = append(plattformTable, byte(plattform[i]))
	}
	for i := 0; i < lengthOfPassword; i++ {
		erg := (seedNum ^ int(plattformTable[i%len(plattformTable)])) % len(chars)
		//fmt.Println(seedNum ^ int(plattformTable[i%len(plattformTable)]))
		result += string(chars[erg])
		seedNum = (erg + 1) ^ seedNumTable[(i+1)%len(seedNumTable)]
		seedNumTable = append(seedNumTable, seedNum)
	}

	return result
}

func readNextString() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input = strings.Replace(input, "\n", "", -1)
	//For Windows
	return strings.Replace(input, "\r", "", -1)
}
