package main

import (
	"database/sql"
	"fmt"
	"library-mngmt/domain/entities"
	"library-mngmt/domain/globals"
	"library-mngmt/pkg"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type Config struct {
	Driver   string
	Name     string
	Host     string
	User     string
	Password string
}

func main() {
	if err := godotenv.Load("../../.env.local"); err != nil {
		log.Fatalf("Error loading .env.local file")
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	fmt.Println("Connected to the database!")

	// add admin details
	text, err := globals.GenerateAPIKey()
	if err != nil {
		log.Fatal(err)
	}
	apiKey, err := pkg.Encrypt(text, os.Getenv("ENCRYPTION_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	user := entities.User{
		Name:     "admin",
		Email:    "admin@gmail.com",
		Password: string(hashedPassword),
		RoleID:   1,
		APIKey:   apiKey,
	}
	query := `INSERT INTO users (name, email, password, role_id, api_key) VALUES ($1, $2, $3, $4, $5)`

	result, err := db.Exec(query, user.Name, user.Email, user.Password, user.RoleID, user.APIKey)
	if err != nil {
		log.Fatalf("Error inserting user: %v", err)
	}
	fmt.Println(result)
	fmt.Println("Admin added successfully!")
}
