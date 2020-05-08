package main

import (
	"github.com/Rodrigolpb/GolangDAO/dao"
	"github.com/Rodrigolpb/GolangDAO/entities"
)

type answerType struct {
	id    int32
	title string
}

func main() {
	atDAO := dao.NewAnswerTypeDAO()
	atDAO.Create(entities.AnswerType{
		ID:    1,
		Title: "Testing",
	})
}
