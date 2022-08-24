package writefile

import (
	Handle "XosoBot/Handle"
	BoxXoso "XosoBot/Structure"
	"encoding/json"
	"io/ioutil"
)

var nameFile = "xoso_data.json"

func WriteFile(boxXoso BoxXoso.BoxXoso) {
	xosoJson, err := json.Marshal(boxXoso) //Parse JSON
	Handle.CheckError(err)
	err = ioutil.WriteFile(nameFile, xosoJson, 0644) //WriteFile
	Handle.CheckError(err)
}
