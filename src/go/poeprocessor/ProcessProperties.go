package poeprocessor

import (
	"regexp"
	"sync"

	"github.com/AdityaHegde/PathOfExileTrade/database"
	poeormmodel "github.com/AdityaHegde/PathOfExileTrade/model/orm/poe"
	"gorm.io/gorm"
)

var propertyRegexs = [...]*regexp.Regexp{
	regexp.MustCompile("(?P<value>\\d+(?:\\.?\\d*))"),
	regexp.MustCompile("(?P<prefix>Occupied by )(?P<value>[a-zA-Z ]*)"),
}
var syncPropertiesMap = sync.Map{}

// ProcessProperties is exported
func ProcessProperties(db *gorm.DB, item poeormmodel.Item, properties []string) []*poeormmodel.PropertyValue {
	var propertyValues = make([]*poeormmodel.PropertyValue, len(properties))

	for i, property := range properties {
		propertyorm := processPropertyValue(db, item, property)
		if propertyorm != nil {
			propertyValues[i] = propertyorm
		}
	}

	return propertyValues
}

func processPropertyValue(
	db *gorm.DB, item poeormmodel.Item, apiProperty string,
) *poeormmodel.PropertyValue {
	for _, propertyRegex := range propertyRegexs {
		matches := propertyRegex.FindAllStringSubmatch(apiProperty, -1)
		if matches != nil && len(matches) > 0 {
			propertyName := "$prefix#$suffix"

			createProperty(db, propertyRegex.ReplaceAllString(apiProperty, propertyName))

			var propertyValue = poeormmodel.PropertyValue{
				ItemID:     item.ID,
				PropertyID: propertyName,
			}

			propertyValue.PropertyMinValue = matches[0][0]
			if len(matches) > 1 {
				propertyValue.PropertyMaxValue = matches[1][0]
			}

			return &propertyValue
		}
	}

	return nil
}

func createProperty(db *gorm.DB, propertyName string) {
	property := poeormmodel.Property{
		PropertyName: propertyName,
	}

	database.FindOrCreateInMap(db, &syncPropertiesMap, propertyName, &property)
}
