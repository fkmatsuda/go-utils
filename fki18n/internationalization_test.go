package fki18n

import "testing"

const (
	TEST_LANGUAGE        LanguageKey = "test"
	NO_PARAMS            TextKey     = 1
	WITH_STRING_PARAM    TextKey     = 2
	WITH_INT_PARAM       TextKey     = 3
	WITH_MULTIPLE_PARAMS TextKey     = 4
	WITH_PLURAL          TextKey     = 5
)

func init() {
	RegisterLanguage(Language{
		Key: TEST_LANGUAGE,
		TextMap: map[TextKey]TextDef{
			NO_PARAMS: TextDef{
				SingularText: "No params text",
			},
			WITH_STRING_PARAM: TextDef{
				SingularText: "With the \"%s\" param value",
			},
			WITH_INT_PARAM: TextDef{
				SingularText: "With the %d param value",
			},
			WITH_MULTIPLE_PARAMS: TextDef{
				SingularText: "With the %d and %s params, %[1]d",
			},
			WITH_PLURAL: TextDef{
				SingularText: "Just one",
				PluralText:   "Has %d",
			},
		},
	})
}

func TestText(t *testing.T) {

	text := TEST_LANGUAGE.Text(NO_PARAMS)

	if "No params text" != text {
		t.Errorf("text = \"%s\"; want \"No params text\"", text)
	}

	text = TEST_LANGUAGE.Text(WITH_STRING_PARAM, "text")
	if "With the \"text\" param value" != text {
		t.Errorf("text = \"%s\"; want \"With the \"text\" param value\"", text)
	}

	text = TEST_LANGUAGE.Text(WITH_INT_PARAM, 1)
	if "With the 1 param value" != text {
		t.Errorf("text = \"%s\"; want \"With the 1 param value\"", text)
	}

	text = TEST_LANGUAGE.Text(WITH_MULTIPLE_PARAMS, 1, "text")
	if "With the 1 and text params, 1" != text {
		t.Errorf("text = \"%s\"; want \"With the 1 and text params, 1\"", text)
	}

	text = TEST_LANGUAGE.Text(WITH_PLURAL, 1)
	if "Just one" != text {
		t.Errorf("text = \"%s\"; want \"Just one\"", text)
	}

	text = TEST_LANGUAGE.Text(WITH_PLURAL, 2)
	if "Has 2" != text {
		t.Errorf("text = \"%s\"; want \"Has 2\"", text)
	}

}
