package controllers

import "github.com/revel/revel"

func init() {
	revel.OnAppStart(InitDB)
	revel.InterceptMethod((*Gorp).Begin, revel.BEFORE)
	revel.InterceptMethod((*Gorp).Commit, revel.AFTER)
	revel.InterceptMethod((*Gorp).Rollback, revel.FINALLY)
}
