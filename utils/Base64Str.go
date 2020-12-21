package utils

import "encoding/base64"

func Base64EncodeString(str string)string  {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
