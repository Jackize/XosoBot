package readfile

import (
	BoxXoso "XosoBot/Structure"
	"encoding/json"
	"io/ioutil"
)

var nameFile = "xoso_data.json"
var URL = "http://api.xosodo.vn/webview/soicaudetail?loto=90&p1=7&p2=19&limit=2&provinceId=1&dayPrize=08/11/2018&type=3"

// Read file in nameFile and return it data parse to struct BoxXoso
func ReadFile(boxXoso BoxXoso.BoxXoso) BoxXoso.BoxXoso {
	file, _ := ioutil.ReadFile(nameFile)
	_ = json.Unmarshal([]byte(file), &boxXoso)
	return boxXoso
}
