package repositories

import (
	"library-mngmt/domain/entities"
)

func (dr *UserRepository) GetUserByEmail(email string) (entities.User, error) {
	var user entities.User
	query := `SELECT * FROM users WHERE email = $1`
	row := dr.Database.QueryRow(query, email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleID, &user.APIKey)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (dr *UserRepository) CreateUser(user entities.User) error {
	query := `INSERT INTO users (name, email, password, role_id, api_key) VALUES ($1, $2, $3, $4, $5)`
	_, err := dr.Database.Exec(query, user.Name, user.Email, user.Password, user.RoleID, user.APIKey)
	if err != nil {
		return err
	}
	return nil
}

// --- need to check updated values
func (dr *UserRepository) UpdateUser(user entities.User) error {
	query := `UPDATE users SET name = $1, email = $2, role_id = $3, WHERE id = $4`
	_, err := dr.Database.Exec(query, user.Name, user.Email, user.RoleID, user.ID)
	if err != nil {
		return err
	}
	return nil
}
func (dr *UserRepository) UpdatePassword(user entities.User) error {
	query := `UPDATE users SET password = $1 WHERE id = $2`
	_, err := dr.Database.Exec(query, user.Password, user.ID)
	if err != nil {
		return err
	}
	return nil
}
func (dr *UserRepository) UpdateRefreshKey(user entities.UserToken) error {
	query := `UPDATE refresh_tokens SET token = $1, expires_at = $2 WHERE id = $3`
	_, err := dr.Database.Exec(query, user.RefreshToken, user.ExpiresAt, user.ID)
	if err != nil {
		return err
	}
	return nil
}
func (dr *UserRepository) DeleteUser(userID uint) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := dr.Database.Exec(query, userID)
	if err != nil {
		return err
	}
	return nil
}

func (dr *UserRepository) GetUserByID(id uint) (entities.User, error) {
	var user entities.User
	query := `SELECT * FROM users WHERE id = $1`
	row := dr.Database.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleID, &user.APIKey)
	if err != nil {
		return user, err
	}
	return user, nil

}
func (dr *UserRepository) GetUserList() ([]entities.User, error) {
	var users []entities.User
	query := `SELECT * FROM users`
	rows, err := dr.Database.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleID, &user.APIKey)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (dr *UserRepository) GetBorrowedHistoryByUserID(id uint) ([]entities.BorrowRecord, error) {
	var borrowedHistories []entities.BorrowRecord
	query := `SELECT * FROM borrowed_histories WHERE user_id = $1`
	rows, err := dr.Database.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var borrowedHistory entities.BorrowRecord
		err := rows.Scan(&borrowedHistory.ID, &borrowedHistory.UserID, &borrowedHistory.BookID, &borrowedHistory.BorrowedAt, &borrowedHistory.ReturnedAt)
		if err != nil {
			return nil, err
		}
		borrowedHistories = append(borrowedHistories, borrowedHistory)
	}
	return borrowedHistories, nil
}
