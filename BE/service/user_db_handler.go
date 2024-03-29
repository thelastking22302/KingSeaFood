package service

import (
	"context"
	"fmt"
	"thelastking/kingseafood/model"
)

func (sql *sql) SignUp(ctx context.Context, data *model.Users) (*model.Users, error) {
	fmt.Printf("sql: %v\n", sql)
	fmt.Printf("sql.db: %v\n", sql.db)
	if err := sql.db.Table("users").FirstOrCreate(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
