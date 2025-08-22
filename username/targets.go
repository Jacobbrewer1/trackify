package username

import (
	"fmt"
	"net/url"
	"strings"

	pkgslices "github.com/jacobbrewer1/web/slices"
)

type target struct {
	url  *url.URL
	name string
}

var allTargets = map[string]*target{
	"facebook": {
		name: strings.ToTitle("facebook"),
		url:  parseTargetURL("https://www.facebook.com/"),
	},
	"twitter": {
		name: strings.ToTitle("twitter"),
		url:  parseTargetURL("https://www.twitter.com/"),
	},
	"instagram": {
		name: strings.ToTitle("instagram"),
		url:  parseTargetURL("https://www.instagram.com/"),
	},
	"linkedin": {
		name: strings.ToTitle("linkedin"),
		url:  parseTargetURL("https://www.linkedin.com/in/"),
	},
	"github": {
		name: strings.ToTitle("github"),
		url:  parseTargetURL("https://www.github.com/"),
	},
	"pinterest": {
		name: strings.ToTitle("pinterest"),
		url:  parseTargetURL("https://www.pinterest.com/"),
	},
	"tumblr": {
		name: strings.ToTitle("tumblr"),
		url:  parseTargetURL("https://www.tumblr.com/"),
	},
	"youtube": {
		name: strings.ToTitle("youtube"),
		url:  parseTargetURL("https://www.youtube.com/"),
	},
	"soundcloud": {
		name: strings.ToTitle("soundcloud"),
		url:  parseTargetURL("https://soundcloud.com/"),
	},
	"snapchat": {
		name: strings.ToTitle("snapchat"),
		url:  parseTargetURL("https://www.snapchat.com/add/"),
	},
	"tiktok": {
		name: strings.ToTitle("tiktok"),
		url:  parseTargetURL("https://www.tiktok.com/@"),
	},
	"behance": {
		name: strings.ToTitle("behance"),
		url:  parseTargetURL("https://www.behance.net/"),
	},
	"medium": {
		name: strings.ToTitle("medium"),
		url:  parseTargetURL("https://www.medium.com/@"),
	},
	"quora": {
		name: strings.ToTitle("quora"),
		url:  parseTargetURL("https://www.quora.com/profile/"),
	},
	"flickr": {
		name: strings.ToTitle("flickr"),
		url:  parseTargetURL("https://www.flickr.com/people/"),
	},
	"periscope": {
		name: strings.ToTitle("periscope"),
		url:  parseTargetURL("https://www.periscope.tv/"),
	},
	"twitch": {
		name: strings.ToTitle("twitch"),
		url:  parseTargetURL("https://www.twitch.tv/"),
	},
	"dribbble": {
		name: strings.ToTitle("dribbble"),
		url:  parseTargetURL("https://www.dribbble.com/"),
	},
	"stumbleupon": {
		name: strings.ToTitle("stumbleupon"),
		url:  parseTargetURL("https://www.stumbleupon.com/stumbler/"),
	},
	"ello": {
		name: strings.ToTitle("ello"),
		url:  parseTargetURL("https://www.ello.co/"),
	},
	"product_hunt": {
		name: strings.ToTitle("Product Hunt"),
		url:  parseTargetURL("https://www.producthunt.com/@"),
	},
	"telegram": {
		name: strings.ToTitle("telegram"),
		url:  parseTargetURL("https://www.telegram.me/"),
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
