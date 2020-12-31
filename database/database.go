package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var dbRef *sql.DB

func Connect(dbDialect string, dbStr string) {
	// Create the database handle, confirm driver is present
	// "<username>:<password>@tcp(<AWSConnectionEndpoint >:<port>)/<dbname>"
	db, _ := sql.Open(dbDialect, dbStr)
	dbRef = db
	// defer db.Close()

	// Connect and check the server version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
}

func GetAll(tableName string) []interface{} {
	// Execute the query
	rows, err := dbRef.Query("SELECT * FROM " + tableName)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var jsonResults []interface{}
	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value interface{}
		var jsonMap map[string]interface{} = make(map[string]interface{}, len(columns))
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				if num, ok := strconv.ParseInt(string(col), 10, 64); ok == nil {
					value = num
				} else if float, ok := strconv.ParseFloat(string(col), 64); ok == nil {
					value = float
				} else if boolean, ok := strconv.ParseBool(string(col)); ok == nil {
					value = boolean
				} else {
					value = string(col)
				}
			}
			jsonMap[columns[i]] = value
		}

		jsonResults = append(jsonResults, jsonMap)
	}

	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	return jsonResults
}

func Insert(tableName string, data map[string]interface{}) {
	// INSERT INTO `go_database`.`users` (`name`) VALUES ('Kris')
	var query string = "INSERT INTO `" + tableName + "` "
	var columns []string
	var values []string

	for key, value := range data {
		columns = append(columns, "`"+key+"`")
		if _, ok := value.(string); ok {
			values = append(values, "'"+fmt.Sprintf("%v", value)+"'")
		} else {
			values = append(values, fmt.Sprintf("%v", value))
		}
	}

	query += "(" + strings.Join(columns, ",") + ")" + " VALUES " + "(" + strings.Join(values, ",") + ")"
	fmt.Println(query)
	rows, err := dbRef.Query(query)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Println(rows)
}
