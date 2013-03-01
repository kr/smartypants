package smartypants

import (
	"bytes"
	"testing"
)

var ts = []struct {
	s string
	w string
	f int
}{
	{
		s: `"word"`,
		w: `&ldquo;word&rdquo;`,
	},
	{
		s: `'word'`,
		w: `&lsquo;word&rsquo;`,
	},
	{
		s: `don't`,
		w: `don&rsquo;t`,
	},
	{
		s: `...`,
		w: `&hellip;`,
	},
	{
		s: `a-b`,
		w: `a-b`,
	},
	{
		s: `a - b`,
		w: `a &ndash; b`,
	},
	{
		s: `a--b`,
		w: `a&mdash;b`,
	},
	{
		s: `a -- b`,
		w: `a &mdash; b`,
	},
	{
		s: `a -- b`,
		w: `a &ndash; b`,
		f: LatexDashes,
	},
	{
		s: `a --- b`,
		w: `a &mdash; b`,
		f: LatexDashes,
	},
	{
		s: `(c)`,
		w: `&copy;`,
	},
	{
		s: `(r)`,
		w: `&reg;`,
	},
	{
		s: `(tm)`,
		w: `&trade;`,
	},
	{
		s: `1/4`,
		w: `&frac14;`,
	},
}

func TestSmartypants(t *testing.T) {
	for i := range ts {
		b := new(bytes.Buffer)
		New(b, ts[i].f).Write([]byte(ts[i].s))
		if g := b.String(); g != ts[i].w {
			t.Errorf("%q => %q != %q", ts[i].s, g, ts[i].w)
		}
	}
}
