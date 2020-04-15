package greeter

import (
	"reflect"
	"testing"
)

func TestNewDB(t *testing.T) {
	tests := []struct {
		name string
		want DB
	}{
		{
			name: "it should create new db instance",
			want: new(db),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_db_FetchDefaultMessage(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "it should return the default message",
			want:    "default message",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &db{}
			got, err := d.FetchDefaultMessage()
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchDefaultMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FetchDefaultMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_db_FetchMessage(t *testing.T) {
	type args struct {
		lang string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "it should return hello when the lang is equals to en",
			args:    args{
				lang: "en",
			},
			want:    "hello",
			wantErr: false,
		},
		{
			name:    "it should return holla when lang is equals to es",
			args:    args{
				lang: "es",
			},
			want:    "holla",
			wantErr: false,
		},
		{
			name:    "it should return selam when language is not equals to en or es",
			args:    args{
				lang: "tr",
			},
			want:    "selam",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &db{}
			got, err := d.FetchMessage(tt.args.lang)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FetchMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}