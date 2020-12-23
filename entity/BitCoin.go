package entity

type Bip struct {
	Type string `json:"type"`
	Active bool `json:"active"`
	Height int `json:"height"`
}

type Block struct {
	Hash string `json:"hash"`
	Confirmations int `json:"confirmations"`
	Srippedsize int `json:"strippedsize"`
	Size int `json:"size"`
	Weight int `json:"weight"`
	Version int `json:"version"`
	VersionHex string `json:"version_hex"`
	Merkleroot string `json:"merkleroot"`
	Height int `json:"height"`
	Time int `json:"time"`
	Mediantime int `json:"mediantime"`
	Nonce int `json:"nonce"`
	Bits string `json:"bits"`
	Difficulty int `json:"difficulty"`
	Chainwork string `json:"chainwork"`
	NTx int `json:"nTx"`
}

type BlockChainInfo struct {
	Headers int `json:"headers"`
	Chain string `json:"chain"`
	Sizeondisk  int `json:"size_on_disk"`
	Mediantime int `json:"mediantime"`
	Blocks int `json:"blocks"`
	Pruned bool `json:"pruned"`
	Warnings string `json:"warnings"`
	Difficulty int `json:"difficulty"`
	Bestblockhash string `json:"bestblockhash"`
	Chainwork string `json:"chainwork"`
	Pruneheight int `json:"pruneheight"`
	Automaticpruning bool `json:"automatic_pruning"`
	Prunetargetsize int `json:"prune_target_size"`
	Softforks SoftForks `json:"softforks"`
}

type RPCRequest struct {
	Id int64 `json:"id"`
	Method string `json:"method"`
	Jsonrpc string `json:"jsonrpc"`
	Params []interface{} `json:"params"`
}

type RPCResult struct {
	Result interface{} `json:"result"`
	Error string `json:"error"`
	Id int64 `json:"id"`
}

type SoftForks struct {
	Bip34 Bip `json:"bip_34"`
	Bip66 Bip `json:"bip_66"`
	Bip65 Bip `json:"bip_65"`
	Csv   Bip `json:"csv"`
	Segwit Bip `json:"segwit"`
}

type BlockHeader struct {
	Hash string `json:"hash"`
	Confirmations int `json:"confirmations"`
	Height int `json:"height"`
	Version int `json:"version"`
	VersionHex string `json:"versionHex"`
	Merkleroot string `json:"merkleroot"`
	Time int `json:"time"`
	Mediantime int `json:"mediantime"`
	Nonce int `json:"nonce"`
	Bits string `json:"bits"`
	Difficulty int `json:"difficulty"`
	Chainwork string `json:"chainwork"`
	NTx int `json:"nTx"`
}

type MemPoolInfo struct {
	Loaded bool `json:"loaded"`
	Size int `json:"size"`
	Bytes int `json:"bytes"`
	Usage int `json:"usage"`
	Maxmempool int `json:"maxmempool"`
	Mempoolminfee float32 `json:"mempoolminfee"`
	Minrelaytxfee float32 `json:"minrelaytxfee"`
}

type TxOutSetInfo struct {
	Height int `json:"height"`
	Bestblock string `json:"bestblock"`
	Transactions int `json:"transactions"`
	Txouts int `json:"txouts"`
	Bogosize int `json:"bogosize"`
	Hash_serialized_2 string `json:"hash_serialized_2"`
	Disk_size int `json:"disk_size"`
	Total_amount float32 `json:"total_amount"`
}

type RpcInfo struct {
	Active_commands []Commands`json:"active_commands"`
	Logpath string `json:"logpath"`
}

type Commands struct {
	Method string `json:"method"`
	Duration int `json:"duration"`
}

type MiningInfo struct {
	Blocks int `json:"blocks"`
	Difficulty int `json:"difficulty"`
	Networkhashps float32 `json:"networkhashps"`
	Pooledtx int `json:"pooledtx"`
	Chain string `json:"chain"`
	Warnings string `json:"warnings"`
}

type MemoryInfo struct {
	Locked Locked `json:"locked"`
}

type Locked struct {
	Used int `json:"used"`
	Free int `json:"free"`
	ToTal int `json:"to_tal"`
	Locked int `json:"locked"`
	Chunks_used int `json:"chunks_used"`
	Chunks_free int `json:"chunks_free"`
}

type DateAddress struct {
	Isvalid  bool `json:"isvalid"`
	Address string `json:"address"`
	ScriptPubKey string `json:"script_pub_key"`
	Isscript bool `json:"isscript"`
	Iswitness bool `json:"iswitness"`
}