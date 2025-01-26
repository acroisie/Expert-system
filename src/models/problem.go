package models

import (
    "expert-system/src/factManager"
    "expert-system/src/rules"
)

type Problem struct {
	Facts []factManager.Fact
	Rules []rules.Rule
	Queries []Query
}
