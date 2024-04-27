package gooqu

type Record map[string]interface{}

type Query struct {
	whereCondition  expression
	tableReferences tableReferences
	limit           limit
}

func Where(record Record) *Query {
	var query Query
	for key, value := range record {
		query.whereCondition = newExpression(key, value)
	}
	return &query
}

func (q *Query) From(table_name string) *Query {
	q.tableReferences = newTableReferences(table_name)
	return q
}

func (q Query) ToSQL() string {
	// TODO: スライスの要素を前から順にpopして、文字列にする、みたいなことをしないといけない。
	//       そうしないと、flagが無限に増えてパターンが指数関数的に増える。
	// return fmt.Sprintf("SELECT * FROM `%s` WHERE %s;", q.tableReferences, q.whereCondition)
	root := newWord("", false)
	root.n(q.tableReferences).n(newWord("WHERE", false)).n(q.whereCondition).n(q.limit).n(newWord(";", false))
	return "SELECT * FROM " + root.String()
}
