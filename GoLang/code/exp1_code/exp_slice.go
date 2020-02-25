package main

import (
	"bytes"
)

// 错误使用 slice 的拼接示例
func main() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/') // 4
	println(sepIndex)

	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]
	println("dir1: ", string(dir1))		// AAAA
	println("dir2: ", string(dir2))		// BBBBBBBBB

	dir1 = append(dir1, "suffix"...)
	println("current path: ", string(path))	// AAAAsuffixBBBB

	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})
	println("dir1: ", string(dir1))		// AAAAsuffix
	println("dir2: ", string(dir2))		// uffixBBBB

	println("new path: ", string(path))	// AAAAsuffix/uffixBBBB	// 错误结果
}

//// 使用 full slice expression
//func main() {
//	path := []byte("AAAA/BBBBBBBBB")
//	sepIndex := bytes.IndexByte(path, '/') // 4
//	dir1 := path[:sepIndex:sepIndex]		// 此时 cap(dir1) 指定为4， 而不是先前的 16
//	dir2 := path[sepIndex+1:]
//	dir1 = append(dir1, "suffix"...)
//
//	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})
//	println("dir1: ", string(dir1))		// AAAAsuffix
//	println("dir2: ", string(dir2))		// BBBBBBBBB
//	println("new path: ", string(path))	// AAAAsuffix/BBBBBBBBB
//}