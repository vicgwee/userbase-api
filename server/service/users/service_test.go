package user

import (
	"context"
	"sort"
	"testing"
	"time"
	"userbase-api/server/dal"
	ue "userbase-api/server/utils/errors"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

var (
	mockDocs = map[string]*dal.UserDocument{
		"id0": {
			Id:          proto.String("id0"),
			Name:        proto.String("name0"),
			DateOfBirth: proto.String("dob0"),
			Address:     proto.String("addr0"),
			Desc:        proto.String("desc0"),
			CreateTs:    proto.Int64(1659312000),
		},
		"id1": {
			Id:          proto.String("id1"),
			Name:        proto.String("name1"),
			DateOfBirth: proto.String("dob1"),
			Address:     proto.String("addr1"),
			Desc:        proto.String("desc1"),
			CreateTs:    proto.Int64(1659398400),
		},
	}

	mockUsers = []*dal.User{
		{
			Id:          proto.String("id0"),
			Name:        proto.String("name0"),
			DateOfBirth: proto.String("dob0"),
			Address:     proto.String("addr0"),
			Desc:        proto.String("desc0"),
			CreateDate:  proto.String("20220801"),
		},
		{
			Id:          proto.String("id1"),
			Name:        proto.String("name1"),
			DateOfBirth: proto.String("dob1"),
			Address:     proto.String("addr1"),
			Desc:        proto.String("desc1"),
			CreateDate:  proto.String("20220802"),
		},
	}

	newUser = &dal.User{
		Id:          proto.String("id2"),
		Name:        proto.String("name2"),
		DateOfBirth: proto.String("dob2"),
		Address:     proto.String("addr2"),
		Desc:        proto.String("desc2"),
	}

	newCreatedUser = &dal.User{
		Id:          proto.String("id2"),
		Name:        proto.String("name2"),
		DateOfBirth: proto.String("dob2"),
		Address:     proto.String("addr2"),
		Desc:        proto.String("desc2"),
		CreateDate:  proto.String(time.Now().Format("20060102")),
	}

	updatedUser = &dal.User{
		Id:          proto.String("id0"),
		Name:        proto.String("newName0"),
		DateOfBirth: proto.String("newDob0"),
		Address:     proto.String("newAddr0"),
		Desc:        proto.String("newDesc0"),
	}

	newUpdatedUser = &dal.User{
		Id:          proto.String("id0"),
		Name:        proto.String("newName0"),
		DateOfBirth: proto.String("newDob0"),
		Address:     proto.String("newAddr0"),
		Desc:        proto.String("newDesc0"),
		CreateDate:  proto.String("20220801"),
	}
)

func getMockDocs() map[string]*dal.UserDocument {
	res := make(map[string]*dal.UserDocument)
	for id, doc := range mockDocs {
		res[id] = &dal.UserDocument{
			Id:          doc.Id,
			Name:        doc.Name,
			DateOfBirth: doc.DateOfBirth,
			Address:     doc.Address,
			Desc:        doc.Desc,
			CreateTs:    doc.CreateTs,
		}
	}
	return res
}

func TestGetAll(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name    string
		want    []*dal.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "GetAll",
			want:    mockUsers,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo = &MockUserRepo{docs: mockDocs}
			got, err := GetAll(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.EqualValues(t, tt.want, got) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	ctx := context.Background()
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *dal.User
		wantErr bool
	}{
		{
			name:    "Get Valid ID",
			args:    args{"id0"},
			want:    mockUsers[0],
			wantErr: false,
		},
		{
			name:    "Get Invalid ID",
			args:    args{"invalid"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo = &MockUserRepo{docs: getMockDocs()}
			got, err := Get(ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	type args struct {
		user *dal.User
	}
	tests := []struct {
		name    string
		args    args
		want    *dal.User
		wantErr bool
	}{
		{
			name:    "Update Existing User",
			args:    args{updatedUser},
			want:    newUpdatedUser,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo = &MockUserRepo{docs: getMockDocs()}
			got, err := Update(ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	ctx := context.Background()
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Delete Existing User",
			args:    args{"id0"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo = &MockUserRepo{docs: getMockDocs()}
			if err := Delete(ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	ctx := context.Background()
	type args struct {
		user *dal.User
	}
	tests := []struct {
		name    string
		args    args
		want    *dal.User
		wantErr bool
	}{
		{
			name:    "Create New User",
			args:    args{newUser},
			want:    newCreatedUser,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo = &MockUserRepo{docs: getMockDocs()}
			got, err := Create(ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

type MockUserRepo struct {
	docs map[string]*dal.UserDocument
}

func (m *MockUserRepo) CreateUser(ctx context.Context, doc *dal.UserDocument) (*dal.UserDocument, error) {
	if _, found := m.docs[*doc.Id]; found {
		return nil, ue.NewUserDuplicateError("")
	}
	m.docs[*doc.Id] = doc
	return m.docs[*doc.Id], nil
}

func (m *MockUserRepo) UpdateUser(ctx context.Context, doc *dal.UserDocument) (*dal.UserDocument, error) {
	if _, found := m.docs[*doc.Id]; !found {
		return nil, ue.NewUserNotFoundError("")
	}
	oldDoc := m.docs[*doc.Id]
	if doc.Name != nil {
		oldDoc.Name = doc.Name
	}
	if doc.DateOfBirth != nil {
		oldDoc.DateOfBirth = doc.DateOfBirth
	}
	if doc.Address != nil {
		oldDoc.Address = doc.Address
	}
	if doc.Desc != nil {
		oldDoc.Desc = doc.Desc
	}
	return m.docs[*doc.Id], nil
}

func (m *MockUserRepo) GetUser(ctx context.Context, id string) (*dal.UserDocument, error) {
	if _, found := m.docs[id]; !found {
		return nil, ue.NewUserNotFoundError("")
	}
	return m.docs[id], nil
}

func (m *MockUserRepo) GetUsers(ctx context.Context) ([]*dal.UserDocument, error) {
	var res []*dal.UserDocument
	var ids []string
	for id := range m.docs {
		ids = append(ids, id)
	}
	sort.Strings(ids) //for stable output
	for _, id := range ids {
		res = append(res, m.docs[id])
	}
	return res, nil
}

func (m *MockUserRepo) DeleteUser(ctx context.Context, id string) error {
	delete(m.docs, id)
	return nil
}
