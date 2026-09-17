package annotations

import (
	"testing"

	o "github.com/onsi/gomega"
)

func TestParseBundleTypesSupported(t *testing.T) {
	g := o.NewWithT(t)

	t.Run("comma-separated", func(t *testing.T) {
		i, p, err := ParseBundleTypesSupported("integration, product")
		g.Expect(err).To(o.Succeed())
		g.Expect(i).To(o.BeTrue())
		g.Expect(p).To(o.BeTrue())
	})

	t.Run("both alias", func(t *testing.T) {
		i, p, err := ParseBundleTypesSupported("both")
		g.Expect(err).To(o.Succeed())
		g.Expect(i).To(o.BeTrue())
		g.Expect(p).To(o.BeTrue())
	})

	t.Run("integration only", func(t *testing.T) {
		i, p, err := ParseBundleTypesSupported("integration")
		g.Expect(err).To(o.Succeed())
		g.Expect(i).To(o.BeTrue())
		g.Expect(p).To(o.BeFalse())
	})

	t.Run("product only", func(t *testing.T) {
		i, p, err := ParseBundleTypesSupported("product")
		g.Expect(err).To(o.Succeed())
		g.Expect(i).To(o.BeFalse())
		g.Expect(p).To(o.BeTrue())
	})

	t.Run("empty defaults to product only", func(t *testing.T) {
		i, p, err := ParseBundleTypesSupported("")
		g.Expect(err).To(o.Succeed())
		g.Expect(i).To(o.BeFalse())
		g.Expect(p).To(o.BeTrue())
	})
}
