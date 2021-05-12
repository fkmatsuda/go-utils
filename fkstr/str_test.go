package fkstr

import "testing"

func TestCountFormatParams(t *testing.T) {

	cnt := CountFormatParams("test")
	if cnt != 0 {
		t.Error("Expecting zero parameters")
	}

	cnt = CountFormatParams("test %")
	if cnt != 0 {
		t.Error("Expecting zero parameters if it at end")
	}

	cnt = CountFormatParams("test %%")
	if cnt != 0 {
		t.Error("Expecting zero parameters if it is a literal at end")
	}

	cnt = CountFormatParams("test %% test")
	if cnt != 0 {
		t.Error("Expecting zero parameters if it is a literal")
	}

	cnt = CountFormatParams("test %s test")
	if cnt != 1 {
		t.Error("Expecting one parameter")
	}

	cnt = CountFormatParams("test %s %d test")
	if cnt != 2 {
		t.Error("Expecting two parameters")
	}

	cnt = CountFormatParams("test %s%s %02d test")
	if cnt != 3 {
		t.Error("Expecting tree parameters")
	}

}
