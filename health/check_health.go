package operate

import (
	"errors"
	"log"

	"github.com/eatmoreapple/openwechat"
)

type Ret int

const (
	ticketError         Ret = -14  // ticket error
	logicError          Ret = -2   // logic error
	sysError            Ret = -1   // sys error
	paramError          Ret = 1    // param error
	failedLoginWarn     Ret = 1100 // failed login warn
	failedLoginCheck    Ret = 1101 // failed login check
	cookieInvalid       Ret = 1102 // cookie invalid
	loginEnvAbnormality Ret = 1203 // login environmental abnormality
	optTooOften         Ret = 1205 // operate too often
)

// SyncCheck错误处理函数
// 心跳 retcode!=0 就会触发这个回调
// 返回false将退出bot
func SyncCheckErrHandler(bot *openwechat.Bot) func(error) bool {
	return func(err error) bool {
		var ret Ret
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
