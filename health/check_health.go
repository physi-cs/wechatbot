package operate

import (
	"errors"
	"log"

	"github.com/eatmoreapple/openwechat"
)


const (
	ticketError         openwechat.Ret = -14  // ticket error
	logicError          openwechat.Ret = -2   // logic error
	sysError            openwechat.Ret = -1   // sys error
	paramError          openwechat.Ret = 1    // param error
	failedLoginWarn     openwechat.Ret = 1100 // failed login warn
	failedLoginCheck    openwechat.Ret = 1101 // failed login check
	cookieInvalid       openwechat.Ret = 1102 // cookie invalid
	loginEnvAbnormality openwechat.Ret = 1203 // login environmental abnormality
	optTooOften         openwechat.Ret = 1205 // operate too often
)

// SyncCheck错误处理函数
// 心跳 retcode!=0 就会触发这个回调
// 返回false将退出bot
func SyncCheckErrHandler(bot *openwechat.Bot) func(error) bool {
	return func(err error) bool {
		var ret openwechat.Ret
		if errors.As(err, &ret) {
			switch ret {
			case failedLoginCheck, cookieInvalid, failedLoginWarn:
				_ = bot.Logout()
				
				// 添加提醒函数
				create_events(err.Error())
				log.Printf("hw event msg sent")

				return false
			}
		}
		return true
	}
}
