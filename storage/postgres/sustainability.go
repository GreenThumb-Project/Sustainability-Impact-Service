package postgres

import (
	"database/sql"
	pb "sustainability-service/generated/sustainability"
)

type SustainabilityRepo struct {
	DB *sql.DB
}

func NewSustainabilityRepo(db *sql.DB) SustainabilityRepo {
	return SustainabilityRepo{DB: db}
}

func (s *SustainabilityRepo) LogImpact(in *pb.LogImpactRequest) (*pb.LogImpactResponse, error) {
	rows, err := s.DB.Exec(`
			INSERT INTO
			impact_logs(
				category,
				amount,
				unit)
			VALUES(
				$1,
				$2,
				$3)
			`, in.Category, in.Amount, in.Unit)

	if err != nil {
		return &pb.LogImpactResponse{Success: false}, err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return &pb.LogImpactResponse{Success: false}, err
	}

	return &pb.LogImpactResponse{Success: true}, err
}

func (s *SustainabilityRepo) GetUserImpact(in *pb.GetUserImpactRequest) (*pb.GetUserImpactResponse, error) {
	rows, err := s.DB.Query(`
			SELECT 
				category,
				amaunt,
				unit
			FROM 
			impact_logs
			WHERE user_id=$1
			`, in.UserId)
	if err != nil {
		return nil, err
	}
	var userImpacts []*pb.UserImpact

	for rows.Next() {
		var userImpact pb.UserImpact
		err := rows.Scan(&userImpact.Category, &userImpact.Amount, &userImpact.Unit)
		if err != nil {
			return nil, err
		}

		userImpacts = append(userImpacts, &userImpact)
	}

	return &pb.GetUserImpactResponse{UserImpact: userImpacts}, nil
}

func (s *SustainabilityRepo) JoinChallenge(in *pb.JoinChallengeRequest) (*pb.JoinChallengeResponse, error) {

	rows, err := s.DB.Exec(`
			INSERT INTO
			user_challenges(
				user_id,
				challenge_id,
				progress
			VALUES(
				$1,
				$2,
				$3)
				)`, in.UserId, in.ChallengeId, in.Progress)

	if err != nil {
		return &pb.JoinChallengeResponse{Success: false}, err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return &pb.JoinChallengeResponse{Success: false}, err
	}

	return &pb.JoinChallengeResponse{Success: true}, nil
}

func (s *SustainabilityRepo) UpdateChallengeProgress(in *pb.UpdateChallengeProgressRequest) (*pb.UpdateChallengeProgressResponse, error) {
	if in.ChallengeId != "" && in.Progress != 0 {
		
		_, err := s.DB.Exec(`
				UPDATE 
				user_challenges
				SET challengeId=$1,
					progress=$2 
				WHERE user_id=$3
					`, in.ChallengeId, in.Progress, in.UserId)

		if err != nil {
			return &pb.UpdateChallengeProgressResponse{Success: false}, err
		}
		return &pb.UpdateChallengeProgressResponse{Success: true}, nil
	}

	return &pb.UpdateChallengeProgressResponse{Success: false}, nil

}
