package pkg

import (
	"strings"

	"github.com/ordinary-xiyv/qzipgo/src/internal"
)

// qzip -k filepath 测试压缩
// output:Executing command: /usr/local/bin/qzip -k /tmp/test.txt
func CompressFile(inputFile string) error {
	cmd := internal.GetDefaultQzipCommand()
	cmd.KeepSource = true
	cmd.InputFile = append(cmd.InputFile, inputFile)

	if err := internal.ExecuteQzipCommand(cmd); err != nil {
		return err
	}
	return nil
}

// CompressWithOutputFile compresses the given inputFile and stores the output in the given outputFile,
// keeping the original file intact.
//
// The function takes two arguments: the input file to be compressed and the output file where the compressed
// content should be stored.
//
// The function returns an error if something goes wrong while compressing the file.
//
// qzip -k -o outputFile inputFile
//
// 单文件压缩并指定输出名称
func CompressWithOutputFile(inputFile, outputFile string) error {
	// Create a new QzipCommand with the default configuration
	cmd := internal.GetDefaultQzipCommand()

	// Set the input file, keep the original file and set the output file
	cmd.KeepSource = true
	cmd.Compression = true
	cmd.InputFile = append(cmd.InputFile, inputFile)
	cmd.OutputFile = outputFile

	// Execute the command and return the error if something goes wrong
	if err := internal.ExecuteQzipCommand(cmd); err != nil {
		return err
	}

	// Return nil if everything went well
	return nil
}

// CompressDictory compresses the given directory, keeping the original directory intact.
//
// The function takes one argument: the directory to be compressed.
//
// The function returns an error if something goes wrong while compressing the directory.
//
// qzip -k -R directory
//
// 目录递归压缩
func CompressDictoryByEveryFile(inputFile string) error {
	// Create a new QzipCommand with the default configuration
	cmd := internal.GetDefaultQzipCommand()

	// Set the input file, keep the original file and set the directory true
	cmd.KeepSource = true
	cmd.Compression = true
	cmd.InputFile = append(cmd.InputFile, inputFile)
	cmd.IsDirctory = true

	// Execute the command and return the error if something goes wrong
	if err := internal.ExecuteQzipCommand(cmd); err != nil {
		return err
	}

	// Return nil if everything went well
	return nil
}

// CompressFiles compresses the given inputFiles, keeping the original files intact.
//
// The function takes a variable number of arguments: the input files to be compressed.
//
// The function returns an error if something goes wrong while compressing the files.
//
// qzip -k file1 file2 ...
//
// 多文件压缩
func CompressFiles(inputFiles ...string) error {
	// Create a new QzipCommand with the default configuration
	cmd := internal.GetDefaultQzipCommand()

	// Set the input files, keep the original files and set the directory false
	cmd.KeepSource = true
	cmd.Compression = true
	cmd.InputFile = append(cmd.InputFile, inputFiles...)
	cmd.IsDirctory = false

	// Execute the command and return the error if something goes wrong
	if err := internal.ExecuteQzipCommand(cmd); err != nil {
		return err
	}

	// Return nil if everything went well
	return nil
}

// CompressFilesWithBusyPool compresses the given inputFiles, keeping the original files intact, using
// busy polling.
//
// The function takes a variable number of arguments: the input files to be compressed.
//
// The function returns an error if something goes wrong while compressing the files.
//
// qzip -k -P busy file1 file2 ...
//
// 目录多文件压缩，忙轮询
func CompressDictoryWithBusyPoll(inputDirectory string) error {
	// Create a new QzipCommand with the default configuration
	cmd := internal.GetDefaultQzipCommand()

	// Set the input files, keep the original files and set the directory false
	cmd.KeepSource = true
	cmd.Compression = true
	cmd.InputFile = append(cmd.InputFile, inputDirectory)
	cmd.IsDirctory = true

	// 设置忙碌轮询
	cmd.BusyPoll = true

	// Execute the command and return the error if something goes wrong
	if err := internal.ExecuteQzipCommand(cmd); err != nil {
		return err
	}

	// Return nil if everything went well
	return nil
}

// eg: tar -cvf mydir.tgz -I "qzip" ./test_files/
//
// output: mydir.tgz
//
// 使用tar归档文件对目录进行  整体 压缩
func CompressDictoryByTar(inputDirectory, outputFile string) error {
	cmd := internal.GetDefaultTarCommand()
	cmd.Compression = true
	if outputFile == "" {
		outputFile = inputDirectory
	}
	// 检查outputFile是否以.tgz或者.tar.gz结尾，如果不是，加上.tgz
	if !strings.HasSuffix(outputFile, ".tgz") && !strings.HasSuffix(outputFile, ".tar.gz") {
		outputFile = outputFile + ".tgz"
	}
	// 这里的归档文件最后会成为压缩后的文件名，故直接将  归档  文件名设置为 输出 文件名
	cmd.ArchiveFile = outputFile
	cmd.InputFile = append(cmd.InputFile, inputDirectory)
	if err := internal.ExecuteTarCommand(cmd); err != nil {
		return err
	}
	return nil
}
