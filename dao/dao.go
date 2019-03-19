package dao

import (
	"database/sql"
	"sync"
)

var daoBufferPool = sync.Pool{
	New: func() interface{} {
		return new(daoBuffer)
	},
}

func (buf *daoBuffer) reset() {
	buf.tx = nil
	buf.stmt = nil
	buf.res = nil
}

func acquireDaoBuffer() *daoBuffer {
	return daoBufferPool.Get().(*daoBuffer)
}

func releaseDaoBuffer(buf *daoBuffer) {
	buf.stmt.Close()
	buf.reset()

	daoBufferPool.Put(buf)
}

// New create new database access object
func New(driver, dsn string) (*Dao, error) {
	db := &Dao{Driver: driver, Dsn: dsn}

	var err error
	db.Connection, err = sql.Open(db.Driver, db.Dsn)

	return db, err
}

func (db *Dao) makeTxStmt(buf *daoBuffer, query string) error {
	var err error

	if buf.tx, err = db.Connection.Begin(); err != nil {
		return err
	}

	if buf.stmt, err = buf.tx.Prepare(query); err != nil {
		return err
	}

	return nil
}

func (db *Dao) makeStmt(buf *daoBuffer, query string) error {
	var err error

	if buf.stmt, err = db.Connection.Prepare(query); err != nil {
		return err
	}

	return nil
}

// Exec insert/update data to/from database
func (db *Dao) Exec(query string, args ...interface{}) (int64, error) {
	buf := acquireDaoBuffer()
	defer releaseDaoBuffer(buf)

	err := db.makeTxStmt(buf, query)
	if err != nil {
		return 0, err
	}

	buf.res, err = buf.stmt.Exec(args...)
	if err != nil {
		return 0, err
	}

	n, err := buf.res.RowsAffected()
	if err != nil {
		return 0, err
	}

	if err = buf.tx.Commit(); err != nil {
		return 0, err
	}

	return n, nil
}

// Query get data from database
func (db *Dao) Query(query string, args ...interface{}) (*sql.Rows, error) {
	buf := acquireDaoBuffer()
	defer releaseDaoBuffer(buf)

	if err := db.makeStmt(buf, query); err != nil {
		return nil, err
	}

	return buf.stmt.Query(args...)
}

// QueryRow get just one data from database
func (db *Dao) QueryRow(query string, args ...interface{}) (*sql.Row, error) {
	buf := acquireDaoBuffer()
	defer releaseDaoBuffer(buf)

	if err := db.makeStmt(buf, query); err != nil {
		return nil, err
	}

	return buf.stmt.QueryRow(args...), nil
}
