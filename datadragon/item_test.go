//go:build integration

package datadragon

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestDataDragon_Items(t *testing.T) {
	type fields struct {
		opts []opt
	}
	type args struct {
		ctx    context.Context
		params []RequestConfig
	}
	type test struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}
	var tests []test
	d := New()
	versions, _ := d.Versions(context.Background())
	for _, version := range versions {
		if version == "7.2.1" {
			break
		}
		tests = append(tests, test{
			name: fmt.Sprintf("Items %s", version),
			args: args{
				ctx: context.Background(),
				params: []RequestConfig{
					{Version: version},
				},
			},
		})
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			d := New(tt.fields.opts...)
			got, err := d.Items(tt.args.ctx, tt.args.params...)

			if err != tt.wantErr {
				t.Errorf("Items() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if got == nil || reflect.DeepEqual(got, &ItemResponse{}) {
				t.Errorf("Items() got empty")
				return
			}
			//for _, item := range got.Items {
			//	indent, _ := json.MarshalIndent(item.Stats, "", "  ")
			//	t.Logf("%s, %s", item.Name, indent)
			//}
		})
	}
}
