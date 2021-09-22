package mwcoreapi

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type GameDetails struct {
	GameCode     string `json:"game_code" db:"game_code"`
	GameName     string `json:"game_name" db:"game_name"`
	GameId       int32  `json:"game_id" db:"game_id"`
	ProviderName string `json:"provider_name" db:"provider_name"`
}

type GameDetailsRepository interface {
	GetGameDetailsByGameCode(game_code string, provider_id int32) (*GameDetails, error)
	GetGameDetailsByTokenId(token_id string, provider_id int32) (*GameDetails, error)
}
type GameDetailsRepo struct {
	db    *sqlx.DB
	cache *Cache
}

func NewGameDetailsRepo(db *sqlx.DB, cache *Cache) *GameDetailsRepo {
	return &GameDetailsRepo{
		db:    db,
		cache: cache,
	}
}

func (gdt *GameDetailsRepo) GetGameDetailsByGameCode(game_code string, provider_id int32) (*GameDetails, error) {
	GameDetails := &GameDetails{}
	if err := gdt.cache.GetKey("game_details:"+game_code, GameDetails); err == nil {
		return GameDetails, nil
	}
	err := gdt.db.Get(GameDetails, "SELECT game_id,game_code,game_name,sub_provider_name as provider_name FROM mwapiv2_main.games inner join mwapiv2_main.sub_providers sp using (sub_provider_id) WHERE game_code = '"+game_code+"' AND sp.sub_provider_id = '"+fmt.Sprint(provider_id)+"' order by sp.sub_provider_id desc")
	if err != nil {
		//logger.InternalError(err.Error(), "GetGameDetailsByGameCode")
		return nil, err
	}
	gdt.cache.SetKey("game_details:"+game_code, GameDetails)

	return GameDetails, nil
}
func (gdt *GameDetailsRepo) GetGameDetailsByTokenId(token_id string, provider_id int32) (*GameDetails, error) {
	return &GameDetails{
		GameCode:     "PROVIDER_TEST_1",
		GameName:     "Provider Spin",
		GameId:       123,
		ProviderName: "PROVIDER TEST",
	}, nil
}
