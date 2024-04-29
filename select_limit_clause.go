package gooqu

import "fmt"

type limitClause struct {
	limit int
	exist bool
}

func (q *Query) Limit(number int) *Query {
	q.setLimitClause = true
	var l limitClause
	l.exist = true
	l.limit = number

	q.limitClause = l
	return q
}

func (l limitClause) String() string {
	if !l.exist {
		return ""
	}

	return fmt.Sprintf("LIMIT %d", l.limit)
}
