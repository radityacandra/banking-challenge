package handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/radityacandra/banking-challenge/api"
	"github.com/radityacandra/banking-challenge/internal/application/user-account/handler"
	userAccount "github.com/radityacandra/banking-challenge/internal/application/user-account/model"
	"github.com/radityacandra/banking-challenge/internal/application/user-account/repository"
	"github.com/radityacandra/banking-challenge/internal/application/user-account/service"
	"github.com/radityacandra/banking-challenge/internal/application/user/model"
	"github.com/radityacandra/banking-challenge/pkg/database"
	"github.com/radityacandra/banking-challenge/pkg/validator"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestIntegrationTransactionCreditPost(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test TestIntegrationTransactionCreditPost")
		return
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
		amountPerTransaction int64
		concurrentProc       int
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
	}{
		{
			name: "should run single transaction",
			args: args{
				amountPerTransaction: 10000,
				concurrentProc:       1,
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name: "should handle race condition",
			args: args{
				amountPerTransaction: 10000,
				concurrentProc:       50,
			},
			wantStatusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userId := uuid.NewString()
			user := model.NewUser(userId, "some name", "081123456789", "1122334455667770", userAccount.UserAccount{})

			db.NamedExecContext(ctx, `
				INSERT INTO public.users(id, name, phone_number, identity_number, created_at, created_by)
				VALUES(:id, :name, :phone_number, :identity_number, :created_at, :created_by)`, user)

			userAccountId := uuid.NewString()
			userAccountNo := uuid.NewString()
			userAccount := userAccount.NewUserAccount(userAccountId, userId, userAccountNo)

			db.NamedExecContext(ctx, `
				INSERT INTO public.user_accounts(id, user_id, account_number, total_balance, created_at, created_by)
				VALUES(:id, :user_id, :account_number, :total_balance, :created_at, :created_by)`, userAccount)

			defer func() {
				db.ExecContext(ctx, "DELETE FROM public.transactions WHERE user_account_id = $1", userAccountId)
				db.ExecContext(ctx, "DELETE FROM public.user_accounts WHERE id = $1", userAccountId)
				db.ExecContext(ctx, "DELETE FROM public.users WHERE id = $1", userId)
			}()

			e := echo.New()
			e.Validator = validator.NewValidator()
			var wg sync.WaitGroup
			for i := 0; i < tt.args.concurrentProc; i++ {
				wg.Add(1)
				go func() {
					reqBody := api.TransactionRequest{
						NoRekening: userAccountNo,
						Nominal:    tt.args.amountPerTransaction,
					}

					bytes, _ := json.Marshal(reqBody)
					req := httptest.NewRequest(http.MethodPost, "/tabung", strings.NewReader(string(bytes)))
					req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
					rec := httptest.NewRecorder()
					c := e.NewContext(req, rec)

					err := handler.TransactionCreditPost(c)
					assert.NoError(t, err)
					assert.Equal(t, tt.wantStatusCode, rec.Code)
					wg.Done()
				}()
			}

			wg.Wait()

			var totalBalance int64
			row := db.QueryRowContext(ctx, `SELECT total_balance FROM public.user_accounts WHERE account_number = $1`, userAccountNo)
			row.Scan(&totalBalance)

			assert.Equal(t, tt.args.amountPerTransaction*int64(tt.args.concurrentProc), totalBalance)
		})
	}
}
