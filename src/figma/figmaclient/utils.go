package figmaclient

import "net/url"

func parseURLOrPanic(rawUrl string) *url.URL {
	result, err := url.Parse(rawUrl)
	if err != nil {
		panic(err)
	}

	return result
}
