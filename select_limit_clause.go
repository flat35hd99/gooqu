package gooqu

import "fmt"

type limit struct {
	limit int
	exist bool
}

func (q *Query) Limit(number int) *Query {
	// overwrite
	// 必要ならログ出すかerror返すかerror設定する
	// if q.hasLimitClause {
	// }
	var l limit
	l.exist = true
	l.limit = number

	q.limit = l
	return q
}

func (l limit) String() string {
	return fmt.Sprintf("LIMIT %d", l.limit)
}

func (l limit) Words() []string {
	if !l.exist {
		return []string{}
	}
	return []string{
		"LIMIT", fmt.Sprint(l.limit),
	}
}
