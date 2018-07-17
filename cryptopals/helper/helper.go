package helper

import b64 "encoding/base64"
import "encoding/hex"
import "log"
import "sort"

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

func SingleXORCipher(str1 string) string {
    sum := 0
	resultStr := ""
	key := make(map[int]map[int]string)
	keysList := make([]int, 255)
    decodedStr1, _ := hex.DecodeString(str1)
    for i := 0; i < 256; i++ {
		var decodedResultStr []byte
		for j := 0; j < len(decodedStr1); j++ {
			result := int(decodedStr1[j]) ^ i
			resultStr = string(result)
            // poor o(n^2) TODO refactor bytes.WriteString maybe
			decodedResultStr = append(decodedResultStr, byte(result))
			if val, ok := distanceCalc[resultStr]; ok {
				sum += val
			} else {
				sum -= 10
			}
		}
		key[sum] = map[int]string{i: string(decodedResultStr)}
		sum = 0
	}
	i := 0
	for k := range key {
		keysList[i] = k
		i++
	}
	sort.Ints(keysList)
	xorKey := key[keysList[len(keysList)-1]]
    var returnKey int
	for k := range xorKey {
        returnKey = k
    }
    return xorKey[returnKey]
}
