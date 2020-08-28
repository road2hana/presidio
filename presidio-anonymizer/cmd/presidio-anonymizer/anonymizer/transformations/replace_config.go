package transformations

import (
	"fmt"
	"strings"
	"regexp"
	types "github.com/Microsoft/presidio-genproto/golang"
)

//ReplaceValue ...
func ReplaceValue(text string, location types.Location, newValue string) (string, error) {
	if location.Length == 0 {
		location.Length = location.End - location.Start
	}
	pos := location.Start + location.Length
	if int32(len(text)) < pos {
		return "", fmt.Errorf("Indexes for values: are out of bounds")
	}
	oldValue := text[location.Start:pos]
	oldValue = strings.ToLower(oldValue)
	// check if the old value is in the white list
	skip_replace := Contains(white_list, oldValue)
	fmt.Println("skip_replace:",skip_replace)
	new := text
	if ! skip_replace {
		new = replaceValueInString(text, newValue, int(location.Start), int(pos))
	}
	return new, nil
}

var white_list = []string {"airfryer","rice", "hurom", "philips","phillips","philli","lazada", "blaze pot","chuck", "grill pan", "cook", "philip", "dayssm", "lamb","wemo", "kolo mee", "roti", "prduct", "glass", "temperaturr", "tong sui", "max", "don''t", "rm1", "sahur", "lang tumagal", "lemang", "nasi minyak","crunchy rice", "frank", "cap", "annuar","mya","everything", "ayam percik", "cheater","lex", "wonderfull", "origina", "sup tulang","severin","lazmall", "seller","goid","looovveee","nasi","renata","terbaik","hpi", "ninja van","mybe", "yang malas", "jug", "mukhang", "kasi namin", "lumalabas na","kala","malas ko", "no crunchy rice", "foc grill", "kpdnkk", "laz","naman", "skynet", "somemore", "midnight", "percikan", "kelangan", "ganda tlaga", "speedy", "php8", "ayamas","peoduct", "rosemary", "blade", "love([ +]|\\w)+","mas mabilis", "pagkaluto", "madali uminit", "nagkadent","basa", "nauna silang", "di ko pa", "maganda", "nung", "gustong","takpelah","yeayyy", "phill","rm39", "rmo", "evthing","tq", "it", "wan tan","harini", "buble", "hangus","deliv", "kurma", "valentine", "kampung", "super lajuu", "pero","masyadong", "napaisip", "kung","di siya", "konti", "awesomeee", "paano na to","maggi", "ayamas", "rosemary", "ady", "pa ibaliktad", "lifetime", "delivery", "caine", " boleh", "lepas", "ang pagkaluto", "boleh", "kasi", "deliv", "madali", "medyo"}
// check if the list contains an element
func Contains(list []string, x string) bool {
	for _, item := range list {
		matched, _ := regexp.MatchString(item, x)
		if matched {
			return true
		}
		
	}
	return false
	
}
