package goal

import (
	"fmt"
	"testing"
)

func TestStatusTodo(t *testing.T) {
	expectedStrings := map[Status]string{
		StatusTodo:       "Todo",
		StatusInprogress: "In progress",
		StatusDone:       "Done",
		StatusBacklog:    "Backlog",
	}

	for status, expectedString := range expectedStrings {
		actualString := fmt.Sprint(status)
		if actualString != expectedString {
			t.Errorf("Expected '%v' for goal status %v. Got '%v'.", expectedString, status, actualString)
		}
	}
}
