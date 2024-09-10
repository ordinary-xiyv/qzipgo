package internal

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type (
	QzipCommand struct {
		IsDirctory bool
		//  输入文件
		InputFile []string
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
		// -r 并发数
		Concurrency int
		// 其他单独选项
		Options []string // 用于存储其他选项
	}

	TarCommand struct {
		//  输入文件
		InputFile []string
		// -o 输出文件
		OutputFile string
		// 归档文件
		ArchiveFile string
		// 是否压缩
		Compression bool
		// 目录层级
		components int
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
		OutputFile:  "",
		InputFile:   nil,
		Concurrency: 10,
	}
}

// 获取默认的tar命令
func GetDefaultTarCommand() TarCommand {
	return TarCommand{
		Compression: true,
		InputFile:   nil,
		OutputFile:  "",
		ArchiveFile: "",
		components:  0,
		Options:     nil,
	}
}

func (q *QzipCommand) BuildQzipCommand() *exec.Cmd {
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
	// 设置并发数
	q.SetConcurrency()
	q.Options = append(q.Options, q.InputFile...)
	return exec.Command("qzip", q.Options...)
}

func ExecuteQzipCommand(cmd QzipCommand) error {
	// 判断传入的文件是否存在
	if condition := len(cmd.InputFile) != 0; condition {
		for _, file := range cmd.InputFile {
			if _, err := os.Stat(file); os.IsNotExist(err) {
				return fmt.Errorf("input file %s does not exist", file)
			}
		}
	}
	// 执行qzip命令
	qzipCmd := cmd.BuildQzipCommand()
	if condition := qzipCmd == nil; condition {
		return errors.New("qzip command or inputFile is nil")
	}
	fmt.Println("Executing command:", qzipCmd.String())
	output, err := qzipCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing command: %s, output: %s", err, output)
	}
	fmt.Printf("Output: %s\n", output)
	return nil
}
func (t *TarCommand) BuildTarCommand() *exec.Cmd {
	// 压缩必须有输入文件
	if len(t.InputFile) == 0 && t.Compression {
		return nil
	}
	// 压缩：-cvf     解压缩 ：-xvf
	t.SetCompressionType()
	// 设置归档文件，必须在-f选项之后
	t.SetArchiveFile()
	t.SetQzipCommand()
	t.SetOutputFile()
	t.SetInputFile()
	t.SetComponents()
	return exec.Command("tar", t.Options...)
}

func ExecuteTarCommand(cmd TarCommand) error {
	// 判断传入的文件是否存在
	// 压缩：输入文件
	// 解压缩：归档文件（输入文件）
	if cmd.Compression {
		if len(cmd.InputFile) != 0 {
			for _, f := range cmd.InputFile {
				if !fileIsExist(f) {
					return fmt.Errorf("file %s does not exist", f)
				}
			}
		} else {
			return errors.New("input file is empty")
		}
	} else {
		if !fileIsExist(cmd.ArchiveFile) {
			return fmt.Errorf("file %s does not exist", cmd.ArchiveFile)
		}
	}
	// ***********************************************************
	// 这段操作是为了保证压缩的目录层级只有一层，如：/tmp/tt/test.txt 解压会得到/path-you-want/tt/test.txt，而不是/path-you-want/tmp/tt/test.txt
	// 因为文件或者目录的层级不定，所以需要进入父目录使用相对路径进行一级层级压缩
	// 记录当前工作目录
	originalDir, err := os.Getwd()
	if err != nil {
		return errors.New("error getting current directory")
	}
	//  提取目录层级：/tmp/test.txt --> /tmp 或/tmp/tt/test.txt --> /tmp/tt
	dataFatherPath := ""
	if cmd.InputFile != nil {
		// 压缩从输入文件中提取目录
		dataFatherPath = filepath.Dir(cmd.InputFile[0])
		// 将输入文件去除父目录，使用相对路径/tmp/tt/test.txt --> test.txt
		modifiedPaths := make([]string, len(cmd.InputFile))
		for i, path := range cmd.InputFile {
			modifiedPaths[i] = filepath.Base(path)
		}
		// 如果需要，可以将 modifiedPaths 赋值回 cmd.InputFile
		cmd.InputFile = modifiedPaths
		fmt.Println("Data father path:", dataFatherPath)
	} else if cmd.ArchiveFile != "" && !cmd.Compression {
		// 解压从归档文件中提取，因为归档文件是解压操作的输入
		dataFatherPath = filepath.Dir(cmd.ArchiveFile)
		fmt.Println("Data father path:", dataFatherPath)
	}
	// 切换到 数据的父目录
	err = os.Chdir(dataFatherPath)
	if err != nil {
		return errors.New("error changing directory")
	}
	// 最终切换到原本的工作目录
	defer func() {
		err := os.Chdir(originalDir)
		if err != nil {
			fmt.Println("Error returning to original directory:", err)
		}
	}()
	// ***********************************************************
	tarCmd := cmd.BuildTarCommand()
	if condition := tarCmd == nil; condition {
		return errors.New("tar command or inputFile is nil")
	}
	fmt.Println("Executing command:", tarCmd.String())
	output, err := tarCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing command: %s, output: %s", err, output)
	}
	fmt.Printf("Output: %s\n", output)
	return nil
}

// 检测文件或者目录是否存在
func fileIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
