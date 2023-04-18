package questionentity

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type QuestionType int

const (
	QuestionTypeMultipleChoice QuestionType = iota
	QuestionTypeTrueFalse
)

var allQuestionTypes = [2]string{"Multiple Choice", "True/False"}

func (qt *QuestionType) String() string {
	return allQuestionTypes[*qt]
}

func parseStringToQuestionType(s string) (QuestionType, error) {
	for i := range allQuestionTypes {
		if allQuestionTypes[i] == s {
			return QuestionType(i), nil
		}
	}

	return QuestionType(0), errors.New("invalid question type string")
}

// Scan read data from SQL to QuestionType
func (qt *QuestionType) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	questionType, err := parseStringToQuestionType(string(bytes))

	if err != nil {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	*qt = questionType

	return nil
}

// Value read data from QuestionType to SQL
func (qt *QuestionType) Value() (driver.Value, error) {
	if qt == nil {
		return nil, nil
	}

	return qt.String(), nil
}

// MarshalJSON support convert from QuestionType to JSON value
func (qt *QuestionType) MarshalJSON() ([]byte, error) {
	if qt == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf("\"%s\"", qt.String())), nil
}

// MarshalJSON support convert from JSON value to QuestionType
func (qt *QuestionType) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	questionType, err := parseStringToQuestionType(str)

	if err != nil {
		return err
	}

	*qt = questionType

	return nil
}
