package Find

import (
	BoxXoso "XosoBot/Structure"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

// Check error
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Check date to update data in file json.
// If date in file json is old, should return false, and vice versa
func CheckUpdateData() bool {
	file, _ := ioutil.ReadFile("xoso_data.json")
	data := BoxXoso.BoxXoso{}
	_ = json.Unmarshal([]byte(file), &data)
	date := time.Now()
	currentDay := "Kết quả xổ số " + date.Format("2/08/2006")
	setTime := "6:00PM"
	return currentDay != data.Xoso[0].Day && setTime == date.Format(time.Kitchen)
}

// Print result kqxs today
func FindResultToday(boxXoso BoxXoso.BoxXoso) string {
	var resutl string
	seperate := "\n-------------------------------------------------------------\n"
	resutl = boxXoso.Xoso[0].Day + "\n" + "DB:\t\t\t" + boxXoso.Xoso[0].Kqxs.DB + seperate + "Nhat:\t\t\t" + boxXoso.Xoso[0].Kqxs.Nhat
	giai := ""
	for _, v := range boxXoso.Xoso[0].Kqxs.Nhi {
		giai += v + "\t\t"
	}
	resutl += seperate + "Nhi:\t\t" + giai
	giai = ""
	for i, v := range boxXoso.Xoso[0].Kqxs.Ba {
		if i == len(boxXoso.Xoso[0].Kqxs.Ba)/2 {
			giai += "\n\t"
		}
		giai += v + "\t\t"
	}
	resutl += seperate + "Ba:\t" + giai
	giai = ""
	for _, v := range boxXoso.Xoso[0].Kqxs.Bon {
		giai += v + "\t\t"
	}
	resutl += seperate + "Bon:\t" + giai
	giai = ""
	for i, v := range boxXoso.Xoso[0].Kqxs.Nam {
		if i == len(boxXoso.Xoso[0].Kqxs.Nam)/2 {
			giai += "\n\t"
		}
		giai += v + "\t\t"
	}
	resutl += seperate + "Nam:\t" + giai
	giai = ""
	for _, v := range boxXoso.Xoso[0].Kqxs.Sau {
		giai += v + "\t\t"
	}
	resutl += seperate + "Sau:\t" + giai
	giai = ""
	for _, v := range boxXoso.Xoso[0].Kqxs.Bay {
		giai += v + "\t\t"
	}
	resutl += seperate + "Bay:\t" + giai
	return resutl
}
