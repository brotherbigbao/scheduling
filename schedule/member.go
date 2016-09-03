package schedule

import "errors"

type Member struct {
	name string
	times int
}

func (this *Member) GetOneTime() (string, error) {
	if this.times > 0 {
		this.times--
		return this.name, nil
	} else {
		return "", errors.New("已经没有了")
	}
}

func (this *Member) GetTimes() int {
	return this.times
}