package kernel

import "testing"

func TestResultFormat(t *testing.T) {
	got, _ := ResultFormat(NewSuccess(nil))
	want := "{\"success\":true}"
	if string(got) != want {
		t.Errorf("got %q; want %q", string(got), want)
	}

	got2, _ := ResultFormat(NewSuccess("abc"))
	want2 := "{\"success\":true,\"data\":\"abc\"}"
	if string(got2) != want2 {
		t.Errorf("got %q; want %q", string(got2), want2)
	}

	got3, _ := ResultFormat(NewSystemRunErr(nil))
	want3 := "{\"code\":500,\"message\":\"system run error\"}"
	if string(got3) != want3 {
		t.Errorf("got %q; want %q", string(got3), want3)
	}

	got4, _ := ResultFormat(NewResult(MethodNotAllowed, GetMessage(MethodNotAllowed), "abc"))
	want4 := "{\"code\":405,\"message\":\"method not allowed\",\"data\":\"abc\"}"
	if string(got4) != want4 {
		t.Errorf("got %q; want %q", string(got4), want4)
	}
}
