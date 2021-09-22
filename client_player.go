package mwcoreapi

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ClientsPlayers struct {
	PlayerName string `json:"player_name" db:"player_name"`
}
type ClientsPlayersItem struct {
	TokenUUID              string  `json:"token_uuid" db:"token_uuid"` //need to change later for token uuid
	PlayerToken            string  `json:"player_token" db:"player_token"`
	PlayerId               int64   `json:"player_id" db:"player_id"`
	Balance                float64 `json:"balance" db:"balance"`
	OperatorID             int64   `json:"operator_id" db:"operator_id"`
	ClientAPIKey           string  `json:"client_api_Key" db:"client_api_key"`
	ClientCode             string  `json:"client_code" db:"client_code"`
	ClientAccessToken      string  `json:"client_access_token" db:"client_access_token"`
	DbConn                 string  `json:"db_conn" db:"db_conn"`
	DbSchema               string  `json:"db_schema" db:"db_schema"`
	ClientID               int64   `json:"client_id" db:"client_id"`
	DefaultCurrency        string  `json:"default_currency" db:"default_currency"`
	PlayerDetailsURL       string  `json:"player_details_url" db:"player_details_url"`
	FundTransferURL        string  `json:"fund_transfer_url" db:"fund_transfer_url"`
	DebitCreditTransferURL string  `json:"debit_credit_transfer_url" db:"debit_credit_transfer_url"`
	TransactionCheckerURL  string  `json:"transaction_checker_url" db:"transaction_checker_url"`
	BalanceURL             string  `json:"balance_url" db:"balance_url"`
	ClientPlayerID         string  `json:"client_player_id" db:"client_player_id"`
	Username               string  `json:"username" db:"username"`
	TestPlayer             bool    `json:"test_player" db:"test_player"`
}
type ClientsPlayersRepository interface {
	GetDetailsByToken(token_uuid string) (*ClientsPlayersItem, error)
	GetDetailsByUserId(user_id int64) (*ClientsPlayersItem, error)
}

type ClientPlayerRepo struct {
	db    *sqlx.DB
	cache *Cache
}

func NewClientPlayerRepo(db *sqlx.DB) *ClientPlayerRepo {
	return &ClientPlayerRepo{
		db: db,
	}
}
func (cp *ClientPlayerRepo) GetDetailsByToken(token_id string) (*ClientsPlayersItem, error) {
	clientPlayerData := &ClientsPlayersItem{}
	qry := "select player_id,token_uuid,player_token,balance,db_schema,o.operator_id,client_api_key,client_code,client_access_token,db_conn,c.client_id,default_currency,player_details_url,fund_transfer_url,debit_credit_transfer_url,transaction_checker_url,balance_url,client_player_id,username,test_player	from (select token_uuid, player_token, balance, operator_id, client_id, player_id from mwapiv2_main.player_session_tokens where token_uuid = ? order by token_uuid desc limit 1) as pst INNER JOIN mwapiv2_main.operator o USING (operator_id) INNER JOIN mwapiv2_main.clients c USING (client_id) INNER JOIN mwapiv2_main.players p USING (player_id)"
	if err := cp.cache.GetKey("client_player:"+token_id, clientPlayerData); err == nil {
		return clientPlayerData, nil
	}
	err := cp.db.Get(clientPlayerData, qry, token_id)
	if err != nil {
		return nil, err
	}
	cp.cache.SetKey("client_player:"+token_id, clientPlayerData)
	return clientPlayerData, nil
}

func (cp *ClientPlayerRepo) GetDetailsByUserId(user_id int64) (*ClientsPlayersItem, error) {
	clientPlayerData := &ClientsPlayersItem{}
	qry := "select player_id,token_uuid,player_token,balance,o.operator_id,client_api_key,client_code,client_access_token,db_conn,c.client_id,default_currency,player_details_url,fund_transfer_url,debit_credit_transfer_url,transaction_checker_url,client_player_id,username,test_player	from (select token_uuid, player_token, balance, operator_id, client_id, player_id from player_session_tokens where player_id = ? order by token_uuid desc limit 1) as pst INNER JOIN operator o USING (operator_id) INNER JOIN clients c USING (client_id) INNER JOIN players p USING (player_id)"
	if err := cp.cache.GetKey("client_player:"+fmt.Sprint(user_id), clientPlayerData); err == nil {
		return clientPlayerData, nil
	}
	err := cp.db.Get(clientPlayerData, qry, user_id)
	if err != nil {
		// logger.InternalError(err.Error(), "GetDetailsByUserId")
		return nil, err
	}
	cp.cache.SetKey("client_player:"+fmt.Sprint(user_id), clientPlayerData)

	return clientPlayerData, nil
}
