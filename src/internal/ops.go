package internal

import (
	"errors"
	"fmt"
)

// qzip -L选项支持的压缩级别
// 在 qzip 中，您可以使用 -L 或 --level 选项来设置压缩级别。具体的级别范围通常从 1 到 9：
// - 1：快速压缩，文件较大，适合需要快速处理的场景。
// - 5：默认压缩级别，平衡了压缩比和速度。
// - 9：最高压缩，文件最小，但压缩速度最慢，适合对压缩比要求极高的场景。
const (
	_ COMPRESSION_LEVEL = iota
	LEVEL_1
	LEVEL_2
	LEVEL_3
	LEVEL_4
	LEVEL_5
	LEVEL_6
	LEVEL_7
	LEVEL_8
	LEVEL_9
)

// qzip -t选项支持的算法
const (
	_ ALGORITHM_TYPE = iota
	LZ4
	LZ4S
	GZIP
	GZIPEXT
)

// qzip -O 选项支持的文件头
const (
	_ FILE_HEADER = iota
	FILE_HEADER_GZIP
	FILE_HEADER_GZIPEXT
	FILE_HEADER_LZ4
	FILE_HEADER_LZ4S
)

// 设置算法
func (q *QzipCommand) SetAlgorithm() {
	switch q.Algorithm {
	case LZ4:
		q.Options = append(q.Options, "-A lz4")
	case GZIP:
		q.Options = append(q.Options, "-A gzip")
	case LZ4S:
		q.Options = append(q.Options, "-A lz4s")
	case GZIPEXT:
		q.Options = append(q.Options, "-A gzipext")
	default:
		// 未知算法或者不指定，均默认使用GZIP，不需要加任何参数
		return
	}
}

// 设置压缩级别
func (q *QzipCommand) SetLevel() {
	if condition := q.Level < LEVEL_1 || q.Level > LEVEL_9; condition {
		// 未知级别或者不指定，均默认使用LEVEL_5
		q.Level = LEVEL_5
	}
	if q.Compression && q.Level != LEVEL_5 {
		q.Options = append(q.Options, "-L", fmt.Sprintf("%d", q.Level))
	}

}

// 设置解压
func (q *QzipCommand) SetDecompression() {
	if !q.Compression {
		q.Options = append(q.Options, "-d")
	}
}

// 设置保留源文件
func (q *QzipCommand) SetKeepSource() {
	if q.KeepSource {
		q.Options = append(q.Options, "-k")
	}
}

// 设置输出文件
func (q *QzipCommand) SetOutputFile() {
	// 如果是目录，则不需要设置-o选项
	if q.IsDirctory {
		return
	}
	if q.OutputFile != "" {
		q.Options = append(q.Options, "-o", q.OutputFile)
	}
}

// 设置目录递归操作
func (q *QzipCommand) SetRecursive() {
	if q.IsDirctory && q.Recursive {
		q.Options = append(q.Options, "-R")
	}
}

// 设置忙轮询
func (q *QzipCommand) SetBusyPoll() {
	if q.BusyPoll {
		q.Options = append(q.Options, "-P busy")
	}
}

// 设置文件头
func (q *QzipCommand) SetFileHeader() error {

	switch q.FileHeader {
	case FILE_HEADER_GZIP:
		if q.Algorithm == GZIP {
			q.Options = append(q.Options, "-O gzip")
			return nil
		}
		return errors.New("算法与指定文件头不匹配，已忽略")
	case FILE_HEADER_GZIPEXT:
		if q.Algorithm == GZIPEXT {
			q.Options = append(q.Options, "-O gzipext")
			return nil
		}
		return errors.New("算法与指定文件头不匹配，已忽略")
	case FILE_HEADER_LZ4:
		if q.Algorithm == LZ4 {
			q.Options = append(q.Options, "-O lz4")
			return nil
		}
		return errors.New("算法与指定文件头不匹配，已忽略")
	case FILE_HEADER_LZ4S:
		if q.Algorithm == LZ4S {
			q.Options = append(q.Options, "-O lz4s")
			return nil
		}
		return errors.New("算法与指定文件头不匹配，已忽略")
	default:
		// 未知算法或者不指定，均默认使用GZIPEXT，不需要加任何参数
		return nil
	}
}
