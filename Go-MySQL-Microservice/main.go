package main // import "github.com/cloudspace/Go-MySQL-Microservice"

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Command-line flags.
var (
	connectionURI = flag.String("connectionURI", "", "Connection URI")
	query         = flag.String("query", "", "Query")
)

func main() {
	flag.Parse()

	db, err := sql.Open("mysql", *connectionURI)
	if err != nil {
		fmt.Println(getJSONError(err))
	}
	defer db.Close()

	result, err := getJSON(*query, db)
	if err != nil {
		fmt.Println(getJSONError(err))
	}
	fmt.Println(result)
}

func getJSON(sqlString string, db *sql.DB) (string, error) {
	rows, err := db.Query(sqlString)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	var tableData []map[string]interface{}
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	result := make(map[string][]map[string]interface{}, 0)
	result["result"] = tableData
	jsonData, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func getJSONError(myError error) string {

	errorJSON := make(map[string]interface{})
	errorJSON["error"] = myError.Error()
	jsonData, err := json.Marshal(errorJSON)
	if err != nil {
		return "{\"error\": \"There was an error generatoring the error.. goodluck\"}"
	}
	return string(jsonData)
}
