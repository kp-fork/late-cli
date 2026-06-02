package mcp

import (
	"strings"
	"testing"
)

// truncateOutput replicates the truncation logic in ToolAdapter.Execute so it
// can be tested independently without a live MCP server.
func truncateOutput(output string) string {
	const maxChars = 32768
	runes := []rune(output)
	if len(runes) > maxChars {
		return string(runes[:maxChars]) + "\n\n[... truncated, output exceeded limit ...]"
	}
	return output
}

func TestTruncateOutput_UnderLimit(t *testing.T) {
	input := strings.Repeat("a", 100)
	got := truncateOutput(input)
	if got != input {
		t.Errorf("expected output to pass through unchanged, got len=%d", len(got))
	}
}

func TestTruncateOutput_AtLimit(t *testing.T) {
	input := strings.Repeat("a", 32768)
	got := truncateOutput(input)
	if got != input {
		t.Errorf("expected output at exactly the limit to pass through unchanged")
	}
}

func TestTruncateOutput_OverLimit(t *testing.T) {
	input := strings.Repeat("a", 32769)
	got := truncateOutput(input)
	if !strings.HasSuffix(got, "\n\n[... truncated, output exceeded limit ...]") {
		t.Errorf("expected truncation marker at end of output, got: %q", got[len(got)-50:])
	}
	runes := []rune(got)
	// Should be exactly maxChars runes plus the marker
	marker := "\n\n[... truncated, output exceeded limit ...]"
	content := string(runes[:len(runes)-len([]rune(marker))])
	if len([]rune(content)) != 32768 {
		t.Errorf("expected 32768 content runes before marker, got %d", len([]rune(content)))
	}
}

// TestTruncateOutput_MultibyteUTF8 verifies that truncation never splits a
// multi-byte UTF-8 character. Each '日' is 3 bytes; byte-slicing at 32768
// would produce invalid UTF-8 if the boundary falls inside a character.
func TestTruncateOutput_MultibyteUTF8(t *testing.T) {
	// Build a string of 32769 Japanese characters (each 3 bytes = 98307 bytes total).
	input := strings.Repeat("日", 32769)
	got := truncateOutput(input)

	// Must be valid UTF-8
	if !strings.ContainsRune(got, '日') {
		t.Error("truncated output contains no valid Japanese characters")
	}

	// The truncated portion must end with the marker, not a broken character
	if !strings.HasSuffix(got, "\n\n[... truncated, output exceeded limit ...]") {
		t.Error("expected truncation marker after multi-byte content")
	}

	// Verify the result decodes cleanly — if any rune is utf8.RuneError it was split
	for i, r := range got {
		if r == '�' {
			t.Errorf("invalid UTF-8 rune at position %d — character was split at boundary", i)
			break
		}
	}
}
