package Find

import (
	Crawler "XosoBot/CrawlData"
	BoxXoso "XosoBot/Structure"
	"fmt"
	"log"
	"strings"
	"time"
)

// Check error
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Print result kqxs today
func FindResultToday(boxXoso BoxXoso.BoxXoso) BoxXoso.Xoso {
	if date := time.Now(); date.Format("31-01-2006") != boxXoso.Xoso[0].Day {
		fmt.Println("Chưa có kqxs " + date.Format("02/01/2006"))
		fmt.Print("Bạn có thể coi kqxs ngày ")
	}
	return boxXoso.Xoso[0]
}

func FindResultFromNumber(boxXoso BoxXoso.BoxXoso) string {
	var kqxs = boxXoso.Xoso[0].Kqxs
	var contex string
	fmt.Print("Nhập số bạn muốn dò: ")
	_, err := fmt.Scanf("%s", &contex)
	CheckError(err)
	if len(contex) == 5 {
		if contex == kqxs.DB {
			return "Bingo!!! Bạn đã trúng giải đặc biệt. Hãy chia cho tôi 50% với nhé"
		} else if contex == kqxs.Nhat {
			return "Bingo!!! Bạn đã trúng giải Nhất. Hãy chia cho tôi 50% với nhé"
		} else {
			for _, v := range kqxs.Nhi {
				if contex == v {
					return "Bingo!!! Bạn đã trúng giải Nhì. Hãy chia cho tôi 50% với nhé"
				}
			}
			for _, v := range kqxs.Ba {
				if contex == v {
					return "Bingo!!! Bạn đã trúng giải Ba. Hãy chia cho tôi 50% với nhé"
				}
			}
		}
	} else if len(contex) == 4 {
		for _, v := range kqxs.Bon {
			if contex == v {
				return "Bingo!!! Bạn đã trúng giải Bốn. Hãy chia cho tôi 50% với nhé"
			}
		}
		for _, v := range kqxs.Nam {
			if contex == v {
				return "Bingo!!! Bạn đã trúng giải Năm. Hãy chia cho tôi 50% với nhé"
			}
		}
	} else if len(contex) == 3 {
		for _, v := range kqxs.Sau {
			if contex == v {
				return "Bingo!!! Bạn đã trúng giải Sáu. Hãy chia cho tôi 50% với nhé"
			}
		}
	} else if len(contex) == 2 {
		for _, v := range kqxs.Bay {
			if contex == v {
				return "Bingo!!! Bạn đã trúng giải Bảy. Hãy chia cho tôi 50% với nhé"
			}
		}
	}
	return "Find not found"
}

func String(boxXoso BoxXoso.Xoso) string {
	var result string
	seperate := "\n-------------------------------------------------------------\n"
	result = boxXoso.Day + "\n" + "DB:\t\t\t" + boxXoso.Kqxs.DB + seperate + "Nhat:\t\t\t" + boxXoso.Kqxs.Nhat
	giai := ""
	for _, v := range boxXoso.Kqxs.Nhi {
		giai += v + "\t\t"
	}
	result += seperate + "Nhi:\t\t" + giai
	giai = ""
	for i, v := range boxXoso.Kqxs.Ba {
		if i == len(boxXoso.Kqxs.Ba)/2 {
			giai += "\n\t"
		}
		giai += v + "\t\t"
	}
	result += seperate + "Ba:\t" + giai
	giai = ""
	for _, v := range boxXoso.Kqxs.Bon {
		giai += v + "\t\t"
	}
	result += seperate + "Bon:\t" + giai
	giai = ""
	for i, v := range boxXoso.Kqxs.Nam {
		if i == len(boxXoso.Kqxs.Nam)/2 {
			giai += "\n\t"
		}
		giai += v + "\t\t"
	}
	result += seperate + "Nam:\t" + giai
	giai = ""
	for _, v := range boxXoso.Kqxs.Sau {
		giai += v + "\t\t"
	}
	result += seperate + "Sau:\t" + giai
	giai = ""
	for _, v := range boxXoso.Kqxs.Bay {
		giai += v + "\t\t"
	}
	result += seperate + "Bay:\t" + giai + seperate
	return result
}

func FindResultFollowingDate() BoxXoso.BoxXoso {
	boxXoso := BoxXoso.BoxXoso{}
	var contex string
	fmt.Println("Nhập ngày bạn muốn tìm kiếm (DD/MM/YYYY): ")
	_, err := fmt.Scanf("%s", &contex)
	CheckError(err)
	day := strings.Split(contex, "/")[0]
	month := strings.Split(contex, "/")[1]
	if len(month) == 1 {
		month = "0" + month
	}
	year := strings.Split(contex, "/")[2]
	if len(year) == 2 {
		year = "20" + year
	}
	contex = "ngay-" + day + "-" + month + "-" + year
	Url := `https://xskt.com.vn/xsmb/` + contex
	boxXoso, err = Crawler.GetXosoByUrlFollowDay(Url, boxXoso)
	CheckError(err)
	return boxXoso
}
