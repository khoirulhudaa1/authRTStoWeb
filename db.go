package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal memuat file .env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Validasi DB_PORT harus berupa angka
	if _, err := strconv.Atoi(dbPort); err != nil {
		log.Fatalf("Port database tidak valid: %v", dbPort)
	}

	// Format DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Koneksi ke database
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Gagal membuka koneksi database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Gagal ping database:", err)
	}

	log.Println("✅ Berhasil koneksi ke database MySQL!")
}

// package main

// import (
// 	"database/sql"
// 	"fmt"

// 	// "log"
// 	"os"
// 	"strconv"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/joho/godotenv"
// )

// var DB *sql.DB

// func InitDB() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Gagal memuat file .env")
// 	}

// 	dbUser := os.Getenv("DB_USER")
// 	dbPassword := os.Getenv("DB_PASSWORD")
// 	dbName := os.Getenv("DB_NAME")
// 	dbHost := os.Getenv("DB_HOST")
// 	dbPort := os.Getenv("DB_PORT")

// 	if _, err := strconv.Atoi(dbPort); err != nil {
// 		log.Fatalf("Port database tidak valid: %v", dbPort)
// 	}

// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
// 		dbUser, dbPassword, dbHost, dbPort, dbName)

// 	DB, err = sql.Open("mysql", dsn)
// 	if err != nil {
// 		log.Fatal("Gagal membuka koneksi database:", err)
// 	}

// 	err = DB.Ping()
// 	if err != nil {
// 		log.Fatal("Gagal ping database:", err)
// 	}

// 	log.Println("✅ Berhasil koneksi ke database MySQL!")
// }

// // Register user baru, jika berhasil langsung bisa login
// func RegisterUser(username, password string) (bool, error) {
// 	// Cek apakah username sudah ada
// 	var exists string
// 	err := DB.QueryRow("SELECT username FROM users WHERE username = ?", username).Scan(&exists)
// 	if err != nil && err != sql.ErrNoRows {
// 		return false, err // Error lain
// 	}

// 	if exists != "" {
// 		return false, fmt.Errorf("❌ Username '%s' sudah terdaftar", username)
// 	}

// 	// Insert user baru
// 	_, err = DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
// 	if err != nil {
// 		return false, err
// 	}

// 	log.Printf("✅ Registrasi berhasil untuk username: %s\n", username)
// 	return true, nil
// }

// // Login user dengan username dan password
// func LoginUser(username, password string) (bool, error) {
// 	var dbUsername string
// 	err := DB.QueryRow("SELECT username FROM users WHERE username = ? AND password = ?", username, password).Scan(&dbUsername)
// 	if err == sql.ErrNoRows {
// 		return false, nil // Username atau password salah
// 	} else if err != nil {
// 		return false, err
// 	}
// 	return true, nil // Login berhasil
// }
