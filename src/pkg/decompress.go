package pkg

import (
	"errors"
	"strings"

	"github.com/ordinary-xiyv/qzipgo/src/internal"
)

// qzip -d -k filepath 测试解压
// output:Executing command: /usr/local/bin/qzip -d -k /tmp/test.txt
func DecompressFile(inputFile string) error {
	cmd := internal.GetDefaultQzipCommand()
	cmd.KeepSource = true
	cmd.Compression = false
	cmd.InputFile = append(cmd.InputFile, inputFile)

	if err := internal.ExecuteQzipCommand(cmd); err != nil {
		return err
	}
	return nil
}

// DecompressWithOutputFile decompresses the given inputFile and stores the output in the given outputFile,
// keeping the original file intact.
//
// The function takes two arguments: the input file to be decompressed and the output file where the decompressed
// content should be stored.
//
// The function returns an error if something goes wrong while decompressing the file.
//
// qzip -k -o outputFile inputFile
func DecompressWithOutputFile(inputFile, outputFile string) error {
	// Create a new QzipCommand with the default configuration
	cmd := internal.GetDefaultQzipCommand()

	// Set the input file, keep the original file and set the output file
	cmd.KeepSource = true
	cmd.Compression = false
	cmd.InputFile = append(cmd.InputFile, inputFile)
	cmd.OutputFile = outputFile

	// Execute the command and return the error if something goes wrong
	if err := internal.ExecuteQzipCommand(cmd); err != nil {
		return err
	}

	// Return nil if everything went well
	return nil
}

// DecompressDictory decompresses the given directory, keeping the original directory intact.
//
// The function takes one argument: the directory to be decompressed.
//
// The function returns an error if something goes wrong while decompressing the directory.
//
// qzip -k -R directory
//
// 解压目录下每一个压缩文件
func DecompressDictoryByEveryFile(inputFile string) error {
	// Create a new QzipCommand with the default configuration
	cmd := internal.GetDefaultQzipCommand()

	// Set the input file, keep the original file and set the directory true
	cmd.KeepSource = false
	cmd.Compression = false
	cmd.InputFile = append(cmd.InputFile, inputFile)
	cmd.IsDirctory = true
	// 解压目录必须使用递归
	cmd.Recursive = true

	// Execute the command and return the error if something goes wrong
	if err := internal.ExecuteQzipCommand(cmd); err != nil {
		return err
	}

	// Return nil if everything went well
	return nil
}

// DecompressFiles decompresses the given inputFiles, keeping the original files intact.
//
// The function takes a variable number of arguments: the input files to be decompressed.
//
// The function returns an error if something goes wrong while decompressing the files.
//
// qzip -k file1 file2 ...
//
// 多文件解压
func DecompressFiles(inputFiles ...string) error {

	// Create a new QzipCommand with the default configuration
	cmd := internal.GetDefaultQzipCommand()

	// Set the input file, keep the original file and set the directory true
	cmd.KeepSource = false
	cmd.Compression = false

	cmd.InputFile = append(cmd.InputFile, inputFiles...)

	// Execute the command and return the error if something goes wrong
	if err := internal.ExecuteQzipCommand(cmd); err != nil {
		return err
	}

	// Return nil if everything went well
	return nil
}

// DecompressDictoryWithBusyPoll decompresses the given directory, keeping the original directory intact and using busy polling.
//
// The function takes one argument: the directory to be decompressed.
//
// The function returns an error if something goes wrong while decompressing the directory.
//
// qzip -k -P busy -R directory
//
// 解压目录下每一个压缩文件 忙轮询
func DecompressDictoryWithBusyPoll(inputDirectory string) error {
	// Create a new QzipCommand with the default configuration
	cmd := internal.GetDefaultQzipCommand()

	// Set the input file, keep the original file and set the directory true
	cmd.KeepSource = false
	cmd.Compression = false
	cmd.InputFile = append(cmd.InputFile, inputDirectory)
	cmd.IsDirctory = true
	// 解压目录必须使用递归
	cmd.Recursive = true

	cmd.BusyPoll = true

	// Execute the command and return the error if something goes wrong
	if err := internal.ExecuteQzipCommand(cmd); err != nil {
		return err
	}

	// Return nil if everything went well
	return nil
}

// DecompressDictoryByTar decompresses the given inputFile using tar, and writes the contents to the given outputDirectory.
//
// The function takes two arguments: the input file to be decompressed, and the output directory where the decompressed
// contents should be written.
//
// The function returns an error if something goes wrong while decompressing the file.
//
// If the input file does not have a .tgz or .tar.gz extension, the function will return an error.
func DecompressDictoryByTar(inputFile, outputDirectory string) error {
	cmd := internal.GetDefaultTarCommand()
	cmd.Compression = false
	if inputFile == "" {
		return errors.New("input file is empty")
	}
	// 检查inputFile是否以.tgz或者.tar.gz结尾，如果不是则退出
	if !strings.HasSuffix(inputFile, ".tgz") && !strings.HasSuffix(inputFile, ".tar.gz") {
		return errors.New("input file is not tgz or tar.gz")
	}
	cmd.ArchiveFile = inputFile
	cmd.OutputFile = outputDirectory
	//cmd.InputFile = append(cmd.InputFile, inputFile)
	if err := internal.ExecuteTarCommand(cmd); err != nil {
		return err
	}
	return nil
}
