package mwcoreapi

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func TestGetDetailsByTokenID(t *testing.T) {
	mwcore := Init(ConnectDB(), NewClient())
	token := "359a63073c1842348e2aaed19aa60d30"
	clientPLayerDetails, err := mwcore.ClientPlayerDetails.GetDetailsByToken(token)
	if err != nil {
		t.Errorf("Expected ClientDetails got %v", err.Error())
	}
	t.Log(clientPLayerDetails)
}

func ConnectDB() *sqlx.DB {
	var db *sqlx.DB
	var err error
	if os.Getenv("DSN") != "" {
		db, err = sqlx.Open("mysql", os.Getenv("DSN"))
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		db, err = sqlx.Open("mysql",
			os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
		if err != nil {
			log.Println(err.Error())
		}
	}

	db.SetConnMaxLifetime(time.Millisecond * 2)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(40)
	return db
}
func NewClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.254.200:6379", //os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",                     //"os.Getenv("REDIS_PASSWORD")", // no password set
		DB:       0,                      // use default DB
	})
	return rdb
}
