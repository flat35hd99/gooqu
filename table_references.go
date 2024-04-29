package gooqu

import (
	"fmt"
	"strings"
)

// もっと込み入ったものにできる
// ref: https://dev.mysql.com/doc/refman/8.3/en/join.html
/*
table_reference: {
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

func newTableReferences(tableName string) tableReferences {
	return tableReferences{
		tableName: tableName,
	}
}

func (tr tableReferences) String() string {
	return fmt.Sprintf(`"%s"`, tr.tableName)
}

// JOIN

type T struct {
	tableName string
}

func (t T) String() string {
	return fmt.Sprintf(`"%s"`, t.tableName)
}

func Table(tableName string) T {
	return T{tableName: tableName}
}

type SearchCondition struct {
	left, right string
}

func (sc SearchCondition) String() string {
	separator := func(v string) string {
		// users -> "users"
		// users.id -> "users"."id"
		// otherwise -> ?
		elements := strings.Split(v, ".")
		if len(elements) == 1 {
			return fmt.Sprintf(`"%s"`, elements[0])
		} else if len(elements) == 2 {
			return fmt.Sprintf(`"%s"."%s"`, elements[0], elements[1])
		}
		return fmt.Sprintf(`"%s"`, v) // FIXME: いい感じにしてほしい
	}
	return fmt.Sprintf(`%s = %s`, separator(sc.left), separator(sc.right))
}

func On(conditions map[string]string) SearchCondition {
	var result SearchCondition
	for left, right := range conditions {
		result = SearchCondition{left: left, right: right} // TODO: 複数個に対応する
	}
	return result
}

type joinClause struct {
	table           T
	searchCondition SearchCondition
}

func (clause joinClause) String() string {
	return fmt.Sprintf(`INNER JOIN %s ON %s`, clause.table, clause.searchCondition)
}

func (q *Query) Join(table T, on SearchCondition) *Query {
	q.fromClause.setJoinClause = true
	q.fromClause.joinClause = joinClause{table: table, searchCondition: on}
	return q
}
