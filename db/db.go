package db

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"midepeter/devtest/config"

	"github.com/go-sql-driver/mysql"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

const migrationsDir = "./db/migrations"

type DB struct {
	config config.Config
	db     *sqlx.DB
}

func NewDB(cfg config.Config) (*DB, error) {
	m := &DB{
		config: cfg,
	}
	m.tryOpenConnection()
	err := m.migrate()
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m *DB) tryOpenConnection() {
	for {
		err := m.OpenConnection()
		if err != nil {
			fmt.Printf("cant open connection to mysql: %s", err.Error())
		} else {
			fmt.Println("mysql connection success")
			return
		}
		time.Sleep(time.Second)
	}
}

func (m *DB) OpenConnection() error {
	dbcfg := config.GetConfig()
	source := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&multiStatements=true&parseTime=true",
		dbcfg.Db.User,
		dbcfg.Db.Password,
		dbcfg.Db.Host,
		dbcfg.Db.Port,
		dbcfg.Db.Name,
	)
	fmt.Println(source)
	var err error
	m.db, err = sqlx.Connect("mysql", source)
	if err != nil {
		return err
	}
	err = m.db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (m DB) migrate() error {
	ex, err := os.Executable()
	if err != nil {
		return err
	}
	dir := filepath.Join(filepath.Dir(ex), migrationsDir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		dir = migrationsDir
		if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
			return errors.New("Migrations dir does not exist: " + dir)
		}
	}
	migrations := &migrate.FileMigrationSource{
		Dir: dir,
	}
	_, err = migrate.Exec(m.db.DB, "mysql", migrations, migrate.Up)
	return err
}

func (m DB) Insert(sql string, args ...interface{}) (id uint64, err error) {
	result, err := m.db.Exec(sql, args...)
	if err != nil {
		mErr, ok := err.(*mysql.MySQLError)
		if ok && mErr.Number == 1062 {
			return 0, errors.New("error")
		}
		return id, err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return id, err
	}
	return uint64(lastID), nil
}
