package errors

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

var wrapperOpt WrapperOpt

// WrapperOpt 是 Wrapper 设置项。
type WrapperOpt struct {
	// ProjectDir 需设置为编译时所在环境的项目路径，
	// 该路径会作为前缀抵消堆栈信息里源码文件路径里的该部分，
	// 如果不设置或者值和编译时项目所在的路径不符，
	// 则堆栈信息里的源码文件路径则会展示执行编译时文件所在的绝对路径。
	//
	// 若项目文件夹名为 app，在执行编译时项目的路径为 /path/to/app/，
	// 则需将该值设置为 "/path/to/"。
	//
	// 注意：若使用 Docker 执行多阶段构建，则需将该值为 Docker 容器内项目的绝对路径。
	ProjectDir string
	// ProjectModule 需设置为项目的 Module 名称，即 go.mod 文件第一行指定的项目名称，
	// 目的同样是为了去除 error 堆栈中函数名称前缀。
	//
	// 若项目 module 名称为 example.com/foo/bar,
	// 则需将该值设置为 "example.com/foo/bar/"。
	ProjectModule string
}

// Wrapper 返回一个包装 err 的 wrapper error，该 wrapper error 携带了调用者调用时的源码文件名、行号和产生错误的函数名称。
// 用 Wrapper 函数包裹错误、向上传递、再包裹、再传递 ... 以这种方式来获取错误冒泡过程中每一层的堆栈信息。
func Wrapper(err error, format string, a ...interface{}) error {
	if err == nil {
		return nil
	}

	if errors.Unwrap(err) == nil {
		err = fmt.Errorf("\t@ [origin]: %w", err)
	}

	pc, file, line, ok := runtime.Caller(1)
	if ok {
		funcName := runtime.FuncForPC(pc).Name()
		if strings.HasPrefix(funcName, wrapperOpt.ProjectModule) {
			funcName = strings.TrimPrefix(funcName, wrapperOpt.ProjectModule)
		}
		if strings.HasPrefix(file, wrapperOpt.ProjectDir) {
			file = strings.TrimPrefix(file, wrapperOpt.ProjectDir)
		}
		return fmt.Errorf("\t@ [%s]: %s (%s:%d)\n%w", funcName, fmt.Sprintf(format, a...), file, line, err)
	}

	return fmt.Errorf("\t@ %s\n%w", fmt.Sprintf(format, a...), err)
}

// SetWrapperOpt 根据 opt 改变 Wrapper 函数的行为， 该函数需在调用 Wrapper 函数前执行。
func SetWrapperOpt(opt WrapperOpt) {
	wrapperOpt.ProjectDir, wrapperOpt.ProjectModule = opt.ProjectDir, opt.ProjectModule
}
