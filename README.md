# Go 语言调用 Qzip

## 简介

QZipGo 是一个基于 QAT（Quick Assist Technology）驱动的高性能压缩/解压缩工具库。本项目利用 [qzip 库](<https://github.com/intel/QATzip> ) 构建命令，通过调用 QAT 驱动实现硬件加速的压缩和解压缩操作，显著提升了数据处理的效率。

## 主要特点

1. **硬件加速**：利用 [Intel QAT 技术](<https://www.intel.com/content/www/us/en/architecture-and-technology/intel-quick-assist-technology-overview.html>)，实现压缩和解压缩操作的硬件加速，大幅提高处理速度。

2. **基于 qzip 库**：在成熟的 qzip 库基础上构建，确保了稳定性和兼容性。

3. **简化的接口**：提供一组精心设计的常用接口，使开发者能够轻松集成和使用硬件加速的压缩/解压缩功能。

4. **灵活性**：支持多种压缩级别和模式，满足不同场景的需求。

5. **高效率**：通过硬件加速，显著减少 CPU 使用率，为系统其他任务释放资源。

## 安装

确保您已经安装了 Go 语言环境。你可以通过以下命令安装 Qzip：

```bash
go get -u github.com/ordinary-xiyv/qzip
```

## 依赖

本仓库下需要以下环境依赖：

1. Go语言版本: **1.21.7**

2. Tar打包工具：**1.30**

3. QAT硬件加速驱动：**2.0**

4. Qzip库：**1.2.0**

## 项目结构

```bash
.
├── go.mod
├── README.md
└── src
    ├── internal
    │   ├── checkqzip.go                                  // 检查 QAT 环境
    │   ├── ops.go                                        // Qzip 命令选项操作
    │   └── qzip.go                                       // Qzip 命令构建与执行
    ├── pkg
    │   ├── compress.go                                   // 对外提供的压缩接口
    │   ├── decompress.go                                 // 对外提供的解压缩接口
    │   └── qzip.go                                       // qzip本地测试和环境检测
    ├── test
    │   └── qzip_test.go                                  // qzip测试用例
    └── testfiles
        └── test_15mb.json                                // 本地测试文件
```

## 使用说明

本库提供了一组常用的压缩和解压缩的接口，它们基于 qzip 库实现，可以对单文件，多文件，文件夹等进行压缩和解压缩。代码均存放在 `pkg` 目录下。

如若您想自定义压缩和解压缩的参数，请参考 [命令选项](<./src/internal/ops.go>)。

接下来我将教您如何完成一个自定义的命令：

```bash
qzip -k ./testfiles/test_15mb.json -o ./testfiles/test_with_output_file.txt
```

1. 首先您需要从代码中获取一个默认的qzip命令：

    ```go
    cmd := internal.GetDefaultQzipCommand()
    ```

2. 然后您需要设置命令的参数：

    ```go
    // KeepSource 设置为 true，保留源文件，对应命令中的 -k 选项
    cmd.KeepSource = true
    // Compression 设置为 true，对应命令为压缩指令（无 -d 选项）
    cmd.Compression = true
    // InputFile 设置为需要压缩的文件路径
    cmd.InputFile = append(cmd.InputFile, "./testfiles/test_15mb.json")
    // OutputFile 设置为压缩后的文件路径以及文件名，对应命令中的 -o 选项，最后文件会带上默认的后缀 .gz
    // 压缩后的文件名为：test_with_output_file.json.gz
    cmd.OutputFile = "./testfiles/test_with_output_file.json"
    ```

3. 最后执行命令：

    ```go
    // 在这一步，代码会根据您的命令选项构建命令，然后执行命令
    if err := internal.ExecuteQzipCommand(cmd); err != nil {
         return err
    }
    ```

## 测试用例

本库目前仅提供压缩和解压缩功能，具体请参考 [测试用例](<./src/test/qzip_test.go>)。

## 常见问题

1. Q: 为什么压缩后的文件名是 test_with_output_file.json.gz，而不是 test_with_output_file.gz？
    A: 因为QAT进行压缩不会更改后缀，而是在原文件名后加上 .gz （gzip算法）后缀。解压亦是如此。

2. Q: 为什么使用QAT压缩目录时无法指定压缩后的文件名？
    A: 因为QAT进行压缩时，会将目录下的所有文件都压缩，即QAT的操作是针对文件，若指定了输出文件名，则会导致目录被压缩为一个文件，解压后丢失目录结构。
