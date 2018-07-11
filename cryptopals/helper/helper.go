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
