package service

import (
	"context"
	"thelastking/kingseafood/model"
)

func (sql *sql) SignUp(ctx context.Context, data *model.Users) (*model.Users, error) {
	if err := sql.db.Table("users").FirstOrCreate(&data, &model.Users{Email: data.Email}).Error; err != nil {
		return nil, err
	}
	return data, nil
}
