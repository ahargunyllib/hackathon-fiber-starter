package test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ahargunyllib/hackathon-fiber-starter/domain/dto"
	"github.com/ahargunyllib/hackathon-fiber-starter/domain/entity"
	userRepo "github.com/ahargunyllib/hackathon-fiber-starter/internal/app/user/repository"
	psqlConnMock "github.com/ahargunyllib/hackathon-fiber-starter/internal/infra/database/mock"
	"github.com/ahargunyllib/hackathon-fiber-starter/pkg/helpers"
	userFixture "github.com/ahargunyllib/hackathon-fiber-starter/tests/unit/user/fixture"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	type params struct {
		ctx   context.Context
		query dto.GetUsersQuery
	}

	tests := []struct {
		name        string
		params      params
		beforeTests func(params params, mock sqlmock.Sqlmock)
		want        []entity.User
		wantErr     error
	}{
		{
			name: "when fetching empty users, it should return empty users",
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
			beforeTests: func(params params, mock sqlmock.Sqlmock) {
				expectedRows := sqlmock.NewRows(userFixture.Rows)

				expectedSQL := `SELECT \* FROM users WHERE 1=1 AND deleted_at IS NULL ORDER BY created_at desc LIMIT \? OFFSET \?`

				limit := params.query.Limit
				offset := params.query.Limit * (params.query.Page - 1)

				mock.ExpectQuery(expectedSQL).
					WithArgs(limit, offset).
					WillReturnRows(expectedRows)
			},
			want:    []entity.User{},
			wantErr: nil,
		},
		{
			name: "when fetching users, it should return users",
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
			beforeTests: func(params params, mock sqlmock.Sqlmock) {
				expectedRows := sqlmock.NewRows(userFixture.Rows).
					AddRows(
						helpers.ConvertToDriverValue(helpers.StructToSlice(userFixture.ActiveUser1)),
						helpers.ConvertToDriverValue(helpers.StructToSlice(userFixture.ActiveUser2)),
					)

				expectedSQL := `SELECT \* FROM users WHERE 1=1 AND deleted_at IS NULL ORDER BY created_at desc LIMIT \? OFFSET \?`

				limit := params.query.Limit
				offset := params.query.Limit * (params.query.Page - 1)

				mock.ExpectQuery(expectedSQL).
					WithArgs(limit, offset).
					WillReturnRows(expectedRows)
			},
			want: []entity.User{
				userFixture.ActiveUser1,
				userFixture.ActiveUser2,
			},
			wantErr: nil,
		},
		{
			name: "when fetching users without deleted ones, it should return active users",
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
			beforeTests: func(params params, mock sqlmock.Sqlmock) {
				expectedRows := sqlmock.NewRows(userFixture.Rows).
					AddRows(
						helpers.ConvertToDriverValue(helpers.StructToSlice(userFixture.ActiveUser1)),
					)

				expectedSQL := `SELECT \* FROM users WHERE 1=1 AND deleted_at IS NULL ORDER BY created_at desc LIMIT \? OFFSET \?`

				limit := params.query.Limit
				offset := params.query.Limit * (params.query.Page - 1)

				mock.ExpectQuery(expectedSQL).
					WithArgs(limit, offset).
					WillReturnRows(expectedRows)
			},
			want: []entity.User{
				userFixture.ActiveUser1,
			},
			wantErr: nil,
		},
		{
			name: "when fetching users with deleted ones, it should return active users",
			params: params{
				ctx: context.TODO(),
				query: dto.GetUsersQuery{
					Limit:          10,
					Page:           1,
					SortBy:         "created_at",
					Order:          "desc",
					IncludeDeleted: true,
					Search:         "",
				},
			},
			beforeTests: func(params params, mock sqlmock.Sqlmock) {
				expectedRows := sqlmock.NewRows(userFixture.Rows).
					AddRows(
						helpers.ConvertToDriverValue(helpers.StructToSlice(userFixture.ActiveUser1)),
						helpers.ConvertToDriverValue(helpers.StructToSlice(userFixture.InactiveUser1)),
					)

				expectedSQL := `SELECT \* FROM users WHERE 1=1 ORDER BY created_at desc LIMIT \? OFFSET \?`

				limit := params.query.Limit
				offset := params.query.Limit * (params.query.Page - 1)

				mock.ExpectQuery(expectedSQL).
					WithArgs(limit, offset).
					WillReturnRows(expectedRows)
			},
			want: []entity.User{
				userFixture.ActiveUser1,
				userFixture.InactiveUser1,
			},
			wantErr: nil,
		},
		{
			name: "when fetching users with search query, it should return searched users",
			params: params{
				ctx: context.TODO(),
				query: dto.GetUsersQuery{
					Limit:          10,
					Page:           1,
					SortBy:         "created_at",
					Order:          "desc",
					IncludeDeleted: false,
					Search:         "activeUser",
				},
			},
			beforeTests: func(params params, mock sqlmock.Sqlmock) {
				expectedRows := sqlmock.NewRows(userFixture.Rows).
					AddRows(
						helpers.ConvertToDriverValue(helpers.StructToSlice(userFixture.ActiveUser1)),
						helpers.ConvertToDriverValue(helpers.StructToSlice(userFixture.ActiveUser2)),
					)

				expectedSQL := `SELECT \* FROM users WHERE 1=1 AND deleted_at IS NULL AND \(name LIKE \? OR email LIKE \?\) ORDER BY created_at desc LIMIT \? OFFSET \?`

				search := "%" + params.query.Search + "%"
				limit := params.query.Limit
				offset := params.query.Limit * (params.query.Page - 1)

				mock.ExpectQuery(expectedSQL).
					WithArgs(search, search, limit, offset).
					WillReturnRows(expectedRows)
			},
			want: []entity.User{
				userFixture.ActiveUser1,
				userFixture.ActiveUser2,
			},
			wantErr: nil,
		},
		{
			name: "when fetching users and database is down, it should return error database down",
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
			beforeTests: func(params params, mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT \* FROM users WHERE 1=1 AND deleted_at IS NULL ORDER BY created_at desc LIMIT \? OFFSET \?`).
					WillReturnError(sql.ErrConnDone)
			},
			want:    nil,
			wantErr: sql.ErrConnDone,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, close, mock := psqlConnMock.NewMockDB(t)
			defer close()

			if test.beforeTests != nil {
				test.beforeTests(test.params, mock)
			}

			repo := userRepo.NewUserRepository(db)
			users, err := repo.GetUsers(test.params.ctx, test.params.query)

			if test.wantErr != nil {
				assert.Equal(t, test.wantErr, err, "error should be equal")
			} else {
				assert.Nil(t, err, "error should be nil")
			}

			assert.Equal(t, test.want, users, "users should be equal")
		})
	}
}

func TestGetUserByField(t *testing.T) {
	type params struct {
		ctx   context.Context
		field string
		value string
	}

	tests := []struct {
		name        string
		params      params
		beforeTests func(params params, mock sqlmock.Sqlmock)
		want        *entity.User
		wantErr     error
	}{
		{
			name: "when fetching user by id field, it should return user",
			params: params{
				ctx:   context.TODO(),
				field: "id",
				value: userFixture.ActiveUser1.ID.String(),
			},
			beforeTests: func(params params, mock sqlmock.Sqlmock) {
				expectedRows := sqlmock.NewRows(userFixture.Rows).
					AddRows(helpers.ConvertToDriverValue(helpers.StructToSlice(userFixture.ActiveUser1)))

				expectedSQL := `SELECT \* FROM users WHERE id = \$1 AND deleted_at IS NULL`

				mock.ExpectQuery(expectedSQL).
					WithArgs(params.value).
					WillReturnRows(expectedRows)
			},
			want:    &userFixture.ActiveUser1,
			wantErr: nil,
		},
		{
			name: "when fetching user by email field, it should return user",
			params: params{
				ctx:   context.TODO(),
				field: "email",
				value: userFixture.ActiveUser1.Email,
			},
			beforeTests: func(params params, mock sqlmock.Sqlmock) {
				expectedRows := sqlmock.NewRows(userFixture.Rows).
					AddRows(helpers.ConvertToDriverValue(helpers.StructToSlice(userFixture.ActiveUser1)))

				expectedSQL := `SELECT \* FROM users WHERE email = \$1 AND deleted_at IS NULL`

				mock.ExpectQuery(expectedSQL).
					WithArgs(params.value).
					WillReturnRows(expectedRows)
			},
			want:    &userFixture.ActiveUser1,
			wantErr: nil,
		},
		{
			name: "when fetching non existent user, it should return no rows error",
			params: params{
				ctx:   context.TODO(),
				field: "id",
				value: uuid.New().String(),
			},
			beforeTests: func(params params, mock sqlmock.Sqlmock) {
				expectedSQL := `SELECT \* FROM users WHERE id = \$1 AND deleted_at IS NULL`

				mock.ExpectQuery(expectedSQL).
					WithArgs(params.value).
					WillReturnError(sql.ErrNoRows)
			},
			want:    nil,
			wantErr: sql.ErrNoRows,
		},
		{
			name: "when fetching user with invalid id, it should return no rows error",
			params: params{
				ctx:   context.TODO(),
				field: "id",
				value: "invalid",
			},
			beforeTests: func(params params, mock sqlmock.Sqlmock) {
				expectedSQL := `SELECT \* FROM users WHERE id = \$1 AND deleted_at IS NULL`

				mock.ExpectQuery(expectedSQL).
					WithArgs(params.value).
					WillReturnError(sql.ErrNoRows)
			},
			want:    nil,
			wantErr: sql.ErrNoRows,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, close, mock := psqlConnMock.NewMockDB(t)
			defer close()

			if test.beforeTests != nil {
				test.beforeTests(test.params, mock)
			}

			repo := userRepo.NewUserRepository(db)
			users, err := repo.GetUserByField(test.params.ctx, test.params.field, test.params.value)

			if test.wantErr != nil {
				assert.Equal(t, test.wantErr, err, "error should be equal")
			} else {
				assert.Nil(t, err, "error should be nil")
			}

			assert.Equal(t, test.want, users, "user should be equal")
		})
	}
}

func TestCreateUser(t *testing.T) {
	type params struct {
		ctx  context.Context
		user *entity.User
	}

	tests := []struct {
		name        string
		params      params
		beforeTests func(params params, mock sqlmock.Sqlmock)
		want        uuid.UUID
		wantErr     error
	}{
		{
			name: "when creating user, it should return user id",
			params: params{
				ctx:  context.TODO(),
				user: &userFixture.ActiveUser1,
			},
			beforeTests: func(params params, mock sqlmock.Sqlmock) {
				expectedSQL := `INSERT INTO users (.+) VALUES (.+)`

				mock.ExpectExec(expectedSQL).
					WithArgs(params.user.ID, params.user.Name, params.user.Password, params.user.Email).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want:    userFixture.ActiveUser1.ID,
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, close, mock := psqlConnMock.NewMockDB(t)
			defer close()

			if test.beforeTests != nil {
				test.beforeTests(test.params, mock)
			}

			repo := userRepo.NewUserRepository(db)
			id, err := repo.CreateUser(test.params.ctx, test.params.user)

			if test.wantErr != nil {
				assert.Equal(t, test.wantErr, err, "error should be equal")
			} else {
				assert.Nil(t, err, "error should be nil")
			}

			assert.Equal(t, test.want, id, "user id should be equal")
		})
	}
}

func TestUpdateUser(t *testing.T) {
	type params struct {
		ctx  context.Context
		user *entity.User
	}

	tests := []struct {
		name        string
		params      params
		beforeTests func(params params, mock sqlmock.Sqlmock)
		want        uuid.UUID
		wantErr     error
	}{
		{
			name: "when updating user, it should return user id",
			params: params{
				ctx:  context.TODO(),
				user: &userFixture.ActiveUser1,
			},
			beforeTests: func(params params, mock sqlmock.Sqlmock) {
				expectedSQL := `UPDATE users SET .+ WHERE id = \?`

				mock.ExpectExec(expectedSQL).
					WithArgs(params.user.Name, params.user.Password, params.user.Email, params.user.ID).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want:    userFixture.ActiveUser1.ID,
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, close, mock := psqlConnMock.NewMockDB(t)
			defer close()

			if test.beforeTests != nil {
				test.beforeTests(test.params, mock)
			}

			repo := userRepo.NewUserRepository(db)
			id, err := repo.UpdateUser(test.params.ctx, test.params.user)

			if test.wantErr != nil {
				assert.Equal(t, test.wantErr, err, "error should be equal")
			} else {
				assert.Nil(t, err, "error should be nil")
			}

			assert.Equal(t, test.want, id, "user id should be equal")
		})
	}
}

func TestSoftDeleteUser(t *testing.T) {
	type params struct {
		ctx context.Context
		id  uuid.UUID
	}

	tests := []struct {
		name        string
		params      params
		beforeTests func(params params, mock sqlmock.Sqlmock)
		want        uuid.UUID
		wantErr     error
	}{
		{
			name: "when soft deleting user, it should return user id",
			params: params{
				ctx: context.TODO(),
				id:  userFixture.ActiveUser1.ID,
			},
			beforeTests: func(params params, mock sqlmock.Sqlmock) {
				expectedSQL := `UPDATE users SET deleted_at = NOW\(\) WHERE id = \$1`

				mock.ExpectExec(expectedSQL).
					WithArgs(params.id).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want:    userFixture.ActiveUser1.ID,
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, close, mock := psqlConnMock.NewMockDB(t)
			defer close()

			if test.beforeTests != nil {
				test.beforeTests(test.params, mock)
			}

			repo := userRepo.NewUserRepository(db)
			id, err := repo.SoftDeleteUser(test.params.ctx, test.params.id)

			if test.wantErr != nil {
				assert.Equal(t, test.wantErr, err, "error should be equal")
			} else {
				assert.Nil(t, err, "error should be nil")
			}

			assert.Equal(t, test.want, id, "user id should be equal")
		})
	}
}

func TestDeleteUser(t *testing.T) {
	type params struct {
		ctx context.Context
		id  uuid.UUID
	}

	tests := []struct {
		name        string
		params      params
		beforeTests func(params params, mock sqlmock.Sqlmock)
		want        uuid.UUID
		wantErr     error
	}{
		{
			name: "when deleting user, it should return user id",
			params: params{
				ctx: context.TODO(),
				id:  userFixture.ActiveUser1.ID,
			},
			beforeTests: func(params params, mock sqlmock.Sqlmock) {
				expectedSQL := `DELETE FROM users WHERE id = \$1`

				mock.ExpectExec(expectedSQL).
					WithArgs(params.id).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want:    userFixture.ActiveUser1.ID,
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, close, mock := psqlConnMock.NewMockDB(t)
			defer close()

			if test.beforeTests != nil {
				test.beforeTests(test.params, mock)
			}

			repo := userRepo.NewUserRepository(db)
			id, err := repo.DeleteUser(test.params.ctx, test.params.id)

			if test.wantErr != nil {
				assert.Equal(t, test.wantErr, err, "error should be equal")
			} else {
				assert.Nil(t, err, "error should be nil")
			}

			assert.Equal(t, test.want, id, "user id should be equal")
		})
	}
}

func TestRestoreUser(t *testing.T) {
	type params struct {
		ctx context.Context
		id  uuid.UUID
	}

	tests := []struct {
		name        string
		params      params
		beforeTests func(params params, mock sqlmock.Sqlmock)
		want        uuid.UUID
		wantErr     error
	}{
		{
			name: "when restoring user, it should return user id",
			params: params{
				ctx: context.TODO(),
				id:  userFixture.InactiveUser1.ID,
			},
			beforeTests: func(params params, mock sqlmock.Sqlmock) {
				expectedSQL := `UPDATE users SET deleted_at = NULL WHERE id = \$1`

				mock.ExpectExec(expectedSQL).
					WithArgs(params.id).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want:    userFixture.InactiveUser1.ID,
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, close, mock := psqlConnMock.NewMockDB(t)
			defer close()

			if test.beforeTests != nil {
				test.beforeTests(test.params, mock)
			}

			repo := userRepo.NewUserRepository(db)
			id, err := repo.RestoreUser(test.params.ctx, test.params.id)

			if test.wantErr != nil {
				assert.Equal(t, test.wantErr, err, "error should be equal")
			} else {
				assert.Nil(t, err, "error should be nil")
			}

			assert.Equal(t, test.want, id, "user id should be equal")
		})
	}
}

func TestCountUsers(t *testing.T) {
	type params struct {
		ctx   context.Context
		query dto.GetUsersStatsQuery
	}

	tests := []struct {
		name        string
		params      params
		beforeTests func(params params, mock sqlmock.Sqlmock)
		want        int64
		wantErr     error
	}{
		{
			name: "when counting users, it should return total users",
			params: params{
				ctx: context.TODO(),
				query: dto.GetUsersStatsQuery{
					IncludeDeleted: false,
				},
			},
			beforeTests: func(params params, mock sqlmock.Sqlmock) {
				expectedRows := sqlmock.NewRows([]string{"count"}).
					AddRow(2)

				expectedSQL := `SELECT COUNT\(\*\) FROM users WHERE 1=1 AND deleted_at IS NULL`

				mock.ExpectQuery(expectedSQL).
					WillReturnRows(expectedRows)
			},
			want:    2,
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, close, mock := psqlConnMock.NewMockDB(t)
			defer close()

			if test.beforeTests != nil {
				test.beforeTests(test.params, mock)
			}

			repo := userRepo.NewUserRepository(db)
			total, err := repo.CountUsers(test.params.ctx, test.params.query)

			if test.wantErr != nil {
				assert.Equal(t, test.wantErr, err, "error should be equal")
			} else {
				assert.Nil(t, err, "error should be nil")
			}

			assert.Equal(t, test.want, total, "total users should be equal")
		})
	}
}
