package username

import (
	"net/http"
)

// isStatusOK returns true if the response status code accepts 2xx codes as successful.
func isStatusOK(resp *http.Response) bool {
	return resp.StatusCode >= 200 && resp.StatusCode < 300
}

// appendUsernameToURL appends the given username to the target URL's path.
func appendUsernameToURL(req *http.Request, username string) {
	req.URL = req.URL.JoinPath(req.URL.Path, username)
}
