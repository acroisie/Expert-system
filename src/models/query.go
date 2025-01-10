package models

import (
	"fmt"
)

type Query struct {
	Letter rune
}

// DISPLAY

func (q Query) String() string {
	return q.getQuery()
}

func (q Query) getQuery() string {
	return string(q.Letter)
}

func DisplayQueries(queries []Query) {
	for i, query := range queries {
		fmt.Printf("%d: %s\n", i+1, query.getQuery())
	}
}