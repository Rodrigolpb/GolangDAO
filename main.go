package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Rodrigolpb/GolangDAO/dao"
	"github.com/Rodrigolpb/GolangDAO/entities"
	_ "github.com/go-sql-driver/mysql"
)

type answerType struct {
	id    int32
	title string
}

func main() {
	db, err := sql.Open("mysql", "root:admin@/retail_chain_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	atDAO := dao.NewAnswerTypeDAO(db)

	rows, err := atDAO.Create(entities.AnswerType{
		Title: "Indestructible",
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows)
}
