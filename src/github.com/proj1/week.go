package main

import (
	"github.com/grsmv/goweek"
)

func days() {
	week, _ := goweek.NewWeek(2015, 46)
	println(week.Days)
}
