package main

import (
	Handle "XosoBot/Handle"
	ReadFile "XosoBot/ReadFile"

	// "io/ioutil"
	"fmt"
	// "time"
)

func Menu() {
	fmt.Println("1. Dò kqxs hôm nay")
	fmt.Println("2. Kiếm kqxs theo ngày")
	fmt.Println("3. Dò kqxs từ bạn? ")
	fmt.Println("4. Exit ")
}

func main() {
	//ReadFile and parse data to boxXoso
	boxXoso := ReadFile.ReadFile()
	var choose int
	for choose != 4 {
		Menu()
		fmt.Print("Nhập số bạn muốn chọn: ")
		_, err := fmt.Scan(&choose)
		Handle.CheckError(err)
		switch choose {
		case 1:
			fmt.Println(Handle.FindResultToday(boxXoso))
		case 2:
			fmt.Println(Handle.FindResultFollowingDate(boxXoso))
		case 3:
			fmt.Println(Handle.FindResultFromNumber(boxXoso))
		case 4:
			fmt.Println("Thanks for trying my demo project")
		}
	}
}
