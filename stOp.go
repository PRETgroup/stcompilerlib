package stcompilerlib

import (
	"github.com/PRETgroup/stcompilerlib/postfixlib"
)

type stOp struct {
	token       string
	precedence  int
	numOperands int
	association postfixlib.Association
}

func (s stOp) GetToken() string {
	return s.token
}

func (s stOp) GetPrecedence() int {
	return s.precedence
}

func (s stOp) GetNumOperands() int {
	return s.numOperands
}

func (s stOp) GetAssociation() postfixlib.Association {
	return s.association
}

func (s stOp) LeftAssociative() bool {
	return s.association == postfixlib.AssociationLeft
}

func (s stOp) MarshalJSON() ([]byte, error) {
	return []byte("\"" + s.token + "\""), nil
}

var stOps = []postfixlib.Operator{
	stOp{stExit, 0, 0, postfixlib.AssociationNone},   //not technically operators but they might appear in expressions (with no other things around them)
	stOp{stReturn, 0, 0, postfixlib.AssociationNone}, //not technically operators but they might appear in expressions (with no other things around them)
	stOp{stNot, 0, 1, postfixlib.AssociationRight},
	stOp{stNegative, 0, 1, postfixlib.AssociationRight},
	stOp{stExponentiation, 1, 2, postfixlib.AssociationRight},
	stOp{stMultiply, 2, 2, postfixlib.AssociationLeft},
	stOp{stDivide, 2, 2, postfixlib.AssociationLeft},
	stOp{stModulo, 2, 2, postfixlib.AssociationLeft},
	stOp{stAdd, 3, 2, postfixlib.AssociationLeft},
	stOp{stSubtract, 3, 2, postfixlib.AssociationLeft},
	stOp{stLessThan, 4, 2, postfixlib.AssociationLeft},
	stOp{stGreaterThan, 4, 2, postfixlib.AssociationLeft},
	stOp{stLessThanEqualTo, 4, 2, postfixlib.AssociationLeft},
	stOp{stGreaterThanEqualTo, 4, 2, postfixlib.AssociationLeft},
	stOp{stEqual, 4, 2, postfixlib.AssociationLeft},
	stOp{stInequal, 4, 2, postfixlib.AssociationLeft},
	stOp{stAnd, 5, 2, postfixlib.AssociationLeft},
	stOp{stExlusiveOr, 5, 2, postfixlib.AssociationLeft},
	stOp{stOr, 5, 2, postfixlib.AssociationLeft},
	stOp{stAssignment, 6, 2, postfixlib.AssociationLeft},
}

//FindOp finds a given operator for a given token
func FindOp(op string) postfixlib.Operator {
	for i := 0; i < len(stOps); i++ {
		if stOps[i].GetToken() == op {
			return stOps[i]
		}
	}
	//still here? might be a function
	if is, fn := postfixlib.IsFunction(op); is {
		return fn
	}
	//still here? not an operator
	return nil
}

//OpTokenIsComparison takes a given operator token and returns a true if it
//is of a comparison type (e.g. ">=")
func OpTokenIsComparison(opTok string) bool {
	if opTok == stGreaterThan ||
		opTok == stGreaterThanEqualTo ||
		opTok == stLessThan ||
		opTok == stLessThanEqualTo ||
		opTok == stEqual ||
		opTok == stInequal {

		return true
	}
	return false
}

//OpTokenIsCombinator takes a given operator token and returns a true if it
//is of a combination type (e.g. "and", "or")
func OpTokenIsCombinator(opTok string) bool {
	if opTok == stAnd ||
		opTok == stOr ||
		opTok == stExlusiveOr {

		return true
	}
	return false
}
