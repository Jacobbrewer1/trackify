package username

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

// TrackUsernames tracks multiple usernames across various platforms.
func TrackUsernames(ctx context.Context, usernames []string) error {
	for _, username := range usernames {
		if err := TrackUsername(ctx, username); err != nil {
			return fmt.Errorf("failed to track username %s: %w", username, err)
		}
	}
	return nil
}

// TrackUsername tracks a single username across various platforms.
func TrackUsername(ctx context.Context, username string) error {
	found := make(map[string]bool)
	for _, t := range targets {
		ok, err := track(ctx, username, t)
		if err != nil {
			fmt.Printf("Error checking %s: %v\n", t.name, err)
			found[t.name] = false
			continue
		}

		found[t.name] = ok
	}

	displayResultTable(username, found)
	return nil
}

// track checks if a username exists on a given target platform.
func track(ctx context.Context, username string, target *target) (bool, error) {
	targetURL := target.url.JoinPath(username)

	req, err := http.NewRequestWithContext(ctx, "GET", targetURL.String(), http.NoBody)
	if err != nil {
		return false, fmt.Errorf("failed to create request for %s: %w", target.name, err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("request failed for %s: %w", target.name, err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	return resp.StatusCode == http.StatusOK, nil
}

// displayResultTable displays the tracking results in a formatted table.
func displayResultTable(username string, results map[string]bool) {
	t := table.NewWriter()
	title := "Tracking results for username: " + username
	t.SetTitle(title)
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Platform", "Found"})
	t.SortBy([]table.SortBy{{Name: "Platform", Mode: table.Asc}})
	t.Style().Size.WidthMin = len(title) + 4

	for platform, found := range results {
		status := "No"
		if found {
			status = "Yes"
		}
		t.AppendRow([]any{platform, status})
	}
	t.Render()
}
