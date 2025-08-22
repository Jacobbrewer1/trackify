package ipaddress

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// TrackIPAddresses tracks multiple IP addresses.
func TrackIPAddresses(ctx context.Context, ips []string) error {
	for _, ip := range ips {
		if err := TrackIPAddress(ctx, ip); err != nil {
			return fmt.Errorf("failed to track IP address %s: %w", ip, err)
		}
	}
	return nil
}

// TrackIPAddress tracks the given IP address.
func TrackIPAddress(ctx context.Context, ip string) error {
	fmt.Printf("Tracking IP address: %s\n", ip)
	// Placeholder for actual tracking logic
	return nil
}

// trackIP tracks an IP address.
func trackIP(ctx context.Context, trackingIP string) (*whoisResponse, error) {
	target := whoIsApiURL().JoinPath(trackingIP)

	req, err := http.NewRequestWithContext(ctx, "GET", target.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for %s: %w", trackingIP, err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %w", trackingIP, err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 response for %s: %d", trackingIP, resp.StatusCode)
	}

	var whoisResp whoisResponse
	if err := json.NewDecoder(resp.Body).Decode(&whoisResp); err != nil {
		return nil, fmt.Errorf("failed to decode response for %s: %w", trackingIP, err)
	}
	return &whoisResp, nil
}

// displayIPResultTable displays the tracking results in a formatted table.
func displayIPResultTable(ip string, result *whoisResponse) {

}
