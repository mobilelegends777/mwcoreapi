package mwcoreapi

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// type GameTransactionDetails struct {
// 	RoundId   int64  `json:"round_id,omitempty"    db:"round_id, omitempty"`
// 	TransID   int64  `json:"trans_id,omitempty"    db:"trans_id,omitempty"`
// 	TransType string `json:"trans_type,omitempty"  db:"trans_type,omitempty"`
// }
type GameTransaction struct {
	RoundId    int64  `json:"round_id,omitempty" db:"round_id"`
	TransID    int64  `json:"trans_id,omitempty"  db:"trans_id"`
	TransType  int    `json:"trans_type,omitempty"  db:"trans_type,omitempty"`
	SchemaName string `json:"schema_name,omitempty"`
}
type GameTransactionRepository interface {
	New() *GameTransaction
	CreateGameTransaction(data GameTransaction) (int64, error)
	// GetGameTransactionByRoundId(round_id int64) (*GameTransactionDetails, error)
	GetGameTransactionByRoundId(data GameTransaction) (*GameTransaction, error)
	UpdateGameTransaction(data GameTransaction) (*GameTransaction, error)
}

type GameTransactionRepo struct {
	db *sqlx.DB
}

func NewGameTransactionRepo(db *sqlx.DB) *GameTransactionRepo {
	return &GameTransactionRepo{
		db: db,
	}
}
func (ggr *GameTransactionRepo) New() *GameTransaction {
	return &GameTransaction{}
}

func (rcvr *GameTransactionRepo) CreateGameTransaction(data GameTransaction) (int64, error) {
	tx := rcvr.db.MustBegin()
	_, err := rcvr.db.NamedExec("INSERT INTO "+fmt.Sprintf(data.SchemaName)+"."+"game_transactions (trans_id,round_id,trans_type) VALUES (:trans_id,:round_id,:trans_type)", &data)
	if err != nil {
		//logger.InternalError(err.Error(), "CreateGameTransaction")
		return 0, err
	}
	tx.Commit()
	return data.TransID, nil
}

func (rcvr *GameTransactionRepo) UpdateGameTransaction(data GameTransaction) (*GameTransaction, error) {
	Details := &GameTransaction{}
	tx := rcvr.db.MustBegin()
	_, err := rcvr.db.NamedExec("UPDATE "+fmt.Sprintf(data.SchemaName)+"."+"game_transactions SET round_id=:round_id WHERE trans_id =:trans_id", &data)
	if err != nil {
		//logger.InternalError(err.Error(), "UpdateGameTransaction")
		return Details, err
	}
	tx.Commit()
	return Details, nil
}

func (rcvr *GameTransactionRepo) GetGameTransactionByRoundId(data GameTransaction) (*GameTransaction, error) {
	Details := &GameTransaction{}
	tx := rcvr.db.MustBegin()
	err := rcvr.db.Get(Details, "select gt.round_id,gt.trans_id from "+fmt.Sprintf(data.SchemaName)+"."+"game_transactions gt where gt.round_id = ?", data.RoundId)
	if err != nil {
		//logger.InternalError(err.Error(), "GetGameTransactionByRoundId")
		return Details, err
	}
	tx.Commit()
	return Details, nil
}
