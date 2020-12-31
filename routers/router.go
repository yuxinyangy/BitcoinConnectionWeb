package routers

import (
	"BitcoinConnectionWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//主界面
    beego.Router("/", &controllers.MainController{})
    //注册界面
    beego.Router("/register.html",&controllers.RegisterController{})
    //注册完成转登入
    beego.Router("/user_register",&controllers.RegisterController{})
    //登入
	beego.Router("/user_login", &controllers.LoginController{})
    //读取左侧菜单栏
    beego.Router("lefthome.html",&controllers.LeftHomeController{})


    //查询最新区块的hash
    beego.Router("GetBestBlockHash.html",&controllers.GetBestBlockHashController{})

    //查询指定高度的hash值
    //高度输入
    beego.Router("GetBlockHashParameter.html",&controllers.GetBlockHashController{})
    //查询
    beego.Router("/height_hash",&controllers.GetBlockHashController{})

    //获取区块的总数
    beego.Router("GetBlockCount.html",&controllers.GetBlockCountController{})

    //获取当前节点所在链的信息
    beego.Router("GetBlockChainInfo.html",&controllers.GetBlockChainInfoController{})

    //根据区块的hash值获取区块的信息
    //hash值输入
    beego.Router("GetBlockParameter.html",&controllers.GetBlockController{})
    //查询
    beego.Router("/hash_block",&controllers.GetBlockController{})

    //生成一个新的比特币的地址
    //标签输入
    beego.Router("GetNewAddressParameter.html",&controllers.GetNewAddressController{})
    //生成地址
    beego.Router("/label_address",&controllers.GetNewAddressController{})

    //获取指定哈希区块头
    //hash值输入
    beego.Router("GetBlockHeaderParameter.html",&controllers.GetBlockHeaderController{})
    //查询
    beego.Router("/hash_blockHeader",&controllers.GetBlockHeaderController{})

    //获取区块难度
    beego.Router("GetDifficulty.html",&controllers.GetDifficultyController{})

    //获取交易池信息
    beego.Router("GetMemPoolInfo.html",&controllers.GetMemPoolInfoController{})

    //获取交易输出集信息
    beego.Router("GetTxOutSetInfo.html",&controllers.GetTxOutSetInfoController{})

    //校验本地区块
    beego.Router("VerifyChain.html",&controllers.VerifyChainController{})

    //获取正常运行时间(秒)
    beego.Router("UpTime.html",&controllers.UpTimeController{})

    //获取RPC服务器信息
    beego.Router("GetRpcInfo.html",&controllers.GetRpcInfoController{})

    //获取挖矿信息
    beego.Router("GetMiningInfo.html",&controllers.GetMiningInfoController{})

    //获取全网哈希生成速率
    //输入
    beego.Router("GetNetworkHashPSParameter.html",&controllers.GetNetworkHashPsController{})
    //查询
    beego.Router("/blocks_height",&controllers.GetNetworkHashPsController{})

    //获取内存利用信息
    beego.Router("GetMemoryInfo.html",&controllers.GetMemoryInfoController{})

    //验证地址有效性
    //输入
    beego.Router("VaLiDateAddressParameter.html",&controllers.VaLiDateAddressController{})
    //查询
    beego.Router("/data_address",&controllers.VaLiDateAddressController{})
}
