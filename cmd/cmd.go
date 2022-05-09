package cmd

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
)

// Exec 启动一个 shell 子进程执行 cmdStr 命令。成功情况下 stdout 对应标准输出，err
// 为 nil。失败情况下 err 为错误原因（包含 stderr 信息）。
func Exec(ctx context.Context, cmdStr string) (stdout string, err error) {
	cmd := exec.Command("sh", "-c", cmdStr)

	stdoutBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)

	cmd.Stdout = stdoutBuf
	cmd.Stderr = stderrBuf

	if err = cmd.Start(); err != nil {
		err = fmt.Errorf("failed to start command: %w", err)
		return
	}

	done := make(chan error, 1) // 容量设 1，防止下面的协程阻塞
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-ctx.Done(): // 超时情况：上面协程执行完毕后，done 管道不被使用，等待回收
		err = fmt.Errorf("command canceled: %v", ctx.Err())
		return
	case err = <-done:
		if err != nil {
			errInfo := fmt.Sprintf("failed to execute command: %v", err)
			if stderr := stderrBuf.String(); stderr != "" {
				errInfo = fmt.Sprintf("%s (stderr: %s)", errInfo, stderr)
			}
			err = errors.New(errInfo)
			return
		}
		stdout = stdoutBuf.String()
		return
	}
}
