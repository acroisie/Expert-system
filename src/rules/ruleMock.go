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

	}
	return rules
}

	func GetRulesMock2() []Rule {
		rules := []Rule{
			// B => A
			Rule{
				Op: IMPLIES,
				LeftExpressionGroup: nil,
				LeftVariable:        &Variable{Letter: 'B', Not: false},
				RightExpressionGroup: nil,
				RightVariable:       &Variable{Letter: 'A', Not: false},
			},

			// D + E => B
			Rule{
				Op: IMPLIES,
				LeftExpressionGroup: &ExpressionGroup{
					Op:           AND,
					LeftVariable: &Variable{Letter: 'D', Not: false},
					RightVariable: &Variable{Letter: 'E', Not: false},
				},
				RightExpressionGroup: nil,
				RightVariable:        &Variable{Letter: 'B', Not: false},
			},

			// G + H => F
			Rule{
				Op: IMPLIES,
				LeftExpressionGroup: &ExpressionGroup{
					Op:           AND,
					LeftVariable: &Variable{Letter: 'G', Not: false},
					RightVariable: &Variable{Letter: 'H', Not: false},
				},
				RightExpressionGroup: nil,
				RightVariable:        &Variable{Letter: 'F', Not: false},
			},

			// I + J => G
			Rule{
				Op: IMPLIES,
				LeftExpressionGroup: &ExpressionGroup{
					Op:           AND,
					LeftVariable: &Variable{Letter: 'I', Not: false},
					RightVariable: &Variable{Letter: 'J', Not: false},
				},
				RightExpressionGroup: nil,
				RightVariable:        &Variable{Letter: 'G', Not: false},
			},

			// G => H
			Rule{
				Op: IMPLIES,
				LeftExpressionGroup: nil,
				LeftVariable:        &Variable{Letter: 'G', Not: false},
				RightExpressionGroup: nil,
				RightVariable:       &Variable{Letter: 'H', Not: false},
			},

			// L + M => K
			Rule{
				Op: IMPLIES,
				LeftExpressionGroup: &ExpressionGroup{
					Op:           AND,
					LeftVariable: &Variable{Letter: 'L', Not: false},
					RightVariable: &Variable{Letter: 'M', Not: false},
				},
				RightExpressionGroup: nil,
				RightVariable:        &Variable{Letter: 'K', Not: false},
			},

			// O + P => L + N
			Rule{
				Op: IMPLIES,
				LeftExpressionGroup: &ExpressionGroup{
					Op:           AND,
					LeftVariable: &Variable{Letter: 'O', Not: false},
					RightVariable: &Variable{Letter: 'P', Not: false},
				},
				RightExpressionGroup: &ExpressionGroup{
					Op:           AND,
					LeftVariable: &Variable{Letter: 'L', Not: false},
					RightVariable: &Variable{Letter: 'N', Not: false},
				},
			},

			// N => M
			Rule{
				Op: IMPLIES,
				LeftVariable: &Variable{Letter: 'N', Not: false},
				RightVariable:        &Variable{Letter: 'M', Not: false},
			},
		}
		return rules
	}

func GetRulesMock3() []Rule {
	rules := []Rule{
		// A => X | Y
		Rule{
			Op: IMPLIES,
			LeftExpressionGroup: nil,
			LeftVariable: &Variable{Letter: 'A', Not: false},
			RightExpressionGroup: &ExpressionGroup{
				Op: OR,
				LeftVariable: &Variable{Letter: 'X', Not: false},
				RightVariable: &Variable{Letter: 'Y', Not: false},
			},
			RightVariable: nil,
		},
		// X => Z
		Rule{
			Op: IMPLIES,
			LeftExpressionGroup: nil,
			LeftVariable: &Variable{Letter: 'X', Not: false},
			RightExpressionGroup: nil,
			RightVariable: &Variable{Letter: 'Z', Not: false},
		},
		// Y => !Z
		Rule{
			Op: IMPLIES,
			LeftExpressionGroup: nil,
			LeftVariable: &Variable{Letter: 'Y', Not: false},
			RightExpressionGroup: nil,
			RightVariable: &Variable{Letter: 'Z', Not: true},
		},
	}
	return rules
}

func GetRulesMock4() []Rule {
	rules := []Rule{
		// B + C => A
		Rule{
			Op: IMPLIES,
			LeftExpressionGroup: &ExpressionGroup{
				Op: AND,
				LeftVariable: &Variable{Letter: 'B', Not: false},
				RightVariable: &Variable{Letter: 'C', Not: false},
			},
			RightExpressionGroup: nil,
			RightVariable: &Variable{Letter: 'A', Not: false},
		},
		// D | E => B
		Rule{
			Op: IMPLIES,
			LeftExpressionGroup: &ExpressionGroup{
				Op: OR,
				LeftVariable: &Variable{Letter: 'D', Not: false},
				RightVariable: &Variable{Letter: 'E', Not: false},
			},
			RightExpressionGroup: nil,
			RightVariable: &Variable{Letter: 'B', Not: false},
		},
		// B => C
		Rule{
			Op: IMPLIES,
			LeftExpressionGroup: nil,
			LeftVariable: &Variable{Letter: 'B', Not: false},
			RightExpressionGroup: nil,
			RightVariable: &Variable{Letter: 'C', Not: false},
		},
	}
	return rules
}