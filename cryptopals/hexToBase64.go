package main

import (
	"github.com/projjol/solutions/cryptopals/helper"
)

func HexToBase64(hexString string) string {
	stringValue := helper.HexToString(hexString)
	base64Value := helper.Base64Converter(stringValue)
	return base64Value
}
