package main

import (
	"github.com/869413421/wechatbot/bootstrap"
)

func main() {
	bootstrap.Run()
}

// func main() {
// 	history := 	bootstrap.HISTORY
// 	history_stack := gtp.New_History_stack(history, 5)
	
// 	_, err := gtp.Completions_with_history("他的妻子是谁", history_stack)
// 	if err != nil {
// 		fmt.Printf("gtp request error: %v \n", err)
// 	}
// 	if len(*history_stack.History) > 0 {
// 		for _, v := range *history_stack.History {
// 			fmt.Printf("now stack history: q%s ,a%s \n", v[0], v[1])
// 			break
// 		}
// 	}
// }