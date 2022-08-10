package database

import (
	"log"

	"gorm.io/gorm"
	"minotauro/app/domain"
)

type DataBase struct {
	Db            *gorm.DB
	Dsn           string
	DsnTes        string
	DbType        string
	DbyTypeTest   string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

func NewDb() *DataBase {
	return &DataBase{}
}

func NewDbTes() *gorm.DB {
	dbIntance := NewDb()
	dbIntance.Env = "test"
	dbIntance.DbyTypeTest = "sqlite3"
	dbIntance.DsnTes = ":memory"
	dbIntance.Debug = true
	
	connection, err :=dbIntance.Connect()
	if err !=nil{
		log.Fatalf("erro no banco de testes : %v", err)
	}
	return connection
}

func (d *DataBase)Connect()(*gorm.DB, error){
	var err error
	
	if d.Env != "teste"{
		d.Db, err = gorm.Open(d.DbType, d.Dsn)
	} else{
		d.Db, err = gorm.Open(d.DbyTypeTest, d.DsnTes)
	}

	if err !=nil{
		return nil, err
	}
	if d.Debug{
		d.Db.Logger.LogMode(true)
	}

	if d.Db.AutoMigrate{
		d.Db.AutoMigrate(
			&domain.Parque{},
			&domain.Pagamento{},
			&domain.TipoProfissional{},
			&domain.Profissionall{},
			&domain.Categoria{},
			&domain.Evento{},
			&domain.Rodizio{},
			&domain.Corrida{},
			
			&domain.


		)

	}
}