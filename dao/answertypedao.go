package dao

import (
	"database/sql"

	"github.com/Rodrigolpb/GolangDAO/entities"
)

// AnswerTypeDAO - AnswerType data access object type, represents answer_types table
type AnswerTypeDAO struct {
	baseDAO
	tableName string
}

// NewAnswerTypeDAO - AnswerTypeDAO constructor
func NewAnswerTypeDAO(db *sql.DB) *AnswerTypeDAO {
	return &AnswerTypeDAO{
		baseDAO: baseDAO{
			db: db,
		},
		tableName: "answer_types",
	}
}

// Create - Adds new value to the configured table
func (at *AnswerTypeDAO) Create(answerType entities.AnswerType) (int64, error) {
	return at.baseDAO.Create(at.tableName, answerType)
}

// ReadOne - queries answertype by ID
func (at *AnswerTypeDAO) ReadOne(id int32) {}
