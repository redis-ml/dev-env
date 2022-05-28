package handler_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/redisliu/dev-env/golang/cmd/sqsfanoutworker/handler"
)

func TestUnmarshal(t *testing.T) {
	type args struct {
		ctx   context.Context
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    handler.TaskSplit
		wantErr bool
	}{
		{
			name: "Empty input",
			args: args{
				ctx:   context.Background(),
				input: "{}",
			},
			want: handler.TaskSplit{},
		},
		{
			name: "Empty all",
			args: args{
				ctx: context.Background(),
				input: `{
					"start": 0,
					"end": 1000000
				}`,
			},
			want: handler.TaskSplit{
				Start: 0,
				End:   1000000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handler.UnmarshalTask(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
