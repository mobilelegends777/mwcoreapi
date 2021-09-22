package mwcoreapi

import "github.com/jmoiron/sqlx"

type PlayerBalanceToken struct {
	TokenId     int     `json:"token_id,omitempty" db:"token_id"`
	Token_UUID  string  `json:"token_uuid,omitempty" db:"token_uuid"`
	OperatorId  int     `json:"operator_id,omitempty" db:"operator_id"`
	ClientId    int     `json:"client_id,omitempty" db:"client_id"`
	PlayerId    int     `json:"player_id,omitempty" db:"player_id"`
	PlayerToken string  `json:"player_token,omitempty" db:"player_token"`
	Balance     float64 `json:"balance,omitempty" db:"balance"`
	StatusId    int     `json:"status_id,omitempty" db:"status_id"`
}

type PlayerBalance struct {
	PlayerId   int     `json:"player_id,omitempty" db:"player_id"`
	Balance    float64 `json:"balance,omitempty" db:"balance"`
	OperatorId string  `json:"operator_id,omitempty" db:"operator_id"`
	ClientId   int     `json:"client_id,omitempty" db:"client_id"`
}

type PlayerBalanceRepository interface {
	NewPlayerBalance() *PlayerBalance
	NewPlayerBalanceToken() *PlayerBalanceToken
	UpdatePlayerTokenBalance(data PlayerBalanceToken) (*PlayerBalanceToken, error)
	UpdatePlayerBalance(data PlayerBalance) (*PlayerBalance, error)
}

type PlayerBalanceRepo struct {
	db *sqlx.DB
}

func NewPlayerBalanceRepo(db *sqlx.DB) *PlayerBalanceRepo {
	return &PlayerBalanceRepo{
		db: db,
	}
}

func (ggr *PlayerBalanceRepo) NewPlayerBalance() *PlayerBalance {
	return &PlayerBalance{}
}
func (ggr *PlayerBalanceRepo) NewPlayerBalanceToken() *PlayerBalanceToken {
	return &PlayerBalanceToken{}
}
func (rcvr *PlayerBalanceRepo) UpdatePlayerTokenBalance(data PlayerBalanceToken) (*PlayerBalanceToken, error) {
	Details := &PlayerBalanceToken{}
	tx := rcvr.db.MustBegin()
	defer tx.Commit()
	_, err := rcvr.db.NamedExec("UPDATE mwapiv2_main.player_session_tokens SET balance=:balance WHERE token_uuid =:token_uuid", &data)
	if err != nil {
		//logger.InternalError(err.Error(), "UpdatePlayerTokenBalance")
		return Details, err
	}
	return Details, nil
}

func (rcvr *PlayerBalanceRepo) UpdatePlayerBalance(data PlayerBalance) (*PlayerBalance, error) {
	Details := &PlayerBalance{}
	tx := rcvr.db.MustBegin()
	defer tx.Commit()
	_, err := rcvr.db.NamedExec("UPDATE mwapiv2_main.players SET balance=:balance WHERE player_id =:player_id", &data)
	if err != nil {
		// log.Fatal(err)
		//logger.InternalError(err.Error(), "UpdatePlayerTokenBalance")
		return Details, err
	}
	return Details, nil
}
