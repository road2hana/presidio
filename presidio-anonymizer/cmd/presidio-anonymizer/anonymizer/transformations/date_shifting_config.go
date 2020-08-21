package transformations

import (
	"fmt"
	"time"
	"regexp"
	"strings"
	"strconv"

	types "github.com/Microsoft/presidio-genproto/golang"
)

var datePatterns map[string]string
var sepSymbols = [4]string {"-"," ","/","."}

//ShiftDate ...
func ShiftDateValue(text string, location types.Location, daysSinceMomentZero int32) (string, error) {
	if location.Length == 0 {
		location.Length = location.End - location.Start
	}
	pos := location.Start + location.Length
	if int32(len(text)) < pos {
		return "", fmt.Errorf("Indexes for values: are out of bounds")
	}
	// get Date value string
	dateValue := text[location.Start:pos]
	shiftedDateValue, err := shiftDate(dateValue,daysSinceMomentZero)
	var textWithDateShifted string
	fmt.Println("err:", err)
	if err == nil {
		textWithDateShifted = replaceValueInString(text, shiftedDateValue,int(location.Start),int(pos))
	}else{
		textWithDateShifted = replaceValueInString(text,"<DATE>", int(location.Start),int(pos))
	}
	fmt.Println("shifted date:", shiftedDateValue)
	return textWithDateShifted, nil
}

//ShiftDateToDaysSinceMomentZeroValue
func ShiftDateToDaysSinceMomentZeroValue(text string, location types.Location, dateOfMomentZero string, dateLayout string) (string, error) {
	if location.Length == 0 {
		location.Length = location.End - location.Start
	}
	pos := location.Start + location.Length
	if int32(len(text)) < pos {
		return "", fmt.Errorf("Indexes for values: are out of bounds")
	}
	// get Date value string
	dateValue := text[location.Start:pos]

	shiftedDateValue, err := shiftDateToDays(dateValue,dateOfMomentZero, dateLayout)
	var textWithDateShifted string
	replacedDateValue := "<DaysSM>" + strconv.FormatInt(shiftedDateValue,10) + "</DaysSM>"
	if err == nil {
		textWithDateShifted = replaceValueInString(text, replacedDateValue,int(location.Start),int(pos))
	}else{
		textWithDateShifted = replaceValueInString(text,"<DATE>", int(location.Start),int(pos))
	}
	fmt.Println("shifted date:", shiftedDateValue)
	return textWithDateShifted, nil
}



func init(){
	datePatterns = make(map[string]string)
	datePatterns["(19|20)\\d\\d([-/.]| +)(0?[1-9]|1[012])([-/.]| +)(0?[1-9]|[12][0-9]|3[01])"] = "2006-1-2"
	//	datePatterns["(0?[1-9]|1[012])([-/.]| +)(0?[1-9]|[12][0-9]|3[01])([-/.]| +)(19|20)\\d\\d"] = "1-2-2006"
	datePatterns["(0?[1-9]|[12][0-9]|3[01])([-/.]| +)(0?[1-9]|1[012])([-/.]| +)(19|20)\\d\\d"] ="2-1-2006"
	datePatterns["(?i)(((31(Jan|Mar|May|Jul|Aug|Oct|Dec))|(0?[1-9]|([12]\\d)|30)((Jan|Mar|May|Apr|Jul|Jun)|Aug|Oct|Sep|Nov|Dec)|((0?[1-9]|1\\d|2[0-8]|(29))Feb))(19|20)\\d\\d)\\b"] = "2Jan2006"
	datePatterns["(?i)(((31(January|March|May|July|August|October|Dec))|(0?[1-9]|([12]\\d)|30)((January|March|May|April|July|June)|August|October|(Sept|Nov|Dec)(ember))|((0?[1-9]|1\\d|2[0-8]|(29))February))(19|20)\\d\\d)\\b"] = "2January2006"
	datePatterns["(?i)(?:(((Jan(uary)?|Ma(r(ch)?|y)|Jul(y)?|Aug(ust)?|Oct(ober)?|Dec(ember)?)([-/.]| +)31)|((Jan(uary)?|Ma(r(ch)?|y)|Apr(il)?|Ju((ly?)|(ne?))|Aug(ust)?|Oct(ober)?|(Sept|Nov|Dec)(ember)?)([-/.]| +)(0?[1-9]|([12]\\d)|30))|(Feb(ruary)?([-/.]| +)(0?[1-9]|1\\d|2[0-8]|(29(([-/.]| +)((1[6-9]|[2-9]\\d)(0[48]|[2468][048]|[13579][26])|((16|[2468][048]|[3579][26])00)))))))([-/.]| +)((1[6-9]|[2-9]\\d)\\d{2}))\\b"] = "Jan-02-2006"
	datePatterns["(?i)(((January|March|May|July|August|October|December)([-/.]| +)31)|((January|March|May|April|July|June)|August|October|(Sept|Nov|Dec)(ember))([-/.]| +)(0[1-9]|([12]\\d)|30)|(February([-/.]| +)(0[1-9]|1\\d|2[0-8]|(29))))\\b"] = "January-02"
	datePatterns["(?i)(((Jan|Mar|May|Jul|Aug|Oct|Dec)([-/.]| +)31)|((Jan|Mar|May|Apr|Jul|Jun)|Aug|Oct|Sep|Nov|Dec)([-/.]| +)(0[1-9]|([12]\\d)|30)|(Feb([-/.]| +)(0[1-9]|1\\d|2[0-8]|(29))))\\b"] = "Jan-06"
	datePatterns["(?i)((31th +of +(January|March|May|July|August|October|December))|(0[1-9]|([12]\\d)|30)th +of +((January|March|May|April|July|June)|August|October|(Sept|Nov|Dec)(ember))|((0[1-9]|1\\d|2[0-8]|(29))th +of + February))\\b"] = "02th of January"
	datePatterns["(?i)((31th +of +(Jan|Mar|May|Jul|Aug|Oct|Dec))|(0[1-9]|([12]\\d)|30)th +of +((Jan|Mar|May|Apr|Jul|Jun)|Aug|Oct|Sep|Nov|Dec)|((0[1-9]|1\\d|2[0-8]|(29))th +of +Feb))\\b"] = "02th of Jan"
	datePatterns["(?i)((31([-/.]| +)(January|March|May|July|August|October|December))|(0?[1-9]|([12]\\d)|30)([-/.]| +)((January|March|May|April|July|June)|August|October|(Sept|Nov|Dec)(ember))|((0?[1-9]|1\\d|2[0-8]|(29))([-/.]| +)February))\\b"] = "2 January"
	datePatterns["(?i)((31([-/.]| +)(Jan|Mar|May|Jul|Aug|Oct|Dec))|(0?[1-9]|([12]\\d)|30)([-/.]| +)((Jan|Mar|May|Apr|Jul|Jun)|Aug|Oct|Sep|Nov|Dec)|((0?[1-9]|1\\d|2[0-8]|(29))([-/.]| +)Feb))\b"] = "2 Jan"
	datePatterns["(0?[1-9]|[12][0-9]|3[01])([-/.]| +)(11|12|9)"] = "2.1"
	//datePatterns["((31([-/.]| +)(10|12|1|3|5|7|8))|(0?[1-9]|([12]\\d)|30)([-/.]| +)((10|11|12)|(1|3|5|4|7|6)|8|9)|((0?[1-9]|1\\d|2[0-8]|(29))([-/.]| +)2))\\b"] = "2.1"


}

func shiftDate(dateValue string, daysSinceMomentZero int32)(string, error){
	// get a date layout matches the date value
	dateLayout, err := parseDateLayout(dateValue)
	fmt.Println("date layout:", dateLayout)
	// parset date
	dateObj, err := time.Parse(dateLayout,dateValue)
	fmt.Println("dateObj:", dateObj)
	// shift date with daysSinceMomentZero
	shiftedDateObj := dateObj.Add(time.Hour * 24 * time.Duration(daysSinceMomentZero))
	shiftedDate := shiftedDateObj.Format(dateLayout)

	return shiftedDate,err
}

func shiftDateToDays(dateValue string, dateOfMomentZero string, momentZeroDateLayout string)(int64 , error){
	// get a dae layout matches the date value
	dateLayout, err := parseDateLayout(dateValue)
	fmt.Println("date layout:", dateLayout)
	// parset date
	dateObj, err := time.Parse(dateLayout, dateValue)
	fmt.Println("dateObj:", dateObj)
	// parset moment zero
	momentZero, err := time.Parse(momentZeroDateLayout, dateOfMomentZero)
	// shift date to days since moment zero
	diff := dateObj.Sub(momentZero)
	days := int64(diff.Hours()/24)
	return days, err
	
}

func parseDateLayout( dateValue string) (string, error){
	// check date pattern that matches the dateValue
	var datePatternMatched string
	for dateReg, dateLayout := range datePatterns {
		var dateValidtor = regexp.MustCompile(dateReg)
		var matched = dateValidtor.MatchString(dateValue)

		fmt.Println("dateReg",dateReg, "dateLayout",dateLayout, "dateValue", dateValue, "matched:", matched)
		if matched {
			datePatternMatched = dateReg
			break
		}
	}
	if datePatternMatched == "" {
		return "",fmt.Errorf("No date pattern matched the date value:"+dateValue)
	}
	// get default dateLayout
	dateLayout := datePatterns[datePatternMatched]
	// get date seperator in dateValue
	defaultSep, newSep := parseDateSeprator(dateLayout,dateValue)
	// convert the default date layout to a new one matching the style of dateValue
	newDateLayout := strings.Replace(dateLayout,defaultSep, newSep,-1)

	return newDateLayout,nil
}

func parseDateSeprator(dateLayout string, dateValue string)(string,string){
	var defaultSep,newSep = "-","-"
	for _, sepSymbol := range sepSymbols {
		// get defaultSep
		if strings.Contains(dateLayout,sepSymbol) {
			defaultSep = sepSymbol
			fmt.Println("defaultSet", defaultSep)
		}
		if strings.Contains(dateValue,sepSymbol) {
			newSep = sepSymbol
			fmt.Println("newSep",newSep)
		}
	}
	return defaultSep, newSep
}
