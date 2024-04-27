package gooqu

import "fmt"

type Record map[string]interface{}

type Query struct {
	whereCondition  expression
	tableReferences tableReferences
}

// もっと込み入ったものにできる
type expression struct {
	columnName  string
	columnValue interface{}
}

// もっと込み入ったものにできる
// ref: https://dev.mysql.com/doc/refman/8.3/en/join.html
/*
able_reference: {
    table_factor
  | joined_table
}

table_factor: {
    tbl_name [PARTITION (partition_names)]
        [[AS] alias] [index_hint_list]
  | [LATERAL] table_subquery [AS] alias [(col_list)]
  | ( table_references )
}

joined_table: {
    table_reference {[INNER | CROSS] JOIN | STRAIGHT_JOIN} table_factor [join_specification]
  | table_reference {LEFT|RIGHT} [OUTER] JOIN table_reference join_specification
  | table_reference NATURAL [INNER | {LEFT|RIGHT} [OUTER]] JOIN table_factor
}
*/
type tableReferences struct {
	tableName string
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
	var columnValue string
	switch v := q.whereCondition.columnValue.(type) {
	case int:
		columnValue = fmt.Sprint(v)
	case string:
		columnValue = v
	}

	tableReferences := q.tableReferences.tableName
	whereCondition := fmt.Sprintf("`%s` = %s", q.whereCondition.columnName, columnValue)
	return fmt.Sprintf("SELECT * FROM `%s` WHERE %s;", tableReferences, whereCondition)
}
