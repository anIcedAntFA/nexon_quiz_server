package questionentity

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type QuestionDifficulty int

const (
	QuestionDifficultyEasy QuestionDifficulty = iota
	QuestionDifficultyNormal
	QuestionDifficultyHard
)

var allQuestionDifficulties = [2]string{"Multiple Choice", "True/False"}

func (qd *QuestionDifficulty) String() string {
	return allQuestionDifficulties[*qd]
}

func parseStringToQuestionDifficulty(s string) (QuestionDifficulty, error) {
	for i := range allQuestionDifficulties {
		if allQuestionDifficulties[i] == s {
			return QuestionDifficulty(i), nil
		}
	}

	return QuestionDifficulty(0), errors.New("invalid question type string")
}

// Scan read data from SQL to QuestionType
func (qd *QuestionDifficulty) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	questionDifficulty, err := parseStringToQuestionDifficulty(string(bytes))

	if err != nil {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	*qd = questionDifficulty

	return nil
}

// Value read data from QuestionType to SQL
func (qt *QuestionDifficulty) Value() (driver.Value, error) {
	if qt == nil {
		return nil, nil
	}

	return qt.String(), nil
}

// MarshalJSON support convert from QuestionType to JSON value
func (qd *QuestionDifficulty) MarshalJSON() ([]byte, error) {
	if qd == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf("\"%s\"", qd.String())), nil
}

// MarshalJSON support convert from JSON value to QuestionType
func (qd *QuestionDifficulty) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	questionDifficulty, err := parseStringToQuestionDifficulty(str)

	if err != nil {
		return err
	}

	*qd = questionDifficulty

	return nil
}
