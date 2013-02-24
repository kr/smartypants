package smartypants

import (
	"bytes"
	"testing"
)

var ts = []struct {
	s string
	w string
}{
	{
		`hello, I "just" don't -- know`,
		`hello, I &ldquo;just&rdquo; don&rsquo;t &mdash; know`,
	},
}

func TestSmartypants(t *testing.T) {
	for i := range ts {
		b := new(bytes.Buffer)
		w := NewEducator(b, 0)
		w.Write([]byte(ts[i].s))
		if g := b.String(); g != ts[i].w {
			t.Errorf("%q != %q", g, ts[i].w)
		}
	}
}
