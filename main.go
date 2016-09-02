package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"bufio"
	"os"
)

func main() {
	// 读取文件解析
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

	// 排除掉不符合条件的成员
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

	// 用户确认是否正确
	confirmMemberMsg := fmt.Sprintf("当前一共有%s个成员,列表如下:")
	fmt.Println(confirmMemberMsg)
	i := 1
	for _, v := range m {
		fmt.Println(strconv.Itoa(i) + ":" + v)
		i++
	}

	if confirmMember() == "n" {
		fmt.Println("Bye-bye.")
		return
	}

	// 用户选择每天值班人数

}

func confirmMember() string {
	fmt.Print("请确认[y/n]:")

	confirm := false
	result := ""

	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	command := string(data)

	for !confirm {
		if command == "y" || command == "Y" {
			confirm, result = true, "y"
			break
		} else if command == "n" || command == "N" {
			confirm, result = true, "n"
			break
		} else {
			fmt.Print("请输入[y/n]:")
			data, _, _ := reader.ReadLine()
			command = string(data)
		}

	}

	return result
}