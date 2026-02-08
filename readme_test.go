package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestREADMEHasFunnyJoke(t *testing.T) {
	content, err := ioutil.ReadFile("README.md")
	if err != nil {
		t.Fatalf("Failed to read README.md: %v", err)
	}

	readmeContent := string(content)
	lines := strings.Split(readmeContent, "\n")

	// Verify expected sections exist
	expectedSections := []string{"# go-web-app-template", "## Structure", "```"}
	for _, section := range expectedSections {
		if !strings.Contains(readmeContent, section) {
			t.Errorf("README.md missing expected section: %s", section)
		}
	}

	// Find the end of structure section
	endOfStructure := findEndOfStructureSection(lines)
	if endOfStructure == -1 {
		t.Error("Could not find end of structure section in README.md")
		return
	}

	// Validate joke content exists after structure section
	validateJokeContent(t, lines[endOfStructure+1:])
}

func findEndOfStructureSection(lines []string) int {
	for i, line := range lines {
		if strings.TrimSpace(line) == "```" && i > 20 {
			return i
		}
	}
	return -1
}

func validateJokeContent(t *testing.T, contentAfterStructure []string) {
	// Filter out empty lines
	var nonEmptyContent []string
	for _, line := range contentAfterStructure {
		if strings.TrimSpace(line) != "" {
			nonEmptyContent = append(nonEmptyContent, line)
		}
	}

	if len(nonEmptyContent) == 0 {
		t.Error("README.md should have content (a joke) after the structure section")
		return
	}

	// Look for joke punctuation
	hasJokePunctuation := false
	for _, line := range nonEmptyContent {
		if strings.Contains(line, "?") || strings.Contains(line, "!") {
			hasJokePunctuation = true
			break
		}
	}

	if !hasJokePunctuation {
		t.Error("README.md should have a funny joke with punctuation (! or ?) at the end")
	}
}

func TestREADMEFileExists(t *testing.T) {
	if _, err := os.Stat("README.md"); os.IsNotExist(err) {
		t.Error("README.md file should exist")
	}
}
