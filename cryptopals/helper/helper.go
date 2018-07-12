package helper

import b64 "encoding/base64"
import "encoding/hex"
import "log"

func Base64Converter(str string) string {
    base64Str := b64.StdEncoding.EncodeToString([]byte(str))
    return base64Str
}

func HexToString(str string) string {
    decoded, err := hex.DecodeString(str)
    if err != nil {
        log.Fatal(err)
    }
    return string(decoded[:])
}

func XORStrings(str1,str2 string) string{
    decodedStr1, _ := hex.DecodeString(str1)
    decodedStr2, _ := hex.DecodeString(str2)
    resultString := make([]byte, len(decodedStr1))

    for i:= 0; i < len(decodedStr1); i++ {
        resultString[i] += decodedStr1[i] ^ decodedStr2[i]
    }
    return hex.EncodeToString(resultString)
}
