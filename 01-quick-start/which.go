package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}
	file := arguments[1]

	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)
		// 是否存在
		fileInfo, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		mode := fileInfo.Mode()
		// 是否为普通文件
		if !mode.IsRegular() {
			continue
		}

		// 是否可执行
		if mode&0111 != 0 {
			fmt.Println(fullPath)
			return
		}
	}
}
