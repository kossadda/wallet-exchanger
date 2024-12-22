package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kossadda/wallet-exchanger/share/pkg/configs"
	_ "github.com/lib/pq"
)

// PostgresDB represents a PostgreSQL database connection.
type PostgresDB struct {
	db *sqlx.DB
}

// New creates a new PostgreSQL database connection using the provided config.
func New(cfg *configs.ConfigDB) (*PostgresDB, error) {
	if cfg.DBPort == "" {
		cfg.DBPort = configs.DefaultPostgresPort
	}

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword, cfg.DBSSLMode))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDB{db: db}, nil
}

// Transaction starts a new transaction, executes the provided function, and commits or rolls back.
func (p *PostgresDB) Transaction(fn func(tx *sqlx.Tx) error) error {
	tx, err := p.db.Beginx()
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("transaction rollback error: %v, original error: %v", rollbackErr, err)
		}
		return err
	}

	return tx.Commit()
}

// Close closes the PostgreSQL database connection.
func (p *PostgresDB) Close() error {
	return p.db.Close()
}
