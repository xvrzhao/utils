package errors

import (
	"io"
	"log"
	"os"
)

var (
	defaultLogger *log.Logger
	logger        *log.Logger
)

// Log 打印错误，配合通过 Wrapper 函数返回的 error 使用。
// 若未通过 SetLoggerWriter 设置过 writer，则默认打印到 stderr。
// 输出的格式为：
//   2020/05/31 16:38:09 an error was caught at runtime:
//   	@ [errors.handler]: can not get user by id (utils/errors/wrapper_test.go:30)
//   	@ [errors.UserModel.GetUserByID]: query fail (utils/errors/wrapper_test.go:39)
//   	@ [origin]: can not connect to db
func Log(err error) {
	if logger == nil {
		defaultLogger.Println(err)
		return
	}
	logger.Println(err)
}

// SetLoggerWriter 改变 Log 函数的输出位置。
func SetLoggerWriter(out io.Writer) {
	logger = newLogger(out)
}

func newLogger(out io.Writer) *log.Logger {
	return log.New(out, "an error was caught at runtime:\n", log.LstdFlags|log.Lmsgprefix)
}

func init() {
	defaultLogger = newLogger(os.Stderr)
}
