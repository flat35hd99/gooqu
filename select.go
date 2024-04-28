package gooqu

import "fmt"

type Record map[string]interface{}

type Query struct {
	selectExpressions selectExpressions
	whereCondition    expression
	tableReferences   tableReferences
	limit             limit
}

func Where(record Record) *Query {
	var query Query
	for key, value := range record {
		query.whereCondition = newExpression(key, value)
	}
	return &query
}

func (q *Query) Select(exprs ...SelectExpression) *Query {
	q.selectExpressions = newSelectExpressions(exprs...)
	return q
}

func (q *Query) From(table_name string) *Query {
	q.tableReferences = newTableReferences(table_name)
	return q
}

func (q Query) ToSQL() string {
	// TODO: スライスの要素を前から順にpopして、文字列にする、みたいなことをしないといけない。
	//       そうしないと、flagが無限に増えてパターンが指数関数的に増える。
	// return fmt.Sprintf("SELECT * FROM `%s` WHERE %s;", q.tableReferences, q.whereCondition)
	// root := newWord("", false)
	// root.n(newWord("SELECT", false)).n(newWord(q.selectExpressions.String(), false)).n(newWord("FROM", false)).n(newWord(q.tableReferences.String(), false)).n(newWord("WHERE", false)).n(newWord(q.whereCondition.String(), false)).n(newWord(q.limit.String(), false)).n(newWord(";", false))
	// return root.String()
	return fmt.Sprintf("SELECT %s FROM %s WHERE %s %s;", q.selectExpressions, q.tableReferences, q.whereCondition, q.limit)
}
