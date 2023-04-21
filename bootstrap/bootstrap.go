package bootstrap

import (
	"log"

	"github.com/869413421/wechatbot/glm"
	"github.com/869413421/wechatbot/handlers"
	operate "github.com/869413421/wechatbot/health"
	"github.com/eatmoreapple/openwechat"
)

func Run() {
	//bot := openwechat.DefaultBot()
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式，上面登录不上的可以尝试切换这种模式
	// 在错误处理函数中添加运维提醒
	bot.MessageErrorHandler = operate.SyncCheckErrHandler(bot)
	// 注册消息处理函数
	bot.MessageHandler = handlers.Handler
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 创建热存储容器对象
	//reloadStorage := openwechat.NewJsonFileHotReloadStorage("storage.json")

	// 执行热登录
	//err := bot.HotLogin(reloadStorage)
	//if err != nil {
	err := bot.Login()
	if err != nil {
		log.Printf("login error: %v \n", err)
		return
	}
	//}
	//异步执行一个定时清空历史消息的函数，周期为24小时
	go glm.ClearHistoryStackSlice()
	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}
