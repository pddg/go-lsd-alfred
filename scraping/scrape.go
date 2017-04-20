package scraping

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/pddg/go-lsd-alfred/models"
)

func ScrapeMeaning(doc *goquery.Document, response *[]models.ResponseItem) error {
	doc.Find("div.caption").Each(func(_ int, caption *goquery.Selection) {
		subtitle := caption.Find("span.headword").Text()
		meanings := findMeaning(caption)
		explanation := meanings.Find("span.explanation").First().Text()
		meanings.Find("span.headword").Each(func(_ int, meaning *goquery.Selection) {
			item := new(models.ResponseItem)
			item.Subtitle = subtitle
			if len(explanation) != 0 {
				item.Text.Largetype = explanation
			}
			headword, url := findHeadword(meaning)
			item.Title, item.Text.Copy, item.Autocomplete, item.Arg = headword, headword, headword, headword
			item.Valid = true
			url = BaseUrl + url
			// when press `shift` key
			createModItem(
				item,
				ShiftKey,
				"Open: "+url,
				url,
				true)
			// when press `cmd` key
			createModItem(
				item,
				CommandKey,
				"Copy '"+headword+"' to Clipboard",
				headword,
				true)
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

func findMeaning(caption *goquery.Selection) *goquery.Selection {
	meanings := caption.Next()
	// if next section is "relword" or something different from "meanings", search "meanings"
	if !meanings.HasClass("meaning") {
		meanings = meanings.Next()
	}
	return meanings
}

func findHeadword(meaning *goquery.Selection) (string, string) {
	var (
		headword string
		url      string
	)
	// find "meaning" Text
	if meaning_a := meaning.Find("a"); meaning_a != nil {
		headword = meaning_a.Text()
		if u, exists := meaning_a.Attr("href"); exists {
			url = u
		}
	} else {
		headword = meaning.Text()
	}
	// find "inflection"
	if inflection := meaning.Next(); inflection.HasClass("inflection") {
		headword = headword + inflection.Text()
	}
	return headword, url
}
