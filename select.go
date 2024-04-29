package gooqu

import (
	"fmt"
	"strings"
)

type Record map[string]interface{}

type Query struct {
	selectClause selectClause
	whereClause  whereClause
	fromClause   fromClause
	limitClause  limitClause
	groupBy      groupByClause

	setWhereClause   bool
	setGroupBYClause bool
	setLimit         bool
}

type selectClause struct {
	selectExpressions selectExpressions
}

func (clause selectClause) String() string {
	return fmt.Sprintf("SELECT %s", clause.selectExpressions.String())
}

type whereClause struct {
	whereCondition expression
}

func (clause whereClause) String() string {
	return fmt.Sprintf("WHERE %s", clause.whereCondition.String())
}

type fromClause struct {
	tableReferences tableReferences
}

func (clause fromClause) String() string {
	return fmt.Sprintf("FROM %s", clause.tableReferences)
}

func Where(record Record) *Query {
	var query Query
	query.setWhereClause = true
	for key, value := range record {
		query.whereClause.whereCondition = newExpression(key, value)
	}
	return &query
}

func Select(exprs ...SelectExpression) *Query {
	var q Query
	q.selectClause.selectExpressions = newSelectExpressions(exprs...)
	return &q
}

func (q *Query) Select(exprs ...SelectExpression) *Query {
	q.selectClause.selectExpressions = newSelectExpressions(exprs...)
	return q
}

func (q *Query) From(table_name string) *Query {
	q.fromClause.tableReferences = newTableReferences(table_name)
	return q
}

func (q *Query) GroupBy(columnName string) *Query {
	q.setGroupBYClause = true
	q.groupBy = newGroupByClause(columnName)
	return q
}

func (q Query) ToSQL() string {
	elements := []fmt.Stringer{
		q.selectClause, q.fromClause,
	}
	if q.setGroupBYClause {
		elements = append(elements, q.groupBy)
	}
	if q.setWhereClause {
		elements = append(elements, q.whereClause)
	}
	if q.setLimit {
		elements = append(elements, q.limitClause)
	}

	hoge := []string{}
	for _, V := range elements {
		hoge = append(hoge, V.String())
	}
	return strings.Join(hoge, " ") + ";"
}
