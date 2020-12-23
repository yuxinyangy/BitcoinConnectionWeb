package getbitcion

import (
	"BitcoinConnectionWeb/entity"
	"encoding/json"
	"fmt"
	"time"
)

//1、准备一个json数据(string)
//json数据:序列化、反序列化
//go:json.Marhxx
func JsonString(methed string,jsonrpc string,params []interface{})[]byte  {
	resquest := entity.RPCRequest{
		Id: time.Now().Unix(),
		Method: methed,
		Jsonrpc: jsonrpc,
		Params: params,
	}
	reqBytes,err := json.Marshal(&resquest)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return reqBytes
}
