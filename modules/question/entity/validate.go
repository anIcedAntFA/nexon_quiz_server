package questionentity

import "fmt"

func checkEmptyField(field string) error {
	if field == "" {
		return fmt.Errorf("%s cannot be blank", field)
	}

	return nil
}
