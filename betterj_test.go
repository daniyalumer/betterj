package betterj

import (
	"testing"
)

func TestMinifyJ(t *testing.T) {
	json_string := `{
        "name": "John",
        "age": "30",
        "city": "New York"
    }`
	expected := `{"age":"30","city":"New York","name":"John"}`
	minified, err := MinifyJ(json_string)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if minified != expected {
		t.Fatalf("expected %s, got %s", expected, minified)
	}
}

func TestBeautifyJ(t *testing.T) {
	json_string := `{"name":"John","age":"30","city":"New York"}`
	expected := `{
  "age": "30",
  "city": "New York",
  "name": "John"
 }`

	beautified, err := BeautifyJ(json_string, " ")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if beautified != expected {
		t.Fatalf("\nexpected:\n%s\ngot:\n%s", expected, beautified)
	}
}
