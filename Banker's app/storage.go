package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)
// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50TnVtYmVyIjoxNjIsImV4cGlyZXNBdCI6MTUwMDB9.OlkMkb-UR2LZr5qY0R7X0NIIquRSQ777auKoEkEZW5M
type Storage interface {
	CreateAccount(*Account) error
	DeleteAccountByID(string) error
	UpdateAccount(*Account) error
	GetAccountByID(string) (*Account, error)
	GetAccounts()([]*Account,error)
	GetAccountByEmail(string)(*Account,error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{db: db}, nil

}
func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	querry := `create table if not exists account(
		id varchar(50) primary key,
		first_name varchar(50),
		last_name varchar(50),
		email varchar(50) unique,
		encrypted_password varchar(200),
		number serial,
		balance serial,
		created_at timestamp
	
	)`
	_, err := s.db.Exec(querry)
	return err

}

func (s *PostgresStore) CreateAccount(acc *Account) error {
	querry := `insert into account
	(id,first_name,last_name,email,encrypted_password,number,balance,created_at)
	values ($1, $2, $3, $4, $5, $6, $7, $8)`
	log.Println()
	log.Println("account:",acc)
	log.Println()
	resp, err := s.db.Query(querry, acc.ID, acc.FirstName, acc.LastName,acc.Email,acc.Password, acc.Number, acc.Balance, acc.CreatedAt)
	if err != nil {
		log.Println("error while inserting...")
		return err
	}
	log.Println(resp)
	return nil
}

func(s *PostgresStore) GetAccounts()([]*Account,error){
	querry := `select * from account`
	rows,err := s.db.Query(querry)
	if err!=nil{
		return []*Account{},err
	}
	accounts := []*Account{}
	
	for rows.Next(){
		account,err := scanEachRow(rows)
		if err!=nil{
			return []*Account{},err
		}
		accounts = append(accounts,account)
	}

	return accounts,nil

}


func (s *PostgresStore) GetAccountByEmail(email string) (*Account,error) {
	querry := `select * from account where email = $1`
	rows,err := s.db.Query(querry,email)
	if err!=nil{
		return &Account{},err
	}

	for rows.Next(){
		return scanEachRow(rows)		
	}
	return nil, fmt.Errorf("email %s not found", email)
}

func (s *PostgresStore) GetAccountByID(id string) (*Account, error) {
	querry := `select * from account where id = $1`
	rows,err := s.db.Query(querry,id)
	if err!=nil{
		return &Account{},err
	}
	for rows.Next(){
		return scanEachRow(rows)	
	}
	return nil, fmt.Errorf("id %s not found", id)
}
func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}
func (s *PostgresStore) DeleteAccountByID(id string) error {
	
	querry := `delete from account where id = $1`
	_,err :=s.db.Exec(querry,id)
	if err!=nil{
		return fmt.Errorf("error occured during deleting account,%s",err)
	}
	return nil
}



func scanEachRow(rows *sql.Rows)(*Account,error){
	account := &Account{}
	err := rows.Scan(&account.ID,&account.FirstName,&account.LastName,&account.Email,&account.Password,&account.Number,&account.Balance,&account.CreatedAt)
	if err!=nil{
		return nil,err
	}
	return account,nil
}