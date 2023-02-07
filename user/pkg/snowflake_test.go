package pkg

import (
	"fmt"
	"testing"
	"time"
)

func TestSnowFlake(t *testing.T) {
	if err := Init("2006-01-02", 1); err != nil {
		fmt.Printf("init failed, err:%v\n", err)
		return
	}
	maps := make(map[int64]interface{}, 0)
	//测试重复性，经过测试，10000000个id生成需要3s左右，且不存在重复
	test1(maps)

	//取余九位数，在1.5ms内并发1w次，不会出现重复(我自己的抖音号是九位数，且九位数能覆盖上亿用户)

	for i := 0; i < 100; i++ {
		test2(maps)
	}
	time.Sleep(1000000000000000)

}
func test1(maps map[int64]interface{}) {
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		id := GenID()
		if _, ok := maps[id]; ok {
			fmt.Printf("出现重复")
		}
		maps[id] = 0
	}
	end := time.Since(start)
	fmt.Printf("容量大小:%v\n", len(maps))
	fmt.Println(end)
}

func test2(maps map[int64]interface{}) {
	//maps := make(map[int64]interface{}, 0)
	start := time.Now()
	for i := 0; i < 100000; i++ {
		id := GenID() % 100000000
		if _, ok := maps[id]; ok {
			fmt.Printf("出现重复\n")
		}
		maps[id] = 0
	}
	end := time.Since(start)
	fmt.Printf("容量大小:%v\n", len(maps))
	fmt.Println(end)
}
