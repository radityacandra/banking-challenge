package handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/banking-challenge/api"
	modelUserAccount "github.com/radityacandra/banking-challenge/internal/application/user-account/model"
	"github.com/radityacandra/banking-challenge/internal/application/user/handler"
	"github.com/radityacandra/banking-challenge/internal/application/user/model"
	"github.com/radityacandra/banking-challenge/internal/application/user/repository"
	"github.com/radityacandra/banking-challenge/internal/application/user/service"
	"github.com/radityacandra/banking-challenge/pkg/database"
	"github.com/radityacandra/banking-challenge/pkg/validator"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestIntegrationUserPost(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping TestIntegrationUserPost")
	}

	ctx := context.Background()
	db, _ := database.Init(ctx, "postgres://some_user:some_password@localhost:5432/some_database?sslmode=disable")
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := &handler.Handler{
		Service: service,
		Logger:  zap.NewNop(),
	}

	type args struct {
		reqBody api.UserPostRequest
	}

	type test struct {
		name           string
		args           args
		wantStatusCode int
	}

	tests := []test{
		{
			name: "should accept request and store to db",
			args: args{
				reqBody: api.UserPostRequest{
					Nama: "some name",
					Nik:  "1122334455667788",
					NoHp: "081123456789",
				},
			},
			wantStatusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				db.ExecContext(ctx, `
					DELETE FROM public.user_accounts 
					WHERE user_id IN (
						SELECT id FROM public.users WHERE name = $1 AND phone_number = $2 AND identity_number = $3
					)`, tt.args.reqBody.Nama, tt.args.reqBody.NoHp, tt.args.reqBody.Nik)

				db.ExecContext(ctx, "DELETE FROM public.users WHERE name = $1 AND phone_number = $2 AND identity_number = $3",
					tt.args.reqBody.Nama, tt.args.reqBody.NoHp, tt.args.reqBody.Nik)
			}()

			e := echo.New()
			e.Validator = validator.NewValidator()

			bytes, _ := json.Marshal(tt.args.reqBody)
			req := httptest.NewRequest(http.MethodPost, "/daftar", strings.NewReader(string(bytes)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := handler.UserPost(c)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantStatusCode, rec.Code)

			var user model.User
			var userAccount modelUserAccount.UserAccount
			row := db.QueryRowxContext(ctx, "SELECT * FROM public.users WHERE name = $1 AND phone_number = $2 AND identity_number = $3",
				tt.args.reqBody.Nama, tt.args.reqBody.NoHp, tt.args.reqBody.Nik)
			row.StructScan(&user)
			assert.NotEmpty(t, user.Id)
			assert.Equal(t, tt.args.reqBody.Nama, user.Name)
			assert.Equal(t, tt.args.reqBody.Nik, user.IdentityNo)
			assert.Equal(t, tt.args.reqBody.NoHp, user.PhoneNo)

			row = db.QueryRowxContext(ctx, "SELECT * FROM public.user_accounts WHERE user_id = $1", user.Id)
			row.StructScan(&userAccount)
			assert.NotEmpty(t, userAccount.Id)
			assert.NotEmpty(t, userAccount.AccountNo)
			assert.Equal(t, int64(0), userAccount.TotalBalance)
		})
	}
}
