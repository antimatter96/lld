package url

import (
	"lld/url_shortner/helper"
)

const DEFAULT_TTL = 60 * 30

type URLObject struct {
	Short       string
	Long        string
	CreatedAt   int64
	KeepForever bool
	ExpireAt    int64
}

func (urlObj *URLObject) Expired() bool {
	if urlObj.KeepForever {
		return false
	}

	return helper.CurrentTimestamp() > urlObj.ExpireAt
}

func NewURLObject(long string, keepForever bool) *URLObject {
	urlObj := &URLObject{
		Long:        long,
		KeepForever: keepForever,
	}

	urlObj.CreatedAt = helper.CurrentTimestamp()

	if !keepForever {
		urlObj.ExpireAt = urlObj.CreatedAt + DEFAULT_TTL
	}

	return urlObj
}
