// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"

	db "github.com/parth-koshta/sparrow/db/sqlc"
	mock "github.com/stretchr/testify/mock"

	pgtype "github.com/jackc/pgx/v5/pgtype"
)

// Querier is an autogenerated mock type for the Querier type
type Querier struct {
	mock.Mock
}

// CreateDraft provides a mock function with given fields: ctx, arg
func (_m *Querier) CreateDraft(ctx context.Context, arg db.CreateDraftParams) (db.Draft, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateDraft")
	}

	var r0 db.Draft
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateDraftParams) (db.Draft, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateDraftParams) db.Draft); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Draft)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateDraftParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePostSuggestion provides a mock function with given fields: ctx, arg
func (_m *Querier) CreatePostSuggestion(ctx context.Context, arg db.CreatePostSuggestionParams) (db.Postsuggestion, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreatePostSuggestion")
	}

	var r0 db.Postsuggestion
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreatePostSuggestionParams) (db.Postsuggestion, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreatePostSuggestionParams) db.Postsuggestion); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Postsuggestion)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreatePostSuggestionParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePrompt provides a mock function with given fields: ctx, arg
func (_m *Querier) CreatePrompt(ctx context.Context, arg db.CreatePromptParams) (db.Prompt, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreatePrompt")
	}

	var r0 db.Prompt
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreatePromptParams) (db.Prompt, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreatePromptParams) db.Prompt); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Prompt)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreatePromptParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateScheduledPost provides a mock function with given fields: ctx, arg
func (_m *Querier) CreateScheduledPost(ctx context.Context, arg db.CreateScheduledPostParams) (db.Scheduledpost, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateScheduledPost")
	}

	var r0 db.Scheduledpost
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateScheduledPostParams) (db.Scheduledpost, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateScheduledPostParams) db.Scheduledpost); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Scheduledpost)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateScheduledPostParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSocialAccount provides a mock function with given fields: ctx, arg
func (_m *Querier) CreateSocialAccount(ctx context.Context, arg db.CreateSocialAccountParams) (db.Socialaccount, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateSocialAccount")
	}

	var r0 db.Socialaccount
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateSocialAccountParams) (db.Socialaccount, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateSocialAccountParams) db.Socialaccount); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Socialaccount)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateSocialAccountParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: ctx, arg
func (_m *Querier) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 db.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateUserParams) (db.User, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateUserParams) db.User); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateUserParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDraft provides a mock function with given fields: ctx, id
func (_m *Querier) DeleteDraft(ctx context.Context, id pgtype.UUID) (db.Draft, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDraft")
	}

	var r0 db.Draft
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) (db.Draft, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) db.Draft); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Draft)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgtype.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePostSuggestion provides a mock function with given fields: ctx, id
func (_m *Querier) DeletePostSuggestion(ctx context.Context, id pgtype.UUID) (db.Postsuggestion, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeletePostSuggestion")
	}

	var r0 db.Postsuggestion
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) (db.Postsuggestion, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) db.Postsuggestion); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Postsuggestion)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgtype.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePrompt provides a mock function with given fields: ctx, id
func (_m *Querier) DeletePrompt(ctx context.Context, id pgtype.UUID) (db.Prompt, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeletePrompt")
	}

	var r0 db.Prompt
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) (db.Prompt, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) db.Prompt); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Prompt)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgtype.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteScheduledPost provides a mock function with given fields: ctx, id
func (_m *Querier) DeleteScheduledPost(ctx context.Context, id pgtype.UUID) (db.Scheduledpost, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteScheduledPost")
	}

	var r0 db.Scheduledpost
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) (db.Scheduledpost, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) db.Scheduledpost); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Scheduledpost)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgtype.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSocialAccount provides a mock function with given fields: ctx, id
func (_m *Querier) DeleteSocialAccount(ctx context.Context, id pgtype.UUID) (db.Socialaccount, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSocialAccount")
	}

	var r0 db.Socialaccount
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) (db.Socialaccount, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) db.Socialaccount); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Socialaccount)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgtype.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDraftByID provides a mock function with given fields: ctx, id
func (_m *Querier) GetDraftByID(ctx context.Context, id pgtype.UUID) (db.Draft, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetDraftByID")
	}

	var r0 db.Draft
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) (db.Draft, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) db.Draft); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Draft)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgtype.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostSuggestionByID provides a mock function with given fields: ctx, id
func (_m *Querier) GetPostSuggestionByID(ctx context.Context, id pgtype.UUID) (db.Postsuggestion, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetPostSuggestionByID")
	}

	var r0 db.Postsuggestion
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) (db.Postsuggestion, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) db.Postsuggestion); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Postsuggestion)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgtype.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPromptByID provides a mock function with given fields: ctx, id
func (_m *Querier) GetPromptByID(ctx context.Context, id pgtype.UUID) (db.Prompt, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetPromptByID")
	}

	var r0 db.Prompt
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) (db.Prompt, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) db.Prompt); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Prompt)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgtype.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetScheduledPostByID provides a mock function with given fields: ctx, id
func (_m *Querier) GetScheduledPostByID(ctx context.Context, id pgtype.UUID) (db.Scheduledpost, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetScheduledPostByID")
	}

	var r0 db.Scheduledpost
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) (db.Scheduledpost, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) db.Scheduledpost); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Scheduledpost)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgtype.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSocialAccountByID provides a mock function with given fields: ctx, id
func (_m *Querier) GetSocialAccountByID(ctx context.Context, id pgtype.UUID) (db.Socialaccount, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetSocialAccountByID")
	}

	var r0 db.Socialaccount
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) (db.Socialaccount, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) db.Socialaccount); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Socialaccount)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgtype.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *Querier) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 db.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (db.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) db.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(db.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: ctx, id
func (_m *Querier) GetUserByID(ctx context.Context, id pgtype.UUID) (db.GetUserByIDRow, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 db.GetUserByIDRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) (db.GetUserByIDRow, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgtype.UUID) db.GetUserByIDRow); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.GetUserByIDRow)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgtype.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDraftsByUserID provides a mock function with given fields: ctx, arg
func (_m *Querier) ListDraftsByUserID(ctx context.Context, arg db.ListDraftsByUserIDParams) ([]db.Draft, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for ListDraftsByUserID")
	}

	var r0 []db.Draft
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.ListDraftsByUserIDParams) ([]db.Draft, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.ListDraftsByUserIDParams) []db.Draft); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Draft)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.ListDraftsByUserIDParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPostSuggestionsByPromptID provides a mock function with given fields: ctx, arg
func (_m *Querier) ListPostSuggestionsByPromptID(ctx context.Context, arg db.ListPostSuggestionsByPromptIDParams) ([]db.Postsuggestion, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for ListPostSuggestionsByPromptID")
	}

	var r0 []db.Postsuggestion
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.ListPostSuggestionsByPromptIDParams) ([]db.Postsuggestion, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.ListPostSuggestionsByPromptIDParams) []db.Postsuggestion); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Postsuggestion)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.ListPostSuggestionsByPromptIDParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPromptsByUserID provides a mock function with given fields: ctx, arg
func (_m *Querier) ListPromptsByUserID(ctx context.Context, arg db.ListPromptsByUserIDParams) ([]db.Prompt, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for ListPromptsByUserID")
	}

	var r0 []db.Prompt
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.ListPromptsByUserIDParams) ([]db.Prompt, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.ListPromptsByUserIDParams) []db.Prompt); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Prompt)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.ListPromptsByUserIDParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListScheduledPostsByUserID provides a mock function with given fields: ctx, arg
func (_m *Querier) ListScheduledPostsByUserID(ctx context.Context, arg db.ListScheduledPostsByUserIDParams) ([]db.Scheduledpost, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for ListScheduledPostsByUserID")
	}

	var r0 []db.Scheduledpost
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.ListScheduledPostsByUserIDParams) ([]db.Scheduledpost, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.ListScheduledPostsByUserIDParams) []db.Scheduledpost); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Scheduledpost)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.ListScheduledPostsByUserIDParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSocialAccountsByUserID provides a mock function with given fields: ctx, arg
func (_m *Querier) ListSocialAccountsByUserID(ctx context.Context, arg db.ListSocialAccountsByUserIDParams) ([]db.Socialaccount, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for ListSocialAccountsByUserID")
	}

	var r0 []db.Socialaccount
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.ListSocialAccountsByUserIDParams) ([]db.Socialaccount, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.ListSocialAccountsByUserIDParams) []db.Socialaccount); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Socialaccount)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.ListSocialAccountsByUserIDParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUsers provides a mock function with given fields: ctx, arg
func (_m *Querier) ListUsers(ctx context.Context, arg db.ListUsersParams) ([]db.User, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for ListUsers")
	}

	var r0 []db.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.ListUsersParams) ([]db.User, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.ListUsersParams) []db.User); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.ListUsersParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDraft provides a mock function with given fields: ctx, arg
func (_m *Querier) UpdateDraft(ctx context.Context, arg db.UpdateDraftParams) (db.Draft, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDraft")
	}

	var r0 db.Draft
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateDraftParams) (db.Draft, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateDraftParams) db.Draft); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Draft)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.UpdateDraftParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePostSuggestion provides a mock function with given fields: ctx, arg
func (_m *Querier) UpdatePostSuggestion(ctx context.Context, arg db.UpdatePostSuggestionParams) (db.Postsuggestion, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePostSuggestion")
	}

	var r0 db.Postsuggestion
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdatePostSuggestionParams) (db.Postsuggestion, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdatePostSuggestionParams) db.Postsuggestion); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Postsuggestion)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.UpdatePostSuggestionParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePrompt provides a mock function with given fields: ctx, arg
func (_m *Querier) UpdatePrompt(ctx context.Context, arg db.UpdatePromptParams) (db.Prompt, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePrompt")
	}

	var r0 db.Prompt
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdatePromptParams) (db.Prompt, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdatePromptParams) db.Prompt); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Prompt)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.UpdatePromptParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateScheduledPost provides a mock function with given fields: ctx, arg
func (_m *Querier) UpdateScheduledPost(ctx context.Context, arg db.UpdateScheduledPostParams) (db.Scheduledpost, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateScheduledPost")
	}

	var r0 db.Scheduledpost
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateScheduledPostParams) (db.Scheduledpost, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateScheduledPostParams) db.Scheduledpost); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Scheduledpost)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.UpdateScheduledPostParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSocialAccount provides a mock function with given fields: ctx, arg
func (_m *Querier) UpdateSocialAccount(ctx context.Context, arg db.UpdateSocialAccountParams) (db.Socialaccount, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSocialAccount")
	}

	var r0 db.Socialaccount
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateSocialAccountParams) (db.Socialaccount, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateSocialAccountParams) db.Socialaccount); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Socialaccount)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.UpdateSocialAccountParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewQuerier creates a new instance of Querier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQuerier(t interface {
	mock.TestingT
	Cleanup(func())
}) *Querier {
	mock := &Querier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
