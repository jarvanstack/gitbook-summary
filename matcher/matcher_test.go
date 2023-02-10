package matcher

import "testing"

func TestRegexMatcher_Match(t *testing.T) {
	type fields struct {
		regexs []string
	}
	type args struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			fields: fields{
				regexs: []string{"_"},
			},
			args: args{
				str: "_sidebar.md",
			},
			want: true,
		},
		{
			fields: fields{
				regexs: []string{".md"},
			},
			args: args{
				str: "test/_sidebar.md",
			},
			want: true,
		},
		{
			fields: fields{
				regexs: []string{"_", ".git"},
			},
			args: args{
				str: ".git/test.md",
			},
			want: true,
		},
		{
			fields: fields{
				regexs: []string{"_"},
			},
			args: args{
				str: "test.md",
			},
			want: false,
		},
		{
			fields: fields{
				regexs: []string{"."},
			},
			args: args{
				str: "test.md",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RegexMatcher{
				regexs: tt.fields.regexs,
			}
			if got := r.Match(tt.args.str); got != tt.want {
				t.Errorf("RegexMatcher.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
