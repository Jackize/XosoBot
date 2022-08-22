package crawler

import (
	Find "XosoBot/Handle"
	BoxXoso "XosoBot/Structure"
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/sync/errgroup"
)

// Crawl all result in a single page
func GetAllXoso(currentUrl string, boxXoso BoxXoso.BoxXoso) (BoxXoso.BoxXoso, error) {
	eg := errgroup.Group{}
	eg.Go(func() error {
		//Now just only update data beacause already get and write file json before
		if Find.CheckUpdateData() {
			updateBox := BoxXoso.BoxXoso{}
			updateBox, err := UpdateData(currentUrl)
			Find.CheckError(err)
			boxXoso.Xoso = append(updateBox.Xoso, boxXoso.Xoso...)
			fmt.Println("Getting data")
		}
		return nil
	})
	err := eg.Wait()
	Find.CheckError(err)
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
