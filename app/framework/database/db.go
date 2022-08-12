package database

import (
	"log"

	"minotauro/app/domain"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DataBase struct {
	Db            *gorm.DB
	Dsn           string
	DsnTes        string
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
	dbIntance.Debug = true

	connection, err := dbIntance.Connect()
	if err != nil {
		log.Fatalf("erro no banco de testes : %v", err)
	}
	return connection
}

func (d *DataBase) Connect() (*gorm.DB, error) {
	var err error

	if d.Env != "teste" {
		d.Db, err = gorm.Open(postgres.Open(d.Dsn))
	} else {
		d.Db, err = gorm.Open(sqlite.Open(d.DsnTes))
	}

	if err != nil {
		return nil, err
	}
	if d.Debug {
		d.Db.Logger.LogMode(logger.Error)
	}

	if d.AutoMigrateDb {
		d.Db.AutoMigrate(
			&domain.Parque{},
			&domain.Pagamento{},
			&domain.TipoProfissional{},
			&domain.Profissionall{},
			&domain.Categoria{},
			&domain.Evento{},
			&domain.Rodizio{},
			&domain.Corrida{},
			&domain.Inscricao{},
		)
		d.Db.Model(domain.Corrida{}).AddFo

	}
	return d.Db, nil
}
