package main

import (
	"container/list"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		return
	}
	file.WriteString("Nitesh Kumar")
	file.Close()

	file, err = os.Open("test.txt")
	if err != nil {
		return
	}
	defer file.Close()

	stat, stat_err := file.Stat()

	if stat_err != nil {
		return
	}

	bs := make([]byte, stat.Size())
	_, read_err := file.Read(bs)

	if read_err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)

	bs, err = ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}
	str = string(bs)
	fmt.Println(str)

	read_dir()

	doubly_linked_list()
}

func read_dir() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.ReadDir(-1)
	if err != nil {
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	walk_dir()

	err1 := errors.New("Error Mesaage")
	fmt.Println(err1)
}

func walk_dir() {
	filepath.Walk("..", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
func doubly_linked_list() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}
