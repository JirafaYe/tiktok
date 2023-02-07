package pkg

import (
	"crypto/md5"
	"crypto/sha512"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"testing"
	"time"
)

func TestEncode(t *testing.T) {
	//// 方式一：使用默认选项
	//salt, encodedPwd := Encode("generic password", nil)
	//check := Verify("generic password", salt, encodedPwd, nil)
	//fmt.Println(encodedPwd)
	//fmt.Println(check) // true
	//
	//// 方式二：使用自定义选项
	//options :=Options{SaltLen: 32, Iterations: 10000, KeyLen: 32, HashFunction: md5.New}
	//salt, encodedPwd = Encode("generic password", options)
	//fmt.Printf("%T\n", encodedPwd)
	//fmt.Println(encodedPwd)
	//fmt.Println(salt)
	//check = Verify("generic password", salt, encodedPwd, options)
	//fmt.Println(check) // true
	options := &Options{SaltLen: 6, Iterations: 10000, KeyLen: 12, HashFunction: md5.New}
	pwd := SaltEncodePwd("1234512334")
	fmt.Println(pwd)
	fmt.Println(len(pwd))
	fmt.Println(options)
	verify := Verify("1234512334", strings.Split(pwd, "$")[1], strings.Split(pwd, "$")[2], options)
	fmt.Println(verify)
}

func TestAb(t *testing.T) {
	options := &Options{SaltLen: 6, Iterations: 10000, KeyLen: 12, HashFunction: sha512.New}
	fmt.Println(options)
	verify := Verify("123", "saWDvH", "609c54ceaaa9c8c53c3ebebf", options)
	fmt.Println(verify)
}

func TestAa(t *testing.T) {
	options := &Options{SaltLen: 6, Iterations: 10000, KeyLen: 12, HashFunction: sha512.New}
	salt, pwd := Encode("123", options)
	fmt.Println(options)
	println(fmt.Sprintf("$%s$%s", salt, pwd))
}

func TestB(t *testing.T) {
	password := []byte("thisIsPassWord")
	//password2 := []byte("thisIsPas1sWord")
	nowG := time.Now()
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	fmt.Println("加密后", string(hashedPassword), "耗时", time.Now().Sub(nowG))
	nowC := time.Now()
	err := bcrypt.CompareHashAndPassword([]byte("$2a$10$TakvaOGlwUVmZwXnzfhecOXsxc/Xoyu7RU5DlBxkvarLb2kPBFr4m"), []byte("123"))
	fmt.Println(err)
	fmt.Println("验证耗费时间", time.Now().Sub(nowC))
	fmt.Println(err)

}
