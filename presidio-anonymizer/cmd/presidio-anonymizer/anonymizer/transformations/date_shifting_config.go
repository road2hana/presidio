package transformations

import (
	"fmt"

        "strconv"

	types "github.com/Microsoft/presidio-genproto/golang"
)

//ShiftDate ...
func ShiftDateValue(text string, location types.Location, daysSinceMomentZero int32) (string, error) {
	if location.Length == 0 {
		location.Length = location.End - location.Start
	}
	pos := location.Start + location.Length
	if int32(len(text)) < pos {
		return "", fmt.Errorf("Indexes for values: are out of bounds")
	}
	new := replaceValueInString(text, strconv.Itoa(int(daysSinceMomentZero)), int(location.Start), int(pos))
	return new, nil
}
