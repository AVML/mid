// Referencing meta documents, or specific parts from meta documents
// mid:<meta-id>[/<content-id>]
// Example mid:960830.1639@XIson.com/partA.960830.1639@XIson.com
package mid

import (
	"errors"
	"net/url"
	"strings"
)

const prefix = "mid://"

func Is(s string) bool {
	return strings.HasPrefix(s, prefix)
}

const sep = "/"

func Encode(metaID, contentID string) string {
	return prefix + strings.Join([]string{
		url.QueryEscape(metaID),
		url.QueryEscape(contentID),
	}, sep)
}

func Decode(s string) (metaID, contentID string, err error) {
	if !strings.HasPrefix(s, prefix) {
		return "", "", errors.New("Invalid mid")
	}
	s = strings.TrimPrefix(s, prefix)
	sa := strings.Split(s, sep)
	if len(sa) != 2 {
		return "", "", errors.New("Invalid mid")
	}
	metaID, err = url.QueryUnescape(sa[0])
	if err != nil {
		return
	}
	contentID, err = url.QueryUnescape(sa[1])
	return
}
