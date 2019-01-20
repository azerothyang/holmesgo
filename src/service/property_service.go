package service

import (
	"context"
	"user"
)


// GreeterService is used to implement GreeterService.
type PropertyService struct{}


func (p *PropertyService) GetAge(context.Context, *user.Empty) (*user.PropertyInfo, error) {
	return &user.PropertyInfo{Age: 28}, nil
}
