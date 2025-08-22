package username

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"sync"

	"github.com/jedib0t/go-pretty/v6/table"

	pkgslices "github.com/jacobbrewer1/web/slices"
	pkgsync "github.com/jacobbrewer1/web/sync"
)

// TrackUsernames tracks multiple usernames across various platforms.
func TrackUsernames(ctx context.Context, usernames, searchTarget []string) error {
	for _, username := range usernames {
		if err := TrackUsername(ctx, username, searchTarget); err != nil {
			return fmt.Errorf("failed to track username %s: %w", username, err)
		}
	}
	return nil
}

// TrackUsername tracks a single username across various platforms.
func TrackUsername(ctx context.Context, username string, searchTargets []string) error {
	filteredTargets := filterTargets(searchTargets)

	found := pkgslices.NewSet[string]()

	var numWorkers uint = 1
	if workerCount := runtime.NumCPU(); workerCount > 1 {
		numWorkers = uint(workerCount)
	}

	var targetCount uint = 1
	if count := len(filteredTargets); count > 1 {
		targetCount = uint(count)
	}

	wp := pkgsync.NewWorkerPool(ctx, "username-tracker", numWorkers, targetCount)
	defer wp.Close()

	wg := &sync.WaitGroup{}
	for _, t := range filteredTargets {
		threadTarget := *t
		wg.Add(1)
		wp.SubmitBlocking(func(ctx context.Context) {
			defer wg.Done()

			ok, err := track(ctx, username, &threadTarget)
			if err != nil {
				fmt.Printf("Error checking %s: %v\n", threadTarget.name, err)
				return
			} else if !ok {
				return
			}

			found.Add(threadTarget.name)
		})
	}
	wg.Wait()

	resultMap := make(map[string]bool)
	for _, t := range filteredTargets {
		resultMap[t.name] = found.Contains(t.name)
	}

	displayResultTable(username, resultMap)
	return nil
}

// track checks if a username exists on a given target platform.
func track(ctx context.Context, username string, target *target) (bool, error) {
	targetURL := target.url.JoinPath(username)

	req, err := http.NewRequestWithContext(ctx, "GET", targetURL.String(), http.NoBody)
	if err != nil {
		return false, fmt.Errorf("failed to create request for %s: %w", target.name, err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("request failed for %s: %w", target.name, err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	return target.isRequestSuccessful(resp), nil
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
