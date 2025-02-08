package database

import (
	"log"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var(
	queries *Queries
	once sync.Once
	pool *pgxpool.Pool
)

func Initialize(conn *pgxpool.Pool){

	once.Do(func(){
		pool = conn
		queries = New(pool)
	})
}

func GetQueries()*Queries{
	if queries == nil{
		log.Fatal("Database not initialized. Call db.Initialize() first.")
	}
	return queries
}

func GetDBPool() *pgxpool.Pool {
	if pool == nil {
		log.Fatal("Database not initialized. Call db.Initialize() first.")
	}
	return pool
}