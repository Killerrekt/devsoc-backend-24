package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	user.Email = email

	var teamID uuid.NullUUID

	err := database.DB.QueryRow("SELECT id, first_name, last_name, reg_no, password, phone, college, gender, role, is_banned, is_added, is_vitian, is_verified, is_profile_complete, is_leader, team_id FROM users WHERE email = $1",
		email).
		Scan(&user.ID, &user.FirstName, &user.LastName, &user.RegNo, &user.Password, &user.Phone,
			&user.College, &user.Gender, &user.Role,
			&user.IsBanned, &user.IsAdded, &user.IsVitian, &user.IsVerified, &user.IsProfileComplete, &user.IsLeader, &teamID)
	if err != nil {
		return nil, err
	}

	if teamID.Valid {
		user.TeamID = teamID.UUID
	} else {
		user.TeamID = uuid.Nil
	}

	return &user, nil
}
