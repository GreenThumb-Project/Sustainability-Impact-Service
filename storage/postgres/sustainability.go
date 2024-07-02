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

func (s *SustainabilityRepo) GetCommunityImpact(req string) (*pb.GetCommunityImpactResponse, error) {
	res := pb.GetCommunityImpactResponse{}
	err := s.DB.QueryRow(`
		SELECT 
			id,
			user_id,
			category,
			amount,
			unit,
			logged_at 
		FROM 
			impact_logs 
		where 
			user_id=$1
		`, req).Scan(&res.Id, &res.UserId, &res.Category, &res.GoalAmount, &res.GoalUnit, &res.LoggedAt)
	return &res, err
}

func (s *SustainabilityRepo) GetChallenges() (*pb.GetChallengesResponse, error) {
	res := []*pb.Challenge{}
	rows, err := s.DB.Query(`
		SELECT 
			id,
			title,
			description,
			goal_amount,
			goal_unit 
		FROM 
			user_challenges
	`)
	if err != nil {
		return &pb.GetChallengesResponse{}, err
	}
	for rows.Next() {
		res1 := pb.Challenge{}
		err = rows.Scan(&res1.Id, &res1.Title, &res1.Description, &res1.GoalAmount, &res1.GoalUnit)
		if err != nil {
			return &pb.GetChallengesResponse{}, err
		}
		res = append(res, &res1)
	}
	return &pb.GetChallengesResponse{Challanges: res}, nil
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