package utility

import (
	"database/sql"
	"fmt"
	"strings"
)

type DatabasePostgreSQLConnector struct {
	connStr string
	db      *sql.DB
}

func NewDatabasePostgreSQLConnector() *DatabasePostgreSQLConnector {
	databasePostgreSQLConnector := &DatabasePostgreSQLConnector{}
	databasePostgreSQLConnector.connStr = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", Configuration.Database.User, Configuration.Database.Password, Configuration.Database.Name)
	return databasePostgreSQLConnector
}

func (databasePostgreSQLConnector *DatabasePostgreSQLConnector) OpenConnection() {
	var err error
	databasePostgreSQLConnector.db, err = sql.Open("postgres", databasePostgreSQLConnector.connStr)
	if err != nil {
		panic(err)
	}
}

func (databasePostgreSQLConnector *DatabasePostgreSQLConnector) CloseConnection() {
	databasePostgreSQLConnector.db.Close()
}

func (databasePostgreSQLConnector *DatabasePostgreSQLConnector) InsertIntoTable(table string, fields map[string]string) int {

	var fieldNames string
	var placeholders string
	var values []interface{}

	for fieldName := range fields {
		fieldNames += fieldName + ","
		placeholders += "?,"
	}
	fieldNames = fieldNames[:len(fieldNames)-1]
	placeholders = placeholders[:len(placeholders)-1]

	for _, field := range fields {
		for _, value := range field {
			values = append(values, value)
		}
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id", table, fieldNames, placeholders)

	stmt, err := databasePostgreSQLConnector.db.Prepare(query)
	if err != nil {
		// Gestisci l'errore
		panic(err)
	}

	defer stmt.Close()

	var id int
	err = databasePostgreSQLConnector.db.QueryRow(query, values...).Scan(&id)
	if err != nil {
		// Gestisci l'errore
		panic(err)
	}

	return id
}

func (databasePostgreSQLConnector *DatabasePostgreSQLConnector) SelectFromTableWhere(table string, whereClause map[string]string, fields ...string) ([]map[string]interface{}, error) {
	var queryBuilder strings.Builder
	var values []interface{}
	queryBuilder.WriteString("SELECT ")

	if len(fields) > 0 {
		queryBuilder.WriteString(strings.Join(fields, ", "))
	} else {
		queryBuilder.WriteString("*")
	}

	queryBuilder.WriteString(" FROM " + table)
	if len(whereClause) > 0 {
		queryBuilder.WriteString(" WHERE ")
		conditions := make([]string, 0, len(whereClause))

		for key, value := range whereClause {
			conditions = append(conditions, fmt.Sprintf("%s = ?", key))
			values = append(values, value)
		}
		queryBuilder.WriteString(strings.Join(conditions, " AND "))
	}

	rows, err := databasePostgreSQLConnector.db.Query(queryBuilder.String(), values...)
	if err != nil {
		return nil, fmt.Errorf("Error during query execution: %v", err)
	}
	defer rows.Close()

	var results []map[string]interface{}
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("Error while fetching column names: %v", err)
	}
	for rows.Next() {
		result := make(map[string]interface{})
		values := make([]interface{}, len(columns))
		valuePointers := make([]interface{}, len(columns))
		for i := range columns {
			valuePointers[i] = &values[i]
		}
		if err := rows.Scan(valuePointers...); err != nil {
			return nil, fmt.Errorf("Error during result scanning: %v", err)
		}
		for i, column := range columns {
			result[column] = values[i]
		}
		results = append(results, result)
	}

	return results, nil
}

func (databasePostgreSQLConnector *DatabasePostgreSQLConnector) UpdateTableWhere(table string, whereClause map[string]string, updateFields map[string]string) (int, error) {
	queryBuilder := "UPDATE " + table + " SET "
	var setValues []string
	var updateValues []interface{}
	for key, value := range updateFields {
		setValues = append(setValues, fmt.Sprintf("%s=?", key))
		updateValues = append(updateValues, value)
	}
	queryBuilder += fmt.Sprintf("%s ", Join(setValues, ", "))
	var whereValues []string
	var whereParams []interface{}
	for key, value := range whereClause {
		whereValues = append(whereValues, fmt.Sprintf("%s=?", key))
		whereParams = append(whereParams, value)
	}
	queryBuilder += fmt.Sprintf("WHERE %s ", Join(whereValues, " AND "))

	result, err := databasePostgreSQLConnector.db.Exec(queryBuilder, append(updateValues, whereParams...)...)
	if err != nil {
		return 0, fmt.Errorf("errore durante l'esecuzione della query di aggiornamento: %v", err)
	}
	numRowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("errore durante l'ottenimento del numero di righe interessate: %v", err)
	}
	return int(numRowsAffected), nil
}

func (databasePostgreSQLConnector *DatabasePostgreSQLConnector) DeleteFromTableWhere(table string, whereClause map[string]string) (int, error) {
	queryBuilder := "DELETE FROM " + table
	var whereValues []string
	var whereParams []interface{}
	for key, value := range whereClause {
		whereValues = append(whereValues, fmt.Sprintf("%s=?", key))
		whereParams = append(whereParams, value)
	}
	if len(whereValues) > 0 {
		queryBuilder += fmt.Sprintf(" WHERE %s", Join(whereValues, " AND "))
	}
	result, err := databasePostgreSQLConnector.db.Exec(queryBuilder, whereParams...)
	if err != nil {
		return 0, fmt.Errorf("errore durante l'esecuzione della query di eliminazione: %v", err)
	}
	numRowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("errore durante l'ottenimento del numero di righe interessate: %v", err)
	}
	return int(numRowsAffected), nil
}
