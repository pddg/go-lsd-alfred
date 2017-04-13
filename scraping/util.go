package scraping

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/pddg/go-lsd-alfred/models"
	"golang.org/x/text/unicode/norm"
	"strings"
)

const BaseUrl string = "https://lsd-project.jp"
const ShiftKey string = "shift"
const CommandKey string = "cmd"

func GetPage(mode string, query string) (*goquery.Document, error) {
	return goquery.NewDocument(BaseUrl + "/weblsd/" + NormalizeQuery(mode+"/"+query))
}

func CreateError(err error, msg string) *models.ResponseItem {
	item := new(models.ResponseItem)
	item.Title = err.Error()
	item.Subtitle = msg
	item.Valid = false
	return item
}

func CreateOrigin(mode string, query string) *models.ResponseItem {
	url := BaseUrl + "/weblsd/" + NormalizeQuery(mode+"/"+query)
	item := new(models.ResponseItem)
	item.Title = "See all results in web site ..."
	item.Subtitle = "Please input `Shift` + `Enter`"
	item.Arg = url
	createModItem(item, ShiftKey, "Open: "+url, url, true)
	return item
}

func createModItem(item *models.ResponseItem, keyname string, subtitle string, arg string, valid bool) {
	switch strings.ToLower(keyname) {
	case ShiftKey:
		item.Mod.Shift.Subtitle = subtitle
		item.Mod.Shift.Arg = arg
		item.Mod.Shift.Valid = valid
	case CommandKey:
		item.Mod.Cmd.Subtitle = subtitle
		item.Mod.Cmd.Arg = arg
		item.Mod.Cmd.Valid = valid
	}
}

func NormalizeQuery(query string) string {
	return norm.NFC.String(query)
}
