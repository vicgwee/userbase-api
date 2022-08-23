package user

import (
	"context"
	"time"
	"userbase-api/server/dal"
	"userbase-api/server/dal/mongoDB"
	ue "userbase-api/server/utils/errors"

	"google.golang.org/protobuf/proto"
)

var userRepo mongoDB.UserRepository

const dateFormat = "20060102"

func init() {
	userRepo = mongoDB.NewUserRepository()
}

func Create(ctx context.Context, user *dal.User) (*dal.User, error) {
	doc, err := validateUserCreate(user)
	if err != nil {
		return nil, err
	}
	doc.CreateTs = proto.Int64(time.Now().Unix())

	newDoc, err := userRepo.CreateUser(ctx, doc)
	if err != nil {
		return nil, err
	}

	return fromDocument(newDoc)
}

func validateUserCreate(user *dal.User) (*dal.UserDocument, error) {
	if user.Id == nil {
		return nil, ue.NewUserInvalidError("create missing userID")
	}
	if user.CreateDate != nil {
		return nil, ue.NewUserInvalidError("create createDate specified")
	}
	return toDocument(user)
}

func Update(ctx context.Context, user *dal.User) (*dal.User, error) {
	doc, err := validateUserUpdate(user)
	if err != nil {
		return nil, err
	}

	newDoc, err := userRepo.UpdateUser(ctx, doc)
	if err != nil {
		return nil, err
	}

	return fromDocument(newDoc)
}

func validateUserUpdate(user *dal.User) (*dal.UserDocument, error) {
	if user.Id == nil {
		return nil, ue.NewUserInvalidError("update missing userID")
	}
	if user.CreateDate != nil {
		return nil, ue.NewUserInvalidError("update createDate specified")
	}
	return toDocument(user)
}

func GetAll(ctx context.Context) ([]*dal.User, error) {
	var users []*dal.User
	docs, err := userRepo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	for _, doc := range docs {
		user, err := fromDocument(doc)
		if err != nil {
			continue
		}
		users = append(users, user)
	}

	return users, nil
}

func Get(ctx context.Context, id string) (*dal.User, error) {
	doc, err := userRepo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return fromDocument(doc)
}

func Delete(ctx context.Context, id string) error {
	return userRepo.DeleteUser(ctx, id)
}

func toDocument(user *dal.User) (*dal.UserDocument, error) {
	doc := &dal.UserDocument{
		Id:          user.Id,
		Name:        user.Name,
		DateOfBirth: user.DateOfBirth,
		Address:     user.Address,
		Desc:        user.Desc,
	}

	if user.CreateDate == nil {
		return doc, nil
	}
	createTime, err := time.Parse(dateFormat, *user.CreateDate)
	if err != nil {
		return nil, err
	}
	doc.CreateTs = proto.Int64(createTime.Unix())

	return doc, nil
}

func fromDocument(doc *dal.UserDocument) (*dal.User, error) {
	user := &dal.User{
		Id:          doc.Id,
		Name:        doc.Name,
		DateOfBirth: doc.DateOfBirth,
		Address:     doc.Address,
		Desc:        doc.Desc,
	}

	if doc.CreateTs == nil {
		return user, nil
	}
	user.CreateDate = proto.String(time.Unix(*doc.CreateTs, 0).Format(dateFormat))

	return user, nil
}
