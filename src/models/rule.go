package models

import (
	"fmt"
)

type Rule struct {
    LeftArgs []Arg
    RightArgs []Arg
    Op ConditionalOperator
}

func (c Rule) GetRule() string {
    var leftArgs, rightArgs string
    for _, arg := range c.LeftArgs {
        leftArgs += arg.Op.getSymbol() + " " + arg.getLetter() + " "
    }
    for _, arg := range c.RightArgs {
        rightArgs += arg.Op.getSymbol() + " " + arg.getLetter() + " "
    }
    return leftArgs + c.Op.getSymbol() + rightArgs
}

func DisplayRules(rules []Rule) {
    for i, rule := range rules {
        fmt.Printf("%d: %s\n", i+1, rule.GetRule())
    }
}

func GetRulesMock() []Rule {
    rules := []Rule{
        // C => E
        {LeftArgs: []Arg{{Op: NOTHING, Letter: 'C'}}, RightArgs: []Arg{{Op: NOTHING, Letter: 'E'}}, Op: IMPLIES},

        // A + B + C => D
        {LeftArgs: []Arg{
            {Op: NOTHING, Letter: 'A'}, 
            {Op: AND, Letter: 'B'}, 
            {Op: AND, Letter: 'C'},
        }, RightArgs: []Arg{{Op: NOTHING, Letter: 'D'}}, Op: IMPLIES},

        // A | B => C
        {LeftArgs: []Arg{
            {Op: NOTHING, Letter: 'A'}, 
            {Op: OR, Letter: 'B'},
        }, RightArgs: []Arg{{Op: NOTHING, Letter: 'C'}}, Op: IMPLIES},

        // A + !B => F
        {LeftArgs: []Arg{
            {Op: NOTHING, Letter: 'A'}, 
            {Op: AND, Letter: 'B', Not: true},
        }, RightArgs: []Arg{{Op: NOTHING, Letter: 'F'}}, Op: IMPLIES},

        // C | !G => H
        {LeftArgs: []Arg{
            {Op: NOTHING, Letter: 'C'}, 
            {Op: OR, Letter: 'G', Not: true},
        }, RightArgs: []Arg{{Op: NOTHING, Letter: 'H'}}, Op: IMPLIES},

        // V ^ W => X
        {LeftArgs: []Arg{
            {Op: NOTHING, Letter: 'V'}, 
            {Op: XOR, Letter: 'W'},
        }, RightArgs: []Arg{{Op: NOTHING, Letter: 'X'}}, Op: IMPLIES},

        // A + B => Y + Z
        {LeftArgs: []Arg{
            {Op: NOTHING, Letter: 'A'}, 
            {Op: AND, Letter: 'B'},
        }, RightArgs: []Arg{
            {Op: NOTHING, Letter: 'Y'}, 
            {Op: AND, Letter: 'Z'},
        }, Op: IMPLIES},

        // C | D => X | V
        {LeftArgs: []Arg{
            {Op: NOTHING, Letter: 'C'}, 
            {Op: OR, Letter: 'D'},
        }, RightArgs: []Arg{
            {Op: NOTHING, Letter: 'X'}, 
            {Op: OR, Letter: 'V'},
        }, Op: IMPLIES},

        // E + F => !V
        {LeftArgs: []Arg{
            {Op: NOTHING, Letter: 'E'}, 
            {Op: AND, Letter: 'F'},
        }, RightArgs: []Arg{{Op: NOTHING, Letter: 'V', Not: true}}, Op: IMPLIES},

        // A + B <=> C
        {LeftArgs: []Arg{
            {Op: NOTHING, Letter: 'A'}, 
            {Op: AND, Letter: 'B'},
        }, RightArgs: []Arg{{Op: NOTHING, Letter: 'C'}}, Op: IFF},

        // A + B <=> !C
        {LeftArgs: []Arg{
            {Op: NOTHING, Letter: 'A'}, 
            {Op: AND, Letter: 'B'},
        }, RightArgs: []Arg{{Op: NOTHING, Letter: 'C', Not: true}}, Op: IFF},
    }

    return rules
}
