package factManager

import (
	"expert-system/src/v"
)

func GetFactsMock() []Fact {
	facts := []Fact{
		{Letter: 'A', Value: v.TRUE, Initial: true, Reason: Reason{Msg: "Initial fact"}},
		{Letter: 'B', Value: v.TRUE, Initial: true, Reason: Reason{Msg: "Initial fact"}},
		{Letter: 'G', Value: v.TRUE, Initial: true, Reason: Reason{Msg: "Initial fact"}},
		{Letter: 'C', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'E', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'D', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'F', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'H', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'V', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'W', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'X', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'Y', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'Z', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
	}
	return facts
}

func GetFactsMock2() []Fact {
	facts := []Fact{
		{Letter: 'A', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'B', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'D', Value: v.TRUE, Initial: true, Reason: Reason{Msg: ""}},
		{Letter: 'E', Value: v.TRUE, Initial: true, Reason: Reason{Msg: ""}},
		{Letter: 'F', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'G', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'H', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'I', Value: v.TRUE, Initial: true, Reason: Reason{Msg: ""}},
		{Letter: 'J', Value: v.TRUE, Initial: true, Reason: Reason{Msg: ""}},
		{Letter: 'K', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'L', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'M', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'N', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'O', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'P', Value: v.TRUE, Initial: true, Reason: Reason{Msg: ""}},
	}
	return facts
}

func GetFactsMock3() []Fact {
	facts := []Fact{
		{Letter: 'A', Value: v.TRUE, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'X', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'Y', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'Z', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
	}
	return facts
}

func GetFactsMock4() []Fact {
	facts := []Fact{
		{Letter: 'A', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'B', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'C', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'D', Value: v.TRUE, Initial: false, Reason: Reason{Msg: ""}},
		{Letter: 'E', Value: v.TRUE, Initial: false, Reason: Reason{Msg: ""}},
	}
	return facts
}
