
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
var dateEndSymbols =[]string{" ",".","!",")"}

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
func ShiftDateToDaysSinceMomentZeroValue(text string, location types.Location, dateOfMomentZero string, dateLayout string, anonymizeTextContext *types.AnonymizeTextContext) (string, error) {
	if location.Length == 0 {
		location.Length = location.End - location.Start
	}
	pos := location.Start + location.Length
	if int32(len(text)) < pos {
		return "", fmt.Errorf("Indexes for values: are out of bounds")
	}
        // get Date value string
	textRune := []rune(text)
	dateValueRune := textRune[location.Start:pos]
	dateValue := string(dateValueRune)

	// skipe none-date(dateLayout is 2.1 which sometime is not a valid date value) determined by the context of the date value
	sufix_trimed_date_value, err := trimDateValueSuffix(dateValue) 
	_, default_date_layout, err := parseDateLayout(sufix_trimed_date_value)
	fmt.Println("default_date_layout:", default_date_layout)
	if default_date_layout == "2.1" || default_date_layout == "2-1-06" {
		pre_text := string(textRune[:location.Start])
		after_text := ""
		after_text = string(textRune[location.End:])
		after_none_date_pattern := "^(?i)( +)?((year|years|yrs)|(day|days)|(hour|hours|hrs)|(minutes|min)|(liter|glasses|liters|capacity)|(pax))"
		none_date_matched, _ := regexp.MatchString(after_none_date_pattern, after_text)
		fmt.Println("pre_text: ", pre_text)
		fmt.Println("after_text:", after_text)
		fmt.Println("none_date_matched:", none_date_matched)
		if none_date_matched {
			return text, nil
		}
	}


	shiftedDateValue, err := shiftDateToDays(sufix_trimed_date_value, dateOfMomentZero, dateLayout, anonymizeTextContext)
	var textWithDateShifted string
	var replacedDateValue string
	if shiftedDateValue >0 {
		replacedDateValue = "<DaysSM>" + strconv.FormatInt(shiftedDateValue,10) + "</DaysSM>"
	}else{
		replacedDateValue = "<DaysSM>Invalid</DaysSM>" 
	}
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

	// Day Month Year
	datePatterns["^(19|20)\\d\\d([-/.]| +)(0?[1-9]|1[012])([-/.]| +)(0?[1-9]|[12][0-9]|3[01])$"] = "2006-1-2"
	//	datePatterns["^(0?[1-9]|1[012])([-/.]| +)(0?[1-9]|[12][0-9]|3[01])([-/.]| +)(19|20)\\d\\d$"] = "1-2-2006"	
	datePatterns["^((0?[1-9]|[12][0-9]|3[01])([-/.]| +)(0?[1-9]|1[012])([-/.]| +)(19|20)\\d\\d)$"] ="2-1-2006"
	datePatterns["^((0?[1-9]|[12][0-9]|3[01])([-/.]| +)(0?[1-9]|1[012])([-/.]| +)\\d\\d)$"] ="2-1-06" 
	datePatterns["^(?i)(((31(Jan|Mar|May|Jul|Aug|Oct|Dec))|(0?[1-9]|([12]\\d)|30)((Jan|Mar|May|Apr|Jul|Jun)|Aug|Oct|Sep|Nov|Dec)|((0?[1-9]|1\\d|2[0-8]|(29))Feb))(19|20)\\d\\d)$"] = "2Jan2006"
	datePatterns["^(?i)(((31th([-/.]| +)(Jan|Mar|May|Jul|Aug|Oct|Dec))|(0?[1-9]|([12]\\d)|30)th([-/.]| +)((Jan|Mar|May|Apr|Jul|Jun)|Aug|Oct|Sep|Nov|Dec)|((0?[1-9]|1\\d|2[0-8]|(29))([-/.]| +)Feb))([-/.]| +)(19|20)\\d\\d)$"] = "2th-Jan-2006"
	datePatterns["^(?i)(((31([-/.]| +)(Jan|Mar|May|Jul|Aug|Oct|Dec))|(0?[1-9]|([12]\\d)|30)([-/.]| +)((Jan|Mar|May|Apr|Jul|Jun)|Aug|Oct|Sep|Nov|Dec)|((0?[1-9]|1\\d|2[0-8]|(29))([-/.]| +)Feb))([-/.]| +)\\d\\d)$"]= "2-Jan-06"
	datePatterns["^(?i)(((31(January|March|May|July|August|October|Dec))|(0?[1-9]|([12]\\d)|30)((January|March|May|April|July|June)|August|October|(Sept|Nov|Dec)(ember))|((0?[1-9]|1\\d|2[0-8]|(29))February))(19|20)\\d\\d)$"] = "2January2006"
	datePatterns["^(?i)(?:(((Jan|Mar|May|Jul|Aug|Oct|Dec)([-/.]| +))(31)|(Jan(uary)?|Mar|May|Apr|Jun|Jul|Aug|Oct|Sep|Nov|Dec)([-/.]| +)(0?[1-9]|([12]\\d)|30)|(Feb([-/.]| +)(0?[1-9]|1\\d|2[0-8]|(29(([-/.]| +)((1[6-9]|[2-9]\\d)(0[48]|[2468][048]|[13579][26])|((16|[2468][048]|[3579][26])00)))))))([-/.]| +)((1[6-9]|[2-9]\\d)\\d{2}))$"] = "Jan-02-2006"
	datePatterns["^(?i)(?:((3([-/.]| +)(January|March|May|July|August|October|December))|((0?[1-9]|([12]\\d)|30)([-/.]| +)(January|March|May|April|June|July|August|October|(Sept|Nov|Dec)(ember)))|((0?[1-9]|1\\d|2[0-8]|(29(([-/.]| +)((1[6-9]|[2-9]\\d)(0[48]|[2468][048]|[13579][26])|((16|[2468][048]|[3579][26])00)))))([-/.]| +)February))([-/.]| +)((1[6-9]|[2-9]\\d)\\d{2}))$"] = "2-January-2006"
	datePatterns["^(?i)(?:((3([-/.]| +)(Jan|Mar|May|Jul|Aug|Oct|Dec))|((0?[1-9]|([12]\\d)|30)([-/.]| +)(Jan|Mar|May|Apr|Jun|Jul|Aug|Oct|Sep|Nov|Dec))|((0?[1-9]|1\\d|2[0-8]|(29(([-/.]| +)((1[6-9]|[2-9]\\d)(0[48]|[2468][048]|[13579][26])|((16|[2468][048]|[3579][26])00)))))([-/.]| +)Feb))([-/.]| +)((1[6-9]|[2-9]\\d)\\d{2}))$"] = "2-Jan-2006"
	datePatterns["^(?i)(((31th([-/.]| +)(January|March|May|July|August|October|December))|(0?[1-9]|([12]\\d)|30)th([-/.]| +)((January|March|May|April|July|June)|August|October|September|November|December)|((0?[1-9]|1\\d|2[0-8]|(29))([-/.]| +)February))([-/.]| +)(19|20)\\d\\d)$"] = "2th-January-2006"

	// Month and Year
	datePatterns["^(?i)((?:Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)([-/.]? +)(19|20)\\d\\d)$"] = "Jan-2006"
	datePatterns["^(?i)(?:January|February|March|April|May|June|July|August|September|October|November|December)([-/.]? +)(19|20)\\d\\d$"] = "January-2006"

	
	// Day Month
	datePatterns["^(?i)(((January|March|May|July|August|October|December)([-/.]| +)31)|((January|March|May|April|July|June)|August|October|(Sept|Nov|Dec)(ember))([-/.]| +)(0?[1-9]|([12]\\d)|30)|(February([-/.]| +)(0?[1-9]|1\\d|2[0-8]|(29))))$"] = "January-2"
	datePatterns["^(?i)(((Jan|Mar|May|Jul|Aug|Oct|Dec)([-/.]| +)31)|((Jan|Mar|May|Apr|Jul|Jun)|Aug|Oct|Sep|Nov|Dec)([-/.]| +)(0?[1-9]|([12]\\d)|30)|(Feb([-/.]| +)(0?[1-9]|1\\d|2[0-8]|(29))))\\b$"] = "Jan-2"

	datePatterns["^(?i)((31th +of +(January|March|May|July|August|October|December))|(0?[1-9]|([12]\\d)|30)th +of +((January|March|May|April|July|June)|August|October|(Sept|Nov|Dec)(ember))|((0?[1-9]|1\\d|2[0-8]|(29))th +of +February))$"] = "2th of January"
	datePatterns["^(?i)((31st +(January|March|May|July|August|October|December))|(0?[1-9]|([12]\\d)|30)st +((January|March|May|April|July|June)|August|October|(Sept|Nov|Dec)(ember))|((0?[1-9]|1\\d|2[0-8]|(29))st +February))$"] = "2st January"
	datePatterns["^(?i)((31th +(Jan|Mar|May|Jul|Aug|Oct|Dec))|(0?[1-9]|([12]\\d)|30)th +((Jan|Mar|May|Apr|Jul|Jun)|Aug|Oct|Sep|Nov|Dec)|((0?[1-9]|1\\d|2[0-8]|(29))th +Feb))$"] = "2th Jan"
	datePatterns["^(?i)((31th +(January|March|May|July|August|October|December))|(0?[1-9]|([12]\\d)|30)th +((January|March|May|April|July|June)|August|October|(Sept|Nov|Dec)(ember))|((0?[1-9]|1\\d|2[0-8]|(29))th +February))$"] = "2th January"
	datePatterns["^(?i)((31th +of +(Jan|Mar|May|Jul|Aug|Oct|Dec))|(0?[1-9]|([12]\\d)|30)th +of +((Jan|Mar|May|Apr|Jul|Jun)|Aug|Oct|Sep|Nov|Dec)|((0?[1-9]|1\\d|2[0-8]|(29))th +of +Feb))$"] = "2th of Jan"
	 datePatterns["^(?i)((31 +of +(Jan|Mar|May|Jul|Aug|Oct|Dec))|(0?[1-9]|([12]\\d)|30) +of +((Jan|Mar|May|Apr|Jul|Jun)|Aug|Oct|Sep|Nov|Dec)|((0?[1-9]|1\\d|2[0-8]|(29)) +of +Feb))$"] = "2 of Jan"
	datePatterns["^(?i)((31st +(Jan|Mar|May|Jul|Aug|Oct|Dec))|(0?[1-9]|([12]\\d)|30)st +((Jan|Mar|May|Apr|Jul|Jun)|Aug|Oct|Sep|Nov|Dec)|((0?[1-9]|1\\d|2[0-8]|(29))st +of +Feb))$"] = "2st Jan"
	datePatterns["^(?i)((31([-/.]| +)(January|March|May|July|August|October|December))|(0?[1-9]|([12]\\d)|30)([-/.]| +)((January|March|May|April|July|June)|August|October|(Sept|Nov|Dec)(ember))|((0?[1-9]|1\\d|2[0-8]|(29))([-/.]| +)February))$"] = "2 January"

	datePatterns["^(?i)((31([-/.]| +)(Jan|Mar|May|Jul|Aug|Oct|Dec))|(0?[1-9]|([12]\\d)|30)([-/.]| +)((Jan|Mar|May|Apr|Jul|Jun)|Aug|Oct|Sep|Nov|Dec)|((0?[1-9]|1\\d|2[0-8]|(29))([-/.]| +)Feb))$"] = "2 Jan"
	datePatterns["^(0?[1-9]|[12][0-9]|3[01])([-/.]| +)(11|12|9)$"] = "2.1"
	datePatterns["^(?:)((31([-/.]| +)(10|12|1|3|5|7|8))|(0?[1-9]|([12]\\d)|30)([-/.]| +)((10|11|12)|0?(1|3|5|4|7|6)|8|9)|((0?[1-9]|1\\d|2[0-8]|(29))([-/.)]| +)2))$"] = "2.1"

	// Month
	datePatterns["^(?i)((?:Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec))$"] = "Jan"


}

func shiftDate(dateValue string, daysSinceMomentZero int32)(string, error){
	// get a date layout matches the date value
	dateLayout, _, err := parseDateLayout(dateValue)
	fmt.Println("date layout:", dateLayout)
	// parset date
	dateObj, err := time.Parse(dateLayout,dateValue)
	fmt.Println("dateObj:", dateObj)
	// shift date with daysSinceMomentZero
	shiftedDateObj := dateObj.Add(time.Hour * 24 * time.Duration(daysSinceMomentZero))
	shiftedDate := shiftedDateObj.Format(dateLayout)

	return shiftedDate,err
}

func shiftDateToDays(dateValue string, dateOfMomentZero string, momentZeroDateLayout string, anonymizeTextContext *types.AnonymizeTextContext)(int64 , error){
	// get a dae layout matches the date value
	dateLayout, _, err := parseDateLayout(dateValue)

	// parset date
	dateObj, err := time.Parse(dateLayout, dateValue)
	//	fmt.Println(dateObj, dateLayout)
	// deal with date in free text without specifing Year, for example, 11/11. The solution is leveraging the AnonymizeTextContext information
	if !dateObj.IsZero() && dateObj.Year() == 0 {
		fmt.Println("Missing Year!")
		//		fmt.Println("text createDate:", anonymizeTextContext)
		if anonymizeTextContext != nil && anonymizeTextContext.CreateDate != "" {
			textCreateDate_layout, _ , _ := parseDateLayout(anonymizeTextContext.CreateDate)
			textCreateDate, _ := time.Parse(textCreateDate_layout, anonymizeTextContext.CreateDate)
			dateObj = dateObj.AddDate(textCreateDate.Year(),0,0)
			fmt.Println("textCreateDate:",textCreateDate)
		}
	}
	fmt.Println("dateObj:", dateObj)
	// parset moment zero
	momentZero, err := time.Parse(momentZeroDateLayout, dateOfMomentZero)
	// shift date to days since moment zero
	diff := dateObj.Sub(momentZero)
	days := int64(diff.Hours()/24)
	return days, err
	
}

func parseDateLayout( dateValue string) (string,string, error){
	// check date pattern that matches the dateValue
	var datePatternMatched string
	for dateReg, dateLayout := range datePatterns {
		var dateValidtor = regexp.MustCompile(dateReg)
		var matched = dateValidtor.MatchString(dateValue)

		fmt.Println("dateReg",dateReg, "dateLayout",dateLayout, "dateValue", dateValue,"length",len(dateValue), "matched:", matched)
		if matched {
			datePatternMatched = dateReg
			break
		}
	}
	if datePatternMatched == "" {
		return "","", fmt.Errorf("No date pattern matched the date value:"+dateValue)
	}
	// get default dateLayout
	dateLayout := datePatterns[datePatternMatched]
	// get date seperator in dateValue
	defaultSep, newSep := parseDateSeprator(dateLayout,dateValue)
	// convert the default date layout to a new one matching the style of dateValue
	newDateLayout := strings.Replace(dateLayout,defaultSep, newSep,-1)

	return newDateLayout,dateLayout,nil
}

func parseDateSeprator(dateLayout string, dateValue string)(string,string){
	var defaultSep,newSep = "-","-"
	for _, sepSymbol := range sepSymbols {
		// get defaultSep
		if strings.Contains(dateLayout,sepSymbol) {
			defaultSep = sepSymbol
			fmt.Println("defaultSep", defaultSep)
		}
		if strings.Contains(dateValue,sepSymbol) {
			newSep = sepSymbol
			fmt.Println("newSep",newSep)
		}
	}
	return defaultSep, newSep
}

func trimDateValueSuffix(dateValue string)(string, error){
	var newDateValue string
	newDateValue = dateValue
	for _, sepSymbol := range dateEndSymbols {
		newDateValue = strings.TrimSuffix(newDateValue, sepSymbol)
	}
	return newDateValue, nil
}
