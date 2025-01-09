package models

type ExpressionGroup struct {
    Op LogicalOperator
	LeftVariable *Variable
	RightVariable *Variable
	LeftExpressionGroup *ExpressionGroup
	RightExpressionGroup *ExpressionGroup
}



// func (arg Arg) solving(previousValue Value, facts []Fact) (bool, error) {
// 	fmt.Println("Arg solving with arg: %s and previousValue: %s", arg, previousValue)
// 	fact, err := GetFactReferenceByLetter(facts, arg.Letter)
// 	fmt.Println("Fact: %s", fact)
// 	if err != nil {
// 		return false, err
// 	}

// 	if previousLetter == 0 {
// 		return c.value(*fact), nil
// 	}

// 	if !arg.Op.isValid() {
// 		return false, errors.New(fmt.Sprintf("Unknown operator: %s", arg.Op.toString()))
// 	}
// 	result, err := arg.Op.solve(previousValue, arg.value(*fact))

// 	return result, nil
// }



// // DISPLAY

// func (arg Arg) String() string {
// 	return fmt.Sprintf("%s %c", arg.Op.getSymbol(), arg.getLetter())
// }

// func (c Arg) getLetter() string {
// 	if (c.Not) {
// 		return "!" + string(c.Letter)
// 	} else {
// 		return string(c.Letter)
// 	}
// }


// func (c Rule) GetRule() string {
//     var leftArgs, rightArgs string
//     for _, arg := range c.LeftArgs {
//         leftArgs += arg.Op.getSymbol() + " " + arg.getLetter() + " "
//     }
//     for _, arg := range c.RightArgs {
//         rightArgs += arg.Op.getSymbol() + " " + arg.getLetter() + " "
//     }
//     return leftArgs + c.Op.getSymbol() + rightArgs
// }

// func (c Rule) InitialSolving(facts []Fact) {
    
// }

// func DisplayRules(rules []Rule) {
//     for i, rule := range rules {
//         fmt.Printf("%d: %s\n", i+1, rule.GetRule())
//     }
// }

// func RulesConditionalOperatorFormatter(rules []Rule) []Rule {
//     var newRules []Rule
//     for _, rule := range rules {
//         if rule.Op == IFF {
//             newRules = append(newRules, Rule{
//                 LeftArgs: rule.LeftArgs,
//                 RightArgs: rule.RightArgs,
//                 Op: IMPLIES,
//             })
//             newRules = append(newRules, Rule{
//                 LeftArgs: rule.RightArgs,
//                 RightArgs: rule.LeftArgs,
//                 Op: IMPLIES,
//             })
//         } else {
//             newRules = append(newRules, rule)
//         }
//     }
//     return newRules
// }

// func GetRulesMock() []Rule {
//     rules := []Rule{
//         // C => E
//         {LeftArgs: []Arg{{Op: NOTHING, Letter: 'C'}}, RightArgs: []Arg{{Op: NOTHING, Letter: 'E'}}, Op: IMPLIES},

//         // A + B + C => D
//         {LeftArgs: []Arg{
//             {Op: NOTHING, Letter: 'A'}, 
//             {Op: AND, Letter: 'B'}, 
//             {Op: AND, Letter: 'C'},
//         }, RightArgs: []Arg{{Op: NOTHING, Letter: 'D'}}, Op: IMPLIES},

//         // A | B => C
//         {LeftArgs: []Arg{
//             {Op: NOTHING, Letter: 'A'}, 
//             {Op: OR, Letter: 'B'},
//         }, RightArgs: []Arg{{Op: NOTHING, Letter: 'C'}}, Op: IMPLIES},

//         // A + !B => F
//         {LeftArgs: []Arg{
//             {Op: NOTHING, Letter: 'A'}, 
//             {Op: AND, Letter: 'B', Not: true},
//         }, RightArgs: []Arg{{Op: NOTHING, Letter: 'F'}}, Op: IMPLIES},

//         // C | !G => H
//         {LeftArgs: []Arg{
//             {Op: NOTHING, Letter: 'C'}, 
//             {Op: OR, Letter: 'G', Not: true},
//         }, RightArgs: []Arg{{Op: NOTHING, Letter: 'H'}}, Op: IMPLIES},

//         // V ^ W => X
//         {LeftArgs: []Arg{
//             {Op: NOTHING, Letter: 'V'}, 
//             {Op: XOR, Letter: 'W'},
//         }, RightArgs: []Arg{{Op: NOTHING, Letter: 'X'}}, Op: IMPLIES},

//         // A + B => Y + Z
//         {LeftArgs: []Arg{
//             {Op: NOTHING, Letter: 'A'}, 
//             {Op: AND, Letter: 'B'},
//         }, RightArgs: []Arg{
//             {Op: NOTHING, Letter: 'Y'}, 
//             {Op: AND, Letter: 'Z'},
//         }, Op: IMPLIES},

//         // C | D => X | V
//         {LeftArgs: []Arg{
//             {Op: NOTHING, Letter: 'C'}, 
//             {Op: OR, Letter: 'D'},
//         }, RightArgs: []Arg{
//             {Op: NOTHING, Letter: 'X'}, 
//             {Op: OR, Letter: 'V'},
//         }, Op: IMPLIES},

//         // E + F => !V
//         {LeftArgs: []Arg{
//             {Op: NOTHING, Letter: 'E'}, 
//             {Op: AND, Letter: 'F'},
//         }, RightArgs: []Arg{{Op: NOTHING, Letter: 'V', Not: true}}, Op: IMPLIES},

//         // A + B <=> C
//         {LeftArgs: []Arg{
//             {Op: NOTHING, Letter: 'A'}, 
//             {Op: AND, Letter: 'B'},
//         }, RightArgs: []Arg{{Op: NOTHING, Letter: 'C'}}, Op: IFF},

//         // A + B <=> !C
//         {LeftArgs: []Arg{
//             {Op: NOTHING, Letter: 'A'}, 
//             {Op: AND, Letter: 'B'},
//         }, RightArgs: []Arg{{Op: NOTHING, Letter: 'C', Not: true}}, Op: IFF},
//     }

//     return rules
// }
