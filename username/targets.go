package username

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	pkgslices "github.com/jacobbrewer1/web/slices"
)

type requestEditorFunc = func(*http.Request, string)

type target struct {
	url                 *url.URL
	name                string
	isRequestSuccessful func(*http.Response) bool
	urlBuilder          func(*url.URL, string) *url.URL
	requestEditor       []requestEditorFunc
}

var allTargets = map[string]*target{
	"facebook": {
		name:                strings.ToTitle("facebook"),
		url:                 parseTargetURL("https://www.facebook.com/"),
		isRequestSuccessful: isStatusOK,
		requestEditor:       []requestEditorFunc{appendUsernameToURL},
	},
	"twitter": {
		name:                strings.ToTitle("twitter"),
		url:                 parseTargetURL("https://x.com/"),
		isRequestSuccessful: isStatusOK,
		requestEditor:       []requestEditorFunc{appendUsernameToURL},
	},
	"instagram": {
		name:                strings.ToTitle("instagram"),
		url:                 parseTargetURL("https://www.instagram.com/"),
		isRequestSuccessful: isStatusOK,
		requestEditor:       []requestEditorFunc{appendUsernameToURL},
	},
	"linkedin": {
		name:                strings.ToTitle("linkedin"),
		url:                 parseTargetURL("https://www.linkedin.com/in/"),
		isRequestSuccessful: isStatusOK,
		requestEditor:       []requestEditorFunc{appendUsernameToURL},
	},
	"github": {
		name:                strings.ToTitle("github"),
		url:                 parseTargetURL("https://www.github.com/"),
		isRequestSuccessful: isStatusOK,
		requestEditor:       []requestEditorFunc{appendUsernameToURL},
	},
	"pinterest": {
		name:                strings.ToTitle("pinterest"),
		url:                 parseTargetURL("https://www.pinterest.com/"),
		isRequestSuccessful: isStatusOK,
		requestEditor:       []requestEditorFunc{appendUsernameToURL},
	},
	"tumblr": {
		name:                strings.ToTitle("tumblr"),
		url:                 parseTargetURL("https://www.tumblr.com/"),
		isRequestSuccessful: isStatusOK,
		requestEditor:       []requestEditorFunc{appendUsernameToURL},
	},
	"youtube": {
		name:                strings.ToTitle("youtube"),
		url:                 parseTargetURL("https://www.youtube.com/"),
		isRequestSuccessful: isStatusOK,
		requestEditor:       []requestEditorFunc{appendUsernameToURL},
	},
	"soundcloud": {
		name:                strings.ToTitle("soundcloud"),
		url:                 parseTargetURL("https://soundcloud.com/"),
		isRequestSuccessful: isStatusOK,
		requestEditor:       []requestEditorFunc{appendUsernameToURL},
	},
	"snapchat": {
		name:                strings.ToTitle("snapchat"),
		url:                 parseTargetURL("https://www.snapchat.com/add/"),
		isRequestSuccessful: isStatusOK,
		requestEditor:       []requestEditorFunc{appendUsernameToURL},
	},
	"tiktok": {
		name:                strings.ToTitle("tiktok"),
		url:                 parseTargetURL("https://www.tiktok.com/"),
		isRequestSuccessful: isStatusOK,
		requestEditor: []requestEditorFunc{func(req *http.Request, username string) {
			appendUsernameToURL(req, "@"+username)
		}},
	},
	"behance": {
		name:                strings.ToTitle("behance"),
		url:                 parseTargetURL("https://www.behance.net/"),
		isRequestSuccessful: isStatusOK,
		requestEditor:       []requestEditorFunc{appendUsernameToURL},
	},
	"medium": {
		name:                strings.ToTitle("medium"),
		url:                 parseTargetURL("https://www.medium.com/"),
		isRequestSuccessful: isStatusOK,
		requestEditor: []requestEditorFunc{func(req *http.Request, username string) {
			appendUsernameToURL(req, "@"+username)
		}},
	},
	"quora": {
		name:                strings.ToTitle("quora"),
		url:                 parseTargetURL("https://www.quora.com/profile/"),
		isRequestSuccessful: isStatusOK,
		requestEditor:       []requestEditorFunc{appendUsernameToURL},
	},
	"flickr": {
		name:                strings.ToTitle("flickr"),
		url:                 parseTargetURL("https://www.flickr.com/people/"),
		isRequestSuccessful: isStatusOK,
		requestEditor:       []requestEditorFunc{appendUsernameToURL},
	},
	"twitch": {
		name:                strings.ToTitle("twitch"),
		url:                 parseTargetURL("https://www.twitch.tv/"),
		isRequestSuccessful: isStatusOK,
		requestEditor:       []requestEditorFunc{appendUsernameToURL},
	},
	"dribbble": {
		name:                strings.ToTitle("dribbble"),
		url:                 parseTargetURL("https://www.dribbble.com/"),
		isRequestSuccessful: isStatusOK,
		requestEditor:       []requestEditorFunc{appendUsernameToURL},
	},
	"ello": {
		name:                strings.ToTitle("ello"),
		url:                 parseTargetURL("https://www.ello.co/"),
		isRequestSuccessful: isStatusOK,
		requestEditor:       []requestEditorFunc{appendUsernameToURL},
	},
	"product_hunt": {
		name:                strings.ToTitle("Product Hunt"),
		url:                 parseTargetURL("https://www.producthunt.com/"),
		isRequestSuccessful: isStatusOK,
		requestEditor: []requestEditorFunc{func(req *http.Request, username string) {
			appendUsernameToURL(req, "@"+username)
		}},
	},
	"telegram": {
		name:                strings.ToTitle("telegram"),
		url:                 parseTargetURL("https://www.telegram.me/"),
		isRequestSuccessful: isStatusOK,
		requestEditor:       []requestEditorFunc{appendUsernameToURL},
	},
}

func parseTargetURL(urlStr string) *url.URL {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		panic(fmt.Sprintf("failed to parse target URL: %v", err))
	}
	return parsedURL
}

// filterTargets filters the available targets based on user input.
func filterTargets(searchTargets []string) []*target {
	if len(searchTargets) == 0 {
		targets := make([]*target, 0, len(allTargets))
		for _, v := range allTargets {
			targets = append(targets, v)
		}
		return targets
	}

	var filtered []*target
	targetSet := pkgslices.NewSet[string]()
	for _, t := range searchTargets {
		targetSet.Add(t)
	}

	for _, t := range targetSet.Items() {
		searchableName := strings.TrimSpace(t)
		searchableName = strings.ToLower(searchableName)
		searchableName = strings.ReplaceAll(searchableName, " ", "_")
		foundTarget, exists := allTargets[searchableName]
		if !exists {
			fmt.Printf("Warning: Unknown target platform %q specified, skipping.\n", t)
			continue
		}
		filtered = append(filtered, foundTarget)
	}
	return filtered
}
