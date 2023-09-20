package sanitizer

import (
	"net/url"
	"os"
	"regexp"
	"strings"
)

func isImageProxyURL(imageUrl string) bool {
	if getOsEnv("YARR_IMG_PROXY") != "Y" {
		return false
	}
	imgProxyDomains := getOsEnv("YARR_IMG_PROXY_EXCLUDE_DOMAINS")
	if imgProxyDomains == "" {
		return true
	}
	imgProxyDomainSlice := strings.Split(imgProxyDomains, ",")
	// get imageUrl's domain from url
	uri, err := url.Parse(imageUrl)
	if err != nil {
		return false
	}
	for _, domain := range imgProxyDomainSlice {
		if matched, _ := regexp.MatchString(strings.ReplaceAll(domain, ".", "\\."), uri.Host); matched {
			return false
		}
	}
	return true
}

func getImageProxyEndpoint() string {
	proxyEndpoint := getOsEnv("YARR_IMG_PROXY_ENDPOINT")
	if proxyEndpoint != "" {
		return proxyEndpoint
	}
	return "https://images.weserv.nl?url="
}

func getImageProxiedEndpoint(imageUrl string) string {
	return getImageProxyEndpoint() + imageUrl
}

func getOsEnv(key string) string {
	return os.Getenv(key)
}
