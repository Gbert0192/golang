package memberships

import (
	"belajar/internal/model/memberships"
	"context"
	"database/sql"
)

func (r *repository) GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error) {
	query := `select id, email, password, username, created_at, updated_at, updated_by from users where email = ? or username = ?`
	row := r.db.QueryRowContext(ctx, query, email, username)
	var response memberships.UserModel // ini sama seperti response := membership.UserModel{}
	err := row.Scan(&response.ID, &response.Email, &response.Password, &response.Username, &response.CreatedAt, &response.CreatedBy, &response.UpdatedAt, &response.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &response, nil
}

func (r *repository) CreateUser(ctx context.Context, model memberships.UserModel) error {
	query := `insert into users (email, password, username, created_at, updated_at, created_by, updated_by) values (?,?,?,?,?,?,?)`
	_, err := r.db.ExecContext(ctx, query, model.Email, model.Password, model.Username, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}
