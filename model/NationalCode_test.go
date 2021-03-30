package model

import (
	"testing"
)

func TestNationalCode_IsValid(t *testing.T) {
	t.Run("most be validate", func(t *testing.T) {
		var element =  NewNatinalCode("2080512031")
		if !element.IsValid(){
			t.Errorf("%s most be validate" , element.code)
		}
	})
	t.Run("validate", func(t *testing.T) {
		var element =  NewNatinalCode("20805120331")
		if element.IsValid(){
			t.Errorf("%s most be not validate" , element.code)
		}
	})
}
