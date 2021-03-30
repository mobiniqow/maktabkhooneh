package model

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

var posistionValues = [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

const lengthOfNationalCode = 10

// NationalCode godoc
type NationalCode struct {
	code       string
	stringCode string
	Valid      bool
	City 	   City
}
// NationalCode constructor
func NewNatinalCode(code string) *NationalCode {
	instance := &NationalCode{}
	instance.code = code
	return instance
}

// IsValid godoc

func (n *NationalCode) IsValid() bool {
	reg, err := regexp.Compile("/[^0-9]/")
	if err != nil {
		log.Fatal(err)
	}
	n.code = reg.ReplaceAllString( n.code, "")
	if len( n.code) != lengthOfNationalCode {
		return false
	}
	codes := strings.Split( n.code, "")
	last, err := strconv.Atoi(codes[9])
	i := lengthOfNationalCode
	sum := 0

	for in, el := range codes {
		temp, err := strconv.Atoi(el)

		if err != nil {
			log.Fatal(err)
		}

		if in == 9 {
			break
		}

		sum += temp * i
		i -= 1
	}

	mod := sum % 11

	if mod >= 2 {
		mod = 11 - mod
	}
	return mod == last
}

