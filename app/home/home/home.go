// Package home
// This package home is an example of GRPC home route
package home

import (
	"GRPC-SERVE/helper"
	"context"
)

type HomeStruct struct {
	UnimplementedLoanServiceServer
}

// Request HomeDataRequest consit all params passed via api
// Response HomedataResponse returns response
func (s *HomeStruct) HomeData(ctx context.Context, req *HomeDataRequest) (*HomedataResponse, error) {
	helper.SugarObj.Info("Inside home route")
	var finalStr = req.City + req.Name
	return &HomedataResponse{
		Result: finalStr,
	}, nil
}
