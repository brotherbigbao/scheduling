package main

import (
	//"os"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	bytes, err := ioutil.ReadFile("./list.txt")
	if err != nil {
		fmt.Println("文件list.txt读取失败.")
		return
	}
	str := string(bytes)
	str = strings.Replace(str, "\r", "", -1)
	str = strings.Trim(str, " ")
	members := strings.Split(str, "\n")
	if len(members) == 0 {
		fmt.Println("文件内容为空")
		return
	}
	var m []string
	for _,v := range members {
		if(len(v) > 0) {
			m = append(m, v)
		}
	}
	if len(m) < 2 {
		notice := fmt.Sprintf("成员数量是:%s, 不符合要求", strconv.Itoa(len(m)))
		fmt.Println(notice)
		return
	}

	fmt.Println("当前成员列表如下:")
	fmt.Println(m)
	fmt.Println("请按回车确认")
}