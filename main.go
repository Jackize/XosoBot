package main

import (
	Find "XosoBot/Handle"
	ReadFile "XosoBot/ReadFile"
	BoxXoso "XosoBot/Structure"
	WriteFile "XosoBot/WriteFile"

	// "io/ioutil"
	"fmt"
	// "time"
)

func Menu() {
	fmt.Println("1. Dò kqxs hôm nay")
	fmt.Println("2. Kiếm kqxs theo ngày")
	fmt.Println("3. Dò kqxs từ bạn? ")
}

func main() {
	//WriteFile
	WriteFile.WriteFile()
	// Create struct BoxXoso
	var boxXoso = BoxXoso.BoxXoso{}
	//ReadFile and parse data to boxXoso
	boxXoso = ReadFile.ReadFile(boxXoso)

	resutl := Find.FindResultToday(boxXoso)
	fmt.Println(resutl)
}
