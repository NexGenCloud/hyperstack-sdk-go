package hyperstack

// Forced Hyperstack client identification headers.
//
// Maintained in sdk-generator (util/go/HyperstackHeaders.go) and copied into the
// generated SDK by the `add-hyperstack-headers` task. Do not edit in the SDK repo.

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

// hyperstackSDKVersion is stamped from $SDK_VERSION at generation time; the Go
// generator emits no version constant of its own.
const hyperstackSDKVersion = "v1.55.7-alpha"

const hyperstackSDKName = "hyperstack-go-sdk"

const (
	hyperstackClientHeader    = "Hyperstack-Client"
	hyperstackUserAgentHeader = "User-Agent"
)

var (
	hyperstackClient    = sanitizeHyperstackHeader(hyperstackSDKName + "/" + hyperstackSDKVersion)
	hyperstackUserAgent = buildHyperstackUserAgent()
)

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

// sanitizeHyperstackHeader keeps printable ASCII only, capped at 256 chars, so a
// stray control character can never reach the wire.
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

// setHyperstackHeaders forces the identification headers onto an outgoing request.
//
// Set rather than Add: prepareRequest already calls Header.Add for User-Agent and
// again for every cfg.DefaultHeader entry, so Add would emit duplicate values.
func setHyperstackHeaders(req *http.Request) {
	if req == nil {
		return
	}
	req.Header.Set(hyperstackClientHeader, hyperstackClient)
	req.Header.Set(hyperstackUserAgentHeader, hyperstackUserAgent)
}
