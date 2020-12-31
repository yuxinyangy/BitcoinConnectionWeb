package getbitcion

import (
	"BitcoinConnectionWeb/entity"
	"BitcoinConnectionWeb/utils"
	"encoding/json"
	"fmt"
)
//1、最新区块的hash
func GetBestBlockHash() string{
	reqBytes := JsonString(GETBESTBLOCKHASH, RPCVERSION,nil)
	body := DoPost(RPCURL,reqBytes)
	RPCResult := utils.GetResult(body)
	return RPCResult.Result.(string)
}

//2、获取指定区块的hash
func GetBlockHash(height int) string {
	reqBytes := JsonString(GETBLOCKHASH, RPCVERSION,[]interface{}{height})
	body := DoPost(RPCURL,reqBytes)
	RPCResult := utils.GetResult(body)
	return RPCResult.Result.(string)
}
//3、获取区块的总数
func GetBlockCount() float64 {
	reqBytes := JsonString(GETBLOCKCOUNT, RPCVERSION,nil)
	body := DoPost(RPCURL,reqBytes)
	RPCResult := utils.GetResult(body)
	return RPCResult.Result.(float64)
}

//4、获取当前节点所在链的信息
func GetBlockChainInfo() *entity.BlockChainInfo {
	reqBytes := JsonString(GETBLOCKCHAININFO, RPCVERSION,nil)
	body := DoPost(RPCURL,reqBytes)
	RPCResult := utils.GetResult(body)
	jsonstr,err :=json.Marshal(RPCResult.Result)
	if err != nil {
		return nil
	}
	BlockChain := new(entity.BlockChainInfo)
	err = json.Unmarshal(jsonstr,BlockChain)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return BlockChain
}

//5、根据区块的hash值获取区块的信息
func GetBlock(hashcode string) *entity.Block {
	reqBytes := JsonString(GETBLOCK, RPCVERSION,[]interface{}{hashcode})
	body := DoPost(RPCURL,reqBytes)
	RPCResult := utils.GetResult(body)
	jsonstr,err :=json.Marshal(RPCResult.Result)
	if err != nil {
		return nil
	}
	Block := new(entity.Block)
	err = json.Unmarshal(jsonstr,Block)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return Block
}

//6、生成一个新的比特币的地址
func GetNewAddress(label string,str string) string {
	reqBytes := JsonString(GETNEWADDRESS, RPCVERSION,[]interface{}{label,str})
	body := DoPost(RPCURL,reqBytes)
	RPCResult := utils.GetResult(body)
	return RPCResult.Result.(string)
}

//7、获取指定哈希区块头
func GetBlockHeader(hashcode string) *entity.BlockHeader {
	reqBytes := JsonString(GETBLOCKHEADER, RPCVERSION,[]interface{}{hashcode})
	body := DoPost(RPCURL,reqBytes)
	RPCResult := utils.GetResult(body)
	jsonstr,err :=json.Marshal(RPCResult.Result)
	if err != nil {
		return nil
	}
	BlockHeader := new(entity.BlockHeader)
	err = json.Unmarshal(jsonstr,BlockHeader)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return BlockHeader
}

//8、获取区块难度
func GetDifficulty() float64 {
	reqBytes := JsonString(GETDIFFICULT, RPCVERSION,nil)
	body := DoPost(RPCURL,reqBytes)
	RPCResult := utils.GetResult(body)
	return RPCResult.Result.(float64)
}

//9、获取交易池信息
func GetMemPoolInfo() *entity.MemPoolInfo {
	reqBytes := JsonString(GETMEMPOOLINFO, RPCVERSION,nil)
	body := DoPost(RPCURL,reqBytes)
	RPCResult := utils.GetResult(body)
	jsonstr,err :=json.Marshal(RPCResult.Result)
	if err != nil {
		return nil
	}
	MemPoolInfo := new(entity.MemPoolInfo)
	err = json.Unmarshal(jsonstr,MemPoolInfo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return MemPoolInfo
}

//10、获取交易输出集信息
func GetTxOutSetInfo() *entity.TxOutSetInfo {
	reqBytes := JsonString(GETTXOUTSETINFO, RPCVERSION,nil)
	body := DoPost(RPCURL,reqBytes)
	RPCResult := utils.GetResult(body)
	jsonstr,err :=json.Marshal(RPCResult.Result)
	if err != nil {
		return nil
	}
	TxOutSetInfo := new(entity.TxOutSetInfo)
	err = json.Unmarshal(jsonstr,TxOutSetInfo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return TxOutSetInfo
}

//11、校验本地区块
func VerifyChain() bool {
	reqBytes := JsonString(VERIFYCHAIN, RPCVERSION,nil)
	body := DoPost(RPCURL,reqBytes)
	RPCResult := utils.GetResult(body)
	return RPCResult.Result.(bool)
}

//12、获取RPC服务器信息
func GetRpcInfo() *entity.RpcInfo {
	reqBytes := JsonString(GETRPCINFO, RPCVERSION,nil)
	body := DoPost(RPCURL,reqBytes)
	RPCResult := utils.GetResult(body)
	jsonstr,err :=json.Marshal(RPCResult.Result)
	if err != nil {
		return nil
	}
	RpcInfo := new(entity.RpcInfo)
	err = json.Unmarshal(jsonstr,RpcInfo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return RpcInfo
}

//13、获取正常运行时间(秒)
func UpTime() float64 {
	reqBytes := JsonString(UPTIME, RPCVERSION,nil)
	body := DoPost(RPCURL,reqBytes)
	RPCResult := utils.GetResult(body)
	return RPCResult.Result.(float64)
}

//14、获取挖矿信息
func GetMiningInfo() *entity.MiningInfo {
	reqBytes := JsonString(GETMININGINFO, RPCVERSION,nil)
	body := DoPost(RPCURL,reqBytes)
	RPCResult := utils.GetResult(body)
	jsonstr,err :=json.Marshal(RPCResult.Result)
	if err != nil {
		return nil
	}
	MiningInfo := new(entity.MiningInfo)
	err = json.Unmarshal(jsonstr,MiningInfo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return MiningInfo
}

//15、获取全网哈希生成速率
func GetNetworkHashPS(nBlocks int, height int) float64 {
	reqBytes := JsonString(GETNETWORKHASHPS, RPCVERSION,[]interface{}{nBlocks,height})
	body := DoPost(RPCURL,reqBytes)
	RPCResult := utils.GetResult(body)
	return RPCResult.Result.(float64)
}

//16、获取内存利用信息
func GetMemoryInfo() *entity.MemoryInfo {
	reqBytes := JsonString(GETMEMORYINFO, RPCVERSION,nil)
	body := DoPost(RPCURL,reqBytes)
	RPCResult := utils.GetResult(body)
	jsonstr,err :=json.Marshal(RPCResult.Result)
	if err != nil {
		return nil
	}
	MemoryInfo := new(entity.MemoryInfo)
	err = json.Unmarshal(jsonstr,MemoryInfo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return MemoryInfo
}

//17、验证地址有效性
func VaLiDateAddress(address string) *entity.DateAddress {
	reqBytes := JsonString(VALIDATEADDRESS, RPCVERSION,[]interface{}{address})
	body := DoPost(RPCURL,reqBytes)
	RPCResult := utils.GetResult(body)
	jsonstr,err :=json.Marshal(RPCResult.Result)
	if err != nil {
		return nil
	}
	DateAddress := new(entity.DateAddress)
	err = json.Unmarshal(jsonstr,DateAddress)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return DateAddress
}



