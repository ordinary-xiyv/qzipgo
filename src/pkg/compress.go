package pkg

import "github.com/ordinary-xiyv/qzipgo/src/internal"

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
func CompressDictoryWithBusyPoll(inputFiles string) error {
	// Create a new QzipCommand with the default configuration
	cmd := internal.GetDefaultQzipCommand()

	// Set the input files, keep the original files and set the directory false
	cmd.KeepSource = true
	cmd.Compression = true
	cmd.InputFile = append(cmd.InputFile, inputFiles)
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
