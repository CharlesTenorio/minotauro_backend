CREATE TABLE parque (
	id_parque UUID not NULL PRIMARY KEY,
	nome VARCHAR ( 50 )  NOT NULL,
	endereco VARCHAR ( 50 ) NOT NULL,
	bairro VARCHAR ( 50 ) not null,
	cidade VARCHAR(40) not null,
	uf varchar(2) not null,
	ativado BOOLEAN NOT null,
	img varchar(300),
	referencia VARCHAR(200) not null,
	data_at TIMESTAMP NOT NULL
       
);

CREATE TABLE categoria (
	id_categoria UUID not NULL PRIMARY KEY,
	id_parque UUID references parque(id_parque),
	nome VARCHAR ( 50 )  NOT NULL,
	valor NUMERIC(9,2) not null,
	ativado BOOLEAN NOT null,
	data_at TIMESTAMP NOT NULL
       
);