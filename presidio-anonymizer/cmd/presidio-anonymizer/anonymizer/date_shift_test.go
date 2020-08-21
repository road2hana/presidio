package anonymizer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	types "github.com/Microsoft/presidio-genproto/golang"
)

var testDatePlans = []struct {
	desc                    string
	text                    string
	expected                string
	analyzeResults          []*types.AnalyzeResult
	fieldTypeTransformation []*types.FieldTypeTransformation
	defaultTransformation   *types.Transformation
}{
	// Replace  date time fields with a shifting transformation which was declared for ALL fields
	{
		desc:     "Shifting date time fields",
		text:     "My name is David and I live in Miami. I traveled on 12/03/2018 to the beach, and returned on April 20. I need to go to school on 13th of June, and back to home on 11.11",
		expected: "My name is <PERSON> and I live in <LOCATION>. I traveled on <DaysSM>315</DaysSM> to the beach, and returned on <DaysSM>-106751</DaysSM>. I need to go to school on <DaysSM>-106751</DaysSM>, and back to home on <DaysSM>-106751</DaysSM>",
		analyzeResults: []*types.AnalyzeResult{{
			Location: &types.Location{
				Start: 52,
				End:   62,
			},
			Field: &types.FieldTypes{
				Name: types.FieldTypesEnum_DATE_TIME.String(),
			},
		},
                {
                        Location: &types.Location{
				Start: 11,
				End:   16,
			},
                        Score: 0.85,
			Field: &types.FieldTypes{
				Name: types.FieldTypesEnum_PERSON.String(),
			},
                },
                {
                        Location: &types.Location{
				Start: 93,
				End:   101,
			},
                        Score: 0.85,
			Field: &types.FieldTypes{
				Name: types.FieldTypesEnum_DATE_TIME.String(),
			},
                },
                {
                        Location: &types.Location{
				Start: 129,
				End:   141,
			},
                        Score: 0.85,
			Field: &types.FieldTypes{
				Name: types.FieldTypesEnum_DATE_TIME.String(),
			},
                },
                {
                        Location: &types.Location{
				Start: 31,
				End:   36,
			},
                        Score: 0.85,
			Field: &types.FieldTypes{
				Name: types.FieldTypesEnum_LOCATION.String(),
			},
                
                },
                {
			Location: &types.Location{
				Start: 93,
				End:   101,
			},
			Field: &types.FieldTypes{
				Name: types.FieldTypesEnum_DATE_TIME.String(),
			},
		},
                {
			Location: &types.Location{
				Start: 163,
				End:   168,
			},
			Field: &types.FieldTypes{
				Name: types.FieldTypesEnum_DATE_TIME.String(),
			},
		},			
                },
		fieldTypeTransformation: []*types.FieldTypeTransformation{{
		        Fields: []*types.FieldTypes{{
                                Name: types.FieldTypesEnum_DATE_TIME.String(),
		        }},
			Transformation: &types.Transformation{
				ShiftDateToDaysSinceMomentZeroValue: &types.ShiftDateToDaysSinceMomentZeroValue{
					DateOfMomentZero: "1/5/2017",
					DateLayout: "2/1/2006",
				},
			},
		},
                {
	        	Fields: []*types.FieldTypes{{
		        	Name: types.FieldTypesEnum_LOCATION.String(),
		                }},
		        Transformation: &types.Transformation{
			        ReplaceValue: &types.ReplaceValue{
				NewValue: "<LOCATION>",
			        },
		        },
	        },
                {
	        	Fields: []*types.FieldTypes{{
			        Name: types.FieldTypesEnum_PERSON.String(),
		                }},
		        Transformation: &types.Transformation{
			        ReplaceValue: &types.ReplaceValue{
				NewValue: "<PERSON>",
			        },
		        },
	        },                
                },
	},
	// Transform date with  customized DATE recognizer and shiftDate value
	{
		desc:     "Bought 29 June. Call 3 time Lazada",
		text:     "Bought 29 June. Call 3 time Lazada",
		expected: "Bought <DaysSM>-106751</DaysSM>. Call 3 time Lazada",
		analyzeResults: []*types.AnalyzeResult{{
			Location: &types.Location{
				Start: 7,
				End:   13,
			},
			Field: &types.FieldTypes{
				Name: "DATE",
			},
		},
		{
			Location: &types.Location{
				Start: 7,
				End:   14,
			},
			Field: &types.FieldTypes{
				Name: "DATE",
			},
		},
		},
		fieldTypeTransformation: []*types.FieldTypeTransformation{{
			Transformation: &types.Transformation{
				ShiftDateToDaysSinceMomentZeroValue: &types.ShiftDateToDaysSinceMomentZeroValue{
					DateOfMomentZero: "1/5/2017",
					DateLayout: "2/1/2006",
				},
			},
		}},
	},
	// Transform date with  customized DATE recognizer with 2 diffent location
	{
		desc:     "2 years duration (10June2015)  I didn''t use it until today (15Oct2016). ",
		text:     "2 years duration (10June2015)  I didn''t use it until today (15Oct2016). ",
		expected: "2 years duration (<DaysSM>-691</DaysSM>)  I didn''t use it until today (<DaysSM>-198</DaysSM>). ",
		analyzeResults: []*types.AnalyzeResult{{
			Location: &types.Location{
				Start: 18,
				End:   28,
			},
			Field: &types.FieldTypes{
				Name: "DATE",
			},
		},
		{
			Location: &types.Location{
				Start: 61,
				End:   70,
			},
			Field: &types.FieldTypes{
				Name: "DATE",
			},
		},
		},
		fieldTypeTransformation: []*types.FieldTypeTransformation{{
			Transformation: &types.Transformation{
				ShiftDateToDaysSinceMomentZeroValue: &types.ShiftDateToDaysSinceMomentZeroValue{
					DateOfMomentZero: "1/5/2017",
					DateLayout: "2/1/2006",
				},
			},
		}},
	},
	// test 11 july
	{
		desc:     "For your information I placed the order on 13/7/2017. My parcel should be arrived on 19/7/2017 but parcel failed to deliver. I was informed by lazada that the items will be refund without \the logic reason. But why until today 7/8/2017 lazada haven’t received any refund yet??? Please call me asap or email the explanation",
		text:     "For your information I placed the order on 13/7/2017. My parcel should be arrived on 19/7/2017 but parcel failed to deliver. I was informed by lazada that the items will be refund without the logic reason. But why until today 7/8/2017 lazada haven’t received any refund yet??? Please call me asap or email the explanation",
		expected: "For your information I placed the order on <DaysSM>73</DaysSM>. My parcel should be arrived on <DaysSM>79</DaysSM> but parcel failed to deliver. I was informed by lazada that the items will be refund without the logic reason. But why until today <DaysSM>98</DaysSM> lazada haven’t received any refund yet??? Please call me asap or email the explanation",
		analyzeResults: []*types.AnalyzeResult{{
			Location: &types.Location{
				Start: 43,
				End:   52,
			},
			Field: &types.FieldTypes{
				Name: "DATE",
			},
		},
		{
			Location: &types.Location{
				Start: 85,
				End:   94,
			},
			Field: &types.FieldTypes{
				Name: "DATE",
			},
		},
		{
			Location: &types.Location{
				Start: 226,
				End:   234,
			},
			Field: &types.FieldTypes{
				Name: "DATE",
			},
		},			
		},
		fieldTypeTransformation: []*types.FieldTypeTransformation{{
			Transformation: &types.Transformation{
				ShiftDateToDaysSinceMomentZeroValue: &types.ShiftDateToDaysSinceMomentZeroValue{
					DateOfMomentZero: "1/5/2017",
					DateLayout: "2/1/2006",
				},
			},
		}},
	},
	
}

func TestDatePlan(t *testing.T) {
	for _, plan := range testDatePlans {
		t.Logf("Testing %s", plan.desc)

		anonymizerTemplate := types.AnonymizeTemplate{
			FieldTypeTransformations: plan.fieldTypeTransformation,
			DefaultTransformation:    plan.defaultTransformation,
		}
		output, err := AnonymizeText(plan.text, plan.analyzeResults, &anonymizerTemplate)
		assert.NoError(t, err)
		assert.Equal(t, plan.expected, output)
	}
}
