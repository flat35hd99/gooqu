package gooqu

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
	tableName *word
}

func newTableReferences(tableName string) tableReferences {
	return tableReferences{
		tableName: newWord(tableName, true),
	}
}

func (tr tableReferences) String() string {
	return tr.tableName.String()
}

func (tr tableReferences) Words() *word {
	return tr.tableName
}
