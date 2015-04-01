/*
Copyright 2014 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// outyet is a web server that announces whether or not a particular Go version
// has been tagged.
package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"flag"
	"log"
	"encoding/json"
	"fmt"
)

// Command-line flags.
var (
	user			= flag.String("username", 			"", "User Name")
	password	= flag.String("password",				"", "Password")
	ip				= flag.String("ip", 						"", "IP Address")
	port			= flag.String("port", 					"", "Port")
	dbName		= flag.String("databasename", 	"", "Database Name")
	query			= flag.String("query", 					"", "Query")
)

func main() {
	flag.Parse()

	log.Println(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", *user, *password, *ip, *port, *dbName))
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", *user, *password, *ip, *port, *dbName))
	if err != nil {
		fmt.Println(getJsonError(err))
	}
	defer db.Close()

	result, err := getJSON(*query, db)
	if err != nil {
		fmt.Println(getJsonError(err))
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
  tableData := make([]map[string]interface{}, 0)
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
  jsonData, err := json.Marshal(tableData)
  if err != nil {
      return "", err
  }
  return string(jsonData), nil
}

func getJsonError(myError error) string {

	errorJson := make(map[string]interface{})
	errorJson["error"] = myError.Error()
  jsonData, err := json.Marshal(errorJson)
  if err != nil {
      return "{\"error\": \"There was an error generatoring the error.. goodluck\"}"
  }
  return string(jsonData)
}
