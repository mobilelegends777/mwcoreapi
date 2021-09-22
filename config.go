package mwcoreapi

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type MwCoreConfig struct {
	ClientPlayerDetails *ClientPlayerRepo
	GameDetails         *GameDetailsRepo
	GameRound           *GameRoundsRepo
	GameTransaction     *GameTransactionRepo
	PlayerBalance       *PlayerBalanceRepo
	ProviderTransaction *ProviderTransactionRepo
}

func Init(db *sqlx.DB, rdb *redis.Client) MwCoreConfig {
	return MwCoreConfig{
		ClientPlayerDetails: &ClientPlayerRepo{db: db, cache: NewCache(rdb)},
		GameDetails:         &GameDetailsRepo{db: db, cache: NewCache(rdb)},
		GameRound:           &GameRoundsRepo{db: db, cache: NewCache(rdb)},
		GameTransaction:     &GameTransactionRepo{db: db},
		PlayerBalance:       &PlayerBalanceRepo{db: db},
		ProviderTransaction: &ProviderTransactionRepo{db: db, cache: NewCache(rdb)},
	}
}
