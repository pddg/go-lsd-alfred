package scraping

import (
	"errors"
	"github.com/pddg/go-lsd-alfred/models"
	"golang.org/x/text/unicode/norm"
	"reflect"
	"testing"
	"unicode/utf8"
)

func TestNormalizeQuery(t *testing.T) {
	query_string := "あかﾊﾟジャマ青パｼﾞｬﾏ"
	normalized := NormalizeQuery(query_string)
	expected := norm.NFC.String(query_string)
	for i, w := 0, 0; i < len(normalized); i += w {
		normalizedRune, width := utf8.DecodeRuneInString(normalized[i:])
		expectedRune, _ := utf8.DecodeRuneInString(expected[i:])
		if normalizedRune != expectedRune {
			t.Errorf("Normalized result is '%U', but expected is '%U'", normalizedRune, expectedRune)
		}
		w += width
	}
}

func TestCreateModItem(t *testing.T) {
	keys := []string{ShiftKey, CommandKey}
	expected_subtitles := []string{"Shift is pressed", "Command is pressed"}
	expected_arg := []string{"shift", "command"}
	expected_valids := []bool{true, false}
	check := func(k string, val interface{}, ex interface{}, res *models.ModItems) {
		result := reflect.ValueOf(val).Interface()
		expect := reflect.ValueOf(ex).Interface()
		if result != expect {
			t.Errorf("'%v' section's value is invalid. The result is '%v', but expected is '%v'. Response is as follows\n%+v",
				k, result, expect, res)
		}
	}
	item := new(models.ResponseItem)
	createModItem(item, keys[0], expected_subtitles[0], expected_arg[0], expected_valids[0])
	check(keys[0], item.Mod.Shift.Subtitle, expected_subtitles[0], &item.Mod.Shift)
	check(keys[0], item.Mod.Shift.Arg, expected_arg[0], &item.Mod.Shift)
	check(keys[0], item.Mod.Shift.Valid, expected_valids[0], &item.Mod.Shift)
	createModItem(item, keys[1], expected_subtitles[1], expected_arg[1], expected_valids[1])
	check(keys[1], item.Mod.Cmd.Subtitle, expected_subtitles[1], &item.Mod.Cmd)
	check(keys[1], item.Mod.Cmd.Arg, expected_arg[1], &item.Mod.Cmd)
	check(keys[1], item.Mod.Cmd.Valid, expected_valids[1], &item.Mod.Cmd)
}

func TestCreateOrigin(t *testing.T) {
	mode, query := "begin", "パネキシン1"
	url := BaseUrl + "/weblsd/" + norm.NFC.String(mode+"/"+query)
	res := CreateOrigin(mode, query)
	check := func(value string, expected string, section string, response *models.ResponseItem) {
		if value != expected {
			t.Errorf("The %v of created origin response is '%v', but expected is '%v'.\nResponse is as follows\n%+v",
				section, value, expected, response)
		}
	}
	check(res.Title, "See all results in web site ...", "title", res)
	check(res.Subtitle, "Please input `Shift` + `Enter`", "subtitle", res)
	check(res.Arg, url, "arg", res)
	check(res.Mod.Shift.Arg, url, "Mod.Shift.Arg", res)
	check(res.Mod.Shift.Subtitle, "Open: "+url, "Mod.Shift.Subtitle", res)
	if res.Valid {
		t.Errorf("Origin response sholud not be valid. Response is as follows\n%+v", res)
	}

}

func TestCreateError(t *testing.T) {
	msg := "This is a unittest"
	err_msg := "Test error"
	err := errors.New(err_msg)
	res := CreateError(err, msg)
	check := func(value string, expected string, section string, response *models.ResponseItem) {
		if value != expected {
			t.Errorf("The %v of created error response is '%v', but expected is '%v'.\nResponse is as follows\n%+v",
				section, value, expected, response)
		}
	}
	check(res.Title, err_msg, "title", res)
	check(res.Subtitle, msg, "subtitle", res)
	if res.Valid {
		t.Errorf("Error response sholud not be valid. Response is as follows\n%+v", res)
	}
}
