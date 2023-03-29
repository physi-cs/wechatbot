package gtp

import (
	"fmt"
)

// TODO: 改造history为map，以wx_id为键
type History_stack struct {
	History   *[][]string
	Max_boxes int
}


func New_History_stack(history *[][]string,max_boxes int) *History_stack{
	history_stack := History_stack{
		History: history,
		Max_boxes: max_boxes,
	}
	return &history_stack
}

func (h *History_stack) clear() {
	// TODO之前的history空间会自行垃圾回收吗？
	h.History = &[][]string{}
}

func (h *History_stack) check_rounds() error {
	len := h.count()
	if len > h.Max_boxes {
		return New("too much round, context cleared, please try later")
	}
	return nil
}

func (h *History_stack) count() int {
	return len(*h.History)
}

// TODO:使用error代替0/1
type tooMuchRound struct {
	msg string
}

func (e tooMuchRound) Error() string {
	return fmt.Sprintf("msg:%v", e.msg)
}

func New(msg string) error {
	return tooMuchRound{
		msg: msg,
	}
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