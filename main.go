package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

// --run 実行
// --recursive, -r サブディレクトリ以下の.ファイルも削除
// --remove-dir .dirなフォルダも削除
func main() {
	if err := cmd(); err != nil {
		log.Fatal(err)
	}
}

func cmd() error {
	flag.Parse()

	inputs := flag.Args()

	for _, input := range inputs {
		dir(input)
	}

	return nil
}

func dir(current string) error {
	entries, err := os.ReadDir(current)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		next := path.Join(current, entry.Name())
		if entry.IsDir() {
			dir(next)
		} else if isDotFile(entry.Name()) {
			fmt.Printf("%s\n", next)
		}
	}

	return nil
}

func isDotFile(name string) bool {
	return strings.HasPrefix(name, ".")
}
