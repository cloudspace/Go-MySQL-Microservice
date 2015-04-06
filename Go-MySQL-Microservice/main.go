package main // import "github.com/cloudspace/Go-MySQL-Microservice"

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println(errorStringAsJSON(fmt.Sprintf("Must have 3 arguments: your are passing %v arguments", len(os.Args)-1)))
		return
	}

	connectionURI := os.Args[1]
	password := os.Args[2]
	query := os.Args[3]

	db, err := sql.Open("mysql", fmt.Sprintf(connectionURI, password))
	if err != nil {
		fmt.Println(getJSONError(err))
	}
	defer db.Close()

	result, err := getJSON(query, db)
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
	result := make(map[string]interface{}, 0)
	result["result"] = tableData
	result["error"] = ""
	return asJSON(result), nil
}

func asJSON(anything interface{}) string {

	jsonData, err := json.Marshal(anything)
	if err != nil {
		return getJSONError(err)
	}
	return string(jsonData)
}

func getJSONError(myError error) string {

	errorJSON := make(map[string]interface{})
	errorJSON["error"] = myError.Error()
	jsonData, err := json.Marshal(errorJSON)
	if err != nil {
		return errorStringAsJSON("There was an error generatoring the error.. goodluck")
	}
	return string(jsonData)
}

func errorStringAsJSON(errorString string) string {

	return "{\"result\": \"\"\n\"error\": \"" + errorString + "\"}"
}
