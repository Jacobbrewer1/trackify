package username

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	pkgslices "github.com/jacobbrewer1/web/slices"
)

func isStatusOK(resp *http.Response) bool {
	return resp.StatusCode >= 200 && resp.StatusCode < 300
}

type target struct {
	url                 *url.URL
	name                string
	isRequestSuccessful func(*http.Response) bool
}

var allTargets = map[string]*target{
	"facebook": {
		name:                strings.ToTitle("facebook"),
		url:                 parseTargetURL("https://www.facebook.com/"),
		isRequestSuccessful: isStatusOK,
	},
	"twitter": {
		name:                strings.ToTitle("twitter"),
		url:                 parseTargetURL("https://www.twitter.com/"),
		isRequestSuccessful: isStatusOK,
	},
	"instagram": {
		name:                strings.ToTitle("instagram"),
		url:                 parseTargetURL("https://www.instagram.com/"),
		isRequestSuccessful: isStatusOK,
	},
	"linkedin": {
		name:                strings.ToTitle("linkedin"),
		url:                 parseTargetURL("https://www.linkedin.com/in/"),
		isRequestSuccessful: isStatusOK,
	},
	"github": {
		name:                strings.ToTitle("github"),
		url:                 parseTargetURL("https://www.github.com/"),
		isRequestSuccessful: isStatusOK,
	},
	"pinterest": {
		name:                strings.ToTitle("pinterest"),
		url:                 parseTargetURL("https://www.pinterest.com/"),
		isRequestSuccessful: isStatusOK,
	},
	"tumblr": {
		name:                strings.ToTitle("tumblr"),
		url:                 parseTargetURL("https://www.tumblr.com/"),
		isRequestSuccessful: isStatusOK,
	},
	"youtube": {
		name:                strings.ToTitle("youtube"),
		url:                 parseTargetURL("https://www.youtube.com/"),
		isRequestSuccessful: isStatusOK,
	},
	"soundcloud": {
		name:                strings.ToTitle("soundcloud"),
		url:                 parseTargetURL("https://soundcloud.com/"),
		isRequestSuccessful: isStatusOK,
	},
	"snapchat": {
		name:                strings.ToTitle("snapchat"),
		url:                 parseTargetURL("https://www.snapchat.com/add/"),
		isRequestSuccessful: isStatusOK,
	},
	"tiktok": {
		name:                strings.ToTitle("tiktok"),
		url:                 parseTargetURL("https://www.tiktok.com/@"),
		isRequestSuccessful: isStatusOK,
	},
	"behance": {
		name:                strings.ToTitle("behance"),
		url:                 parseTargetURL("https://www.behance.net/"),
		isRequestSuccessful: isStatusOK,
	},
	"medium": {
		name:                strings.ToTitle("medium"),
		url:                 parseTargetURL("https://www.medium.com/@"),
		isRequestSuccessful: isStatusOK,
	},
	"quora": {
		name:                strings.ToTitle("quora"),
		url:                 parseTargetURL("https://www.quora.com/profile/"),
		isRequestSuccessful: isStatusOK,
	},
	"flickr": {
		name:                strings.ToTitle("flickr"),
		url:                 parseTargetURL("https://www.flickr.com/people/"),
		isRequestSuccessful: isStatusOK,
	},
	"periscope": {
		name:                strings.ToTitle("periscope"),
		url:                 parseTargetURL("https://www.periscope.tv/"),
		isRequestSuccessful: isStatusOK,
	},
	"twitch": {
		name:                strings.ToTitle("twitch"),
		url:                 parseTargetURL("https://www.twitch.tv/"),
		isRequestSuccessful: isStatusOK,
	},
	"dribbble": {
		name:                strings.ToTitle("dribbble"),
		url:                 parseTargetURL("https://www.dribbble.com/"),
		isRequestSuccessful: isStatusOK,
	},
	"stumbleupon": {
		name:                strings.ToTitle("stumbleupon"),
		url:                 parseTargetURL("https://www.stumbleupon.com/stumbler/"),
		isRequestSuccessful: isStatusOK,
	},
	"ello": {
		name:                strings.ToTitle("ello"),
		url:                 parseTargetURL("https://www.ello.co/"),
		isRequestSuccessful: isStatusOK,
	},
	"product_hunt": {
		name:                strings.ToTitle("Product Hunt"),
		url:                 parseTargetURL("https://www.producthunt.com/@"),
		isRequestSuccessful: isStatusOK,
	},
	"telegram": {
		name:                strings.ToTitle("telegram"),
		url:                 parseTargetURL("https://www.telegram.me/"),
		isRequestSuccessful: isStatusOK,
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
