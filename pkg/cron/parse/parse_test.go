package parse

import (
	"testing"

	"github.com/perithompson/cron-parser/pkg/cron"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	type args struct {
		cronStr string
	}
	tests := []struct {
		name    string
		args    args
		want    cron.Cron
		wantErr bool
	}{
		{
			name: "missing command",
			args: args{
				cronStr: "*/15 0 1,15 * 1-5",
			},
			want:    cron.Cron{},
			wantErr: true,
		},
		{
			name: "example",
			args: args{
				cronStr: "*/15 0 1,15 * 1-5 /usr/bin/find",
			},
			want: cron.Cron{
				Minute:     []int{0, 15, 30, 45},
				Hour:       []int{0},
				DayOfMonth: []int{1, 15},
				Month:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				DayOfWeek:  []int{1, 2, 3, 4, 5},
				Command:    "/usr/bin/find",
			},
			wantErr: false,
		},
		{
			name: "duplicate day of month",
			args: args{
				cronStr: "*/15 0 1,15,15 * 1-5 /usr/bin/find",
			},
			want:    cron.Cron{},
			wantErr: true,
		},
		{
			name: "string days",
			args: args{
				cronStr: "*/15 0 1,15 * Monday-Friday /usr/bin/find",
			},
			want: cron.Cron{
				Minute:     []int{0, 15, 30, 45},
				Hour:       []int{0},
				DayOfMonth: []int{1, 15},
				Month:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				DayOfWeek:  []int{1, 2, 3, 4, 5},
				Command:    "/usr/bin/find",
			},
			wantErr: false,
		},
		{
			name: "invalid hour",
			args: args{
				cronStr: "*/15 25 1,15 * Monday-Friday /usr/bin/find",
			},
			want:    cron.Cron{},
			wantErr: true,
		},
		{
			name: "invalid mins",
			args: args{
				cronStr: "365 0 1,15 * Monday-Friday /usr/bin/find",
			},
			want:    cron.Cron{},
			wantErr: true,
		},
		{
			name: "range mins",
			args: args{
				cronStr: "0-5 0 1,15 * Monday-Friday /usr/bin/find",
			},
			want: cron.Cron{
				Minute:     []int{0, 1, 2, 3, 4, 5},
				Hour:       []int{0},
				DayOfMonth: []int{1, 15},
				Month:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				DayOfWeek:  []int{1, 2, 3, 4, 5},
				Command:    "/usr/bin/find",
			},
			wantErr: false,
		},
		{
			name: "short day letters",
			args: args{
				cronStr: "0-5 0 1,15 * MON-FRI /usr/bin/find",
			},
			want: cron.Cron{
				Minute:     []int{0, 1, 2, 3, 4, 5},
				Hour:       []int{0},
				DayOfMonth: []int{1, 15},
				Month:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				DayOfWeek:  []int{1, 2, 3, 4, 5},
				Command:    "/usr/bin/find",
			},
			wantErr: false,
		},
		{
			name: "short month letters",
			args: args{
				cronStr: "0-5 0 1,15 JAN 1-5 /usr/bin/find",
			},
			want: cron.Cron{
				Minute:     []int{0, 1, 2, 3, 4, 5},
				Hour:       []int{0},
				DayOfMonth: []int{1, 15},
				Month:      []int{1},
				DayOfWeek:  []int{1, 2, 3, 4, 5},
				Command:    "/usr/bin/find",
			},
			wantErr: false,
		},
		{
			name: "At minute 23 past every 2nd hour from 0 through 20.",
			args: args{
				cronStr: "23 0-20/2 * * * /usr/bin/find",
			},
			want: cron.Cron{
				Minute:     []int{23},
				Hour:       []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
				DayOfMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
				Month:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				DayOfWeek:  []int{0, 1, 2, 3, 4, 5, 6},
				Command:    "/usr/bin/find",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.cronStr)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, *got)
			}
		})
	}
}
