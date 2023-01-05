package config

import (
	"database/sql"
	"fmt"
)

func NewMysqlConn(c *DatabaseConfig) *sql.DB {
	conn, err := sql.Open(c.Driver, fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:%s)/%s",
		c.Username,
		c.Password,
		c.Port,
		c.Database))
	if err != nil {
		//log.Fatalf("Can't open database connection, %v", err)
		return nil
	}
	if err := conn.Ping(); err != nil {
		return nil
	}
	return conn
}
