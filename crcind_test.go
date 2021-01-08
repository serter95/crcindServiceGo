package crcindservicego

import (
	"fmt"
	"reflect"
	"testing"
)

func EvaluateResult(results []StandardResponse, err error, t *testing.T) {

	if err != nil {
		fmt.Println("Fallo la Peticion", err)
	}

	if reflect.TypeOf(results).Kind() != reflect.Slice {
		t.Errorf("results not is Slice, contain = %q", results)
	}

	for _, iterator := range results {
		if reflect.TypeOf(iterator.Category).Kind() != reflect.String {
			t.Errorf("Category not is String, contain = %q", iterator.Category)
		}

		if reflect.TypeOf(iterator.Name).Kind() != reflect.String {
			t.Errorf("Name not is String, contain = %q", iterator.Name)
		}

		if reflect.TypeOf(iterator.Author).Kind() != reflect.String {
			t.Errorf("Author not is String, contain = %q", iterator.Author)
		}

		if reflect.TypeOf(iterator.PreviewURL).Kind() != reflect.String {
			t.Errorf("PreviewURL not is String, contain = %q", iterator.PreviewURL)
		}

		if reflect.TypeOf(iterator.Origin).Kind() != reflect.String {
			t.Errorf("Origin not is String, contain = %q", iterator.Origin)
		}
	}
}

func TestFindResultsUsingPeterResult(t *testing.T) {
	results, err := FindResults("  PeTeR  ")
	EvaluateResult(results, err, t)
}

func TestFindResultsUsingAdamResult(t *testing.T) {
	results, err := FindResults("adam")
	EvaluateResult(results, err, t)
}

func TestFindResultsWithoutResult(t *testing.T) {
	results, err := FindResults("")
	EvaluateResult(results, err, t)
}
