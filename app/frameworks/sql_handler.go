package frameworks

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/NasSilverBullet/twitter-clone-api/app/interfaces"
	"github.com/NasSilverBullet/twitter-clone-api/app/usecases"
	_ "github.com/go-sql-driver/mysql"
)

type SQLHandler struct {
	Conn *sql.DB
}

type Tx struct {
	Tx *sql.Tx
}

type Result struct {
	Result sql.Result
}

type Row struct {
	Rows *sql.Rows
}

func NewSQLHandler(logger usecases.Logger) (interfaces.SQLHandler, error) {
	logger.Info("Start opening a database specified by its database driver")
	sqlHandler := &SQLHandler{}
	dataSourceName := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?parseTime=true&loc=Asia%%2FTokyo&multiStatements=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	conn, err := sql.Open(os.Getenv("DB_DRIVER"), dataSourceName)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	sqlHandler.Conn = conn

	logger.Info("Finish opening a database specified by its database driver")

	return sqlHandler, nil
}

func (s *SQLHandler) Begin() (interfaces.Tx, error) {
	t, err := s.Conn.Begin()

	if err != nil {
		return nil, err
	}

	tx := &Tx{}
	tx.Tx = t

	return tx, nil
}

func (s *SQLHandler) Query(query string, args ...interface{}) (interfaces.Row, error) {
	rows, err := s.Conn.Query(query, args...)

	if err != nil {
		return nil, err
	}

	row := &Row{}
	row.Rows = rows

	return row, nil
}

func (s *SQLHandler) Exec(query string, args ...interface{}) (interfaces.Result, error) {
	result, err := s.Conn.Exec(query, args...)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t Tx) Commit() error {
	err := t.Tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func (t Tx) Rollback() error {
	err := t.Tx.Rollback()

	if err != nil {
		return err
	}

	return nil
}

func (t Tx) Exec(query string, args ...interface{}) (interfaces.Result, error) {
	result, err := t.Tx.Exec(query, args...)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r Result) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r Result) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

func (r Row) Scan(value ...interface{}) error {
	return r.Rows.Scan(value...)
}

func (r Row) Next() bool {
	return r.Rows.Next()
}

func (r Row) Close() error {
	return r.Rows.Close()
}

func (r Row) Err() error {
	return r.Rows.Err()
}
