package reader_test

import (
	"os"
	"testing"

	reader "github.com/ei-sugimoto/9ick/internal/template"
)

func TestList(t *testing.T) {
	want := "test"
	r := reader.NewReader()
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(originalDir)

	//cd root directory
	testDir := "../../"
	if err := os.Chdir(testDir); err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}
	got, err := r.List()
	if err != nil {
		t.Fatalf("Error listing templates: %v", err)
	}
	if len(got) == 0 {
		t.Fatalf("List() returned empty list")
	}

	found := false
	for _, tmpl := range got {
		if tmpl == want {
			found = true
			return
		}
	}

	if !found {
		t.Fatalf("List() did not return %s", want)

	}
}
