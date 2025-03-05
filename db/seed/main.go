package main

import (
	"database/sql"
	"fmt"
	"library-mngmt/app/config"
	"library-mngmt/domain/entities"
	"library-mngmt/domain/globals"
	"library-mngmt/pkg"
	"log"

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

	cfg := config.Parse("/config")

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Database)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
		return
	}
	fmt.Println("Connected to the database!")

	// add admin details
	text, err := globals.GenerateAPIKey()
	if err != nil {
		log.Fatal(err)
	}
	apiKey, err := pkg.Encrypt(text, cfg.Service.EncrytionKey)
	if err != nil {
		log.Fatal(err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	users := []entities.User{
		{
			Name:     "member",
			Email:    "member@gmail.com",
			Password: string(hashedPassword),
			RoleID:   3,
			APIKey:   apiKey,
		},
		{
			Name:     "admin",
			Email:    "admin@gmail.com",
			Password: string(hashedPassword),
			RoleID:   1,
			APIKey:   apiKey,
		},
		{Name: "librarian",
			Email:    "lib@gmail.com",
			Password: string(hashedPassword),
			RoleID:   2,
			APIKey:   apiKey,
		},
	}
	// // insert user details
	for _, user := range users {
		query := `INSERT INTO users (name, email, password, role_id, api_key) VALUES ($1, $2, $3, $4, $5)`
		_, err = db.Exec(query, user.Name, user.Email, user.Password, user.RoleID, user.APIKey)
		if err != nil {
			log.Fatalf("Error inserting user: %v", err)
		}
	}
	// insert permissions
	permissions := []entities.Permission{
		{Action: "user"},
		{Action: "book"},
		{Action: "permission"},
	}
	for _, permission := range permissions {
		query := `INSERT INTO permissions (action) VALUES ($1)`
		_, err = db.Exec(query, permission.Action)
		if err != nil {
			log.Fatalf("Error inserting permission: %v", err)
		}
	}
	// insert role permission

	rolePermission := []entities.RolePermission{
		{RoleID: 1, PermissionID: 1},
		{RoleID: 1, PermissionID: 2},
		{RoleID: 1, PermissionID: 3},
		{RoleID: 2, PermissionID: 1},
		{RoleID: 2, PermissionID: 2},
		{RoleID: 3, PermissionID: 1},
	}
	for _, rolePermission := range rolePermission {
		query := `INSERT INTO role_permissions (role_id, permission_id) VALUES ($1, $2)`
		_, err = db.Exec(query, rolePermission.RoleID, rolePermission.PermissionID)
		if err != nil {
			log.Fatalf("Error inserting role permission: %v", err)
		}
	}

	fmt.Println("added successfully!")
}
