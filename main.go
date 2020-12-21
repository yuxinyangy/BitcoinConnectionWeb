package main

import (
	"BitcoinConnectionWeb/getbitcion"
	_ "BitcoinConnectionWeb/routers"
	"fmt"
)

func main() {
	//beego.Run()
	fmt.Println("最新区间的hash:",getbitcion.GetBestBlockHash())
	fmt.Println("指定区块的hash值:",getbitcion.GetBlockHash(0))
	fmt.Println("区块的总数:",getbitcion.GetBlockCount())
	fmt.Println("获取当前节点所在链的信息")
	fmt.Println("区块的总数:",getbitcion.GetBlockChainInfo().Blocks)
	fmt.Println("根据区块的hash值获取区块的信息")
	fmt.Println("区块的hash值:",getbitcion.GetBlock(getbitcion.Hash0).Hash)
	address := getbitcion.GetNewAddress("y",getbitcion.LEGACY)
	fmt.Println("新生成的地址:",address)
	fmt.Println("获取指定哈希区块头")
	fmt.Println("区块的hash值:",getbitcion.GetBlockHeader(getbitcion.Hash0).Hash)
	fmt.Println("区块的难度:",getbitcion.GetDifficulty())
	fmt.Println("获取交易池信息")
	fmt.Println("交易池大小:",getbitcion.GetMemPoolInfo().Size)
	fmt.Println("获取交易输出集信息")
	fmt.Println("最新区块hash值:",getbitcion.GetTxOutSetInfo().Bestblock)
	fmt.Println("校验本地区块:",getbitcion.VerifyChain())
	fmt.Println("获取RPC服务器信息")
	fmt.Println("方法:",getbitcion.GetRpcInfo().Active_commands[0].Method)
	fmt.Println("获取正常运行时间为:",getbitcion.UpTime(),"秒")
	fmt.Println("获取挖矿信息")
	fmt.Println("当前节点的区块数:",getbitcion.GetMiningInfo().Blocks)
	fmt.Println("全网哈希生成速率为:",getbitcion.GetNetworkHashPS(getbitcion.N_Block,getbitcion.GetMiningInfo().Blocks))
	fmt.Println("获取内存利用信息")
	fmt.Println("已使用字节数:",getbitcion.GetMemoryInfo().Locked.Used)
	fmt.Println("验证地址有效性")
	fmt.Println("地址是否有效:",getbitcion.VaLiDateAddress(address).Isvalid)
}

