package rules

func GetRulesMock() []Rule {
	rules := []Rule{
		// C => E
		Rule{
			Op:                   IMPLIES,
			LeftExpressionGroup:  nil,
			LeftVariable:         &Variable{Letter: 'C', Not: false},
			RightExpressionGroup: nil,
			RightVariable:        &Variable{Letter: 'E', Not: false},
		},

		// A + B + C => D
		Rule{
			Op: IMPLIES,
			LeftExpressionGroup: &ExpressionGroup{
				Op:           AND,
				LeftVariable: &Variable{Letter: 'A', Not: false},
				RightExpressionGroup: &ExpressionGroup{
					Op:            AND,
					LeftVariable:  &Variable{Letter: 'B', Not: false},
					RightVariable: &Variable{Letter: 'C', Not: false},
				},
			},
			RightExpressionGroup: nil,
			RightVariable:        &Variable{Letter: 'D', Not: false},
		},

		// A | B => C
		Rule{
			Op: IMPLIES,
			LeftExpressionGroup: &ExpressionGroup{
				Op:            OR,
				LeftVariable:  &Variable{Letter: 'A', Not: false},
				RightVariable: &Variable{Letter: 'B', Not: false},
			},
			RightExpressionGroup: nil,
			RightVariable:        &Variable{Letter: 'C', Not: false},
		},

		// A + !B => F
		Rule{
			Op: IMPLIES,
			LeftExpressionGroup: &ExpressionGroup{
				Op:            AND,
				LeftVariable:  &Variable{Letter: 'A', Not: false},
				RightVariable: &Variable{Letter: 'B', Not: true}, // NOT appliqué
			},
			RightExpressionGroup: nil,
			RightVariable:        &Variable{Letter: 'F', Not: false},
		},

		// C | !G => H
		Rule{
			Op: IMPLIES,
			LeftExpressionGroup: &ExpressionGroup{
				Op:            OR,
				LeftVariable:  &Variable{Letter: 'C', Not: false},
				RightVariable: &Variable{Letter: 'G', Not: true},
			},
			RightExpressionGroup: nil,
			RightVariable:        &Variable{Letter: 'H', Not: false},
		},

		// V ^ W => X
		Rule{
			Op: IMPLIES,
			LeftExpressionGroup: &ExpressionGroup{
				Op:            XOR,
				LeftVariable:  &Variable{Letter: 'V', Not: false},
				RightVariable: &Variable{Letter: 'W', Not: false},
			},
			RightExpressionGroup: nil,
			RightVariable:        &Variable{Letter: 'X', Not: false},
		},

		// A + B => Y + Z
		Rule{
			Op: IMPLIES,
			LeftExpressionGroup: &ExpressionGroup{
				Op:            AND,
				LeftVariable:  &Variable{Letter: 'A', Not: false},
				RightVariable: &Variable{Letter: 'B', Not: false},
			},
			RightExpressionGroup: &ExpressionGroup{
				Op:            AND,
				LeftVariable:  &Variable{Letter: 'Y', Not: false},
				RightVariable: &Variable{Letter: 'Z', Not: false},
			},
		},

		// C | D => X | V
		Rule{
			Op: IMPLIES,
			LeftExpressionGroup: &ExpressionGroup{
				Op:            OR,
				LeftVariable:  &Variable{Letter: 'C', Not: false},
				RightVariable: &Variable{Letter: 'D', Not: false},
			},
			RightExpressionGroup: &ExpressionGroup{
				Op:            OR,
				LeftVariable:  &Variable{Letter: 'X', Not: false},
				RightVariable: &Variable{Letter: 'V', Not: false},
			},
		},

		// E + F => !V
		Rule{
			Op: IMPLIES,
			LeftExpressionGroup: &ExpressionGroup{
				Op:            AND,
				LeftVariable:  &Variable{Letter: 'E', Not: false},
				RightVariable: &Variable{Letter: 'F', Not: false},
			},
			RightExpressionGroup: nil,
			RightVariable:        &Variable{Letter: 'V', Not: true},
		},

		// A + B <=> C
		Rule{
			Op: IFF,
			LeftExpressionGroup: &ExpressionGroup{
				Op:            AND,
				LeftVariable:  &Variable{Letter: 'A', Not: false},
				RightVariable: &Variable{Letter: 'B', Not: false},
			},
			RightExpressionGroup: nil,
			RightVariable:        &Variable{Letter: 'C', Not: false},
		},

		// A + B <=> !C
		// Rule{
		//     Op: IFF,
		//     LeftExpressionGroup: &ExpressionGroup{
		//         Op: AND,
		//         LeftVariable: &Variable{Letter: 'A', Not: false},
		//         RightVariable: &Variable{Letter: 'B', Not: false},
		//     },
		//     RightExpressionGroup: nil,
		//     RightVariable: &Variable{Letter: 'C', Not: true},
		// },
	}
	return rules
}
