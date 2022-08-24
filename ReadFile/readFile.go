package readfile

import (
	Crawler "XosoBot/CrawData"
	Handle "XosoBot/Handle"
	BoxXoso "XosoBot/Structure"
	WriteFile "XosoBot/WriteFile"
	"encoding/json"
	"io/ioutil"
)

var nameFile = "xoso_data.json"
var URL = "http://api.xosodo.vn/webview/soicaudetail?loto=90&p1=7&p2=19&limit=2&provinceId=1&dayPrize=08/11/2018&type=3"

// Read file in nameFile and return it data parse to struct BoxXoso
func ReadFile() BoxXoso.BoxXoso {
	boxXoso := BoxXoso.BoxXoso{}
	file, _ := ioutil.ReadFile(nameFile)
	_ = json.Unmarshal([]byte(file), &boxXoso)
	updateXoso, err := Crawler.UpdateData(URL)
	Handle.CheckError(err)
	if boxXoso.Xoso[0].Day != updateXoso.Xoso[0].Day {
		boxXoso, err = Crawler.GetAllXoso(boxXoso, updateXoso)
		Handle.CheckError(err)
		WriteFile.WriteFile(boxXoso)
	}
	return boxXoso
}
