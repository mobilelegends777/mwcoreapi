package mwcoreapi

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type GameRound struct {
	RoundId     int64   `json:"round_id,omitempty" db:"round_id,omitempty"`
	BetAmount   float64 `json:"bet_amount,omitempty" db:"bet_amount,omitempty"`
	StatusId    int     `json:"status_id,omitempty" db:"status_id,omitempty"`
	PayAmount   float64 `json:"pay_amount,omitempty" db:"pay_amount,omitempty"`
	Income      float64 `json:"income,omitempty" db:"income,omitempty"`
	EntryId     int     `json:"entry_id,omitempty" db:"entry_id,omitempty"`
	OperatorId  int     `json:"operator_id,omitempty" db:"operator_id,omitempty"`
	ClientId    int     `json:"client_id,omitempty" db:"client_id,omitempty"`
	PlayerId    int     `json:"player_id,omitempty" db:"player_id,omitempty"`
	ProviderId  int     `json:"provider_id,omitempty" db:"provider_id,omitempty"`
	GameId      int     `json:"game_id,omitempty" db:"game_id,omitempty"`
	TransStatus int     `json:"trans_status,omitempty" db:"trans_status,omitempty"`
	SchemaName  string  `json:"schema_name,omitempty"`
}

type GameRoundRepository interface {
	New() *GameRound
	GetGameRoundByRoundId(round_id int64, schema_name string) (*GameRound, error)
	CreateGameRound(data GameRound) (int64, error)
	UpdateGameRound(data GameRound) (*GameRound, error)
	UpdateGameRoundWithBet(data GameRound) (*GameRound, error)
}

type GameRoundsRepo struct {
	db    *sqlx.DB
	cache *Cache
}

func NewGameRoundsRepo(db *sqlx.DB, cache *Cache) *GameRoundsRepo {
	return &GameRoundsRepo{
		db:    db,
		cache: cache,
	}
}
func (ggr *GameRoundsRepo) New() *GameRound {
	return &GameRound{}
}
func (ggr *GameRoundsRepo) CreateGameRound(data GameRound) (int64, error) {
	// GameDetails := &models.GameRoundDetails{}
	tx := ggr.db.MustBegin()
	_, err := ggr.db.NamedExec("INSERT INTO "+fmt.Sprintf(data.SchemaName)+"."+"game_rounds (round_id,bet_amount,status_id,pay_amount,income,entry_id,operator_id,client_id,player_id,provider_id,game_id,trans_status) VALUES (:round_id,:bet_amount,:status_id,:pay_amount,:income,:entry_id,:operator_id,:client_id,:player_id,:provider_id,:game_id,:trans_status)", &data)
	if err != nil {
		//logger.InternalError(err.Error(), "CreateGameRound")
		return 0, err
	}
	tx.Commit()
	// roundId, _ := res.LastInsertId()
	return data.RoundId, nil
}

func (ggr *GameRoundsRepo) UpdateGameRound(data GameRound) (*GameRound, error) {
	// fmt.Println("CALLED")
	GameDetails := &GameRound{}
	tx := ggr.db.MustBegin()

	_, err := ggr.db.NamedExec("UPDATE "+fmt.Sprintf(data.SchemaName)+"."+"game_rounds SET pay_amount=:pay_amount, income=:income, entry_id=:entry_id, status_id=:status_id, trans_status=:trans_status WHERE round_id =:round_id", &data)
	if err != nil {
		//logger.InternalError(err.Error(), "UpdateGameRound")
		return GameDetails, err
	}
	tx.Commit()
	return GameDetails, nil
}

func (ggr *GameRoundsRepo) UpdateGameRoundWithBet(data GameRound) (*GameRound, error) {
	// fmt.Println("CALLED")
	GameDetails := &GameRound{}
	tx := ggr.db.MustBegin()
	defer tx.Commit()
	_, err := ggr.db.NamedExec("UPDATE "+fmt.Sprintf(data.SchemaName)+"."+"game_rounds SET bet_amount=:bet_amount,pay_amount=:pay_amount, income=:income, entry_id=:entry_id, status_id=:status_id, trans_status=:trans_status WHERE round_id =:round_id", &data)
	if err != nil {
		//logger.InternalError(err.Error(), "UpdateGameRoundWithBet")
		return GameDetails, err
	}
	return GameDetails, nil
}

func (ggr *GameRoundsRepo) GetGameRoundByRoundId(round_id int64, schema_name string) (*GameRound, error) {
	GameDetails := &GameRound{}
	tx := ggr.db.MustBegin()
	defer tx.Commit()
	err := ggr.db.Get(GameDetails, "SELECT round_id,bet_amount,status_id,pay_amount,income,entry_id,operator_id,client_id,player_id,provider_id,game_id,trans_status FROM "+schema_name+"."+"game_rounds WHERE round_id = ?", round_id)
	if err != nil {
		return GameDetails, err
	}
	return GameDetails, nil
	// if err := ggr.cache.GetKey("client_player:", 1); err == nil {
	// 	return GameDetails, nil
	// } else {
	// 	err := ggr.db.Get(GameDetails, "SELECT round_id,bet_amount,status_id,pay_amount,income,entry_id,operator_id,client_id,player_id,provider_id,game_id,trans_status FROM game_rounds WHERE round_id = ?", round_id)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		return nil, err
	// 	}
	// 	ggr.cache.SetKey("sample men", GameDetails)
	// }

	// return GameDetails, nil
}
