package scraping

import (
	"github.com/pddg/go-lsd-alfred/models"
	"github.com/PuerkitoBio/goquery"
)

func GetPage(mode string, query string) (*goquery.Document, error) {
	return goquery.NewDocument("https://lsd-project.jp/weblsd/" + mode + "/" + query)
}

func CreateError(err error, msg string) *models.ResponseItem {
	item := new(models.ResponseItem)
	item.Title = err.Error()
	item.Subtitle = msg
	item.Valid = false
	return item
}

func CreateOrigin(mode string, query string) *models.ResponseItem {
	url := "https://lsd-project.jp/weblsd/" + mode + "/" + query
	item := new(models.ResponseItem)
	item.Title = "See all results in web site ..."
	item.Subtitle = "Please input `Shift` + `Enter`"
	item.Arg = url
	item.Mod.Shift.Arg = url
	item.Mod.Shift.Valid = true
	item.Mod.Shift.Subtitle = "Open: " + url
	return item
}