package gooqu

import "fmt"

type groupByClause struct {
	columnName string
}

func newGroupByClause(columnName string) groupByClause {
	return groupByClause{columnName: columnName}
}

func (clause groupByClause) String() string {
	return fmt.Sprintf(`GROUP BY "%s"`, clause.columnName)
}
