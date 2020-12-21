package getbitcion

/**
 *RPC通信的配置
 */
const RPCURL = "http://127.0.0.1:8332"
const RPCUSER = "user"
const RPCPASSWORD = "pwd"
const RPCVERSION = "2.0"
const Hash0 = "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f"

/**
 *比特币节点的命令
 */
const GETBLOCKCOUNT = "getblockcount"
const GETBESTBLOCKHASH = "getbestblockhash"
const GETBLOCK = "getblock"
const GETBLOCKHASH = "getblockhash"
const GETBLOCKCHAININFO = "getblockchaininfo"
const GETNEWADDRESS = "getnewaddress"
const GETBLOCKHEADER = "getblockheader"
const GETDIFFICULT = "getdifficulty"
const GETMEMPOOLINFO = "getmempoolinfo"
const GETTXOUTSETINFO = "gettxoutsetinfo"
const VERIFYCHAIN = "verifychain"
const GETRPCINFO = "getrpcinfo"
const UPTIME = "uptime"
const GETMININGINFO = "getmininginfo"
const GETNETWORKHASHPS = "getnetworkhashps"
const GETMEMORYINFO = "getmemoryinfo"
const VALIDATEADDRESS = "validateaddress"

/**
 * 其他
 */
const LEGACY = "legacy"
const P2SH_SEGWIT = "p2sh-segwit"
const BECH32 = "bench32"
const N_Block = -1