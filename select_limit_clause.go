package gooqu

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

func (l limit) Words() *word {
	if !l.exist {
		return newWord("", false)
	}
	root := newWord("LIMIT", false)
	root.n(newWord(l.limit, false))
	return root
}
