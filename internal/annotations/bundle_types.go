package annotations

import (
	"fmt"
	"strings"
)

// ParseBundleTypesSupported reports whether a chart declares product and/or
// integration placement support. Integration Helm installs are unused; prefer
// listing charts under installer.products only.
//
// Annotation helmet.../bundle-types-supported:
//   - "product" — may appear under installer.products
//   - "integration" — legacy token (charts are not installed via installer.integrations)
//   - "integration,product" or "both" — product listing allowed; integration token ignored for topology
//   - empty — treated as product-only
func ParseBundleTypesSupported(bundleTypesSupported string) (integration, product bool, err error) {
	s := strings.TrimSpace(bundleTypesSupported)
	if s == "" {
		return false, true, nil
	}
	return parseBundleTypesSupported(s)
}

func parseBundleTypesSupported(s string) (integration, product bool, err error) {
	parts := strings.Split(s, ",")
	for _, p := range parts {
		tok := strings.ToLower(strings.TrimSpace(p))
		if tok == "" {
			continue
		}
		switch tok {
		case "integration":
			integration = true
		case "product":
			product = true
		case "both":
			integration, product = true, true
		default:
			return false, false, fmt.Errorf("unknown %s token %q", BundleTypesSupported, tok)
		}
	}
	if !integration && !product {
		return false, false, fmt.Errorf("%s must list at least one of integration, product, or both", BundleTypesSupported)
	}
	return integration, product, nil
}
