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
	fmt.Println("4. Dò kqxs trúng được bao nhiêu lần từ 1 quãng thời gian ")
	fmt.Println("5. Exit")
}

func main() {
	//ReadFile and parse data to boxXoso
	boxXoso := ReadFile.ReadFile()
	var choose int
	for choose != 5 {
		Menu()
		fmt.Print("Nhập số bạn muốn chọn: ")
		_, err := fmt.Scan(&choose)
		Handle.CheckError(err)
		switch choose {
		case 1:
			fmt.Println(Handle.String(Handle.FindResultToday(boxXoso)))
		case 2:
			fmt.Scanln()
			fmt.Println(Handle.String(Handle.FindResultFollowingDate().Xoso[0]))
		case 3:
			fmt.Scanln()
			fmt.Println(Handle.FindResultFromNumber(boxXoso))
		case 4:
			fmt.Scanln()
			Handle.FindTimes()
		case 5:
			fmt.Println("Thanks for trying my demo Project!!!")
		default:
			choose = 0
		}
	}
}
