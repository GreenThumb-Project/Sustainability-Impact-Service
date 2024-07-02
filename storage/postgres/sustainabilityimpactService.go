package postgres

import (
	"database/sql"
	"sustainability-service/generated/sustainability"
)

type SustainabilityImpactService struct {
	db *sql.DB
}

func ConnectSustainability(db *sql.DB) *SustainabilityImpactService {
	return &SustainabilityImpactService{db: db}
}

func (S *SustainabilityImpactService) GetCommunityImpact(req string) (*sustainability.GetCommunityImpactResponse, error) {
	res := sustainability.GetCommunityImpactResponse{}
	err := S.db.QueryRow("SELECT id,user_id,category,amount,unit,logged_at FROM impact_logs where user_id=$1", req).Scan(&res.Id, &res.UserId, &res.Category, &res.GoalAmount, &res.GoalUnit, &res.LoggedAt)
	return &res, err
}

func (S *SustainabilityImpactService) GetChallenges() (*sustainability.GetChallengesResponse, error) {
	res := []*sustainability.Challenge{}
	rows, err := S.db.Query(`SELECT id,title,description,goal_amount,goal_unit FROM user_challenges`)
	if err != nil {
		return &sustainability.GetChallengesResponse{}, err
	}
	for rows.Next() {
		res1 := sustainability.Challenge{}
		err = rows.Scan(&res1.Id, &res1.Title, &res1.Description, &res1.GoalAmount, &res1.GoalUnit)
		if err != nil {
			return &sustainability.GetChallengesResponse{}, err
		}
		res = append(res, &res1)
	}
	return &sustainability.GetChallengesResponse{Challanges: res}, nil
}
