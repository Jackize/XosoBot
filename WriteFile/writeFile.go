package writefile

import (
	Crawler "XosoBot/CrawData"
	ReadFile "XosoBot/ReadFile"
	BoxXoso "XosoBot/Structure"
	"encoding/json"
	"io/ioutil"
	"log"
)

var URL = "http://api.xosodo.vn/webview/soicaudetail?loto=90&p1=7&p2=19&limit=2&provinceId=1&dayPrize=08/11/2018&type=3"

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var nameFile = "xoso_data.json"

func WriteFile() {
	boxXoso := BoxXoso.BoxXoso{}
	boxXoso = ReadFile.ReadFile(boxXoso)             //Lấy dữ liệu từ file cũ
	boxXoso, err := Crawler.GetAllXoso(URL, boxXoso) //Lấy dữ liệu từ api
	checkError(err)
	xosoJson, err := json.Marshal(boxXoso) //Parse JSON
	checkError(err)
	err = ioutil.WriteFile(nameFile, xosoJson, 0644) //WriteFile
	checkError(err)
}
