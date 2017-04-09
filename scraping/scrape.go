package scraping

import (
	"github.com/pddg/go-lsd-alfred/models"
	"github.com/PuerkitoBio/goquery"
)

func ScrapeMeaning(doc *goquery.Document, response *[]models.ResponseItem) error {
	base_url := "https://lsd-project.jp"
	doc.Find("div.caption").Each(func(_ int, caption *goquery.Selection) {
		subtitle := caption.Find("span.headword").Text()
		meanings := caption.Next()
		explanation := meanings.Find("span.explanation").First().Text()
		meanings.Find("span.headword").Each(func(_ int, meaning *goquery.Selection) {
			item := new(models.ResponseItem)
			item.Subtitle = subtitle
			if len(explanation) != 0 {
				item.Text.Largetype = explanation
			}
			meaning_a := meaning.Find("a")
			var headword string
			if meaning_a != nil {
				headword = meaning_a.Text()
				item.Title, item.Text.Copy, item.Autocomplete = headword, headword, headword
				url, link_exists := meaning_a.Attr("href")
				if link_exists {
					item.Arg = headword
					item.Valid = true
					// when press `shift` key
					item.Mod.Shift.Arg = base_url + url
					item.Mod.Shift.Valid = true
					item.Mod.Shift.Subtitle = "Open: " + base_url + url
				}
			} else {
				headword = meaning.Text()
				item.Title, item.Text.Copy, item.Autocomplete = headword, headword, headword
			}
			// when press `cmd` key
			item.Mod.Cmd.Valid = true
			item.Mod.Cmd.Subtitle = "Copy '" + headword + "' to Clipboard"
			item.Mod.Cmd.Arg = headword
			*response = append(*response, *item)
		})
	})
	if len(*response) == 0 {
		item := new(models.ResponseItem)
		item.Title = "No result found."
		*response = append(*response, *item)
	}
	return nil
}
