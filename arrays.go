package main

import (
	"fmt"
)

var (
	winList = [6]string{"Windows xp", "Linux 1", "Raspbian Teensy", "MacOS", "iOS", "Google Android"}
)

func main() {

	winList[0] = "windows10"
	winList[1] = "linux 16.04"
	winList[2] = "Raspbian Buster"

	fmt.Println(winList)

	newList := [9]string{}

	for i := 0; i < len(winList); i++ {
		newList[i] = winList[i]
	}

	newList[6] = "Ubuntu"
	newList[7] = "MS-Dos"
	newList[8] = "Solaris"

	fmt.Println(newList)
	fmt.Println(newList[:3])
	fmt.Println(newList[3:6])
	fmt.Println(newList[6:])

}
