package main
import "testing"

func TestTest(t *testing.T){
	got := Test(5)
	want := 5
	if got != want {
		t.Error("Damn",got,want)
	}
}