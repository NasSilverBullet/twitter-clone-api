package interfaces

type SQLHandler interface {
	Begin() (Tx, error)
	Query(string, ...interface{}) (Row, error)
	Exec(string, ...interface{}) (Result, error)
}

type Tx interface {
	Commit() error
	Rollback() error
	Exec(string, ...interface{}) (Result, error)
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
	Err() error
}
