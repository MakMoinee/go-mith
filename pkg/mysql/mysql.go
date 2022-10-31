package mysql

import (
	"database/sql"
	"log"
)

type service struct {
	DatabaseName     string
	DatabaseServer   string
	DatabaseUser     string
	DatabasePassword string
	ConnectionString string
	Db               *sql.DB
	DbDriver         string
	maxConn          int
}

type GoMithMysql interface {
	GetDBConnection() *sql.DB
	SetConnectionString(conn string, maxConn int)
}

func NewGoMithMysql(dbName, dbServer, dbUser, dbPassword, dbDriver string) GoMithMysql {
	svc := service{}
	svc.DbDriver = dbDriver
	return &svc
}

func (s *service) GetDBConnection() *sql.DB {
	db, err := sql.Open(s.DbDriver, s.ConnectionString)
	if err != nil {
		log.Println(err.Error())
	}
	db.SetMaxOpenConns(s.maxConn)
	return db
}

func (s *service) SetConnectionString(conn string, maxConn int) {
	s.ConnectionString = conn
	s.maxConn = maxConn
}
