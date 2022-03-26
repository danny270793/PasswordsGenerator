package passwordsgenerator

import (
	"strconv"
	"strings"

	"github.com/danny270793/passwords-generator/src/faker"
)

func Generate(characters int, containSymbols bool, containNumbers bool) string {
	password := ""
	for i := 0; i < characters; i++ {
		if containSymbols && faker.Boolean(0.25) {
			symbol := faker.Symbol()
			password += symbol
		} else if containNumbers && faker.Boolean(0.25) {
			number := faker.NumberBetween(0, 9)
			password += strconv.Itoa(number)
		} else {
			character := faker.Character()
			if faker.Boolean(0.5) {
				character = strings.ToUpper(character)
			}
			password += character
		}
	}
	return password
}
