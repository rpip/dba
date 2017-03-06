package dba

import (
	"github.com/hashicorp/hil/ast"
	"github.com/icrowley/fake"
)

var (
	funcMap = map[string]interface{}{
		"brand":                 fake.Brand,
		"character":             fake.Character,
		"characters":            fake.Characters,
		"city":                  fake.City,
		"color":                 fake.Color,
		"company":               fake.Company,
		"continent":             fake.Continent,
		"country":               fake.Country,
		"credit_card_type":      fake.CreditCardType,
		"currency":              fake.Currency,
		"currency_code":         fake.CurrencyCode,
		"day":                   fake.Day,
		"digits":                fake.Digits,
		"domain_name":           fake.DomainName,
		"domain_zone":           fake.DomainZone,
		"email":                 fake.EmailAddress,
		"first_name":            fake.FirstName,
		"last_name":             fake.LastName,
		"full_name":             fake.FullName,
		"full_name_with_prefix": fake.FullNameWithPrefix,
		"full_name_with_suffix": fake.FullNameWithSuffix,
		"gender":                fake.Gender,
		"gender_abbrev":         fake.GenderAbbrev,
		"hex_color":             fake.HexColor,
		"hex_color_short":       fake.HexColorShort,
		"ipv4":                  fake.IPv4,
		"ipv6":                  fake.IPv6,
		"industry":              fake.Industry,
		"job_title":             fake.JobTitle,
		"language":              fake.Language,
		"latitude":              fake.Latitute,
		"latitude_degrees":      fake.LatitudeDegress,
		"latitude_direction":    fake.LatitudeDirection,
		"latitude_minutes":      fake.LatitudeMinutes,
		"latitude_seconds":      fake.LatitudeSeconds,
		"longitude":             fake.Longitude,
		"longitude_degrees":     fake.LongitudeDegrees,
		"longitude_direction":   fake.LongitudeDirection,
		"longitude_minutes":     fake.LongitudeMinutes,
		"longitude_seconds":     fake.LongitudeSeconds,
		"model":                 fake.Model,
		"month":                 fake.Month,
		"month_num":             fake.MonthNum,
		"month_short":           fake.MonthShort,
		"paragraph":             fake.Paragraph,
		"phone":                 fake.Phone,
		"product":               fake.Product,
		"product_name":          fake.ProductName,
		"sentence":              fake.Sentence,
		"sentences":             fake.Sentences,
		"state":                 fake.State,
		"state_abbrev":          fake.StateAbbrev,
		"street":                fake.Street,
		"street_address":        fake.StreetAddress,
		"title":                 fake.Title,
		"top_level_domain":      fake.TopLevelDomain,
		"useragent":             fake.UserAgent,
		"username":              fake.UserName,
		"weekday":               fake.WeekDay,
		"weekday_short":         fake.WeekDayShort,
		"word":                  fake.Word,
		"words":                 fake.Words,
		"zip":                   fake.Zip,
		"simple_password":       fake.SimplePassword,
	}
)

var (
	year = ast.Function{
		ArgTypes:   []ast.Type{ast.TypeInt, ast.TypeInt},
		ReturnType: ast.TypeInt,
		Variadic:   false,
		Callback: func(inputs []interface{}) (interface{}, error) {
			from := inputs[0].(int)
			to := inputs[1].(int)
			return fake.Year(from, to), nil
		},
	}

	sentencesN = ast.Function{
		ArgTypes:   []ast.Type{ast.TypeInt},
		ReturnType: ast.TypeString,
		Variadic:   false,
		Callback: func(inputs []interface{}) (interface{}, error) {
			n := inputs[0].(int)
			return fake.SentencesN(n), nil
		},
	}

	wordsN = ast.Function{
		ArgTypes:   []ast.Type{ast.TypeInt},
		ReturnType: ast.TypeString,
		Variadic:   false,
		Callback: func(inputs []interface{}) (interface{}, error) {
			n := inputs[0].(int)
			return fake.WordsN(n), nil
		},
	}

	charactersN = ast.Function{
		ArgTypes:   []ast.Type{ast.TypeInt},
		ReturnType: ast.TypeString,
		Variadic:   false,
		Callback: func(inputs []interface{}) (interface{}, error) {
			n := inputs[0].(int)
			return fake.CharactersN(n), nil
		},
	}

	digitsN = ast.Function{
		ArgTypes:   []ast.Type{ast.TypeInt},
		ReturnType: ast.TypeString,
		Variadic:   false,
		Callback: func(inputs []interface{}) (interface{}, error) {
			n := inputs[0].(int)
			return fake.DigitsN(n), nil
		},
	}

	paragraphsN = ast.Function{
		ArgTypes:   []ast.Type{ast.TypeInt},
		ReturnType: ast.TypeString,
		Variadic:   false,
		Callback: func(inputs []interface{}) (interface{}, error) {
			n := inputs[0].(int)
			return fake.ParagraphsN(n), nil
		},
	}

	creditCardNum = ast.Function{
		ArgTypes:   []ast.Type{ast.TypeString},
		ReturnType: ast.TypeString,
		Variadic:   false,
		Callback: func(inputs []interface{}) (interface{}, error) {
			vendor := inputs[0].(string)
			return fake.CreditCardNum(vendor), nil
		},
	}

	password = ast.Function{
		ArgTypes:   []ast.Type{ast.TypeInt, ast.TypeInt, ast.TypeBool, ast.TypeBool, ast.TypeBool},
		ReturnType: ast.TypeString,
		Variadic:   false,
		Callback: func(inputs []interface{}) (interface{}, error) {
			atLeast := inputs[0].(int)
			atMost := inputs[1].(int)
			allowUpper := inputs[2].(bool)
			allowNumeric := inputs[3].(bool)
			allowSpecial := inputs[4].(bool)
			return fake.Password(atLeast, atMost, allowUpper, allowNumeric, allowSpecial), nil
		},
	}
)
