package dao

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql" //
)

// BaseDAO - Data Access Object base structure to execute CRUD operations into an SQL database
type BaseDAO struct{}

func toSnakeCase(s string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")

	s = matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	s = matchAllCap.ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(s)
}

func getEntityProperties(object interface{}, withID bool) ([]map[string]reflect.Value, []string, interface{}) {
	properties := []map[string]reflect.Value{}
	keys := []string{}
	var values []interface{}

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

	for value := range values {
		fmt.Println(reflect.TypeOf(value))
	}

	return properties, keys, values
}

// Create - Executes insert operation into database and returns the number of affected rows
func (dao *BaseDAO) Create(tableName string, object interface{}) {
	db, err := sql.Open("mysql", "root:admin@/retail_chain_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, keys, values := getEntityProperties(object, false)
	fmt.Println(values)

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s?)", tableName, strings.Join(keys, ", "), strings.Repeat("?, ", len(keys)-1))
	fmt.Println(query)

	// stmt, err := db.Prepare(query)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer stmt.Close()

	// fmt.Println(stmt)

	// _, err = stmt.Exec(values)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

// ReadOne - Executes select operation one specific register from database and returns the selected data
func (dao *BaseDAO) ReadOne(id int32) {}

// ReadList - Executes select operation into database and returns the selected data
func (dao *BaseDAO) ReadList() {}

// Update - Executes update operation into database and returns the number of affected rows
func (dao *BaseDAO) Update() {}

// Delete - Executes delete operation into database and returns the number of affected rows
func (dao *BaseDAO) Delete() {}
