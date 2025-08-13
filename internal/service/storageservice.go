package service

import (
	"database/sql"
	"fmt"
	"log"
	"questbotAI/internal/config"

	_ "modernc.org/sqlite"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite", "../../internal/repository/database.db")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ping DataBase Succsesful")
	return db
}
func CreateTable(db *sql.DB) {
	query := fmt.Sprintf("%s %s %s", config.CreateTablePlayer, config.CreateTableMoves, config.CreateTableInventory)

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All table created")
}

func InsertUser(db *sql.DB, chatID int, mainPromt string) {
	query := config.InsertUser
	res, err := db.Exec(query, chatID, mainPromt)
	if err != nil {
		log.Fatal(err)
	}

	id, _ := res.LastInsertId()
	fmt.Printf("Add User id - %d", id)
}

func InsertMove(db *sql.DB, chatID int, description string) {
	query := config.InsertMove
	res, err := db.Exec(query, chatID, description)
	if err != nil {
		log.Fatal(err)
	}

	id, _ := res.LastInsertId()
	fmt.Printf("Add Move id - %d", id)
}

func GetMove(db *sql.DB, chatID int) ([]string, error) {
	var description []string

	query := config.GetMove

	rows, err := db.Query(query, chatID)
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var desc string
		if err := rows.Scan(&desc); err != nil {
			return nil, fmt.Errorf("scan failed: %v", err)
		}
		description = append(description, desc)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return description, nil
}

func GetMainPromt(db *sql.DB, chatID int) (string, error) {
	var description string

	quert := config.GetMainPromt

	err := db.QueryRow(quert, chatID).Scan(&description)
	if err != nil {

		if err == sql.ErrNoRows {
			return "", err
		}
		return "", fmt.Errorf("failed to fetch move: %v", err)

	}
	return description, nil
}
