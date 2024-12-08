package dao

import (
	"content_system/internal/model"
	"testing"
)

func TestContentDao_Create(t *testing.T) {
	type args struct {
		detail model.ContentDetail
	}
	tests := []struct {
		name    string
		c       *ContentDao
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Create(tt.args.detail); (err != nil) != tt.wantErr {
				t.Errorf("ContentDao.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
