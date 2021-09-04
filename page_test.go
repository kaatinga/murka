package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/julienschmidt/httprouter"
	localErrors "github.com/kaatinga/3lines.club/errors"
	"testing"
)

func Test_extractPages(t *testing.T) {

	tests := []struct {
		path    string
		ps      httprouter.Params
		wantErr error
	}{
		{"pages/about/1", httprouter.Params{
			0: httprouter.Param{
				Key:   "page",
				Value: "pages/about/1",
			},
		}, nil},
		{"/pages/about/1", httprouter.Params{
			0: httprouter.Param{
				Key:   "page",
				Value: "/pages/about/1",
			},
		}, nil},
		{"/", httprouter.Params{
			0: httprouter.Param{
				Key:   "page",
				Value: "/",
			},
		}, localErrors.ErrIncorrectPageURI},
		{"empty Params", httprouter.Params{}, localErrors.ErrIncorrectPageURI},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			pageTree, err := extractPages(tt.ps)
			if err != tt.wantErr {
				t.Errorf("extractPages() error = %v, wantErr %v", err, tt.wantErr)
			}
			spew.Dump(pageTree)
		})
	}
}

func Test_validatePagePath(t *testing.T) {

	tests := []struct {
		pagePath string
		wantErr  error
	}{
		{"test_page", nil},
		{"1", nil},
		{"p1", nil},
		{"az", nil},
		{"AZ", nil},
		{"19", nil},
		{"p_1", nil},
		{"ы", localErrors.ErrIncorrectPagePathSymbol},
		{"-", localErrors.ErrIncorrectPagePathSymbol},
		{".", localErrors.ErrIncorrectPagePathSymbol},
	}
	for _, tt := range tests {
		t.Run(tt.pagePath, func(t *testing.T) {
			if err := validatePagePath(tt.pagePath); err != tt.wantErr {
				t.Errorf("validatePagePath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}


func Test_validatePagePath2(t *testing.T) {

	tests := []struct {
		pagePath string
		wantErr  error
	}{
		{"test_page", nil},
		{"1", nil},
		{"p1", nil},
		{"az", nil},
		{"AZ", nil},
		{"19", nil},
		{"p_1", nil},
		{"ы", localErrors.ErrIncorrectPagePathSymbol},
		{"-", localErrors.ErrIncorrectPagePathSymbol},
		{".", localErrors.ErrIncorrectPagePathSymbol},
	}
	for _, tt := range tests {
		t.Run(tt.pagePath, func(t *testing.T) {
			if err := validatePagePath2(tt.pagePath); err != tt.wantErr {
				t.Errorf("validatePagePath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validatePagePath3(t *testing.T) {

	tests := []struct {
		pagePath string
		wantErr  error
	}{
		{"test_page", nil},
		{"1", nil},
		{"p1", nil},
		{"aZ", nil},
		{"AZ", nil},
		{"19", nil},
		{"p_1", nil},
		{"ы", localErrors.ErrIncorrectPagePathSymbol},
		{"-", localErrors.ErrIncorrectPagePathSymbol},
		{".", localErrors.ErrIncorrectPagePathSymbol},
	}
	for _, tt := range tests {
		t.Run(tt.pagePath, func(t *testing.T) {
			if err := validatePagePath3(tt.pagePath); err != tt.wantErr {
				t.Errorf("validatePagePath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}