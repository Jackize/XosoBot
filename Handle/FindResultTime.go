package Find

import (
	crawler "XosoBot/CrawlData"
	structure "XosoBot/Structure"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var mu sync.Mutex

func Total_Day(timeBegin string, timeEnd string) int {
	var data = map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
	dayB, err := strconv.Atoi(strings.Split(timeBegin, "/")[0])
	CheckError(err)
	monthB, err := strconv.Atoi(strings.Split(timeBegin, "/")[1])
	CheckError(err)
	yearB, err := strconv.Atoi(strings.Split(timeBegin, "/")[2])
	CheckError(err)

	dayE, err := strconv.Atoi(strings.Split(timeEnd, "/")[0])
	CheckError(err)
	monthE, err := strconv.Atoi(strings.Split(timeEnd, "/")[1])
	CheckError(err)
	yearE, err := strconv.Atoi(strings.Split(timeEnd, "/")[2])
	CheckError(err)
	if dayB == dayE && monthB == monthE && yearE == yearB {
		return 1
	} else if yearE < yearB {
		return 0
	}
	total_day := (yearE - yearB) * 365
	for i := yearB; i <= yearE; i++ {
		if i%4 == 0 && monthB <= 2 {
			total_day++
		}
	}
	if monthB <= monthE {
		for i := monthB; i <= monthE; i++ {
			total_day += data[i]
		}
	} else if monthB > monthE {
		for i := monthE + 1; i < monthB; i++ {
			total_day -= data[i]
		}
	}
	total_day = total_day - dayB - (data[monthE] - dayE)
	if total_day < 0 {
		return 0
	}
	return total_day
}

func GetUrlEachDay(timeBegin string, total_day int, UrlCh chan string, wg *sync.WaitGroup) {
	dayB, err := strconv.Atoi(strings.Split(timeBegin, "/")[0])
	CheckError(err)
	monthB, err := strconv.Atoi(strings.Split(timeBegin, "/")[1])
	CheckError(err)
	yearB, err := strconv.Atoi(strings.Split(timeBegin, "/")[2])
	CheckError(err)
	day, month, year := dayB, monthB, yearB
	for i := 0; i < total_day; i++ {
		if month == 12 && day == 31 {
			year++
			month = 1
			day = 1
		} else if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 {
			if day == 31 {
				day = 1
				month++
			} else {
				day++
			}
		} else if month == 4 || month == 6 || month == 9 || month == 11 {
			if day == 30 {
				day = 1
				month++
			} else {
				day++
			}
		} else if month == 2 {
			if year%4 == 0 && day == 29 {
				day = 1
				month++
			} else if year%4 != 0 && day == 28 {
				day = 1
				month++
			} else {
				day++
			}
		}
		url := strconv.Itoa(day) + "-" + strconv.Itoa(month) + "-" + strconv.Itoa(year)
		UrlCh <- url
	}
	close(UrlCh)
	wg.Done()
}

func GetDataFromURL(path string, Urlch chan string, Xoso chan structure.Kqxs, wg *sync.WaitGroup) {
	for querry := range Urlch {
		url := path + querry
		boxXoso := structure.BoxXoso{}
		nboxXoso, err := crawler.GetXosoByUrlFollowDay(url, boxXoso)
		if err != nil {
			fmt.Println(err)
		} else {
			Xoso <- nboxXoso.Xoso[0].Kqxs
		}
	}
	close(Xoso)
	wg.Done()
}
func FindResult(number string, KqCh chan structure.Kqxs, counter *int, wg *sync.WaitGroup) {
	len := len(number)
	count := counter
	for i := range KqCh {
		mu.Lock()
		if len == 5 {
			if number == i.DB {
				*count++
			} else if number == i.Nhat {
				*count++
			} else {
				for _, v := range i.Nhi {
					if number == v {
						*count++
					}
				}
				for _, v := range i.Ba {
					if number == v {
						*count++
					}
				}
			}
		} else if len == 4 {
			for _, v := range i.Bon {
				if number == v {
					*count++
				}
			}
			for _, v := range i.Nam {
				if number == v {
					*count++
				}
			}
		} else if len == 3 {
			for _, v := range i.Sau {
				if number == v {
					*count++
				}
			}
		} else if len == 2 {
			for _, v := range i.Bay {
				if number == v {
					*count++
				}
			}
		}
		mu.Unlock()
	}
	wg.Done()
}

// Find the number of occurrences of a number in an interval
func FindTimes() {
	var number, daybegin, dayend string

	fmt.Print("Nhập số bạn muốn dò: ")
	fmt.Scanln(&number)
	fmt.Print("From (DD/MM/YYYY): ")
	fmt.Scanln(&daybegin)
	fmt.Print("To (DD/MM/YYYY): ")
	fmt.Scanln(&dayend)
	total_day := Total_Day(daybegin, dayend)
	Url := `https://xskt.com.vn/xsmb/ngay-`
	count := 0

	t := time.Now()

	UrlCh := make(chan string, total_day)
	KqCh := make(chan structure.Kqxs, total_day)

	var wg sync.WaitGroup
	if total_day > 0 {
		wg.Add(3)
		go GetUrlEachDay(daybegin, total_day, UrlCh, &wg)
		go GetDataFromURL(Url, UrlCh, KqCh, &wg)
		go FindResult(number, KqCh, &count, &wg)
	}
	wg.Wait()
	fmt.Printf("Kết quả %s trúng được %d lần\n", number, count)
	cntime := time.Since(t)
	fmt.Println("Concurrent operation time: ", cntime)

	t = time.Now()
	count = 0
	NotUseGoroutines(daybegin, total_day, Url, number, &count)
	fmt.Printf("Kết quả %s trúng được %d lần\n", number, count)
	sqtime := time.Since(t)
	fmt.Println("Sequential operation time: ", sqtime)

}

func NotUseGoroutines(timeBegin string, total_day int, path, number string, counter *int) {
	dayB, err := strconv.Atoi(strings.Split(timeBegin, "/")[0])
	CheckError(err)
	monthB, err := strconv.Atoi(strings.Split(timeBegin, "/")[1])
	CheckError(err)
	yearB, err := strconv.Atoi(strings.Split(timeBegin, "/")[2])
	CheckError(err)
	day, month, year := dayB, monthB, yearB
	UrlSlice := make([]string, 0)
	for i := 0; i < total_day; i++ {
		if month == 12 && day == 31 {
			year++
			month = 1
			day = 1
		} else if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 {
			if day == 31 {
				day = 1
				month++
			} else {
				day++
			}
		} else if month == 4 || month == 6 || month == 9 || month == 11 {
			if day == 30 {
				day = 1
				month++
			} else {
				day++
			}
		} else if month == 2 {
			if year%4 == 0 && day == 29 {
				day = 1
				month++
			} else if year%4 != 0 && day == 28 {
				day = 1
				month++
			} else {
				day++
			}
		}
		url := strconv.Itoa(day) + "-" + strconv.Itoa(month) + "-" + strconv.Itoa(year)
		UrlSlice = append(UrlSlice, url)
	}
	XosoSlice := make([]structure.Kqxs, 0)
	for _, v := range UrlSlice {
		url := path + v
		boxXoso := structure.BoxXoso{}
		nboxXoso, err := crawler.GetXosoByUrlFollowDay(url, boxXoso)
		if err != nil {
			fmt.Println(err)
		} else {
			XosoSlice = append(XosoSlice, nboxXoso.Xoso[0].Kqxs)
		}
	}
	len := len(number)
	count := counter
	for _, i := range XosoSlice {
		if len == 5 {
			if number == i.DB {
				*count++
			} else if number == i.Nhat {
				*count++
			} else {
				for _, v := range i.Nhi {
					if number == v {
						*count++
					}
				}
				for _, v := range i.Ba {
					if number == v {
						*count++
					}
				}
			}
		} else if len == 4 {
			for _, v := range i.Bon {
				if number == v {
					*count++
				}
			}
			for _, v := range i.Nam {
				if number == v {
					*count++
				}
			}
		} else if len == 3 {
			for _, v := range i.Sau {
				if number == v {
					*count++
				}
			}
		} else if len == 2 {
			for _, v := range i.Bay {
				if number == v {
					*count++
				}
			}
		}
	}
	return
}
