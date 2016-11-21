package gygax

import (
  "reflect"
  "testing"
)


func TestParseTextForDice(t *testing.T) {
  t.Parallel()
  if !()
  if !(reflect.DeepEqual(parseTextForDice("d10 d20"), []int{10, 20})) {
    t.Fail()
  }
  if !(reflect.DeepEqual(parseTextForDice("d4"), []int{4})) {
    t.Fail()
  }
}

func TestDiceRollText(t *testing.T) {

}
