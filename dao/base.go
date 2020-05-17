package dao

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

type baseDAO struct {
	db *sql.DB
}

func getEntityProperties(object interface{}, isInsert bool) ([]string, []interface{}) {
	keys := []string{}
	values := make([]interface{}, 0, 0)

	objectType := reflect.TypeOf(object)
	for i := 0; i < objectType.NumField(); i++ {
		field := reflect.ValueOf(object).Field(i)
		if objectType.Field(i).Tag.Get("ignoreoninsert") == "" || !isInsert {
			keys = append(keys, objectType.Field(i).Tag.Get("column"))
			values = append(values, field.Interface())
		}
	}

	return keys, values
}

// Create - Executes insert operation into database and returns the number of affected rows
func (dao *baseDAO) Create(tableName string, object interface{}) (int64, error) {

	keys, values := getEntityProperties(object, true)
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s?)", tableName, strings.Join(keys, ", "), strings.Repeat("?, ", len(keys)-1))
	stmt, _ := dao.db.Prepare(query)
	defer stmt.Close()

	res, err := stmt.Exec(values...)
	rows, _ := res.RowsAffected()
	return rows, err
}

// ReadOne - Executes select operation one specific register from database and returns the selected data
func (dao *baseDAO) ReadOne(id int32) {}

// ReadList - Executes select operation into database and returns the selected data
func (dao *baseDAO) ReadList() {}

// Update - Executes update operation into database and returns the number of affected rows
func (dao *baseDAO) Update() {}

// Delete - Executes delete operation into database and returns the number of affected rows
func (dao *baseDAO) Delete() {}
