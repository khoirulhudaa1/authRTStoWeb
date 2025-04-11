package config

import (
	"database/sql"
	"fmt"
	"log"
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
