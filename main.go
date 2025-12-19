package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const changelogURL = "https://raw.githubusercontent.com/anthropics/claude-code/main/CHANGELOG.md"

var version = "dev"

// ChangelogEntry represents a single version entry in the changelog
type ChangelogEntry struct {
	Version string   `json:"version"`
	Changes []string `json:"changes"`
}

var (
	jsonOutput     = flag.Bool("json", false, "Output as JSON")
	markdownOutput = flag.Bool("md", false, "Output raw markdown")
	targetVersion  = flag.String("version", "", "Fetch specific version (e.g., 2.0.70)")
	listVersions   = flag.Bool("list", false, "List all available versions")
	showVersion    = flag.Bool("v", false, "Show cc-latest version")
	showHelp       = flag.Bool("h", false, "Show help")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "cc-latest - Fetch the latest Claude Code changelog entry\n\n")
		fmt.Fprintf(os.Stderr, "Usage: cc-latest [flags]\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  cc-latest              # Latest entry as plain text\n")
		fmt.Fprintf(os.Stderr, "  cc-latest -json        # Latest entry as JSON\n")
		fmt.Fprintf(os.Stderr, "  cc-latest -md          # Latest entry as raw markdown\n")
		fmt.Fprintf(os.Stderr, "  cc-latest -version 2.0.70  # Specific version\n")
		fmt.Fprintf(os.Stderr, "  cc-latest -list        # List all versions\n")
	}

	flag.Parse()

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}

	if *showVersion {
		fmt.Printf("cc-latest version %s\n", version)
		os.Exit(0)
	}

	changelog, err := fetchChangelog()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching changelog: %v\n", err)
		os.Exit(1)
	}

	entries := parseChangelog(changelog)
	if len(entries) == 0 {
		fmt.Fprintf(os.Stderr, "Error: No changelog entries found\n")
		os.Exit(1)
	}

	if *listVersions {
		for _, entry := range entries {
			fmt.Println(entry.Version)
		}
		os.Exit(0)
	}

	var entry *ChangelogEntry
	if *targetVersion != "" {
		for i := range entries {
			if entries[i].Version == *targetVersion {
				entry = &entries[i]
				break
			}
		}
		if entry == nil {
			fmt.Fprintf(os.Stderr, "Error: Version %s not found\n", *targetVersion)
			os.Exit(1)
		}
	} else {
		entry = &entries[0]
	}

	if *jsonOutput {
		outputJSON(entry)
	} else if *markdownOutput {
		outputMarkdown(entry)
	} else {
		outputPlainText(entry)
	}
}

func fetchChangelog() (string, error) {
	resp, err := http.Get(changelogURL)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	return string(body), nil
}

func parseChangelog(content string) []ChangelogEntry {
	var entries []ChangelogEntry

	// Split by version headers (## X.X.X)
	versionRegex := regexp.MustCompile(`(?m)^## (\d+\.\d+\.\d+)\s*$`)
	matches := versionRegex.FindAllStringSubmatchIndex(content, -1)

	for i, match := range matches {
		versionEnd := match[1]
		version := content[match[2]:match[3]]

		var contentEnd int
		if i+1 < len(matches) {
			contentEnd = matches[i+1][0]
		} else {
			contentEnd = len(content)
		}

		sectionContent := content[versionEnd:contentEnd]
		changes := parseChanges(sectionContent)

		entries = append(entries, ChangelogEntry{
			Version: version,
			Changes: changes,
		})
	}

	return entries
}

func parseChanges(content string) []string {
	var changes []string
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "- ") {
			change := strings.TrimPrefix(trimmed, "- ")
			changes = append(changes, change)
		}
	}

	return changes
}

func outputJSON(entry *ChangelogEntry) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(entry); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
		os.Exit(1)
	}
}

func outputMarkdown(entry *ChangelogEntry) {
	fmt.Printf("## %s\n\n", entry.Version)
	for _, change := range entry.Changes {
		fmt.Printf("- %s\n", change)
	}
}

func outputPlainText(entry *ChangelogEntry) {
	fmt.Printf("Claude Code %s\n", entry.Version)
	fmt.Println(strings.Repeat("-", 40))
	for _, change := range entry.Changes {
		fmt.Printf("  * %s\n", change)
	}
}
