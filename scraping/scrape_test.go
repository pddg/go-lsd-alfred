package scraping

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/pddg/go-lsd-alfred/models"
	"io/ioutil"
	"strings"
	"testing"
)

func getDoc(t *testing.T) *goquery.Document {
	html, err := ioutil.ReadFile("./sample.html")
	if err != nil {
		t.Fatalf(err.Error())
	}
	stringReader := strings.NewReader(string(html))
	doc, err := goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		t.Fatalf(err.Error())
	}
	return doc
}

func TestFindMeaning(t *testing.T) {
	expected := []string{
		"/weblsd/c/begin/%E7%B4%B0%E8%83%9E",
		"/weblsd/c/begin/%E3%82%BB%E3%83%AB",
		"/weblsd/c/begin/%E7%B4%B0%E8%83%9E%E7%B5%84%E7%B9%94%E7%99%82%E6%B3%95",
		"/weblsd/c/begin/%E6%9A%97%E9%BB%99",
		"/weblsd/c/begin/%E6%BD%9C%E5%9C%A",
	}
	caption_count := 0
	mean_count := 0
	href_count := 0
	doc := getDoc(t)
	doc.Find("div.caption").Each(func(_ int, caption *goquery.Selection) {
		caption_count++
		findMeaning(caption).Each(func(_ int, meaning *goquery.Selection) {
			mean_count++
			meaning.Find("a").Each(func(_ int, link *goquery.Selection) {
				url, exists := link.Attr("href")
				if !exists {
					t.Errorf("href couldn't find in 'meaning' section.")
				}
				if url != expected[href_count] {
					t.Errorf("Scraped URL is %v, expected is %v", url, expected[href_count])
				}
				href_count++
			})
		})
	})
	if caption_count != 2 {
		t.Errorf("Number of scraped 'div.caption' is %v, but expected is %v", caption_count, 2)
	}
	if mean_count != 2 {
		t.Errorf("Number of scraped 'div.meaning' is %v, but expected is %v", mean_count, 2)
	}
	if href_count != 3 {
		t.Errorf("Number of scraped 'a[href]' is %v, but expected is %v", href_count, 3)
	}
}

func TestFindHeadword(t *testing.T) {
	expected_headword := []string{"細胞", "セル", "細胞組織療法", "暗黙の", "潜在している"}
	expected_url := []string{
		"/weblsd/c/begin/%E7%B4%B0%E8%83%9E",
		"/weblsd/c/begin/%E3%82%BB%E3%83%AB",
		"/weblsd/c/begin/%E7%B4%B0%E8%83%9E%E7%B5%84%E7%B9%94%E7%99%82%E6%B3%95",
		"/weblsd/c/begin/%E6%9A%97%E9%BB%99",
		"/weblsd/c/begin/%E6%BD%9C%E5%9C%A",
	}
	doc := getDoc(t)
	count := 0
	doc.Find("div.caption").Each(func(_ int, caption *goquery.Selection) {
		findMeaning(caption).Find("span.headword").Each(func(_ int, meaning *goquery.Selection) {
			headword, url := findHeadword(meaning)
			if headword != expected_headword[count] {
				t.Errorf("Scraped 'span.headword' is %v, but expected is %v", headword, expected_headword[count])
			}
			if url != expected_url[count] {
				t.Errorf("Scraped 'a[href]' is %v, but expected is %v", url, expected_url[count])
			}
			count++
		})
	})
	if count != 3 {
		t.Errorf("Number of scraped results is %v, but expected is %v", count, 3)
	}
}

func TestScrapeMeaning(t *testing.T) {
	doc := getDoc(t)
	res := new(models.Response)
	ScrapeMeaning(doc, &res.Items)
	expected_items := createTestItems()
	for i := 0; i < len(res.Items); i++ {
		item := res.Items[i]
		expected := expected_items[i]
		if item != expected {
			t.Errorf("Response item is \n%+v\nbut expected is \n%+v", item, expected)
		}
	}
}

func createTestItems() []models.ResponseItem {
	items := []models.ResponseItem{}
	headwords := []string{"細胞", "セル", "細胞組織療法", "暗黙の", "潜在している"}
	cap_headwords := []string{"cell", "cell", "cell- and tissue-based therapy", "implicit"}
	urls := []string{
		BaseUrl + "/weblsd/c/begin/%E7%B4%B0%E8%83%9E",
		BaseUrl + "/weblsd/c/begin/%E3%82%BB%E3%83%AB",
		BaseUrl + "/weblsd/c/begin/%E7%B4%B0%E8%83%9E%E7%B5%84%E7%B9%94%E7%99%82%E6%B3%95",
		BaseUrl + "/weblsd/c/begin/%E6%9A%97%E9%BB%99",
		BaseUrl + "/weblsd/c/begin/%E6%BD%9C%E5%9C%A",
	}
	for i := 0; i < 3; i++ {
		item := new(models.ResponseItem)
		item.Title, item.Arg, item.Autocomplete = headwords[i], headwords[i], headwords[i]
		item.Valid = true
		item.Subtitle = cap_headwords[i]
		item.Text.Copy = headwords[i]
		item.Mod.Cmd.Subtitle = "Copy '" + headwords[i] + "' to Clipboard"
		item.Mod.Cmd.Arg = headwords[i]
		item.Mod.Cmd.Valid = true
		item.Mod.Shift.Subtitle = "Open: " + urls[i]
		item.Mod.Shift.Arg = urls[i]
		item.Mod.Shift.Valid = true
		items = append(items, *item)
	}
	return items
}
