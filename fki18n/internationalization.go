package fki18n

import (
	"fmt"

	"github.com/fkmatsuda/go-utils/fkstr"
)

// LanguageKey is a type definition for a language indentification
type LanguageKey string

// TextKey is a type definition for a text identification
type TextKey int32

// TextDef defines the text pattern in a language for a given key
type TextDef struct {

	// SingularText define text pattern for singular form
	SingularText string

	// PluralText (optional) define text pattern for plural form
	PluralText string

	// ParamSelectIndex (required if PluralText is defined) defines the parameter index that will be used to choose the text pattern in singular or plural form
	ParamSelectIndex int8
}

// Language is a type definition to map texts to the specified language
type Language struct {

	// Key is the key selector for this language
	Key LanguageKey

	// TextMap is the TextDef map
	TextMap map[TextKey]TextDef
}

var (
	registeredLanguages map[LanguageKey]Language = map[LanguageKey]Language{}
)

// RegisterLanguage registers a language text map
func RegisterLanguage(language Language) {
	registeredLanguages[language.Key] = language
}

// Text assembles the text into a language according to the parameters entered
func (key LanguageKey) Text(text TextKey, params ...interface{}) string {

	textDef := registeredLanguages[key].TextMap[text]

	if len(params) == 0 {
		return textDef.SingularText
	}

	var textPattern string
	if textDef.PluralText != "" && params[textDef.ParamSelectIndex].(int) > 1 {
		textPattern = textDef.PluralText
	} else {
		textPattern = textDef.SingularText
	}

	n := fkstr.CountFormatParams(textPattern)

	if len(params) > n {
		params = params[:n]
	}

	return fmt.Sprintf(textPattern, params...)

}
