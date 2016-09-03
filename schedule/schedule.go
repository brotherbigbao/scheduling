package schedule

import (
	"fmt"
	//"strconv"
	//"strconv"
	"math/rand"
)

type Schedule struct {
	member []string
	personEveryDay int
	totalDay int
}

func (this *Schedule) Create(member []string, personEveryDay int, totalDay int) {
	this.member, this.personEveryDay, this.totalDay = member, personEveryDay, totalDay
	this.create()
}

func (this *Schedule) create() {
	memberNum := len(this.member) // 成员数量
	personNums := this.totalDay * this.personEveryDay // 总人次
	timesEveryPerson := personNums/memberNum //每人总执班天数

	memberBucket := make([]Member, memberNum)
	for i, name := range this.member {
		memberBucket[i] = Member{name, timesEveryPerson}
	}
	for i, obj := range memberBucket {
		fmt.Println(i, obj)
	}

	var result = make([]interface{}, this.totalDay)
	for index := 0; index < this.totalDay; index++ {
		result[index] = this.pick(memberBucket)
	}
}

func (this *Schedule) pick(bucket []Member) []string {
	var pickedNum int   // 已提取数量
	var pickedKey []int // 已提取过的member key
	var maxkey int      // 剩余次数最大的那个
	var result []string // 提取结果

	for key, member := range bucket {
		if pickedNum == this.personEveryDay {
			break
		}
		if member.GetTimes() > bucket[maxkey].GetTimes() {
			maxkey = key
		}
	}
	exclude := make([]int, this.personEveryDay)
	for index2 := 0; index2 < this.personEveryDay; index2++ {
		random := rand.Intn(memberNum-1)
		name, err := memberBucket[index].GetOneTime()
	}
}