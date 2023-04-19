package questionentity

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type QuestionDifficulty int

const (
	QuestionDifficultyEasy QuestionDifficulty = iota + 1
	QuestionDifficultyNormal
	QuestionDifficultyHard
)

func (qt QuestionDifficulty) String() string {
	switch qt {
	case QuestionDifficultyEasy:
		return "easy"
	case QuestionDifficultyNormal:
		return "normal"
	default:
		return "hard"
	}
}

func parseStringToQuestionDifficulty(s string) (QuestionDifficulty, error) {
	switch s {
	case "easy":
		return QuestionDifficultyEasy, nil
	case "normal":
		return QuestionDifficultyNormal, nil
	case "hard":
		return QuestionDifficultyHard, nil
	default:
		return QuestionDifficulty(1), errors.New("invalid question difficulty string")
	}
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
