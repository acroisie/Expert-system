package algo

import (
	"fmt"
	"expert/models"
)

func Solve(rules []models.Rule, facts []models.Fact) {
	fmt.Println("Solving...")
	for {
		updated := false
		for _, rule := range rules {
			if rule.Evaluate(facts) {
				updated = true
				facts = append(facts, models.Fact{Letter: rule.Conclusion.Letter, Value: rule.Conclusion.Value, Initial: false, Reason: models.Reason{Msg: rule.String()}})
			}
		}
		if !updated {
			break
		}
	}
	fmt.Println("Solved!")
}
