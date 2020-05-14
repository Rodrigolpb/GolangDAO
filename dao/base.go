package dao

import (
	"database/sql"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type baseDAO struct {
	db *sql.DB
}

func toSnakeCase(s string) string {
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")
	return strings.ToLower(matchAllCap.ReplaceAllString(s, "${1}_${2}"))
}

func getEntityProperties(object interface{}, withID bool) ([]map[string]reflect.Value, []string, []interface{}) {
	properties := []map[string]reflect.Value{}
	keys := []string{}
	values := make([]interface{}, 0, 0)

	objectType := reflect.TypeOf(object)
	for i := 0; i < objectType.NumField(); i++ {
		fieldName := toSnakeCase(objectType.Field(i).Name)
		if withID || fieldName != "id" {
			field := reflect.ValueOf(object).Field(i)
			properties = append(properties, map[string]reflect.Value{
				fieldName: field,
			})
			keys = append(keys, fieldName)
			values = append(values, field.Interface())
		}
	}

	return properties, keys, values
}

// Create - Executes insert operation into database and returns the number of affected rows
func (dao *baseDAO) Create(tableName string, object interface{}) (int64, error) {

	_, keys, values := getEntityProperties(object, false)
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
