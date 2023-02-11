package local

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestInsert(t *testing.T) {
	manager, err := New()
	if err != nil {
		log.Fatal(err)
	}
	c := Comment{
		AuthorId: 1,
		VideoId:  1,
		Msg:      "hello test",
		IsTopped: false,
	}

	err = manager.InsertComment(&c)

	if err != nil {
		log.Println(err)
	}

	marshal, _ := json.Marshal(c)
	fmt.Printf("%s", string(marshal))

	err = manager.DeleteComment(c)
	if err != nil {
		fmt.Println(err)
	}

}

func TestSelect(t *testing.T) {
	manager, _ := New()
	i, err := manager.SelectCommentNumsByVideoId(1000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("有%d条评论", i)
}

func TestList(t *testing.T) {
	manager, _ := New()
	comment, err := manager.SelectCommentListByVideoId(1)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range comment {
		fmt.Println(v)
	}
}
