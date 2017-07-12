package persistence

import (
	"database/sql"
	"fmt"

	"github.com/DaveBlooman/go-app/config"
	_ "github.com/lib/pq"
)

type Persistence struct {
	db *sql.DB
}

type Scanner interface {
	Scan(dest ...interface{}) error
}

func (p *Persistence) Init(dbDetails *config.Database) error {
	dbInfo := fmt.Sprintf("user=%s password=%s database=%s sslmode=%s host=%s port=%d",
		dbDetails.Username, dbDetails.Password, dbDetails.Name,
		dbDetails.SslMode, dbDetails.Host, dbDetails.Port)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}

	p.db = db

	return nil
}
