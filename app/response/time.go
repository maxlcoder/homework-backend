package response

import (
	"fmt"
	"time"
)

type JSONTime time.Time

const format = "2006-01-02 15:04:05"

func (t JSONTime) MarshalJSON() ([]byte, error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return []byte(fmt.Sprintf(`"%s"`, time.Time(t).In(loc).Format(format))), nil
}

func (t JSONTime) String() string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.Time(t).In(loc).Format(format)
}
