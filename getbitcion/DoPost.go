package getbitcion

import (
	"BitcoinConnectionWeb/utils"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func DoPost(url string,reqBytes []byte)[]byte {
	client := &http.Client{}
	request,err := http.NewRequest("POST",url,bytes.NewBuffer(reqBytes))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	//设置请求头
	request.Header.Add("Encoding","UTF-8")
	request.Header.Add("Content-Type","application/json")
	request.Header.Add("Authorization", "Basic " + utils.Base64EncodeString(RPCUSER+ ":" +RPCPASSWORD))


	//java:client.execute
	response,err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	body,err := ioutil.ReadAll(response.Body)
	//code := response.StatusCode
	//if code == 200{
	//	fmt.Println("请求成功")
	//}else {
	//	fmt.Println("请求失败")
	//}
	return body
}
