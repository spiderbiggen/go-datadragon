//go:build integration

package datadragon

import (
	"context"
	"testing"
)

func TestDataDragon_Realm(t *testing.T) {
	type fields struct {
		Region Region
	}
	type args struct {
		ctx       context.Context
		optRegion []Region
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name:   "Default",
			fields: fields{},
			args: args{
				ctx: context.Background(),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New(WithRegion(tt.fields.Region))
			got, err := d.Realm(tt.args.ctx, tt.args.optRegion...)
			if err != tt.wantErr {
				t.Errorf("Realm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got == nil {
				t.Errorf("Realm() got = %v", got)
				return
			}
		})
	}
}
