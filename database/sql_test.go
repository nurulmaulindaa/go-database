package database

// import (
// 	"testing"
// 	"fmt"
// 	"context"
// 	"time"
// 	"database/sql"
// 	"strconv"
// )

// func TestExecSql(t *testing.T) {
// 	db := GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()

// 	script := "INSERT INTO customer(id, name) VALUES('kyile', 'Kyile')"
// 	_, err := db.ExecContext(ctx, script)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Success insert new customer")
// }

// func TestQuerySql(t *testing.T) {
// 	db := GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()

// 	script := "SELECT id, name FROM customer"
// 	rows, err := db.QueryContext(ctx, script)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var id, name string
// 		err = rows.Scan(&id, &name)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println("Id", id)
// 		fmt.Println("Name", name)
// 	}
// }

// func TestQuerySqlComplex(t *testing.T) {
// 	db := GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()

// 	script := "SELECT id, name, email, balance, rating, birth_date, merried, created_at FROM customer"
// 	rows, err := db.QueryContext(ctx, script)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var id, name string
// 		var email sql.NullString
// 		var balance int32
// 		var rating float64
// 		var birthDate sql.NullTime
// 		var createdAt time.Time
// 		var merried bool

// 		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &merried ,&createdAt)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println("=================")
// 		fmt.Println("Id", id)
// 		fmt.Println("Name", name)
// 		if email.Valid {
// 		fmt.Println("Email", email.String)
// 		}
// 		fmt.Println("Balance", balance)
// 		fmt.Println("Rating", rating)
// 		if birthDate.Valid {
// 		fmt.Println("Birth Date", birthDate.Time)
// 		}
// 		fmt.Println("Merried", merried)
// 		fmt.Println("Created At", createdAt)
// 	}
// }

// func TestSqlInjection(t *testing.T)  {
// 	db := GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()

// 	username := "admin"
// 	password := "admin"

// 	script := "SELECT username FROM users WHERE username = '" + username + "' AND pass = '" + password + "' LIMIT 1"
// 	// script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
// 	// script := "SELECT * FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
// 	rows, err := db.QueryContext(ctx, script)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	if rows.Next() {
// 		var username string
// 		err := rows.Scan(&username)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println("Sukses Login", username)
// 	}else {
// 		fmt.Println("Gagal login")
// 	}
// }

// func TestSqlInjectionSafe(t *testing.T)  {
// 	db := GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()

// 	username := "admin"
// 	password := "admin"

// 	script := "SELECT username FROM users WHERE username = $1 AND pass = $2 LIMIT 1"
// 	// script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
// 	// script := "SELECT * FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
// 	rows, err := db.QueryContext(ctx, script, username, password)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	if rows.Next() {
// 		var username string
// 		err := rows.Scan(&username)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println("Sukses Login", username)
// 	}else {
// 		fmt.Println("Gagal login")
// 	}
// }

// func TestExecSqlParameter(t *testing.T) {
// 	db := GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()

// 	username := "maulinda"
// 	password := "maulinda"

// 	script := "INSERT INTO users(username, pass) VALUES($1, $2)"
// 	_, err := db.ExecContext(ctx, script, username, password)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Success insert new users")
// }

// func TestAutoIncrement(t *testing.T)  {
// 	db := GetConnection()
// 	defer db.Close()

// 	// ctx := context.Background()

// 	email := "maulinda@gmail.com"
// 	comment := "Test Comment"

// 	script := "INSERT INTO comments(email, comment) VALUES($1, $2) returning id"
// 	result, err := db.Prepare(script)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer result.Close()

// 	var insertId int
// 	err = result.QueryRow(email, comment).Scan(&insertId)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Success insert new comment with id", insertId)
// }

// func TestPrepareStatement(t *testing.T)  {
// 	db := GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()
// 	script := "INSERT INTO comments(email, comment) VALUES($1, $2) returning id"
// 	statement, err := db.PrepareContext(ctx, script)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer statement.Close()

// 	for i := 0; i < 10; i++ {
// 		email := "nurul" + strconv.Itoa(i) + "@gmail.com"
// 		comment := "Komentar ke " + strconv.Itoa(i)

// 		var insertId int
// 		err := statement.QueryRowContext(ctx, email, comment).Scan(&insertId)
// 		if err != nil {
// 			panic(err)
// 		}

// 		fmt.Println("Coment Id ", insertId)
// 	}
// }

// func TestTransaction(t *testing.T) {
// 	db := GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()
// 	tx, err := db.Begin()
// 	if err != nil {
// 		panic(err)
// 	}

// 	script := "INSERT INTO comments(email, comment) VALUES($1, $2) returning id"
// 	//do transaction
// 	for i := 0; i < 10; i++ {
// 		email := "nurul" + strconv.Itoa(i) + "@gmail.com"
// 		comment := "Komentar ke " + strconv.Itoa(i)

// 		var insertId int
// 		err := tx.QueryRowContext(ctx, script, email, comment).Scan(&insertId)
// 		if err != nil {
// 			panic(err)
// 		}

// 		fmt.Println("Coment Id ", insertId)
// 	}

// 	err = tx.Rollback()
// 	if err != nil {
// 		panic(err)
// 	}
// }
