package util

import (
	"math/rand"
	"time"
	"fmt"
)

type DiceRoll struct {
	numberOfDice int
	typeOfDice int
	rolls []int
}

func (dr DiceRoll) String() string {
	return fmt.Sprintf("Rolled %dD%d. Results %v, for a total of %d", dr.numberOfDice, dr.typeOfDice, dr.rolls, dr.Total())
	
}

func (dr DiceRoll) IndividualRolls() []int {
	return dr.rolls	
}

func (dr DiceRoll) Total() (t int) {
	t = 0
	for _, r := range(dr.rolls) {
		t += r
	}
	return
}

func NewDiceRoll(n, t int, r []int) *DiceRoll {
	return &DiceRoll{
		numberOfDice: n,
		typeOfDice: t,
		rolls: r,
	}
}

func RollDSix() int {
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(6 - 1 + 1) + 1
}


func RollDice(numDice, typeDice int) (*DiceRoll, error) {
	if (numDice < 1) {
		return nil, GenericError("Please specify a number of dice") 
	}
	switch {
	case typeDice == 4:
		return nil, NotYetImplemented("d4 rolling")
	case typeDice == 6:
		rolls := rollDice(numDice, RollDSix)
		return NewDiceRoll(numDice, typeDice, rolls), nil
	case typeDice == 8:
		return nil, NotYetImplemented("d8 rolling")
	case typeDice == 10:
		return nil, NotYetImplemented("d10 rolling")
	case typeDice == 12:
		return nil, NotYetImplemented("d12 rolling")
	case typeDice == 20:
		return nil, NotYetImplemented("d20 rolling")
	case typeDice == 100:
		return nil, NotYetImplemented("d100 rolling")
	default:
		return nil, GenericError("Unknown dice type")
	}
}


func rollDice(num int, diceRoll func() int) (rolls []int) {
	for i := 0; i < num; i++ {
		rolls = append(rolls, diceRoll())
	}  
	return 
}