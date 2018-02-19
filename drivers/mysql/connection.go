package mysql

import (
	"database/sql"

	"github.com/eaciit/dbflex"
	"github.com/eaciit/toolkit"

	"github.com/eaciit/dbflex/drivers/rdbms"
	_ "github.com/go-sql-driver/mysql"
)

// Connection implementation of dbflex.IConnection
type Connection struct {
	rdbms.Connection
	db *sql.DB
}

func init() {
	dbflex.RegisterDriver("mysql", func(si *dbflex.ServerInfo) dbflex.IConnection {
		c := new(Connection)
		c.ServerInfo = *si
		return c
	})
}

// Connect to database instance
func (c *Connection) Connect() error {
	sqlconnstring := toolkit.Sprintf("tcp(%s:%d)/%s", c.Host, c.Port, c.Database)
	if c.User != "" {
		sqlconnstring = toolkit.Sprintf("%s:%s@%s", c.User, c.Password, sqlconnstring)
	}
	db, err := sql.Open("mysql", sqlconnstring)
	c.db = db
	return err
}

// Close database connection
func (c *Connection) Close() {
	if c.db != nil {
		c.db.Close()
	}
}

// NewQuery generates new query object to perform query action
func (c *Connection) NewQuery() dbflex.IQuery {
	q := new(Query)
	q.SetThis(q)
	q.db = c.db
	return q
}
