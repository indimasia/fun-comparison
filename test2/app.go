package main

import (
    "database/sql"
    "fmt"
    "math/rand"
    "strings"
    "time"
    _ "github.com/go-sql-driver/mysql"
)

func randomString(length int) string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    sb := strings.Builder{}
    sb.Grow(length)
    for i := 0; i < length; i++ {
        sb.WriteByte(charset[rand.Intn(len(charset))])
    }
    return sb.String()
}

func main() {
    // Set up database configuration
    dbUser := "root"     // Replace with your database username
    dbPass := ""     // Replace with your database password
    dbHost := "localhost"         // Replace with your database host
    dbPort := "33061"              // Replace with your database port
    dbName := "comparison"     // Replace with your database name

    // Start timing
    startTime := time.Now()

    // Create the database connection string
    connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

    // Open a connection to the database
    db, err := sql.Open("mysql", connectionString)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    

    // Prepare the statement for inserting data
    stmt, err := db.Prepare("INSERT INTO employees (name, salary, greeting) VALUES (?, ?, ?)")
    if err != nil {
        panic(err.Error())
    }
    defer stmt.Close()

    // Insert 100,000 records
    for i := 0; i < 100000; i++ {
        name := randomString(10)
        salary := rand.Intn(90000) + 10000 // Random salary
        greeting := "Mr"
        if rand.Float64() < 0.5 {
            greeting = "Ms"
        }
        _, err := stmt.Exec(name, salary, greeting)
        if err != nil {
            panic(err.Error())
        }
    }
    // Prepare the statement for inserting data
    stmt2, err2 := db.Prepare("INSERT INTO employees2 (name, salary, greeting) VALUES (?, ?, ?)")
    if err2 != nil {
        panic(err2.Error())
    }
    defer stmt2.Close()

    // Insert 100,000 records
    for i := 0; i < 100000; i++ {
        name := randomString(10)
        salary := rand.Intn(90000) + 10000 // Random salary
        greeting := "Mr"
        if rand.Float64() < 0.5 {
            greeting = "Ms"
        }
        _, err2 := stmt2.Exec(name, salary, greeting)
        if err2 != nil {
            panic(err2.Error())
        }
    }

    // End timing
    endTime := time.Now()

    fmt.Printf("Go script execution time: %v seconds\n", endTime.Sub(startTime).Seconds())
}
