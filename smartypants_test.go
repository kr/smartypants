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
		w: `“word”`,
	},
	{
		s: `'word'`,
		w: `‘word’`,
	},
	{
		s: `don't`,
		w: `don’t`,
	},
	{
		s: `...`,
		w: `…`,
	},
	{
		s: `a-b`,
		w: `a-b`,
	},
	{
		s: `a - b`,
		w: `a – b`,
	},
	{
		s: `a--b`,
		w: `a—b`,
	},
	{
		s: `a -- b`,
		w: `a — b`,
	},
	{
		s: `a -- b`,
		w: `a – b`,
		f: LatexDashes,
	},
	{
		s: `a --- b`,
		w: `a — b`,
		f: LatexDashes,
	},
	{
		s: `(c)`,
		w: `©`,
	},
	{
		s: `(r)`,
		w: `®`,
	},
	{
		s: `(tm)`,
		w: `™`,
	},
	{
		s: `1/4`,
		w: `¼`,
	},
}

func TestSmartypants(t *testing.T) {
	for i := range ts {
		b := new(bytes.Buffer)
		New(b, ts[i].f).Write([]byte(ts[i].s))
		if g := b.String(); g != ts[i].w {
			t.Errorf("%#q => %#q != %#q", ts[i].s, g, ts[i].w)
		}
	}
}
