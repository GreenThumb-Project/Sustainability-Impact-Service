package postgres

import (
	"database/sql"
	pb "sustainability-service/generated/sustainability"
)

type SustainabilityRepo struct {
	DB *sql.DB
}

func (s *SustainabilityRepo) GetUsersChallenges(id string) (*pb.GetUserChallengesResponse, error) {
	var userChallenges []*pb.UserChallenge
	rows, err := s.DB.Query(`
		SELECT
			user_id,
			challenge_id,
			progress,
			completed_at
		FROM
			user_challenges
		WHERE 
			id = $1
	`, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var userChallenge pb.UserChallenge

		err = rows.Scan(userChallenge.UserId, &userChallenge.ChallengeId, &userChallenge.Progress, &userChallenge.CompletedAt)

		if err != nil {
			return nil, err
		}

		userChallenges = append(userChallenges, &userChallenge)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &pb.GetUserChallengesResponse{UserChallenges: userChallenges}, nil
}

func (s *SustainabilityRepo) GetUsersLeaderboard() ([]*pb.UsersLeaderboard, error) {
	var lederboards []*pb.UsersLeaderboard

	rows, err := s.DB.Query(`
		SELECT
			user_id,
			sum(progress) AS points
		FROM
			user_challenges
		GROUP BY
			user_id
		ORDER BY
			points DESC
	`, )

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var lederboard pb.UsersLeaderboard

		err = rows.Scan(lederboard.UserId, lederboard.Points)
		if err != nil {
			return nil, err
		}

		lederboards = append(lederboards, &lederboard)
	}

	return lederboards, nil
}