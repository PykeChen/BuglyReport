package api

import (
	"strings"
	"testing"
)


func TestRequest2(t *testing.T) {
	var version = strings.Replace(strings.Replace("2.9.5_alpha", ".", "", -1), "_alpha", "", -1)
	t.Log(version)
}
