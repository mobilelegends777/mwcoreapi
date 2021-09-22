package mwcoreapi

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProviderTransaction struct {
	Amount          float64 `json:"amount,omitempty"  db:"amount,omitempty"`
	ProvTransId     int64   `json:"prov_trans_id,omitempty"  db:"prov_trans_id,omitempty"`
	ProviderName    string  `json:"provider_name,omitempty" db:"provider_name,omitempty"`
	ProviderTransId string  `json:"provider_trans_id,omitempty"  db:"provider_trans_id,omitempty"`
	ProviderRoundId string  `json:"provider_round_id,omitempty"  db:"provider_round_id,omitempty"`
	RoundId         int64   `json:"round_id,omitempty"  db:"round_id,omitempty"`
	TransId         int64   `json:"trans_id,omitempty"  db:"trans_id,omitempty"`
	TokenUUID       string  `json:"token_uuid,omitempty"  db:"token_uuid,omitempty"`
	TransType       int     `json:"trans_type,omitempty"  db:"trans_type,omitempty"`
	SchemaName      string  `json:"schema_name,omitempty"`
}

type AllProviderTransactionResponseBodyItems struct {
	Amount          float64 `json:"amount,omitempty"  db:"amount,omitempty"`
	ProvTransId     int64   `json:"prov_trans_id,omitempty"  db:"prov_trans_id,omitempty"`
	ProviderName    string  `json:"provider_name,omitempty" db:"provider_name,omitempty"`
	ProviderTransId string  `json:"provider_trans_id,omitempty"  db:"provider_trans_id,omitempty"`
	ProviderRoundId string  `json:"provider_round_id,omitempty"  db:"provider_round_id,omitempty"`
	RoundId         int64   `json:"round_id,omitempty"  db:"round_id,omitempty"`
	TransId         int64   `json:"trans_id,omitempty"  db:"trans_id,omitempty"`
	TokenUUID       string  `json:"token_uuid,omitempty"  db:"token_uuid,omitempty"`
	TransType       int     `json:"trans_type,omitempty"  db:"trans_type,omitempty"`
}

type ProviderTransactionRepository interface {
	New() *ProviderTransaction
	CreateProviderTransaction(data ProviderTransaction) (int64, error)
	UpdateProviderTransaction(data ProviderTransaction) (*ProviderTransaction, error)
	GetProviderTransactionByProviderRoundId(data ProviderTransaction) (*ProviderTransaction, error)
	GetProviderTransactionByProviderRoundIdAndTransType(data ProviderTransaction) (*ProviderTransaction, error)
	GetProviderTransactionByProviderTransId(data ProviderTransaction) (*ProviderTransaction, error)
	GetProviderTransactionByProviderTransIdAndTransType(data ProviderTransaction) (*ProviderTransaction, error)
	GetAllProviderTransactionByProviderRoundId(data ProviderTransaction) (*AllProviderTransactionResponseBodyItems, error)
}
type ProviderTransactionRepo struct {
	db    *sqlx.DB
	cache *Cache
}

func NewProviderTransactionRepo(db *sqlx.DB, cache *Cache) *ProviderTransactionRepo {
	return &ProviderTransactionRepo{
		db:    db,
		cache: cache,
	}
}
func (rcvr *ProviderTransactionRepo) New() *ProviderTransaction {
	return &ProviderTransaction{}
}

/**
 * @return [int64][prov_trans_id]
 *
 **/
func (rcvr *ProviderTransactionRepo) CreateProviderTransaction(data ProviderTransaction) (int64, error) {
	tx := rcvr.db.MustBegin()
	if err := rcvr.cache.GetKey("provider_transaction_"+fmt.Sprint(data.ProviderName)+":"+fmt.Sprint(data.ProviderName)+"_"+fmt.Sprint(data.ProviderTransId), fmt.Sprint(data.ProviderName)+"_"+fmt.Sprint(data.ProviderTransId)); err == nil {
		// if key exist on redis throw duplicate
		err := errors.New("duplicate unique provider id already exist" + "provider_transaction_" + fmt.Sprint(data.ProviderName) + ":" + fmt.Sprint(data.ProviderName) + "_" + fmt.Sprint(data.ProviderTransId))
		//logger.Debug(err.Error(), "CreateProviderTransaction")
		return 0, err
	}
	res, err := rcvr.db.NamedExec("INSERT INTO `mwapiv2_provider_transaction`.`provider_transaction_"+fmt.Sprintf(data.ProviderName)+"` (provider_round_id,provider_trans_id,token_uuid,trans_id,round_id, trans_type,amount) VALUES (:provider_round_id,:provider_trans_id,:token_uuid,:trans_id,:round_id,:trans_type,:amount)", &data)
	if err != nil {
		// if err not nil means duplicate unique constraint column provider_trans_id
		//logger.InternalError(err.Error(), "CreateProviderTransaction")
		return 0, err
	}
	tx.Commit()
	transactionId, _ := res.LastInsertId()
	rcvr.cache.SetKey("provider_transaction_"+fmt.Sprint(data.ProviderName)+":"+fmt.Sprint(data.ProviderName)+"_"+fmt.Sprint(data.ProviderTransId), fmt.Sprint(data.ProviderName)+"_"+fmt.Sprint(data.ProviderTransId))
	return transactionId, nil

}
func (rcvr *ProviderTransactionRepo) UpdateProviderTransaction(data ProviderTransaction) (*ProviderTransaction, error) {
	Details := &ProviderTransaction{}
	tx := rcvr.db.MustBegin()
	defer tx.Commit()
	// fmt.Println(&data)
	_, err := rcvr.db.NamedExec("UPDATE "+fmt.Sprintf(data.SchemaName)+"."+"`provider_transaction_"+fmt.Sprintf(data.ProviderName)+"` SET round_id=:round_id,token_id=:token_id,trans_id=:trans_id WHERE  prov_trans_id =:prov_trans_id", &data)
	if err != nil {
		//logger.InternalError(err.Error(), "UpdateProviderTransaction")
		return Details, err
	}
	return Details, nil
}

func (rcvr *ProviderTransactionRepo) GetProviderTransactionByProviderRoundId(data ProviderTransaction) (*ProviderTransaction, error) {
	Details := &ProviderTransaction{}
	tx := rcvr.db.MustBegin()
	defer tx.Commit()
	err := rcvr.db.Get(Details, "SELECT prov_trans_id,round_id, provider_round_id,provider_trans_id, trans_id, token_uuid,trans_type,amount FROM "+fmt.Sprintf(data.SchemaName)+"."+"`provider_transaction_"+fmt.Sprintf(data.ProviderName)+"` WHERE provider_round_id = ?", data.ProviderRoundId)
	//fmt.Println("SELECT prov_trans_id,round_id, provider_round_id,provider_trans_id, trans_id, token_id FROM `mwapiv2_provider_transaction`.`provider_transaction_" + fmt.Sprintf(data.ProviderName) + "` WHERE provider_round_id = ?")
	if err != nil {
		//logger.InternalError(err.Error(), "GetProviderTransactionByProviderRoundId")
		return Details, err
	}
	return Details, nil
}
func (rcvr *ProviderTransactionRepo) GetProviderTransactionByProviderRoundIdAndTransType(data ProviderTransaction) (*ProviderTransaction, error) {
	Details := &ProviderTransaction{}
	tx := rcvr.db.MustBegin()
	defer tx.Commit()
	err := rcvr.db.Get(Details, "SELECT prov_trans_id,round_id, provider_round_id,provider_trans_id, trans_id, token_uuid,trans_type,amount FROM "+fmt.Sprintf(data.SchemaName)+"."+"`provider_transaction_"+fmt.Sprintf(data.ProviderName)+"` WHERE provider_round_id = ? and trans_type = ?", data.ProviderRoundId, data.TransType)
	//fmt.Println("SELECT prov_trans_id,round_id, provider_round_id,provider_trans_id, trans_id, token_id FROM `mwapiv2_provider_transaction`.`provider_transaction_" + fmt.Sprintf(data.ProviderName) + "` WHERE provider_round_id = ?")
	if err != nil {
		// logger.InternalError(err.Error(), "GetProviderTransactionByProviderRoundId")
		return Details, err
	}
	return Details, nil
}
func (rcvr *ProviderTransactionRepo) GetProviderTransactionByProviderTransId(data ProviderTransaction) (*ProviderTransaction, error) {
	Details := &ProviderTransaction{}
	tx := rcvr.db.MustBegin()
	defer tx.Commit()
	err := rcvr.db.Get(Details, "SELECT prov_trans_id,round_id, provider_round_id,provider_trans_id, trans_id, token_uuid,trans_type,amount FROM "+fmt.Sprintf(data.SchemaName)+"."+"`provider_transaction_"+fmt.Sprintf(data.ProviderName)+"` WHERE provider_trans_id = ?", data.ProviderTransId)
	//fmt.Println("SELECT prov_trans_id,round_id, provider_round_id,provider_trans_id, trans_id, token_id FROM `mwapiv2_provider_transaction`.`provider_transaction_" + fmt.Sprintf(data.ProviderName) + "` WHERE provider_trans_id_id = ?")
	if err != nil {
		//logger.InternalError(err.Error(), "GetProviderTransactionByProviderTransId")
		return Details, err
	}
	return Details, nil
}
func (rcvr *ProviderTransactionRepo) GetProviderTransactionByProviderTransIdAndTransType(data ProviderTransaction) (*ProviderTransaction, error) {
	Details := &ProviderTransaction{}
	tx := rcvr.db.MustBegin()
	defer tx.Commit()
	err := rcvr.db.Get(Details, "SELECT prov_trans_id,round_id, provider_round_id,provider_trans_id, trans_id, token_uuid,trans_type,amount FROM "+fmt.Sprintf(data.SchemaName)+"."+"`provider_transaction_"+fmt.Sprintf(data.ProviderName)+"` WHERE provider_trans_id = ? and trans_type = ?", data.ProviderTransId, data.TransType)
	//fmt.Println("SELECT prov_trans_id,round_id, provider_round_id,provider_trans_id, trans_id, token_id FROM `mwapiv2_provider_transaction`.`provider_transaction_" + fmt.Sprintf(data.ProviderName) + "` WHERE provider_trans_id_id = ?")
	if err != nil {
		// logger.InternalError(err.Error(), "GetProviderTransactionByProviderTransId")
		return Details, err
	}
	return Details, nil
}
func (rcvr *ProviderTransactionRepo) GetAllProviderTransactionByProviderRoundId(data ProviderTransaction) ([]AllProviderTransactionResponseBodyItems, error) {
	var ProviderBag []AllProviderTransactionResponseBodyItems
	ProviderDatas := &AllProviderTransactionResponseBodyItems{}
	tx := rcvr.db.MustBegin()
	defer tx.Commit()
	provider_data, err := rcvr.db.Queryx("SELECT prov_trans_id,round_id, provider_round_id,provider_trans_id, trans_id, token_uuid,trans_type,amount FROM " + fmt.Sprintf(data.SchemaName) + "." + "provider_transaction_" + fmt.Sprintf(data.ProviderName) + " WHERE provider_round_id = " + "'" + fmt.Sprintf(data.ProviderRoundId) + "'" + "")
	fmt.Println("SELECT prov_trans_id,round_id, provider_round_id,provider_trans_id, trans_id, token_uuid,trans_type FROM " + fmt.Sprintf(data.SchemaName) + "." + "provider_transaction_" + fmt.Sprintf(data.ProviderName) + " WHERE provider_round_id = " + "'" + fmt.Sprintf(data.ProviderRoundId) + "'" + "")
	if err != nil {
		// log.Fatal(err)
		return ProviderBag, err
	}
	for provider_data.Next() {
		err := provider_data.StructScan(ProviderDatas)
		if err != nil {
			return ProviderBag, err
		}
		ProviderBag = append(ProviderBag, AllProviderTransactionResponseBodyItems{
			ProvTransId:     ProviderDatas.ProvTransId,
			ProviderName:    ProviderDatas.ProviderName,
			ProviderTransId: ProviderDatas.ProviderTransId,
			ProviderRoundId: ProviderDatas.ProviderRoundId,
			RoundId:         ProviderDatas.RoundId,
			Amount:          ProviderDatas.Amount,
			TransId:         ProviderDatas.TransId,
			TokenUUID:       ProviderDatas.TokenUUID,
			TransType:       ProviderDatas.TransType,
		})
	}
	if err != nil {
		// logger.InternalError(err.Error(), "GetAllProviderTransactionByProviderRoundId")
		return ProviderBag, err
	}
	return ProviderBag, nil
}
