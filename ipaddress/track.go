package ipaddress

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"go.uber.org/multierr"
)

// TrackIPAddresses tracks multiple IP addresses.
func TrackIPAddresses(ctx context.Context, ips []string) error {
	var merr error
	for _, ip := range ips {
		if err := TrackIPAddress(ctx, ip); err != nil {
			merr = multierr.Append(merr, err)
		}
	}
	if merr != nil {
		return fmt.Errorf("one or more IP tracking operations failed: %w", merr)
	}
	return nil
}

// TrackIPAddress tracks the given IP address.
func TrackIPAddress(ctx context.Context, ip string) error {
	result, err := trackIP(ctx, ip)
	if err != nil {
		return fmt.Errorf("failed to track IP %s: %w", ip, err)
	}
	displayIPResultTable(ip, result)
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
	if !whoisResp.Success {
		return nil, fmt.Errorf("unsuccessful response for %s: %s", trackingIP, whoisResp.Message)
	}
	return &whoisResp, nil
}

// displayIPResultTable displays the tracking results in a formatted table.
func displayIPResultTable(ip string, result *whoisResponse) {
	t := table.NewWriter()
	title := "Tracking results for IP Address: " + ip
	t.SetTitle(title)
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Field", "Value"})
	t.SortBy([]table.SortBy{{Name: "Field", Mode: table.Asc}})
	t.Style().Size.WidthMin = len(title) + 4

	appendWhoisData(t, result)

	t.Render()
}

// appendWhoisData appends the whois data to the table.
func appendWhoisData(t table.Writer, data *whoisResponse) {
	appendTableRow(t, "ip type", data.Type)
	appendTableRow(t, "country", data.Country)
	appendTableRow(t, "country code", data.CountryCode)
	appendTableRow(t, "continent", data.Continent)
	appendTableRow(t, "continent code", data.ContinentCode)
	appendTableRow(t, "region", data.Region)
	appendTableRow(t, "region code", data.RegionCode)
	appendTableRow(t, "location latitude", fmt.Sprintf("%f", data.Latitude))
	appendTableRow(t, "location longitude", fmt.Sprintf("%f", data.Longitude))
	appendTableRow(t, "location link", googleMapsLink(data.Latitude, data.Longitude))
	appendTableRow(t, "zip", data.Postal)
	appendTableRow(t, "calling code", data.CallingCode)
	appendTableRow(t, "capital", data.Capital)
	appendTableRow(t, "boarders", data.Borders)
	appendTableRow(t, "connection asn", data.Connection.Asn)
	appendTableRow(t, "connection org", data.Connection.Org)
	appendTableRow(t, "connection isp", data.Connection.Isp)
	appendTableRow(t, "connection domain", data.Connection.Domain)
	appendTableRow(t, "timezone id", data.Timezone.Id)
	appendTableRow(t, "timezone abbr", data.Timezone.Abbr)
	appendTableRow(t, "timezone is dst", data.Timezone.IsDst)
	appendTableRow(t, "timezone offset", data.Timezone.Offset)
	appendTableRow(t, "timezone utc", data.Timezone.Utc)
	appendTableRow(t, "timezone current time", data.Timezone.CurrentTime)
}

// googleMapsLink generates a Google Maps link for the given latitude and longitude.
func googleMapsLink(lat, lon float64) string {
	return fmt.Sprintf("https://www.google.com/maps/@%f,%f", lat, lon)
}

// appendTableRow appends a row to the table, ensuring that string values are not empty or whitespace.
func appendTableRow[T comparable](t table.Writer, field string, value T) {
	if reflect.TypeOf(value).Kind() == reflect.String {
		str, ok := any(value).(string)
		if !ok || strings.TrimSpace(str) == "" {
			value, _ = any("Failed casting to string").(T)
		}
	}

	t.AppendRow([]any{strings.ToTitle(field), value})
}
