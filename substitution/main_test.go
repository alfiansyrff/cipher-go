package main

import "testing"

func TestMainTableDriven(t *testing.T) {
	var tests = []struct {
		text      string
		isEncrypt bool
		subs      int64
		want      string
	}{
		{"alfin", true, 5, "fqkns"},
		{"polstat STIS", true, 8, "xwtabib ABQA"},
		{"saya nonton tv", true, 12, "emkm zazfaz fh"},
		{"tidur aja deh", true, 15, "ixsjg pyp stw"},

		{"ixsjg pyp stw", false, 15, "tidur aja deh"},
		{"xwtabib ABQA", false, 8, "polstat STIS"},
		{"fqkns", false, 5, "alfin"},
		{"emkm zazfaz fh", false, 12, "saya nonton tv"},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			if tt.isEncrypt {
				ans := Encrypt(tt.text, tt.subs)
				if ans != tt.want {
					t.Errorf("got %s, want %s", ans, tt.want)
				}
			} else {
				ans := Decrypt(tt.text, tt.subs)
				if ans != tt.want {
					t.Errorf("got %s, want %s", ans, tt.want)
				}
			}
		})
	}
}
