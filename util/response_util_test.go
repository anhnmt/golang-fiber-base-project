package util

import "testing"

func TestJsonResponse(t *testing.T) {
	type args struct {
		status  int
		message string
		data    []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := JsonResponse(tt.args.status, tt.args.message, tt.args.data...); (err != nil) != tt.wantErr {
				t.Errorf("JsonResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResponseBadRequest(t *testing.T) {
	type args struct {
		message string
		data    []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ResponseBadRequest(tt.args.message, tt.args.data...); (err != nil) != tt.wantErr {
				t.Errorf("ResponseBadRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResponseError(t *testing.T) {
	type args struct {
		message string
		data    []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ResponseError(tt.args.message, tt.args.data...); (err != nil) != tt.wantErr {
				t.Errorf("ResponseError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResponseNotFound(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ResponseNotFound(tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("ResponseNotFound() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResponseSuccess(t *testing.T) {
	type args struct {
		message string
		data    []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ResponseSuccess(tt.args.message, tt.args.data...); (err != nil) != tt.wantErr {
				t.Errorf("ResponseSuccess() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResponseUnauthorized(t *testing.T) {
	type args struct {
		message string
		data    []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ResponseUnauthorized(tt.args.message, tt.args.data...); (err != nil) != tt.wantErr {
				t.Errorf("ResponseUnauthorized() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
