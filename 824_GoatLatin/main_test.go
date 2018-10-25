package main

import "testing"

func TestToGoatLatin(t *testing.T) {
	tcs := []struct {
		inStr  string
		outStr string
	}{
		{"I speak Goat Latin", "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"},
		{"The quick brown fox jumped over the lazy dog", "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"},

		{"a", "amaa"},
		{"A", "Amaa"},
		{"b", "bmaa"},
		{"ab", "abmaa"},
		{"ba", "abmaa"},
		{"a b", "amaa bmaaa"},
		{"a b c d e f", "amaa bmaaa cmaaaa dmaaaaa emaaaaaa fmaaaaaaa"},
		{"ab cd ef", "abmaa dcmaaa efmaaaa"},
		{"aa ee uu ii oo", "aamaa eemaaa uumaaaa iimaaaaa oomaaaaaa"},
		{"AA EE UU II OO", "AAmaa EEmaaa UUmaaaa IImaaaaa OOmaaaaaa"},
		{"AA be UU cC OO De", "AAmaa ebmaaa UUmaaaa Ccmaaaaa OOmaaaaaa eDmaaaaaaa"},
	}

	for idx, tc := range tcs {
		out := toGoatLatin(tc.inStr)
		if out != tc.outStr {
			t.Errorf("%d. toGoatLatin(%q) = %q. Expected: %q", idx, tc.inStr, out, tc.outStr)
		}
	}
}
