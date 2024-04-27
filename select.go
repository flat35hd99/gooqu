package gooqu

import "strings"

type Record map[string]interface{}

type Query struct {
	whereCondition  expression
	tableReferences tableReferences
	limit           limit
	hasLimitClause  bool
}

func Where(record Record) *Query {
	var query Query
	for key, value := range record {
		var exp expression
		exp.columnName = key
		exp.columnValue = value
		query.whereCondition = exp
	}
	return &query
}

func (q *Query) From(table_name string) *Query {
	q.tableReferences.tableName = table_name
	return q
}

func (q Query) ToSQL() string {
	// TODO: スライスの要素を前から順にpopして、文字列にする、みたいなことをしないといけない。
	//       そうしないと、flagが無限に増えてパターンが指数関数的に増える。
	// return fmt.Sprintf("SELECT * FROM `%s` WHERE %s;", q.tableReferences, q.whereCondition)
	selectClause := []string{"SELECT", "*"}
	fromClause := append([]string{"FROM"}, q.tableReferences.Words()...)
	whereClause := append([]string{"WHERE"}, q.whereCondition.Words()...)
	limitClause := q.limit.Words()

	result := append(append(append(selectClause, fromClause...), whereClause...), limitClause...)
	return strings.Join(result, " ") + ";"
}
