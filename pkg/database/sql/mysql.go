package sql

import (
	"time"

	"github.com/zdao-pro/sky_blue/pkg/log"

	// database driver
	_ "github.com/go-sql-driver/mysql"
)

// Config mysql config.
type Config struct {
	DSN          string        `yaml:"DSN"`          // write data source name.
	ReadDSN      []string      `yaml:"ReadDSN"`      // read data source name.
	Active       int           `yaml:"Active"`       // pool
	Idle         int           `yaml:"Idle"`         // pool
	IdleTimeout  time.Duration `yaml:"IdleTimeout"`  // connect max life time.
	QueryTimeout time.Duration `yaml:"QueryTimeout"` // query sql timeout
	ExecTimeout  time.Duration `yaml:"ExecTimeout"`  // execute sql timeout
	TranTimeout  time.Duration `yaml:"TranTimeout"`  // transaction sql timeout
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) (db *DB) {
	if c.QueryTimeout == 0 || c.ExecTimeout == 0 || c.TranTimeout == 0 {
		panic("mysql must be set query/execute/transction timeout")
	}
	db, err := Open(c)
	if err != nil {
		log.Error("open mysql error(%v)", err)
		panic(err)
	}
	return
}
