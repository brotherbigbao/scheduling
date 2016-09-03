package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"bufio"
	"os"
	"github.com/dabao1989/scheduling/schedule"
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
	var totalM int = len(m)
	if totalM < 2 {
		notice := fmt.Sprintf("成员数量是:%s, 不符合要求", strconv.Itoa(totalM))
		fmt.Println(notice)
		return
	}

	// 用户确认是否正确
	confirmMemberMsg := fmt.Sprintf("当前一共有%s个成员,列表如下:", strconv.Itoa(totalM))
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
	numsOneDay := setNumsOneDay(totalM)

	// 执班总天数
	totalDay := setTotalDay(numsOneDay, totalM)

	// 用户确认
	numsOneDayNotice := fmt.Sprintf("你输入的每天值班人数是:%s, 总值班天数是:%s", strconv.Itoa(numsOneDay), strconv.Itoa(totalDay))
	fmt.Println(numsOneDayNotice)
	fmt.Println("正在生成排班计划,请稍等...")

	schedule := schedule.Schedule{}
	result := schedule.Create(m, numsOneDay, totalDay)

	for _, v := range result {
		fmt.Println(v)
	}
	//fmt.Println(result)
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
			data, _, _ = reader.ReadLine()
			command = string(data)
		}

	}

	return result
}

func setNumsOneDay(max int) int {
	fmt.Print("请输入每天值班人数:")

	confirm := false
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	command := string(data)


	var result int
	var num int

	for !confirm {
		nums, err := strconv.ParseInt(command, 10, 8)
		num = int(nums)
		if err != nil || num < 1 || num >= max {
			fmt.Print("请输入每天值班人数,需大于0小于成员数:")
			data, _, _ = reader.ReadLine()
			command = string(data)
		} else {
			confirm, result = true, num
			break
		}

	}

	return result

}

//func totalDay
func setTotalDay(numsOneDay, totalM int) int {
	fmt.Print("请输入值班总天数:")
	confirm := false
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	command := string(data)

	var result int
	var total int

	for !confirm {
		totalT, err := strconv.ParseInt(command, 10, 8)
		total = int(totalT)
		ava := (total*numsOneDay)/totalM //平均每人要值班几天
		if err == nil && ava > 0 {
			result = (ava*totalM)/numsOneDay
			confirm = true
			break
		} else {
			fmt.Print("请输入值班总天数:")
			data, _, _ = reader.ReadLine()
			command = string(data)
		}
	}

	if total != result {
		noticeMsg := fmt.Sprintf("由于你输入的天数%s不能保证每人值班天数相同,已自动设为%s天", strconv.Itoa(total), strconv.Itoa(result))
		fmt.Println(noticeMsg)
	}

	return result
}