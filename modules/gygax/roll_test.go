package gygax

import (
  "reflect"
  "testing"
)


func TestParseTextForDice(t *testing.T) {
  t.Parallel()
  if !(reflect.DeepEqual(parseTextForDice("d4"), []int{4})) {
    t.Fail()
  }
}

func TestDiceRollText(t *testing.T) {

}
