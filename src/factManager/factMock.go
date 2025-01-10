package factManager

import (
    "expert/v"
)

func GetFactsMock() []Fact {
	facts := []Fact{
		{Letter: 'A', Value: v.TRUE, Initial: true, Reason: Reason{Msg: "Initial fact"}},
		{Letter: 'B', Value: v.TRUE, Initial: true, Reason: Reason{Msg: "Initial fact"}},
		{Letter: 'G', Value: v.TRUE, Initial: true, Reason: Reason{Msg: "Initial fact"}},
        {Letter: 'C', Value: v.UNKNOW, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'E', Value: v.UNKNOW, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'D', Value: v.UNKNOW, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'F', Value: v.UNKNOW, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'G', Value: v.UNKNOW, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'H', Value: v.UNKNOW, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'V', Value: v.UNKNOW, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'W', Value: v.UNKNOW, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'X', Value: v.UNKNOW, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'Y', Value: v.UNKNOW, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'Z', Value: v.UNKNOW, Initial: false, Reason: Reason{Msg: ""}},
	}
	return facts
}
