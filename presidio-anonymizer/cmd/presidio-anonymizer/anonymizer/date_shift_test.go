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
		text:     "My name is David and I live in Miami. I traveled on 12/03/2018 to the beach, and returned on April 20. I need to go to school on 13th of June",
		expected: "My name is <PERSON> and I live in <LOCATION>. I traveled on 11/03/2018 to the beach, and returned on March 21. I need to go to school on 14th of May",
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
                },
		fieldTypeTransformation: []*types.FieldTypeTransformation{{
		        Fields: []*types.FieldTypes{{
                                Name: types.FieldTypesEnum_DATE_TIME.String(),
		        }},                
			Transformation: &types.Transformation{
				ShiftDateValue: &types.ShiftDateValue{
					DaysSinceMomentZero: -30,
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
