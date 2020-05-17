package entities

// AnswerType - Type that represents checklists questions group of possible answers
type AnswerType struct {
	ID    int    `column:"id" ignoreOnWrite:"true"`
	Title string `column:"title"`
}
