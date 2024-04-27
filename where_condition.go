package gooqu

import "fmt"

// もっと込み入ったものにできる
type expression struct {
	columnName  string
	columnValue interface{}
}

func (exp expression) String() string {
	var columnValue string
	switch v := exp.columnValue.(type) {
	case int:
		columnValue = fmt.Sprint(v)
	case string:
		columnValue = v
	}
	return fmt.Sprintf("`%s` = %s", exp.columnName, columnValue)
}

func (exp expression) Words() []string {
	var columnValue string
	switch v := exp.columnValue.(type) {
	case int:
		columnValue = fmt.Sprint(v)
	case string:
		columnValue = v
	}
	return []string{
		"`" + exp.columnName + "`", "=", columnValue,
	}
}
