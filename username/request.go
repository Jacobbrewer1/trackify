package username

import (
	"net/http"
	"net/url"
)

// isStatusOK returns true if the response status code accepts 2xx codes as successful.
func isStatusOK(resp *http.Response) bool {
	return resp.StatusCode >= 200 && resp.StatusCode < 300
}

// appendUsernameToURL appends the given username to the target URL's path.
func appendUsernameToURL(targetURL *url.URL, username string) *url.URL {
	targetURL.JoinPath(username)
}
