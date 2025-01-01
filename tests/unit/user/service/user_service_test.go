package test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/ahargunyllib/hackathon-fiber-starter/domain"
	"github.com/ahargunyllib/hackathon-fiber-starter/domain/dto"
	"github.com/ahargunyllib/hackathon-fiber-starter/domain/entity"
	userSvc "github.com/ahargunyllib/hackathon-fiber-starter/internal/app/user/service"
	uuidMock "github.com/ahargunyllib/hackathon-fiber-starter/pkg/uuid/mock"
	validatorMock "github.com/ahargunyllib/hackathon-fiber-starter/pkg/validator/mock"
	"github.com/ahargunyllib/hackathon-fiber-starter/tests/unit/user/fixture"
	userRepoMock "github.com/ahargunyllib/hackathon-fiber-starter/tests/unit/user/repository/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type mockObjects struct {
	userRepo  *userRepoMock.MockUserRepository
	validator *validatorMock.MockValidatorInterface
	uuid      *uuidMock.MockUUIDInterface
}

func TestGetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := userRepoMock.NewMockUserRepository(ctrl)
	validator := validatorMock.NewMockValidatorInterface(ctrl)
	uuid := uuidMock.NewMockUUIDInterface(ctrl)

	mockObjs := mockObjects{
		userRepo:  userRepo,
		validator: validator,
		uuid:      uuid,
	}

	type params struct {
		ctx   context.Context
		query dto.GetUsersQuery
	}

	tests := []struct {
		name        string
		params      params
		beforeTests func(params params, mockObjects mockObjects)
		want        dto.GetUsersResponse
		wantErr     error
	}{
		{
			name: "when getting users, it should return the users",
			params: params{
				ctx: context.TODO(),
				query: dto.GetUsersQuery{
					Limit:          10,
					Page:           1,
					SortBy:         "created_at",
					Order:          "desc",
					IncludeDeleted: false,
					Search:         "",
				},
			},
			beforeTests: func(params params, mockObjects mockObjects) {
				mockObjects.validator.EXPECT().Validate(params.query).Return(nil)
				mockObjects.userRepo.EXPECT().GetUsers(params.ctx, params.query).Return([]entity.User{fixture.ActiveUser1, fixture.ActiveUser2}, nil)
			},
			want: dto.GetUsersResponse{
				Users: []entity.User{fixture.ActiveUser1, fixture.ActiveUser2},
			},
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userService := userSvc.NewUserService(mockObjs.userRepo, mockObjs.validator, mockObjs.uuid)

			if test.beforeTests != nil {
				test.beforeTests(test.params, mockObjs)
			}

			users, err := userService.GetUsers(test.params.ctx, test.params.query)

			if test.wantErr != nil {
				assert.Equal(t, test.wantErr, err, "error should be equal")
			} else {
				assert.Nil(t, err, "error should be nil")
			}

			assert.Equal(t, test.want, users, "users should be equal")
		})
	}

}

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := userRepoMock.NewMockUserRepository(ctrl)
	validator := validatorMock.NewMockValidatorInterface(ctrl)
	uuid := uuidMock.NewMockUUIDInterface(ctrl)

	mockObjs := mockObjects{
		userRepo:  userRepo,
		validator: validator,
		uuid:      uuid,
	}

	type params struct {
		ctx context.Context
		req dto.GetUserByIDRequest
	}

	tests := []struct {
		name        string
		params      params
		beforeTests func(params params, mockObjects mockObjects)
		want        dto.GetUserByIDResponse
		wantErr     error
	}{
		{
			name: "when getting user by id, it should return the user",
			params: params{
				ctx: context.TODO(),
				req: dto.GetUserByIDRequest{
					ID: fixture.ActiveUser1.ID,
				},
			},
			beforeTests: func(params params, mockObjects mockObjects) {
				mockObjects.validator.EXPECT().Validate(params.req).Return(nil)
				mockObjects.userRepo.EXPECT().GetUserByField(params.ctx, "id", params.req.ID.String()).Return(&fixture.ActiveUser1, nil)
			},
			want: dto.GetUserByIDResponse{
				User: fixture.ActiveUser1,
			},
			wantErr: nil,
		},
		{
			name: "when getting user by id and user not found, it should return user not found error",
			params: params{
				ctx: context.TODO(),
				req: dto.GetUserByIDRequest{
					ID: fixture.ActiveUser1.ID,
				},
			},
			beforeTests: func(params params, mockObjects mockObjects) {
				mockObjects.validator.EXPECT().Validate(params.req).Return(nil)
				mockObjects.userRepo.EXPECT().GetUserByField(params.ctx, "id", params.req.ID.String()).Return(nil, sql.ErrNoRows)
			},
			want:    dto.GetUserByIDResponse{},
			wantErr: domain.ErrUserNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userService := userSvc.NewUserService(mockObjs.userRepo, mockObjs.validator, mockObjs.uuid)

			if test.beforeTests != nil {
				test.beforeTests(test.params, mockObjs)
			}

			user, err := userService.GetUserByID(test.params.ctx, test.params.req)

			if test.wantErr != nil {
				assert.Equal(t, test.wantErr, err, "error should be equal")
			} else {
				assert.Nil(t, err, "error should be nil")
			}

			assert.Equal(t, test.want, user, "user should be equal")
		})
	}
}

func TestGetUsersStats(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := userRepoMock.NewMockUserRepository(ctrl)
	validator := validatorMock.NewMockValidatorInterface(ctrl)
	uuid := uuidMock.NewMockUUIDInterface(ctrl)

	mockObjs := mockObjects{
		userRepo:  userRepo,
		validator: validator,
		uuid:      uuid,
	}

	type params struct {
		ctx context.Context
	}

	tests := []struct {
		name        string
		params      params
		beforeTests func(params params, mockObjects mockObjects)
		want        dto.GetUsersStatsResponse
		wantErr     error
	}{
		{
			name: "when getting users stats, it should return the stats",
			params: params{
				ctx: context.TODO(),
			},
			beforeTests: func(params params, mockObjects mockObjects) {
				mockObjects.userRepo.EXPECT().CountUsers(params.ctx, gomock.Any()).Return(int64(3), nil)
				mockObjects.userRepo.EXPECT().CountUsers(params.ctx, gomock.Any()).Return(int64(2), nil)
			},
			want: dto.GetUsersStatsResponse{
				TotalNonDeletedUsers: 2,
				TotalDeletedUsers:    1,
				TotalUsers:           3,
			},
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userService := userSvc.NewUserService(mockObjs.userRepo, mockObjs.validator, uuid)

			if test.beforeTests != nil {
				test.beforeTests(test.params, mockObjs)
			}

			stats, err := userService.GetUsersStats(test.params.ctx)

			if test.wantErr != nil {
				assert.Equal(t, test.wantErr, err, "error should be equal")
			} else {
				assert.Nil(t, err, "error should be nil")
			}

			assert.Equal(t, test.want, stats, "stats should be equal")
		})
	}
}

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := userRepoMock.NewMockUserRepository(ctrl)
	validator := validatorMock.NewMockValidatorInterface(ctrl)
	uuid := uuidMock.NewMockUUIDInterface(ctrl)

	mockObjs := mockObjects{
		userRepo:  userRepo,
		validator: validator,
		uuid:      uuid,
	}

	type params struct {
		ctx context.Context
		req dto.CreateUserRequest
	}

	tests := []struct {
		name        string
		params      params
		beforeTests func(params params, mockObjects mockObjects)
		want        dto.CreateUserResponse
		wantErr     error
	}{
		{
			name: "when creating user, it should return the user id",
			params: params{
				ctx: context.TODO(),
				req: dto.CreateUserRequest{
					Name:     fixture.ActiveUser1.Name,
					Password: fixture.ActiveUser1.Password,
					Email:    fixture.ActiveUser1.Email,
				},
			},
			beforeTests: func(params params, mockObjects mockObjects) {
				mockObjects.validator.EXPECT().Validate(params.req).Return(nil)
				mockObjects.uuid.EXPECT().NewV7().Return(fixture.ActiveUser1.ID, nil)
				mockObjects.userRepo.EXPECT().GetUserByField(params.ctx, "email", params.req.Email).Return(nil, sql.ErrNoRows)
				mockObjects.userRepo.EXPECT().CreateUser(params.ctx, gomock.Any()).Return(fixture.ActiveUser1.ID, nil)
			},
			want: dto.CreateUserResponse{
				ID: fixture.ActiveUser1.ID,
			},
		},
		{
			name: "when creating user and user email already exists, it should return user email already exists error",
			params: params{
				ctx: context.TODO(),
				req: dto.CreateUserRequest{
					Name:     fixture.ActiveUser1.Name,
					Password: fixture.ActiveUser1.Password,
					Email:    fixture.ActiveUser1.Email,
				},
			},
			beforeTests: func(params params, mockObjects mockObjects) {
				mockObjects.validator.EXPECT().Validate(params.req).Return(nil)
				mockObjects.uuid.EXPECT().NewV7().Return(fixture.ActiveUser1.ID, nil)
				mockObjects.userRepo.EXPECT().GetUserByField(params.ctx, "email", params.req.Email).Return(&fixture.ActiveUser1, nil)
			},
			want:    dto.CreateUserResponse{},
			wantErr: domain.ErrUserEmailAlreadyExists,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userService := userSvc.NewUserService(mockObjs.userRepo, mockObjs.validator, mockObjs.uuid)

			if test.beforeTests != nil {
				test.beforeTests(test.params, mockObjs)
			}

			user, err := userService.CreateUser(test.params.ctx, test.params.req)

			if test.wantErr != nil {
				assert.Equal(t, test.wantErr, err, "error should be equal")
			} else {
				assert.Nil(t, err, "error should be nil")
			}

			assert.Equal(t, test.want, user, "user should be equal")
		})
	}
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := userRepoMock.NewMockUserRepository(ctrl)
	validator := validatorMock.NewMockValidatorInterface(ctrl)
	uuid := uuidMock.NewMockUUIDInterface(ctrl)

	mockObjs := mockObjects{
		userRepo:  userRepo,
		validator: validator,
		uuid:      uuid,
	}

	type params struct {
		ctx context.Context
		req dto.UpdateUserRequest
	}

	tests := []struct {
		name        string
		params      params
		beforeTests func(params params, mockObjects mockObjects)
		want        dto.UpdateUserResponse
		wantErr     error
	}{
		{
			name: "when updating user, it should return the user id",
			params: params{
				ctx: context.TODO(),
				req: dto.UpdateUserRequest{
					ID:       fixture.ActiveUser1.ID,
					Name:     fixture.ActiveUser1.Name,
					Password: fixture.ActiveUser1.Password,
					Email:    fixture.ActiveUser1.Email,
				},
			},
			beforeTests: func(params params, mockObjects mockObjects) {
				mockObjects.validator.EXPECT().Validate(params.req).Return(nil)
				mockObjects.userRepo.EXPECT().GetUserByField(params.ctx, "id", params.req.ID.String()).Return(&fixture.ActiveUser1, nil)
				mockObjects.userRepo.EXPECT().GetUserByField(params.ctx, "email", params.req.Email).Return(nil, sql.ErrNoRows)
				mockObjects.userRepo.EXPECT().UpdateUser(params.ctx, gomock.Any()).Return(fixture.ActiveUser1.ID, nil)
			},
			want: dto.UpdateUserResponse{
				ID: fixture.ActiveUser1.ID,
			},
			wantErr: nil,
		},
		{
			name: "when updating user and user not found, it should return user not found error",
			params: params{
				ctx: context.TODO(),
				req: dto.UpdateUserRequest{
					ID:       fixture.ActiveUser1.ID,
					Name:     fixture.ActiveUser1.Name,
					Password: fixture.ActiveUser1.Password,
					Email:    fixture.ActiveUser1.Email,
				},
			},
			beforeTests: func(params params, mockObjects mockObjects) {
				mockObjects.validator.EXPECT().Validate(params.req).Return(nil)
				mockObjects.userRepo.EXPECT().GetUserByField(params.ctx, "id", params.req.ID.String()).Return(nil, sql.ErrNoRows)
			},
			want:    dto.UpdateUserResponse{},
			wantErr: domain.ErrUserNotFound,
		},
		{
			name: "when updating user and user email already exists, it should return user email already exists error",
			params: params{
				ctx: context.TODO(),
				req: dto.UpdateUserRequest{
					ID:       fixture.ActiveUser1.ID,
					Name:     fixture.ActiveUser1.Name,
					Password: fixture.ActiveUser1.Password,
					Email:    fixture.ActiveUser1.Email,
				},
			},
			beforeTests: func(params params, mockObjects mockObjects) {
				mockObjects.validator.EXPECT().Validate(params.req).Return(nil)
				mockObjects.userRepo.EXPECT().GetUserByField(params.ctx, "id", params.req.ID.String()).Return(&fixture.ActiveUser1, nil)
				mockObjects.userRepo.EXPECT().GetUserByField(params.ctx, "email", params.req.Email).Return(nil, nil)
			},
			want:    dto.UpdateUserResponse{},
			wantErr: domain.ErrUserEmailAlreadyExists,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userService := userSvc.NewUserService(mockObjs.userRepo, mockObjs.validator, mockObjs.uuid)

			if test.beforeTests != nil {
				test.beforeTests(test.params, mockObjs)
			}

			user, err := userService.UpdateUser(test.params.ctx, test.params.req)

			if test.wantErr != nil {
				assert.Equal(t, test.wantErr, err, "error should be equal")
			} else {
				assert.Nil(t, err, "error should be nil")
			}

			assert.Equal(t, test.want, user, "user should be equal")
		})
	}
}

func TestSoftDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := userRepoMock.NewMockUserRepository(ctrl)
	validator := validatorMock.NewMockValidatorInterface(ctrl)
	uuid := uuidMock.NewMockUUIDInterface(ctrl)

	mockObjs := mockObjects{
		userRepo:  userRepo,
		validator: validator,
		uuid:      uuid,
	}

	type params struct {
		ctx context.Context
		req dto.SoftDeleteUserRequest
	}

	tests := []struct {
		name        string
		params      params
		beforeTests func(params params, mockObjects mockObjects)
		want        dto.SoftDeleteUserResponse
		wantErr     error
	}{
		{
			name: "when soft deleting user, it should return the user id",
			params: params{
				ctx: context.TODO(),
				req: dto.SoftDeleteUserRequest{
					ID: fixture.ActiveUser1.ID,
				},
			},
			beforeTests: func(params params, mockObjects mockObjects) {
				mockObjects.validator.EXPECT().Validate(params.req).Return(nil)
				mockObjects.userRepo.EXPECT().SoftDeleteUser(params.ctx, params.req.ID).Return(params.req.ID, nil)
			},
			want: dto.SoftDeleteUserResponse{
				ID: fixture.ActiveUser1.ID,
			},
			wantErr: nil,
		},
		{
			name: "when soft deleting user and user not found, it should return user not found error",
			params: params{
				ctx: context.TODO(),
				req: dto.SoftDeleteUserRequest{
					ID: fixture.ActiveUser1.ID,
				},
			},
			beforeTests: func(params params, mockObjects mockObjects) {
				mockObjects.validator.EXPECT().Validate(params.req).Return(nil)
				mockObjects.userRepo.EXPECT().SoftDeleteUser(params.ctx, params.req.ID).Return(params.req.ID, sql.ErrNoRows)
			},
			want:    dto.SoftDeleteUserResponse{},
			wantErr: domain.ErrUserNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userService := userSvc.NewUserService(mockObjs.userRepo, mockObjs.validator, mockObjs.uuid)

			if test.beforeTests != nil {
				test.beforeTests(test.params, mockObjs)
			}

			user, err := userService.SoftDeleteUser(test.params.ctx, test.params.req)

			if test.wantErr != nil {
				assert.Equal(t, test.wantErr, err, "error should be equal")
			} else {
				assert.Nil(t, err, "error should be nil")
			}

			assert.Equal(t, test.want, user, "user should be equal")
		})
	}
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := userRepoMock.NewMockUserRepository(ctrl)
	validator := validatorMock.NewMockValidatorInterface(ctrl)
	uuid := uuidMock.NewMockUUIDInterface(ctrl)

	mockObjs := mockObjects{
		userRepo:  userRepo,
		validator: validator,
		uuid:      uuid,
	}

	type params struct {
		ctx context.Context
		req dto.DeleteUserRequest
	}

	tests := []struct {
		name        string
		params      params
		beforeTests func(params params, mockObjects mockObjects)
		want        dto.DeleteUserResponse
		wantErr     error
	}{
		{
			name: "when deleting user, it should return the user id",
			params: params{
				ctx: context.TODO(),
				req: dto.DeleteUserRequest{
					ID: fixture.ActiveUser1.ID,
				},
			},
			beforeTests: func(params params, mockObjects mockObjects) {
				mockObjects.validator.EXPECT().Validate(params.req).Return(nil)
				mockObjects.userRepo.EXPECT().DeleteUser(params.ctx, params.req.ID).Return(params.req.ID, nil)
			},
			want: dto.DeleteUserResponse{
				ID: fixture.ActiveUser1.ID,
			},
			wantErr: nil,
		},
		{
			name: "when deleting user and user not found, it should return user not found error",
			params: params{
				ctx: context.TODO(),
				req: dto.DeleteUserRequest{
					ID: fixture.ActiveUser1.ID,
				},
			},
			beforeTests: func(params params, mockObjects mockObjects) {
				mockObjects.validator.EXPECT().Validate(params.req).Return(nil)
				mockObjects.userRepo.EXPECT().DeleteUser(params.ctx, params.req.ID).Return(params.req.ID, sql.ErrNoRows)
			},
			want:    dto.DeleteUserResponse{},
			wantErr: domain.ErrUserNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userService := userSvc.NewUserService(mockObjs.userRepo, mockObjs.validator, mockObjs.uuid)

			if test.beforeTests != nil {
				test.beforeTests(test.params, mockObjs)
			}

			user, err := userService.DeleteUser(test.params.ctx, test.params.req)

			if test.wantErr != nil {
				assert.Equal(t, test.wantErr, err, "error should be equal")
			} else {
				assert.Nil(t, err, "error should be nil")
			}

			assert.Equal(t, test.want, user, "user should be equal")
		})
	}
}

func TestRestoreUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := userRepoMock.NewMockUserRepository(ctrl)
	validator := validatorMock.NewMockValidatorInterface(ctrl)
	uuid := uuidMock.NewMockUUIDInterface(ctrl)

	mockObjs := mockObjects{
		userRepo:  userRepo,
		validator: validator,
		uuid:      uuid,
	}

	type params struct {
		ctx context.Context
		req dto.RestoreUserRequest
	}

	tests := []struct {
		name        string
		params      params
		beforeTests func(params params, mockObjects mockObjects)
		want        dto.RestoreUserResponse
		wantErr     error
	}{
		{
			name: "when restoring user, it should return the user id",
			params: params{
				ctx: context.TODO(),
				req: dto.RestoreUserRequest{
					ID: fixture.ActiveUser1.ID,
				},
			},
			beforeTests: func(params params, mockObjects mockObjects) {
				mockObjects.validator.EXPECT().Validate(params.req).Return(nil)
				mockObjects.userRepo.EXPECT().RestoreUser(params.ctx, params.req.ID).Return(params.req.ID, nil)
			},
			want: dto.RestoreUserResponse{
				ID: fixture.ActiveUser1.ID,
			},
			wantErr: nil,
		},
		{
			name: "when restoring user and user not found, it should return user not found error",
			params: params{
				ctx: context.TODO(),
				req: dto.RestoreUserRequest{
					ID: fixture.ActiveUser1.ID,
				},
			},
			beforeTests: func(params params, mockObjects mockObjects) {
				mockObjects.validator.EXPECT().Validate(params.req).Return(nil)
				mockObjects.userRepo.EXPECT().RestoreUser(params.ctx, params.req.ID).Return(params.req.ID, sql.ErrNoRows)
			},
			want:    dto.RestoreUserResponse{},
			wantErr: domain.ErrUserNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userService := userSvc.NewUserService(mockObjs.userRepo, mockObjs.validator, mockObjs.uuid)

			if test.beforeTests != nil {
				test.beforeTests(test.params, mockObjs)
			}

			user, err := userService.RestoreUser(test.params.ctx, test.params.req)

			if test.wantErr != nil {
				assert.Equal(t, test.wantErr, err, "error should be equal")
			} else {
				assert.Nil(t, err, "error should be nil")
			}

			assert.Equal(t, test.want, user, "user should be equal")
		})
	}
}
