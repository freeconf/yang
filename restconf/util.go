package restconf

import (
	"net/http"
	"net/url"

	"strings"

	"github.com/freeconf/gconf/c2"
)

// SplitAddress takes a complete address and breaks it into pieces according
// to RESTCONF standards so you can use each piece in appropriate API call
// Example:
//   http://server[:port]/restconf[=device]/module:path/here
//
func SplitAddress(fullurl string) (address string, module string, path string, err error) {
	eoSlashSlash := strings.Index(fullurl, "//") + 2
	if eoSlashSlash < 2 {
		err = badAddressErr
		return
	}
	eoSlash := eoSlashSlash + strings.IndexRune(fullurl[eoSlashSlash:], '/') + 1
	if eoSlash <= eoSlashSlash {
		err = badAddressErr
		return
	}
	colon := eoSlash + strings.IndexRune(fullurl[eoSlash:], ':')
	if colon <= eoSlash {
		err = badAddressErr
		return
	}
	moduleBegin := strings.LastIndex(fullurl[:colon], "/")
	address = fullurl[:moduleBegin+1]
	module = fullurl[moduleBegin+1 : colon]
	path = fullurl[colon+1:]
	return
}

func handleErr(err error, w http.ResponseWriter) bool {
	if err == nil {
		return false
	}
	if httpErr, ok := err.(c2.HttpError); ok {
		if httpErr.HttpCode() >= 500 {
			c2.Err.Print(httpErr.Error())
		}
		http.Error(w, httpErr.Error(), httpErr.HttpCode())
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return true
}

func ipAddrSplitHostPort(addr string) (host string, port string) {
	bracket := strings.IndexRune(addr, ']')
	dblColon := strings.Index(addr, "::")
	isIpv6 := (bracket >= 0 || dblColon >= 0)
	if isIpv6 {
		if bracket > 0 {
			host = addr[:bracket+1]
			if len(addr) > bracket+2 {
				port = addr[bracket+2:]
			}
		} else {
			host = addr
		}
	} else {
		colon := strings.IndexRune(addr, ':')
		if colon < 0 {
			host = addr
		} else {
			host = addr[:colon]
			port = addr[colon+1:]
		}
	}
	return
}

func appendUrlSegment(a string, b string) string {
	if a == "" || b == "" {
		return a + b
	}
	slashA := a[len(a)-1] == '/'
	slashB := b[0] == '/'
	if slashA != slashB {
		return a + b
	}
	if slashA && slashB {
		return a + b[1:]
	}
	return a + "/" + b
}

func shift(orig *url.URL, delim rune) (string, *url.URL) {
	if orig.Path == "" {
		return "", orig
	}
	copy := *orig
	var segment string
	segment, copy.Path = shiftInString(copy.Path, delim)
	_, copy.RawPath = shiftInString(copy.RawPath, delim)
	return segment, &copy
}

func shiftInString(orig string, delim rune) (string, string) {
	termPos := strings.IndexRune(orig, delim)

	// deisgn decision : ignore when path starts with the delim
	if termPos == 0 {
		orig = orig[1:]
		termPos = strings.IndexRune(orig, delim)
	}

	var shifted string
	var segment string
	if termPos < 0 {
		segment = orig
		// shifted = empty
	} else {
		segment = orig[:termPos]
		shifted = orig[termPos+1:]
	}
	return segment, shifted
}

func shiftOptionalParamWithinSegment(orig *url.URL, optionalDelim rune, segDelim rune) (string, string, *url.URL) {
	copy := *orig
	var segment, optional string
	// trickery here - mutating a copy of the URL
	segment, optional, copy.Path = shiftOptionalParamWithinSegmentInString(copy.Path, optionalDelim, segDelim)

	// NOTE: the segment and optional param are returned unescaped presumably because caller
	// would want that.  If not, keep these results and not the ones from above
	_, _, copy.RawPath = shiftOptionalParamWithinSegmentInString(copy.RawPath, optionalDelim, segDelim)

	return segment, optional, &copy
}

// this will not work of unescaped paths that contain optionalDelim or segDelim in the part of the
// url it's trying to shift.
func shiftOptionalParamWithinSegmentInString(orig string, optionalDelim rune, segDelim rune) (string, string, string) {
	termPos := strings.IndexRune(orig, segDelim)

	// design decision : ignore when path starts with the delim
	if termPos == 0 {
		orig = orig[1:]
		termPos = strings.IndexRune(orig, segDelim)
	}

	// find the next segment first...
	var shifted string
	var segment string
	if termPos < 0 {
		segment = orig
		// shifted = empty
	} else {
		segment = orig[:termPos]
		shifted = orig[termPos+1:]
	}

	// ...now look for optional param in the found segment
	optPos := strings.IndexRune(segment, optionalDelim)
	if optPos < 0 {
		return segment, "", shifted
	}
	var optional string
	if len(segment) > optPos+1 {
		optional = segment[optPos+1:]
	}
	segment = segment[:optPos]

	return segment, optional, shifted
}
