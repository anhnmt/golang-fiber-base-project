package util

import (
	"github.com/bwmarrin/snowflake"
	"reflect"
	"testing"
)

func TestGenerateSnowflakeID(t *testing.T) {
	type args struct {
		node []int64
	}
	tests := []struct {
		name string
		args args
		want snowflake.ID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateSnowflakeID(tt.args.node...); got != tt.want {
				t.Errorf("GenerateSnowflakeID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateUUID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateUUID(); got != tt.want {
				t.Errorf("GenerateUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFirst(t *testing.T) {
	type args struct {
		args interface{}
		num  []int
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFirst(tt.args.args, tt.args.num...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}
