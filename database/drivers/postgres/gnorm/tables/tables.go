// Code generated by gnorm, DO NOT EDIT!

package tables

import (
	"database/sql"

	"gnorm.org/gnorm/database/drivers/postgres/gnorm"
)

// Row represents a row from 'tables'.
type Row struct {
	TableCatalog              sql.NullString // table_catalog
	TableSchema               sql.NullString // table_schema
	TableName                 sql.NullString // table_name
	TableType                 sql.NullString // table_type
	SelfReferencingColumnName sql.NullString // self_referencing_column_name
	ReferenceGeneration       sql.NullString // reference_generation
	UserDefinedTypeCatalog    sql.NullString // user_defined_type_catalog
	UserDefinedTypeSchema     sql.NullString // user_defined_type_schema
	UserDefinedTypeName       sql.NullString // user_defined_type_name
	IsInsertableInto          sql.NullString // is_insertable_into
	IsTyped                   sql.NullString // is_typed
	CommitAction              sql.NullString // commit_action
}

// Field values for every column in Tables.
var (
	TableCatalogCol              gnorm.SqlNullStringField = "table_catalog"
	TableSchemaCol               gnorm.SqlNullStringField = "table_schema"
	TableNameCol                 gnorm.SqlNullStringField = "table_name"
	TableTypeCol                 gnorm.SqlNullStringField = "table_type"
	SelfReferencingColumnNameCol gnorm.SqlNullStringField = "self_referencing_column_name"
	ReferenceGenerationCol       gnorm.SqlNullStringField = "reference_generation"
	UserDefinedTypeCatalogCol    gnorm.SqlNullStringField = "user_defined_type_catalog"
	UserDefinedTypeSchemaCol     gnorm.SqlNullStringField = "user_defined_type_schema"
	UserDefinedTypeNameCol       gnorm.SqlNullStringField = "user_defined_type_name"
	IsInsertableIntoCol          gnorm.SqlNullStringField = "is_insertable_into"
	IsTypedCol                   gnorm.SqlNullStringField = "is_typed"
	CommitActionCol              gnorm.SqlNullStringField = "commit_action"
)

// Query retrieves rows from 'tables' as a slice of Row.
func Query(db gnorm.DB, where gnorm.WhereClause) ([]*Row, error) {
	const origsqlstr = `SELECT 
		table_catalog, table_schema, table_name, table_type, self_referencing_column_name, reference_generation, user_defined_type_catalog, user_defined_type_schema, user_defined_type_name, is_insertable_into, is_typed, commit_action
		FROM information_schema.tables WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") "

	var vals []*Row
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, err
	}
	for q.Next() {
		r := Row{}

		err = q.Scan(&r.TableCatalog, &r.TableSchema, &r.TableName, &r.TableType, &r.SelfReferencingColumnName, &r.ReferenceGeneration, &r.UserDefinedTypeCatalog, &r.UserDefinedTypeSchema, &r.UserDefinedTypeName, &r.IsInsertableInto, &r.IsTyped, &r.CommitAction)
		if err != nil {
			return nil, err
		}

		vals = append(vals, &r)
	}
	return vals, nil
}