package internal

import (
	"errors"
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
		BusyPoll:    false,
	}
}

func (q *QzipCommand) BuildQzipCommand() *exec.Cmd {
	if q.InputFile == "" {
		return nil
	}
	// =============================
	// 如果不是压缩，那么需要设置解压选项
	if !q.Compression {
		q.SetDecompression()
	}
	// 如果勾选了保留源文件，那么需要设置保留源文件选项
	q.SetKeepSource()
	// 目录递归：需要传入一个目录并且勾选了递归操作   作用： 单独为目录下全部文件生成压缩包
	q.SetRecursive()
	// 忙轮询，通常用于处理并发的压缩或解压缩请求
	q.SetBusyPoll()
	// 设置压缩算法 一般使用默认
	q.SetAlgorithm()
	// 设置文件头 一般使用默认
	q.SetFileHeader()
	// 设置压缩级别
	q.SetLevel()
	// 设置输出文件名称 目录则不需要
	q.SetOutputFile()

	q.Options = append(q.Options, q.InputFile)
	return exec.Command("qzip", q.Options...)
}

func ExecuteQzipCommand(cmd QzipCommand) error {
	qzipCmd := cmd.BuildQzipCommand()
	if condition := qzipCmd == nil; condition {
		return errors.New("qzip command is nil")
	}
	fmt.Println("Executing command:", qzipCmd.String())
	output, err := qzipCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing command: %s, output: %s", err, output)
	}
	fmt.Printf("Output: %s\n", output)
	return nil
}
