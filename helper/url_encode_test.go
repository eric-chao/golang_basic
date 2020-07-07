package helper

import (
	"net/url"
	"testing"
)

func TestUrlEncode(t *testing.T) {

	t.Logf("/m/hyt/bulletin?id=953261469 \n %s", url.QueryEscape("/m/hyt/bulletin?id=953261469"))
}
