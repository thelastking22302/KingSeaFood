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

func (sql *sql) SignIn(ctx context.Context, data *model.Users) error {
	var user model.Users
	if err := sql.db.Table("users").Where("email = ?", data.Email).First(&user).Error; err != nil {
		return err
	}
	data.Email = user.Email
	data.Password = user.Password
	return nil
}
