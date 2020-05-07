package timeular

import (
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestBuildURL(t *testing.T) {
	baseurlstr := "https://example.com"
	expected := baseurlstr
	u, _ := url.Parse(baseurlstr)

	t.Run("plainURL", func(t *testing.T) {
		computed := BuildURL(u)

		if computed != expected {
			t.Errorf("Got '%s' but expected '%s'", computed, expected)
		}
	})

	t.Run("relPath", func(t *testing.T) {
		relpath := "/dev/signin"
		computed := BuildURL(u, relpath)
		expected := u.String() + relpath
		if computed != expected {
			t.Errorf("Got '%s' but expected '%s'", computed, expected)
		}
	})

	t.Run("relPath", func(t *testing.T) {
		relpath := "/dev/signin"
		start := time.Now().UTC().Format(TimeFormat)
		stop := time.Now().UTC().Format(TimeFormat)

		computed := BuildURL(u, relpath, stop, start)
		expected := u.String() + strings.Join([]string{relpath, stop, start}, "/")
		if computed != expected {
			t.Errorf("Got '%s' but expected '%s'", computed, expected)
		}
	})

}
