package services

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
	"github.com/PuerkitoBio/goquery"
)

func GetInstagramMedia(postUrl string) (string, error) {
	re := regexp.MustCompile(`/(?:p|reel)/([A-Za-z0-9_-]+)`)
	matches := re.FindStringSubmatch(postUrl)
	if len(matches) < 2 {
		return "", errors.New("invalid Instagram URL format - must be /p/ or /reel/")
	}

	postId := matches[1]
	embedUrl := fmt.Sprintf("https://www.instagram.com/p/%s/embed/", postId)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", embedUrl, nil)
	if err != nil {
		return "", err
	}

	// to avoid blockage we are using the headers here just in case :)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Referer", "https://www.instagram.com/")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("embed request failed with status %d", resp.StatusCode)
	}

	return extractMediaFromResponse(resp.Body)
}

func extractMediaFromResponse(body io.Reader) (string, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return "", err
	}

	// selectors in order of preference
	selectors := []string{
		"img[src*='fbcdn.net']",
		"img[src*='cdninstagram.com']",
		"video[src*='fbcdn.net']",
		"meta[property='og:image']",
		"meta[property='og:video']",
	}

	for _, selector := range selectors {
		var mediaUrl string
		doc.Find(selector).Each(func(i int, s *goquery.Selection) {
			if mediaUrl != "" {
				return
			}

			var src string
			var exists bool

			if selector == "meta[property='og:image']" || selector == "meta[property='og:video']" {
				src, exists = s.Attr("content")
			} else {
				src, exists = s.Attr("src")
			}

			if exists && isValidMediaUrl(src) {
				mediaUrl = src
			}
		})

		if mediaUrl != "" {
			return mediaUrl, nil
		}
	}

	return "", errors.New("could not extract media URL from Instagram post")
}

func isValidMediaUrl(url string) bool {
	if url == "" {
		return false
	}

	url = strings.ToLower(url)

	// these are instragram cdm patterns
	validPatterns := []string{
		"fbcdn.net",
		"cdninstagram.com",
		"instagram.f",
	}

	hasValidPattern := false
	for _, pattern := range validPatterns {
		if strings.Contains(url, pattern) {
			hasValidPattern = true
			break
		}
	}

	if !hasValidPattern {
		return false
	}

	// to avoid any small image urls
	// can be profiel pcitures, logos etc
	invalidPatterns := []string{
		"profile",
		"150x150",
		"44x44",
		"s150x150",
	}

	for _, pattern := range invalidPatterns {
		if strings.Contains(url, pattern) {
			return false
		}
	}

	return true
}
