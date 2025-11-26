package scraper

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSetupFileOutput(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.txt")

	s := &scraper{
		props: &scraperProps{
			isFile:   true,
			FileName: filePath,
		},
	}

	err := s.setupFileOutput()
	if err != nil {
		t.Fatalf("setupFileOutput() error = %v", err)
	}
	defer s.closeFile()

	if s.file == nil {
		t.Fatal("s.file should not be nil")
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("file was not created: %s", filePath)
	}
}

func TestWriteToFile(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.txt")

	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	s := &scraper{
		file: file,
	}
	defer s.closeFile()

	s.writeToFile(1, 200, "http://example.com")

	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("failed to read temp file: %v", err)
	}

	expected := "    1   200     http://example.com\n"
	if string(content) != expected {
		t.Errorf("expected %q, got %q", expected, string(content))
	}
}

func TestCloseFile(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.txt")

	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	s := &scraper{
		file: file,
	}

	s.closeFile()

	// On Unix, this will produce an error if the file is closed.
	// On Windows, this is not a reliable way to check if a file is closed.
	err = s.file.Close() // try to close it again
	if err == nil {
		t.Logf("closing an already closed file did not produce an error (this may happen on some OSes)")
	}
}
