package murka

import "testing"

func TestHighlight(t *testing.T) {
	type args struct {
		text   string
		left   string
		right  string
		sample string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Highlight(tt.args.text, tt.args.left, tt.args.right, tt.args.sample)
			if (err != nil) != tt.wantErr {
				t.Errorf("Highlight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Highlight() got = %v, want %v", got, tt.want)
			}
		})
	}
}
