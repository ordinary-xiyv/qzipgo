package internal

import (
	"fmt"
	"os/exec"
)

type (
	QzipCommand struct {
		IsDirctory bool
		//  输入文件
		InputFile string
		// -o 输出文件
		OutputFile string
		// -O 压缩文件头格式
		FileHeader FILE_HEADER
		// 是否压缩 否：-d 解压缩
		Compression bool
		// -A 算法
		Algorithm ALGORITHM_TYPE
		// -L 压缩级别
		Level COMPRESSION_LEVEL
		// -k 保留源文件
		KeepSource bool
		// -R 目录递归
		Recursive bool
		// -P 忙轮询
		BusyPoll bool
		// 其他单独选项
		Options []string // 用于存储其他选项
	}

	COMPRESSION_LEVEL int
	ALGORITHM_TYPE    int
	FILE_HEADER       int
)

// 获取默认的qzip命令
func GetDefaultQzipCommand() QzipCommand {
	return QzipCommand{
		IsDirctory:  false,
		Compression: true,
		Algorithm:   0, // 默认：GZIPEXT
		FileHeader:  0, // 默认：不添加参数
		Level:       LEVEL_5,
		KeepSource:  true,
		Recursive:   false,
	}
}

func (q *QzipCommand) BuildQzipCommand() *exec.Cmd {
	args := []string{}

	if !q.Compression {
		args = append(args, "-d") // 默认是压缩，如果是解压，需要-d选项
	}

	args = append(args, q.InputFile)

	if q.OutputFile != "" {
		args = append(args, "-o", q.OutputFile) // 假设有一个输出文件选项
	}

	args = append(args, q.Options...) // 添加其他选项

	return exec.Command("qzip", args...)
}

func ExecuteQzipCommand(cmd QzipCommand) error {
	qzipCmd := cmd.BuildQzipCommand()
	output, err := qzipCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing command: %s, output: %s", err, output)
	}
	fmt.Printf("Output: %s\n", output)
	return nil
}
