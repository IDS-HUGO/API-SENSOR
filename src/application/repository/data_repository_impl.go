package repository

import (
    "go-hexagonal-api/src/domain/entities"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type MySQLRepository struct {
    DB *sql.DB
}

func NewMySQLRepository(dsn string) (*MySQLRepository, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    return &MySQLRepository{DB: db}, nil
}

func (r *MySQLRepository) SaveData(data *entities.Data) error {
    query := "INSERT INTO data (message) VALUES (?)"
    _, err := r.DB.Exec(query, data.Message)
    return err
}