package local

import (
	"fmt"
	"github.com/JirafaYe/user/pkg"
	"strconv"
	"testing"
)

func TestA(t *testing.T) {
	if err := pkg.Init("2006-01-02", 1); err != nil {
		fmt.Printf("init failed, err:%v\n", err)
		return
	}
	id := pkg.GenID()
	num := id >> 48
	fmt.Println(num)
	name := "未命令" + strconv.Itoa(int(num))
	fmt.Println(name)
}

func TestB(t *testing.T) {
	if err := pkg.Init("2006-01-02", 1); err != nil {
		fmt.Printf("init failed, err:%v\n", err)
		return
	}

	for i := 0; i < 1000; i++ {
		id := pkg.GenID()
		num := id >> 52
		fmt.Println(num)
	}
}

func TestC(t *testing.T) {
	if err := pkg.Init("2006-01-02", 1); err != nil {
		fmt.Printf("init failed, err:%v\n", err)
		return
	}

	for i := 0; i < 1000; i++ {
		id := pkg.GenID()
		num := strconv.FormatInt(id, 10)
		fmt.Println(id)
		fmt.Println(num)
		fmt.Println(num[15:19])
	}
}
