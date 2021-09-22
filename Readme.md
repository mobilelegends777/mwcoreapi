# MWCOREAPI

mwcoreapi package is use to implement the data manipulation of the tigergames and it is cosume trough package
to include this package in you golang working directory you need to follow the following steps :
1. first you need to configure your git in your machine to pull using ssh connection you can follow this steps [here](https://docs.github.com/en/github/authenticating-to-github/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent).
2. you need to ask permission and add you as colaborator from the private repo.
3. you need to configure your git machine to use the ssh as link to pull not the http you can refer [here](https://stackoverflow.com/questions/27500861/whats-the-proper-way-to-go-get-a-private-repository).
    ```cmd
        $ git config --global url.git@github.com:.insteadOf https://github.com/
        $ cat ~/.gitconfig
        [url "git@github.com:"]
            insteadOf = https://github.com/
        $ go get github.com/private/repo
    ```
4. configure the goenv to use a private repo.NOTE:private repo means the repo link of the package. 
    ```cmd
        export GOPRIVATE="github.com/private/repo"
    ```
5. execute this command to include the package to you local machine.
    ```cmd
        go get github.com/janjoven/mwcoreapi
    ```
To use the package in your code just follow this steps.
1. first we need to import the package initialize the package and it take 2 parameter as an argument which is the db *sqlx.DB and the *redis.Client 
    and it will return an object type of mwcoreapi.MwcoreConfig
    ```go
        import (
            "log"
            "your_package_name/cache"
            "your_package_name/controller"
            "your_package_name/db"
            "your_package_name/helpers"
            "net/http"

            _ "github.com/go-sql-driver/mysql"
            "github.com/gorilla/mux"
            "github.com/janjoven/mwcoreapi"
            "github.com/joho/godotenv"
        )

        func main() {
            err := godotenv.Load()
            if err != nil {
                log.Fatal("Error loading .env file")
            }
            db := db.ConnectDB() // this is the connection to databse which return the reference of the *sqlx.DB
            rdb := cache.NewClient() // this is the connection to the redis server which return the reference of the *redis.Client
            mc := mwcoreapi.Init(db, rdb) // MWCOREAPI initialization and return a reference of the mwcoreapi.MwcoreConfig
            controller := controller.Handler{Mwcore: &mc} // example on how to pass the mwcoreapi object to the controller
            router := mux.NewRouter()
            router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
                helpers.WritterResponse(w, "Evolution API V2", 200)
            })
            router.HandleFunc("/api/debit", controller.Debit)
            log.Fatal(http.ListenAndServe(":8090", router))
        }
    ```
    for the full code design pattern please see [here](https://gitlab.betrnk.games/tigergamesv2providerportal/evolution/tree/develop-merge)
------
## list of available function:
- [ClientPlayerDetails](https://github.com/janjoven/mwcoreapi#clientplayerdetails)
    - [GetDetailsByToken](https://github.com/janjoven/mwcoreapi#getdetailsbytoken)
    - [GetDetailsByUserId](https://github.com/janjoven/mwcoreapi#getdetailsbyuserid)
- [GameDetails](https://github.com/janjoven/mwcoreapi#gamedetails) 
    - [GetGameDetailsByGameCode](https://github.com/janjoven/mwcoreapi#getgamedetailsbygamecode)
    - [GetGameDetailsByTokenId](https://github.com/janjoven/mwcoreapi#getgamedetailsbytokenid)
- [GameRounds](https://github.com/janjoven/mwcoreapi#gameround)
    - [CreateGameRound](https://github.com/janjoven/mwcoreapi#creategameround)
    - [UpdateGameRound](https://github.com/janjoven/mwcoreapi#updategameround)
    - [UpdateGameRoundWithBet](https://github.com/janjoven/mwcoreapi#updategameroundwithbet)
    - [GetGameRoundByRoundId](https://github.com/janjoven/mwcoreapi#getgameroundbyroundid)
- [GameTransactions](https://github.com/janjoven/mwcoreapi#gameround)
    - [CreateGameTransaction](https://github.com/janjoven/mwcoreapi#creategametransaction)
    - [UpdateGameTransaction](https://github.com/janjoven/mwcoreapi#updategametransaction)
    - [GetGameTransactionByRoundId](https://github.com/janjoven/mwcoreapi#getgametransactionbyroundid)
- [ProviderTransactions](https://github.com/janjoven/mwcoreapi#provider-transactions)
    - [CreateProviderTransaction](https://github.com/janjoven/mwcoreapi#createprovidertransaction)
    - [UpdateProviderTransaction](https://github.com/janjoven/mwcoreapi#updateprovidertransaction)
    - [GetProviderTransactionByProviderRoundId](https://github.com/janjoven/mwcoreapi#getprovidertransactionbyproviderroundid)
    - [GetProviderTransactionByProviderTransId](https://github.com/janjoven/mwcoreapi#getprovidertransactionbyprovidertransid)
- [PlayerBalance](https://github.com/janjoven/mwcoreapi#player-balance)
    - [UpdatePlayerTokenBalance](https://github.com/janjoven/mwcoreapi#updateplayertokenbalance)
    - [UpdatePlayerBalance](https://github.com/janjoven/mwcoreapi#updateplayerbalance)


----
## ClientPlayerDetails

This type is coming from the type ClientPlayerDetails and having 2 method available and return a struct of *mwcore.ClientsPlayersItem.
```go
   type ClientsPlayersItem struct {
	TokenUUID              string  `json:"token_uuid" db:"token_uuid"` 
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
```
#### GetDetailsByToken()
1. This method takes token_id string  as parameter and return an object type  of mwcore.ClientsPlayersItem.

- `func (cp *ClientPlayerRepo) GetDetailsByToken(token_id string) (*ClientsPlayersItem, error)`
#### GetDetailsByUserId()
2. This method take user_id int64 as parameter and return an object type  of mwcore.ClientsPlayersItem.
- `func (cp *ClientPlayerRepo) GetDetailsByUserId(user_id int64) (*ClientsPlayersItem, error)`

> Sample Code:
```go
    type Handler struct {
        Mwcore *mwcoreapi.MwCoreConfig
    }


    func (handler *Handler) SampleFunc(w http.ResponseWriter, r *http.Request) {
        token := "359a63073c1842348e2aaed19aa60d30" // sample token_uuid only
        clientPlayerDetails, err := handler.Mwcore.ClientPlayerDetails.GetDetailsByToken(token)
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        marshaledData ,_ := json.Marshal(clientPlayerDetails)
        fmt.Println(marshaledData)
    }
```
------
## GameDetails
This type  is coming from the type GameDetails and having 2 methods and return a type of *mwcoreapi.GameDetails.
```go
    type GameDetails struct {
	GameCode     string `json:"game_code" db:"game_code"`
	GameName     string `json:"game_name" db:"game_name"`
	GameId       int32  `json:"game_id" db:"game_id"`
	ProviderName string `json:"provider_name" db:"provider_name"`
    }
```
#### GetGameDetailsByGameCode()
1. This method takes game_code string and provider_id int32 and return an object type of *mwcoreapi.GameDetails and error.
- `func (gdt *GameDetailsRepo) GetGameDetailsByGameCode(game_code string, provider_id int32) (*GameDetails, error)`
#### GetGameDetailsByTokenId()
2. This method takes token_uuid string ,provider_id int32 and return an object type of *mwcoreapi.GameDetails and error.
- `func (gdt *GameDetailsRepo) GetGameDetailsByTokenId(token_id string, provider_id int32) (*GameDetails, error) `
> Sample Code :
```go
    type Handler struct {
        Mwcore *mwcoreapi.MwCoreConfig
    }


    func (handler *Handler) SampleFunc(w http.ResponseWriter, r *http.Request) {
        game_code = "sample_code"
        provider_id = 42
        gameDetails, err := handler.Mwcore.GameDetails.GetGameDetailsByGameCode(game_code,provider_id)
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        marshaledData ,_ := json.Marshal(gameDetails)
        fmt.Println(marshaledData)
    }
```
-------------------------------------------------------------------------------------------------------------------
## GameRound
This type is coming from the type GameRounds and having 4 methods and a type of *mwcoreqpi.GameRound.
```go
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
```
#### CreateGameRound()
1. This method takes an object struct type of mwcoreapi.GameRound as parameter  and return an error
- `func (ggr *GameRoundsRepo) CreateGameRound(data GameRound) (int64, error)`
> Sample Code:
```go
    type Handler struct {
        Mwcore *mwcoreapi.MwCoreConfig
    }


    func (handler *Handler) SampleFunc(w http.ResponseWriter, r *http.Request) {
            data := mwcoreapi.GameRound{
            RoundId:     1231211321312313213,
            BetAmount:   10.00,
            StatusId:    1,
            PayAmount:   10.00,
            Income:      10.00,
            EntryId:     10.00,
            OperatorId:  1,
            GameId:      123,
            TransStatus: 1,
            SchemaName:  "sample_schema_name",
            }// need to use the type from the package to create the GameRound Object
            _, err := handler.Mwcore.GameRound.CreateGameRound(data)
            if err != nil {
                log.Println(err.Error())
            }
            marshaledData ,_ := json.Marshal(gameDetails)
            fmt.Println(marshaledData)
    }
```
#### UpdateGameRound()
2. This method take an object struct type of mwcoreapi.GameRound as parameter and return the object updated and error
- `func (ggr *GameRoundsRepo) UpdateGameRound(data GameRound) (*GameRound, error)`
```go
    type Handler struct {
        Mwcore *mwcoreapi.MwCoreConfig
    }


    func (handler *Handler) SampleFunc(w http.ResponseWriter, r *http.Request) {
            data := mwcoreapi.GameRound{
            RoundId:     1231211321312313213,
            BetAmount:   10.00,
            StatusId:    1,
            PayAmount:   10.00,
            Income:      10.00,
            EntryId:     10.00,
            OperatorId:  1,
            GameId:      123,
            TransStatus: 1,
            SchemaName:  "sample_schema_name",
            }// need to use the type from the package to create the GameRound Object
            _, err := handler.Mwcore.GameRound.UpdateGameRound(data)
            if err != nil {
                log.Println(err.Error())
            }
            marshaledData ,_ := json.Marshal(gameDetails)
            fmt.Println(marshaledData)
    }
```
#### UpdateGameRoundWithBet()
3. This method take an object struct type of mwcoreapi.GameRound as parameter and return the object updated and error
- `func (ggr *GameRoundsRepo) UpdateGameRoundWithBet(data GameRound) (*GameRound, error)`
```go
    type Handler struct {
        Mwcore *mwcoreapi.MwCoreConfig
    }


    func (handler *Handler) SampleFunc(w http.ResponseWriter, r *http.Request) {
            data := mwcoreapi.GameRound{
            RoundId:     1231211321312313213,
            BetAmount:   10.00,
            StatusId:    1,
            PayAmount:   10.00,
            Income:      10.00,
            EntryId:     10.00,
            OperatorId:  1,
            GameId:      123,
            TransStatus: 1,
            SchemaName:  "sample_schema_name",
            }// need to use the type from the package to create the GameRound Object
            _, err := handler.Mwcore.GameRound.UpdateGameRoundWithBet(data)
            if err != nil {
                log.Println(err.Error())
            }
            marshaledData ,_ := json.Marshal(gameDetails)
            fmt.Println(marshaledData)
    }
```
#### GetGameRoundByRoundId()
4. This method take round_id int64 and schema_name string as parameter and return an object type of *mwcoreapi.GameRound and error
- `func (ggr *GameRoundsRepo) GetGameRoundByRoundId(round_id int64, schema_name string) (*GameRound, error)`
```go
    type Handler struct {
        Mwcore *mwcoreapi.MwCoreConfig
    }


    func (handler *Handler) SampleFunc(w http.ResponseWriter, r *http.Request) {
        round_id = 1231211321312313213
        schema_name = "sample_schema_name"
        gameDetails, err := handler.Mwcore.GameDetails.GetGameRoundByRoundId(round_id,schema_name)
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        marshaledData ,_ := json.Marshal(gameDetails)
        fmt.Println(marshaledData)
    }
```
-------
## GameTransaction
This type is coming from type GameTransaction and having 4 methods and a type of *mwcoreapi.GameTransaction.
```go
    type GameTransaction struct {
	RoundId    int64  `json:"round_id,omitempty" db:"round_id"`
	TransID    int64  `json:"trans_id,omitempty"  db:"trans_id"`
	TransType  int    `json:"trans_type,omitempty"  db:"trans_type,omitempty"`
	SchemaName string `json:"schema_name,omitempty"`
    }
```
#### CreateGameTransaction()
1. This method take an object struct type of mwcoreapi.GameTransaction as parameter and return the id of the created game_transaction as int64 and error
- `func (rcvr *GameTransactionRepo) CreateGameTransaction(data GameTransaction) (int64, error)`
```go
    type Handler struct {
        Mwcore *mwcoreapi.MwCoreConfig
    }


    func (handler *Handler) SampleFunc(w http.ResponseWriter, r *http.Request) {
            data := mwcoreapi.GameTransaction{
            RoundId:     1231211321312313213,
            TransId: 1231211321312313213,
            TransType : 1,
            SchemaName:  "sample_schema_name",
            }// need to use the type from the package to create the GameRound Object
            _, err := handler.Mwcore.GameTransaction.CreateGameTransaction(data)
            if err != nil {
                log.Println(err.Error())
            }
            marshaledData ,_ := json.Marshal(gameDetails)
            fmt.Println(marshaledData)
    }
```
#### UpdateGameTransaction()
2. This method take an object struct type of mwcoreapi.GameTransaction as parameter and return the id of the updated object  and error
- `func (rcvr *GameTransactionRepo) UpdateGameTransaction(data GameTransaction) (*GameTransaction, error)`
```go
    type Handler struct {
        Mwcore *mwcoreapi.MwCoreConfig
    }


    func (handler *Handler) SampleFunc(w http.ResponseWriter, r *http.Request) {
            data := mwcoreapi.GameTransaction{
            RoundId:     1231211321312313213,
            TransId: 1231211321312313213,
            TransType : 1,
            SchemaName:  "sample_schema_name",
            }// need to use the type from the package to create the GameRound Object
            _, err := handler.Mwcore.GameTransaction.UpdateGameTransaction(data)
            if err != nil {
                log.Println(err.Error())
            }
            marshaledData ,_ := json.Marshal(gameDetails)
            fmt.Println(marshaledData)
    }
```
#### GetGameTransactionByRoundId()
3. This method take an object struct type of mwcoreapi.GameTransaction as parameter and return the id of the  object reference *GameTransaction and error
- `func (rcvr *GameTransactionRepo) GetGameTransactionByRoundId(data GameTransaction) (*GameTransaction, error)`
```go
    type Handler struct {
        Mwcore *mwcoreapi.MwCoreConfig
    }


    func (handler *Handler) SampleFunc(w http.ResponseWriter, r *http.Request) {
            data := mwcoreapi.GameTransaction{
            RoundId:     1231211321312313213,
            TransId: 1231211321312313213,
            TransType : 1,
            SchemaName:  "sample_schema_name",
            }// need to use the type from the package to create the GameRound Object
            _, err := handler.Mwcore.GameTransaction.GetGameTransactionByRoundId(data)
            if err != nil {
                log.Println(err.Error())
            }
            marshaledData ,_ := json.Marshal(gameDetails)
            fmt.Println(marshaledData)
    }
```

------
## Provider Transactions

This type is coming from type ProviderTrasaction and having 4 methods and a type of *mwcoreapi.ProviderTransaction.

```go
    type ProviderTransaction struct {
        Amount          float64 `json:"amount,omitempty"  db:"amount,omitempty"`
        ProvTransId     int64  `json:"prov_trans_id,omitempty"  db:"prov_trans_id,omitempty"`
        ProviderName    string `json:"provider_name,omitempty" db:"provider_name,omitempty"`
        ProviderTransId string `json:"provider_trans_id,omitempty"  db:"provider_trans_id,omitempty"`
        ProviderRoundId string `json:"provider_round_id,omitempty"  db:"provider_round_id,omitempty"`
        RoundId         int64  `json:"round_id,omitempty"  db:"round_id,omitempty"`
        TransId         int64  `json:"trans_id,omitempty"  db:"trans_id,omitempty"`
        TokenUUID       string `json:"token_uuid,omitempty"  db:"token_uuid,omitempty"`
        TransType       int    `json:"trans_type,omitempty"  db:"trans_type,omitempty"`
        SchemaName      string `json:"schema_name,omitempty"`
    }
```
#### CreateProviderTransaction()
1. This method take an object struct type of mwcoreapi.ProviderTransaction as parameter and return the created object and error
- `func (rcvr *ProviderTransactionRepo) CreateProviderTransaction(data ProviderTransaction) (int64, error)`
```go
    type Handler struct {
        Mwcore *mwcoreapi.MwCoreConfig
    }


    func (handler *Handler) SampleFunc(w http.ResponseWriter, r *http.Request) {
            data := mwcoreapi.ProviderTransaction{
                Amount:          27,
                ProvTransId:     1231211321312313213,
                ProviderName: "sample_name",
                ProviderTransId : "sample_transId",
                ProviderRoundId:  "sample+provider_id",
                RoundId: 1231211321312313213,
                TransId : 1231211321312313213,
                TokenUUID:  "token_uuid",
                TransType : 1,
                SchemaName:  "sample_schema_name",
            }// need to use the type from the package to create the GameRound Object
            _, err := handler.Mwcore.GameTransaction.CreateProviderTransaction(data)
            if err != nil {
                log.Println(err.Error())
            }
            marshaledData ,_ := json.Marshal(gameDetails)
            fmt.Println(marshaledData)
    }
```
#### UpdateProviderTransaction()
2. This method take an object struct type of mwcoreapi.ProviderTransaction as parameter and return the id of the updated object  and error
- `func (rcvr *ProviderTransactionRepo) UpdateProviderTransaction(data ProviderTransaction) (*ProviderTransaction, error)`
```go
    type Handler struct {
        Mwcore *mwcoreapi.MwCoreConfig
    }


    func (handler *Handler) SampleFunc(w http.ResponseWriter, r *http.Request) {
            data := mwcoreapi.ProviderTransaction{
                ProvTransId:     1231211321312313213,
                ProviderName: "sample_name",
                ProviderTransId : "sample_transId",
                ProviderRoundId:  "sample+provider_id",
                RoundId: 1231211321312313213,
                TransId : 1231211321312313213,
                TokenUUID:  "token_uuid",
                TransType : 1,
                SchemaName:  "sample_schema_name",
            }// need to use the type from the package to create the GameRound Object
            _, err := handler.Mwcore.GameTransaction.UpdateProviderTransaction(data)
            if err != nil {
                log.Println(err.Error())
            }
            marshaledData ,_ := json.Marshal(gameDetails)
            fmt.Println(marshaledData)
    }
```
#### GetProviderTransactionByProviderRoundId()
3. This method take an object struct type of mwcoreapi.ProviderTransaction as parameter and return the id of the  object reference *ProviderTransaction and error
- `func (rcvr *ProviderTransactionRepo) GetProviderTransactionByProviderRoundId(data ProviderTransaction) (*ProviderTransaction, error)`
```go
    type Handler struct {
        Mwcore *mwcoreapi.MwCoreConfig
    }


    func (handler *Handler) SampleFunc(w http.ResponseWriter, r *http.Request) {
             data := mwcoreapi.ProviderTransaction{
                RoundId: 1231211321312313213,
            }// need to use the type from the package to create the GameRound Object
            _, err := handler.Mwcore.GameTransaction.GetProviderTransactionByProviderRoundId(data)
            if err != nil {
                log.Println(err.Error())
            }
            marshaledData ,_ := json.Marshal(gameDetails)
            fmt.Println(marshaledData)
    }
```
#### GetProviderTransactionByProviderTransId()
4. This method take an object struct type of mwcoreapi.ProviderTransaction as parameter and return the id of the  object reference *ProviderTransaction and error
- `func (rcvr *ProviderTransactionRepo) GetProviderTransactionByProviderTransId(data ProviderTransaction) (*ProviderTransaction, error)`
```go
    type Handler struct {
        Mwcore *mwcoreapi.MwCoreConfig
    }


    func (handler *Handler) SampleFunc(w http.ResponseWriter, r *http.Request) {
             data := mwcoreapi.ProviderTransaction{
                TransId : 1231211321312313213,
            }// need to use the type from the package to create the GameRound Object
            _, err := handler.Mwcore.GameTransaction.GetProviderTransactionByProviderTransId(data)
            if err != nil {
                log.Println(err.Error())
            }
            marshaledData ,_ := json.Marshal(gameDetails)
            fmt.Println(marshaledData)
    }
```
#### GetProviderTransactionByProviderTransIdAndTransType()
5. This method take an object struct type of mwcoreapi.ProviderTransaction as parameter and return the id of the  object reference *ProviderTransaction and error
- `func (rcvr *ProviderTransactionRepo) GetProviderTransactionByProviderTransIdAndTransType(data ProviderTransaction) (*ProviderTransaction, error)`
```go
    type Handler struct {
        Mwcore *mwcoreapi.MwCoreConfig
    }


    func (handler *Handler) SampleFunc(w http.ResponseWriter, r *http.Request) {
             data := mwcoreapi.ProviderTransaction{
                ProviderName:    "SampleProviderName",
                ProviderTransId: 11111111111,
                TransType:       1,
                SchemaName:      "SampleSchema",
            }// need to use the type from the package to create the GameRound Object
            _, err := handler.Mwcore.GameTransaction.GetProviderTransactionByProviderTransIdAndTransType(data)
            if err != nil {
                log.Println(err.Error())
            }
            marshaledData ,_ := json.Marshal(gameDetails)
            fmt.Println(marshaledData)
    }
```
#### GetProviderTransactionByProviderRoundIdAndTransType()
6. This method take an object struct type of mwcoreapi.ProviderTransaction as parameter and return the id of the  object reference *ProviderTransaction and error
- `func (rcvr *ProviderTransactionRepo) GetProviderTransactionByProviderRoundIdAndTransType(data ProviderTransaction) (*ProviderTransaction, error)`
```go
    type Handler struct {
        Mwcore *mwcoreapi.MwCoreConfig
    }


    func (handler *Handler) SampleFunc(w http.ResponseWriter, r *http.Request) {
             data := mwcoreapi.ProviderTransaction{
                ProviderName:    "SampleProviderName",
                ProviderRoundId: 11111111111,
                TransType:       1,
                SchemaName:      "SampleSchema",
            }// need to use the type from the package to create the GameRound Object
            _, err := handler.Mwcore.GameTransaction.GetProviderTransactionByProviderRoundIdAndTransType(data)
            if err != nil {
                log.Println(err.Error())
            }
            marshaledData ,_ := json.Marshal(gameDetails)
            fmt.Println(marshaledData)
    }
```
#### GetAllProviderTransactionByProviderRoundId()
7. This method take an object struct type of mwcoreapi.ProviderTransaction as parameter and return the id of the  object reference *ProviderTransaction and error
- `func (rcvr *ProviderTransactionRepo) GetAllProviderTransactionByProviderRoundId(data ProviderTransaction) ([]AllProviderTransactionResponseBodyItems, error)`
```go
    type Handler struct {
        Mwcore *mwcoreapi.MwCoreConfig
    }


    func (handler *Handler) SampleFunc(w http.ResponseWriter, r *http.Request) {
             data := mwcoreapi.ProviderTransaction{
                ProviderName:    "SampleProviderName",
                ProviderRoundId: 11111111111,
                TransType:       1,
                SchemaName:      "SampleSchema",
            }// need to use the type from the package to create the GameRound Object
            _, err := handler.Mwcore.GameTransaction.GetAllProviderTransactionByProviderRoundId(data)
            if err != nil {
                log.Println(err.Error())
            }
            marshaledData ,_ := json.Marshal(gameDetails)
            fmt.Println(marshaledData)
    }
```
----
## Player Balance
This type is coming from type ProviderTrasaction and having 2 methods and 2 type *mwcoreapi.PlayerBalanceToken , *mwcoreapi.PlayerBalance.
```go
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
```
#### UpdatePlayerTokenBalance()
1. This method take an object struct type of mwcoreapi.PlayerBalanceToken as parameter and return the id of the updated object  and error
- `func (rcvr *PlayerBalanceRepo) UpdatePlayerTokenBalance(data PlayerBalanceToken) (*PlayerBalanceToken, error)`
```go
    type Handler struct {
        Mwcore *mwcoreapi.MwCoreConfig
    }


    func (handler *Handler) SampleFunc(w http.ResponseWriter, r *http.Request) {
            data := mwcoreapi.PlayerBalanceToken{
                TokenId    : 123,
                Token_UUID : "sample_token_uuid",
                OperatorId  : 1,
                ClientId    : 1,
                PlayerId    : 1,
                PlayerToken : "sample_player_token",
                Balance     : 10.00,
                StatusId    : 1,
            }// need to use the type from the package to create the GameRound Object
            _, err := handler.Mwcore.GameTransaction.UpdatePlayerTokenBalance(data)
            if err != nil {
                log.Println(err.Error())
            }
            marshaledData ,_ := json.Marshal(gameDetails)
            fmt.Println(marshaledData)
    }
```
#### UpdatePlayerBalance()
2. This method take an object struct type of mwcoreapi.PlayerBalanceToken as parameter and return the id of the updated object  and error
- `func (rcvr *PlayerBalanceRepo) UpdatePlayerBalance(data PlayerBalance) (*PlayerBalance, error)`
```go
    type Handler struct {
        Mwcore *mwcoreapi.MwCoreConfig
    }


    func (handler *Handler) SampleFunc(w http.ResponseWriter, r *http.Request) {
            data := mwcoreapi.PlayerBalance{
                OperatorId  : 1,
                ClientId    : 1,
                PlayerId    : 1,
                Balance     : 10.00,
            }// need to use the type from the package to create the GameRound Object
            _, err := handler.Mwcore.GameTransaction.UpdatePlayerBalance(data)
            if err != nil {
                log.Println(err.Error())
            }
            marshaledData ,_ := json.Marshal(gameDetails)
            fmt.Println(marshaledData)
    }
```