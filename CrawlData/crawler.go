package crawler

import (
	BoxXoso "XosoBot/Structure"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/sync/errgroup"
)

// Crawl all result in a single page
func GetAllXoso(boxXoso BoxXoso.BoxXoso, updateBox BoxXoso.BoxXoso) (BoxXoso.BoxXoso, error) {
	eg := errgroup.Group{}
	eg.Go(func() error {
		boxXoso.Xoso = append(updateBox.Xoso, boxXoso.Xoso...)
		fmt.Println("Getting data")
		return nil
	})
	err := eg.Wait()
	if err != nil {
		log.Fatal(err)
	}
	return boxXoso, nil
}

// Get each result xoso and parse it to struct BoxXoso
func GetXosoByUrl(currentUrl string, boxXoso BoxXoso.BoxXoso) (BoxXoso.BoxXoso, error) {
	doc, err := goquery.NewDocument(currentUrl)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("div .boxKQ").Each(func(index int, tableHtml *goquery.Selection) {
		var xoso BoxXoso.Xoso
		Day := tableHtml.Find("div:nth-child(2) span:not([class])").Text() //Correct way
		xoso.Day = Day
		DB := tableHtml.Find(".kqxs div:nth-child(1) .no-detail p").Text() //Correct way
		xoso.Kqxs.DB = DB
		Nhat := tableHtml.Find(".kqxs div:nth-child(2) .no-detail p").Text() //Correct way
		xoso.Kqxs.Nhat = Nhat
		tableHtml.Find(".kqxs div:nth-child(3) .no-detail .col-6").Each(func(i int, row *goquery.Selection) {
			xoso.Kqxs.Nhi = append(xoso.Kqxs.Nhi, row.Find("p").Text())
		})
		tableHtml.Find(".kqxs div:nth-child(4) .no-detail .col-4").Each(func(i int, row *goquery.Selection) {
			xoso.Kqxs.Ba = append(xoso.Kqxs.Ba, row.Find("p").Text())
		})
		tableHtml.Find(".kqxs div:nth-child(5) .no-detail .col-3").Each(func(i int, row *goquery.Selection) {
			xoso.Kqxs.Bon = append(xoso.Kqxs.Bon, row.Find("p").Text())
		})
		tableHtml.Find(".kqxs div:nth-child(6) .no-detail .col-4").Each(func(i int, row *goquery.Selection) {
			xoso.Kqxs.Nam = append(xoso.Kqxs.Nam, row.Find("p").Text())
		})
		tableHtml.Find(".kqxs div:nth-child(7) .no-detail .col-4").Each(func(i int, row *goquery.Selection) {
			xoso.Kqxs.Sau = append(xoso.Kqxs.Sau, row.Find("p").Text())
		})
		tableHtml.Find(".kqxs div:nth-child(8) .no-detail .col-3").Each(func(i int, row *goquery.Selection) {
			xoso.Kqxs.Bay = append(xoso.Kqxs.Bay, row.Find("p").Text())
		})
		boxXoso.Xoso = append(boxXoso.Xoso, xoso)
	})
	return boxXoso, nil
}

func GetXosoByUrlFollowDay(currentUrl string, boxXoso BoxXoso.BoxXoso) (BoxXoso.BoxXoso, error) {
	doc, err := goquery.NewDocument(currentUrl)
	if err != nil {
		log.Fatal(err)
	}
	var xoso BoxXoso.Xoso
	date := strings.Split(currentUrl, "-")
	Day := "Kết quả xổ số " + date[1] + "/" + date[2] + "/" + date[3] //Correct way
	xoso.Day = Day

	DB := doc.Find(".result  tr:nth-child(2) td:nth-child(2) em").Text() //Correct way
	xoso.Kqxs.DB = DB

	Nhat := doc.Find(".result  tr:nth-child(3) td:nth-child(2) p").Text() //Correct way
	xoso.Kqxs.Nhat = Nhat

	giaiNhi := strings.Split(doc.Find(".result  tr:nth-child(4) td:nth-child(2) p").Text(), " ")
	xoso.Kqxs.Nhi = append(xoso.Kqxs.Nhi, giaiNhi[0], giaiNhi[1])

	giaiBa := strings.Split(doc.Find(".result  tr:nth-child(5) td:nth-child(2) p").Text(), " ")
	xoso.Kqxs.Ba = GetString(giaiBa, 5)

	giaiTu := strings.Split(doc.Find(".result  tr:nth-child(7) td:nth-child(2) p").Text(), " ")
	xoso.Kqxs.Bon = GetString(giaiTu, 4)

	giaiNam := strings.Split(doc.Find(".result  tr:nth-child(8) td:nth-child(2) p").Text(), " ")
	xoso.Kqxs.Nam = GetString(giaiNam, 4)

	giaiSau := strings.Split(doc.Find(".result  tr:nth-child(10) td:nth-child(2) p").Text(), " ")
	xoso.Kqxs.Sau = GetString(giaiSau, 3)

	giaiBay := strings.Split(doc.Find(".result  tr:nth-child(11) td:nth-child(2) p").Text(), " ")
	xoso.Kqxs.Bay = GetString(giaiBay, 2)

	boxXoso.Xoso = append(boxXoso.Xoso, xoso)
	return boxXoso, nil
}

// Seperate concatenation of strings each result Text got from array in func GetXosoByUrlFollowDay
func GetString(doc []string, numberString int) []string {
	a := []rune(doc[len(doc)/2])
	b := doc[:len(doc)/2]
	c := doc[len(doc)/2+1:]
	doc = b
	res := ""
	for i, v := range a {
		res = res + string(v)
		if i > 0 && (i+1)%numberString == 0 {
			doc = append(doc, res)
			res = ""
		}
	}
	doc = append(doc, c...)
	return doc
}

// Get the first result xoso and parse it to struct BoxXoso
func UpdateData(currentUrl string) (BoxXoso.BoxXoso, error) {
	var boxXoso BoxXoso.BoxXoso
	doc, err := goquery.NewDocument(currentUrl)
	if err != nil {
		log.Fatal(err)
	}
	var xoso BoxXoso.Xoso
	Day := doc.Find("#boxKQ_0 div:nth-child(2) span:not([class])").Text() //Correct way
	xoso.Day = Day
	DB := doc.Find("#boxKQ_0 .kqxs div:nth-child(1) .no-detail p").Text() //Correct way
	xoso.Kqxs.DB = DB
	Nhat := doc.Find("#boxKQ_0 .kqxs div:nth-child(2) .no-detail p").Text() //Correct way
	xoso.Kqxs.Nhat = Nhat
	doc.Find("#boxKQ_0 .kqxs div:nth-child(3) .no-detail .col-6").Each(func(i int, row *goquery.Selection) {
		xoso.Kqxs.Nhi = append(xoso.Kqxs.Nhi, row.Find("p").Text())
	})
	doc.Find("#boxKQ_0 .kqxs div:nth-child(4) .no-detail .col-4").Each(func(i int, row *goquery.Selection) {
		xoso.Kqxs.Ba = append(xoso.Kqxs.Ba, row.Find("p").Text())
	})
	doc.Find("#boxKQ_0 .kqxs div:nth-child(5) .no-detail .col-3").Each(func(i int, row *goquery.Selection) {
		xoso.Kqxs.Bon = append(xoso.Kqxs.Bon, row.Find("p").Text())
	})
	doc.Find("#boxKQ_0 .kqxs div:nth-child(6) .no-detail .col-4").Each(func(i int, row *goquery.Selection) {
		xoso.Kqxs.Nam = append(xoso.Kqxs.Nam, row.Find("p").Text())
	})
	doc.Find("#boxKQ_0 .kqxs div:nth-child(7) .no-detail .col-4").Each(func(i int, row *goquery.Selection) {
		xoso.Kqxs.Sau = append(xoso.Kqxs.Sau, row.Find("p").Text())
	})
	doc.Find("#boxKQ_0 .kqxs div:nth-child(8) .no-detail .col-3").Each(func(i int, row *goquery.Selection) {
		xoso.Kqxs.Bay = append(xoso.Kqxs.Bay, row.Find("p").Text())
	})
	boxXoso.Xoso = append(boxXoso.Xoso, xoso)
	return boxXoso, nil
}
