package service

import (
	"context"
	pb "sustainability-service/generated/sustainability"
	"sustainability-service/storage/postgres"
)

type SustainabilityService struct {
	pb.UnimplementedSustainabilityimpactServiceServer

	Sustainability *postgres.SustainabilityRepo
}

func (s *SustainabilityService) LogImpact(ctx context.Context, in *pb.LogImpactRequest) (*pb.LogImpactResponse, error) {
	resp, err := s.Sustainability.LogImpact(in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}