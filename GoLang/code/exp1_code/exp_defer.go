package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	dir := os.Args[1]
	start, err := os.Stat(dir)
	if err != nil || !start.IsDir() {
		os.Exit(2)
	}

	var targets []string
	filepath.Walk(dir, func(fPath string, fInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fInfo.Mode().IsRegular() {
			return nil
		}

		targets = append(targets, fPath)
		return nil
	})

	for _, target := range targets {
		f, err := os.Open(target)
		if err != nil {
			fmt.Println("bad target:", target, "error:", err) //error:too many open files
			break
		}
		func() {
			fmt.Println("test")
			defer f.Close() }()
		//func() {
		//	f, err := os.Open(target)
		//	if err != nil {
		//		fmt.Println("bad target:", target, "error:", err)
		//		return	// 在匿名函数内使用 return 代替 break 即可
		//	}
		//	defer f.Close()	// 匿名函数执行结束，调用关闭文件资源
		//
		//	// 使用 f 资源
		//}()
	}
}
