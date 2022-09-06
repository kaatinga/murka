package murka

import (
	"errors"
	"testing"
)

func TestHighlight(t *testing.T) {
	tests := []struct {
		text    string
		left    string
		right   string
		sample  string
		want    string
		wantErr error
	}{
		{"christmas tree", "<strong>", "</strong>", "mas", "christ<strong>mas</strong> tree", nil},
		{"christmas tree mas", "<b>", "</b>", "mas", "christ<b>mas</b> tree <b>mas</b>", nil},
		{"mamase", "<b>", "</b>", "mas", "ma<b>mas</b>e", nil},
		{"mase", "<b>", "</b>", "mas", "<b>mas</b>e", nil},
		{"mase", "<b>", "</b>", "ggg", "mase", nil},
		{"mase", "", "", "mas", "mase", nil},
		{"test string", "<b>", "</b>", "str", "test <b>str</b>ing", nil},
		{"test string test string test string", "<b>", "</b>", "str", "test <b>str</b>ing test <b>str</b> <b>str</b>", nil},
	}
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			got, err := Highlight(tt.text, tt.left, tt.right, tt.sample)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Highlight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Highlight() got =\n%v\nwant =\n%v", got, tt.want)
			}
		})
	}
}
