package user_consent_events

// Forced Hyperstack client identification headers.
//
// Template maintained in sdk-generator (go-sdk-generator/data/HyperstackHeaders.go).
// The build task writes one copy into each generated tag package, substituting the
// package name and the SDK version. Do not edit the generated copies.
//
// This lives in its own file rather than inside <Tag>_client.go so the generated
// client needs no extra imports; same package, so hyperstackSetHeaders is visible.

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

const hyperstackSDKName = "hyperstack-go-sdk"

// hyperstackSDKVersion is substituted from $SDK_VERSION at generation time.
const hyperstackSDKVersion = "v1.55.4-alpha"

var hyperstackClient = sanitizeHyperstackHeader(hyperstackSDKName + "/" + hyperstackSDKVersion)

var hyperstackUserAgent = buildHyperstackUserAgent()

func buildHyperstackUserAgent() string {
	return sanitizeHyperstackHeader(fmt.Sprintf(
		"%s (Go/%s; %s/%s)",
		hyperstackClient,
		strings.TrimPrefix(runtime.Version(), "go"),
		normalizeHyperstackOS(runtime.GOOS),
		normalizeHyperstackArch(runtime.GOARCH),
	))
}

func normalizeHyperstackOS(value string) string {
	v := strings.ToLower(strings.TrimSpace(value))
	switch {
	case strings.HasPrefix(v, "darwin"), strings.HasPrefix(v, "mac"):
		return "darwin"
	case strings.HasPrefix(v, "win"):
		return "windows"
	case strings.HasPrefix(v, "linux"):
		return "linux"
	case v == "":
		return "unknown"
	}
	return strings.ReplaceAll(v, " ", "-")
}

func normalizeHyperstackArch(value string) string {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "x86_64", "x64", "amd64":
		return "x86_64"
	case "aarch64", "arm64":
		return "arm64"
	case "i386", "i686", "x86", "386":
		return "386"
	case "":
		return "unknown"
	}
	return strings.ToLower(strings.TrimSpace(value))
}

// sanitizeHyperstackHeader keeps printable ASCII only, capped at 256 chars.
func sanitizeHyperstackHeader(value string) string {
	var b strings.Builder
	for _, r := range value {
		if r >= 0x20 && r < 0x7f {
			b.WriteRune(r)
		}
	}
	out := b.String()
	if len(out) > 256 {
		out = out[:256]
	}
	return out
}

// hyperstackSetHeaders forces the identification headers onto an outgoing request.
//
// Called via `defer` at the top of applyEditors, so it runs after every configured
// RequestEditorFn and before the request is sent -- a caller-supplied editor cannot
// override it. Set rather than Add, so no duplicate values are emitted.
func hyperstackSetHeaders(req *http.Request) {
	if req == nil {
		return
	}
	req.Header.Set("Hyperstack-Client", hyperstackClient)
	req.Header.Set("User-Agent", hyperstackUserAgent)
}
