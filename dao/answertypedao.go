package dao

import "github.com/Rodrigolpb/GolangDAO/entities"

// AnswerTypeDAO - AnswerType data access object type, represents answer_types table
type AnswerTypeDAO struct {
	BaseDAO
	tableName string
}

// NewAnswerTypeDAO - AnswerTypeDAO constructor
func NewAnswerTypeDAO() *AnswerTypeDAO {
	return &AnswerTypeDAO{
		tableName: "answer_types",
	}
}

// Create - Adds new value to the configured table
func (at *AnswerTypeDAO) Create(answerType entities.AnswerType) {
	at.BaseDAO.Create(at.tableName, answerType)
}
