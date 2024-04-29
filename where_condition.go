package gooqu

import "fmt"

// もっと込み入ったものにできる
type expression struct {
	columnName  string
	columnValue interface{}
}

func newExpression(columnName string, columnValue interface{}) expression {
	return expression{
		columnName:  columnName,
		columnValue: columnValue,
	}
}

func (expr expression) String() string {
	var columnValue string
	switch v := expr.columnValue.(type) {
	case string:
		columnValue = fmt.Sprintf(`"%s"`, v)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		columnValue = fmt.Sprintf("%d", v)
	case float32, float64:
		columnValue = fmt.Sprintf("%f", v)
	default:
		columnValue = fmt.Sprintf(`"%v"`, v)
	}
	return fmt.Sprintf(`"%s" = %s`, expr.columnName, columnValue)
}
