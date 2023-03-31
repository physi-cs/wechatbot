package glm

import (
	"fmt"
)

// TODO: 改造history为map，以wx_id为键
type History_stack struct {
	Message_sender string
	History        *[][]string
	Max_boxes      int
}

// history 全局变量
/*
var HISTORY_STACK = &[][]string{{"", ""}}

func History_init() {
	*HISTORY_STACK = (*HISTORY_STACK)[:0]
}
*/
var Max_boxes = 5

func New_History_stack(sender string, history *[][]string, max_boxes int) *History_stack {
	history_stack := History_stack{
		Message_sender: sender,
		History:        history,
		Max_boxes:      max_boxes,
	}
	return &history_stack
}

func (h *History_stack) clear() {
	*h.History = (*h.History)[:0]
}

func (h *History_stack) check_rounds() error {
	len := h.count()
	if len > h.Max_boxes {
		// 删除历史
		h.clear()
		return NewError("轮次过多，已清空上下文。请重新提问。")
	}
	return nil
}

func (h *History_stack) count() int {
	return len(*h.History)
}

var History_stack_slice []*History_stack

// GetHistoryStack 函数根据 Message_sender 查找历史栈，如果不存在则创建一个新的，并返回 History_stack 对象
func GetHistoryStack(sender string) *History_stack {
    for _, stack := range History_stack_slice {
        if stack.Message_sender == sender {
            return stack
        }
    }

    // 如果历史栈不存在，则创建一个新的
    historyStack := New_History_stack(sender, &[][]string{}, Max_boxes)
    History_stack_slice = append(History_stack_slice, historyStack)
    return historyStack
}

type TooMuchRound struct {
	msg string
}

func (e TooMuchRound) Error() string {
	return fmt.Sprintf("msg:%v", e.msg)
}

func NewError(msg string) error {
	return TooMuchRound{
		msg: msg,
	}
}

func (e TooMuchRound) GetMessage() string {
	return e.msg
}

// // TODO 现以json文件存储对话历史数组，可以考虑替换为 ThreadLocal 类

// // 默认以json文件的形式存储
// // JsonFileHistoryStorage 实现 io.ReadWriter
// type JsonFileHistoryStorage struct {
// 	FileName string
// 	file     *os.File
// }

// func (j *JsonFileHistoryStorage) Read(p []byte) (n int, err error) {
// 	if j.file == nil {
// 		j.file, err = os.Open(j.FileName)
// 		if err != nil {
// 			return 0, err
// 		}
// 	}
// 	return j.file.Read(p)
// }

// func (j *JsonFileHistoryStorage) Write(p []byte) (n int, err error) {
// 	if j.file == nil {
// 		j.file, err = os.Create(j.FileName)
// 		if err != nil {
// 			return 0, err
// 		}
// 	}
// 	return j.file.Write(p)
// }

// // 根据文件路径创建存储文件
// func NewJsonFileHistoryStorage(filename string) io.ReadWriter {
// 	return &JsonFileHistoryStorage{FileName: filename}
// }

// // var _ HotReloadStorage = (*JsonFileHotReloadStorage)(nil)

// // 从json文件重载
// func ReloadHistoryStorage(storage io.ReadWriter) (*History_stack, error) {
// 	if storage == nil {
// 		return nil, errors.New("storage can't be nil")
// 	}
// 	var history History_stack

// 	if err := json.NewDecoder(storage).Decode(&history); err != nil {
// 		return nil, err
// 	}
// 	return &history, nil
// }
