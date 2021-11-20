package main

import (
	"container/list"
	"fmt"
	"strconv"
)

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Enter(val interface{}) {
	s.v.PushBack(val)
}

func (s* Stack) Pop() interface{} {
	back := s.v.Back()
	if back != nil {
		return s.v.Remove(back)
	}
	return nil
}

var stack = NewStack()

type Calc interface {
	Add()
	Sub()
	Mult()
	Div()
}

type Integer struct {
	value string
}

type Float struct {
	value string
}

func Add() {
	val1 := stack.Pop()
	val2 := stack.Pop()
	res, newType := Addition(val1, val2)

	switch newType {
	case "Integer":
		stack.Enter(Integer{strconv.FormatFloat(res, 'f', -1, 64)})
	case "Float":
		stack.Enter(Float{strconv.FormatFloat(res, 'f', -1, 64)})
	}
	fmt.Println(res)
}

func Sub() {
	val1 := stack.Pop()
	val2 := stack.Pop()
	res, newType := Subtraction(val1, val2)
	switch newType {
	case "Integer":
		stack.Enter(Integer{strconv.FormatFloat(res, 'f', -1, 64)})
	case "Float":
		stack.Enter(Float{strconv.FormatFloat(res, 'f', -1, 64)})
	}
	fmt.Println(res)
}

func Mult() {
	val1 := stack.Pop()
	val2 := stack.Pop()
	res := Multiplication(val1, val2)
	stack.Enter(Float{strconv.FormatFloat(res, 'f', -1, 64)})
	fmt.Println(res)
}

func Div() {
	val1 := stack.Pop()
	val2 := stack.Pop()
	res := Division(val1, val2)
	stack.Enter(Float{strconv.FormatFloat(res, 'f', -1, 64)})
	fmt.Println(res)
}


func Addition(arg1, arg2 interface{}) (float64, string) {
	switch arg1.(type) {
	case Integer:
		switch arg2.(type) {
		case Integer:
			p1:= arg1.(Integer)
			p2:= arg2.(Integer)
			p3,_ := strconv.Atoi(p1.value)
			p4,_ := strconv.Atoi(p2.value)
			res:=float64(p3+p4)
			return res, "Integer"

		case Float:
			p1:= arg1.(Integer)
			p2:= arg2.(Float)
			p3,_ := strconv.Atoi(p1.value)
			p4,_ := strconv.ParseFloat(p2.value, 64)
			res := float64(p3)+p4
			return res, "Float"

		default:
			return 0, "Err"
	
		}
	case Float:
		switch arg2.(type) {
		case Integer:
			p1:= arg1.(Float)
			p2:= arg2.(Integer)
			p3,_ := strconv.ParseFloat(p1.value, 64)
			p4,_ := strconv.Atoi(p2.value)
			res := p3 + float64(p4)
			return res, "Float"
	
		case Float:
			p1:= arg1.(Float)
			p2:= arg2.(Float)
			p3,_ := strconv.ParseFloat(p1.value, 64)
			p4,_ := strconv.ParseFloat(p2.value, 64)
			res := p3 + p4
			return res, "Float"

		default:
			return 0, "Err"
		}
	default: 
	return 0, "Err"
	}
}

func Subtraction(arg1, arg2 interface{}) (float64, string) {
	switch arg1.(type) {
	case Integer:
		switch arg2.(type) {
		case Integer:
			p1:= arg1.(Integer)
			p2:= arg2.(Integer)
			p3,_ := strconv.Atoi(p1.value)
			p4,_ := strconv.Atoi(p2.value)
			res:=float64(p4-p3)
			return res, "Int"

		case Float:
			p1:= arg1.(Integer)
			p2:= arg2.(Float)
			p3,_ := strconv.Atoi(p1.value)
			p4,_ := strconv.ParseFloat(p2.value, 64)
			res := p4-float64(p3)
			return res, "Float"

		default:
			return 0, "Err"
	
		}
	case Float:
		switch arg2.(type) {
		case Integer:
			p1:= arg1.(Float)
			p2:= arg2.(Integer)
			p3,_ := strconv.ParseFloat(p1.value, 64)
			p4,_ := strconv.Atoi(p2.value)
			res := float64(p4) - p3
			return res, "Float"
	
		case Float:
			p1:= arg1.(Float)
			p2:= arg2.(Float)
			p3,_ := strconv.ParseFloat(p1.value, 64)
			p4,_ := strconv.ParseFloat(p2.value, 64)
			res := p4 - p3
			return res, "Float"

		default:
			return 0, "Err"
		}
	default: 
	return 0, "Err"
	}
}

func Multiplication(arg1, arg2 interface{}) float64 {
	switch arg1.(type) {
	case Integer:
		switch arg2.(type) {
		case Integer:
			p1:= arg1.(Integer)
			p2:= arg2.(Integer)
			p3,_ := strconv.Atoi(p1.value)
			p4,_ := strconv.Atoi(p2.value)
			res:=float64(p3*p4)
			return res

		case Float:
			p1:= arg1.(Integer)
			p2:= arg2.(Float)
			p3,_ := strconv.Atoi(p1.value)
			p4,_ := strconv.ParseFloat(p2.value, 64)
			res := float64(p3)*p4
			return res

		default:
			return 0
	
		}
	case Float:
		switch arg2.(type) {
		case Integer:
			p1:= arg1.(Float)
			p2:= arg2.(Integer)
			p3,_ := strconv.ParseFloat(p1.value, 64)
			p4,_ := strconv.Atoi(p2.value)
			res := p3 * float64(p4)
			return res
	
		case Float:
			p1:= arg1.(Float)
			p2:= arg2.(Float)
			p3,_ := strconv.ParseFloat(p1.value, 64)
			p4,_ := strconv.ParseFloat(p2.value, 64)
			res := p3 * p4
			return res

		default:
			return 0
		}
	default: 
	return 0
	}
}

func Division(arg1, arg2 interface{}) float64 {
	switch arg1.(type) {
	case Integer:
		switch arg2.(type) {
		case Integer:
			p1:= arg1.(Integer)
			p2:= arg2.(Integer)
			p3,_ := strconv.Atoi(p1.value)
			p4,_ := strconv.Atoi(p2.value)
			res:=float64(p4/p3)
			return res

		case Float:
			p1:= arg1.(Integer)
			p2:= arg2.(Float)
			p3,_ := strconv.Atoi(p1.value)
			p4,_ := strconv.ParseFloat(p2.value, 64)
			res := p4/float64(p3)
			return res

		default:
			return 0
	
		}
	case Float:
		switch arg2.(type) {
		case Integer:
			p1:= arg1.(Float)
			p2:= arg2.(Integer)
			p3,_ := strconv.ParseFloat(p1.value, 64)
			p4,_ := strconv.Atoi(p2.value)
			res := float64(p4)/p3
			return res
	
		case Float:
			p1:= arg1.(Float)
			p2:= arg2.(Float)
			p3,_ := strconv.ParseFloat(p1.value, 64)
			p4,_ := strconv.ParseFloat(p2.value, 64)
			res := p4/p3
			return res

		default:
			return 0
		}
	default: 
	return 0
	}
}


func main() {
	stack.Enter(Integer{"7"})
	stack.Enter(Float{"2.3"})
	Div()
	stack.Enter(Integer{"2"})
	Add()

}
