package data

import (
	"context"

	"user/ent/user"
	"user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *UserRepo) Save(ctx context.Context, g *biz.User) (*biz.User, error) {
	user, err := r.data.entClient.User.
		Create().
		SetName(g.Name).
		SetPhone(g.Phone).
		SetPassword(g.Password).
		SetSex(user.Sex(g.Sex)).
		SetAge(int8(g.Age)).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &biz.User{
		ID:    int64(user.ID),
		Name:  user.Name,
		Phone: user.Phone,
		Sex:   user.Sex.String(),
		Age:   int(user.Age),
	}, nil
}

func (r *UserRepo) Update(ctx context.Context, g *biz.User) (*biz.User, error) {
	user, err := r.data.entClient.User.
		UpdateOneID(int(g.ID)).
		SetName(g.Name).
		SetPassword(g.Password).
		SetPhone(g.Phone).
		SetSex(user.Sex(g.Sex)).
		SetAge(int8(g.Age)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		ID:    int64(user.ID),
		Name:  user.Name,
		Phone: user.Phone,
		Sex:   user.Sex.String(),
		Age:   int(user.Age),
	}, nil
}

func (r *UserRepo) FindByPhone(ctx context.Context, phone string) (*biz.User, error) {
	user, err := r.data.entClient.User.
		Query().
		Where(user.Phone(phone)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		ID:    int64(user.ID),
		Name:  user.Name,
		Phone: user.Phone,
		Sex:   user.Sex.String(),
		Age:   int(user.Age),
	}, nil
}

func (r *UserRepo) ListAll(ctx context.Context) ([]*biz.User, error) {
	users, err := r.data.entClient.User.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	busers := make([]*biz.User, 0)
	for _, v := range users {
		busers = append(busers, &biz.User{
			ID:    int64(v.ID),
			Name:  v.Name,
			Phone: v.Phone,
			Sex:   v.Sex.String(),
			Age:   int(v.Age),
		})

	}
	return busers, nil
}

func (r *UserRepo) Delete(ctx context.Context, id int64) error {
	err := r.data.entClient.User.
		DeleteOneID(int(id)).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
