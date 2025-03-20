package handler_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/banking-challenge/api"
	"github.com/radityacandra/banking-challenge/internal/application/user-account/handler"
	"github.com/radityacandra/banking-challenge/internal/application/user-account/service"
	"github.com/radityacandra/banking-challenge/internal/application/user-account/types"
	mockService "github.com/radityacandra/banking-challenge/mocks/internal_/application/user-account/service"
	"github.com/radityacandra/banking-challenge/pkg/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestTransactionCreditPost(t *testing.T) {
	type fields struct {
		Service service.IService
	}

	type args struct {
		reqBody string
	}

	type expected struct {
		statusCode int
		body       interface{}
	}

	type test struct {
		name     string
		fields   fields
		args     args
		mock     func(test) test
		expected expected
	}

	tests := []test{
		{
			name: "should return error if failed to bind request",
			args: args{
				reqBody: `{"no_rekening": "2342342","nominal": "asdfwer"}`,
			},
			mock: func(tt test) test {
				return tt
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
				body: api.DefaultErrorResponse{
					Remarks: "code=400, message=Unmarshal type error: expected=int64, got=string, field=nominal, offset=46, internal=json: cannot unmarshal string into Go struct field TransactionRequest.nominal of type int64",
				},
			},
		},
		{
			name: "should return error if failed to perform validation",
			args: args{
				reqBody: `{"no_rekening": "2342342","nominal": -1}`,
			},
			mock: func(tt test) test {
				return tt
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
				body: api.DefaultErrorResponse{
					Remarks: "nominal must be greater than 0",
				},
			},
		},
		{
			name: "should return error if service return error",
			args: args{
				reqBody: `{"no_rekening": "2342342","nominal": 10000}`,
			},
			mock: func(tt test) test {
				mockService := mockService.NewMockIService(t)
				tt.fields.Service = mockService

				mockService.EXPECT().StoreCash(mock.Anything, types.TransactionInput{
					AccountNo: "2342342",
					Amount:    10000,
				}).Return(types.GetBalanceOutput{}, types.ErrUserAccountNotFound).Once()

				return tt
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
				body: api.DefaultErrorResponse{
					Remarks: "user account is not found",
				},
			},
		},
		{
			name: "should return success if service return no error",
			args: args{
				reqBody: `{"no_rekening": "2342342","nominal": 10000}`,
			},
			mock: func(tt test) test {
				mockService := mockService.NewMockIService(t)
				tt.fields.Service = mockService

				mockService.EXPECT().StoreCash(mock.Anything, types.TransactionInput{
					AccountNo: "2342342",
					Amount:    10000,
				}).Return(types.GetBalanceOutput{
					Balance:   50000,
					AccountNo: "2342342",
				}, nil).Once()

				return tt
			},
			expected: expected{
				statusCode: http.StatusOK,
				body: api.TransactionResponse{
					Saldo: 50000,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt = tt.mock(tt)
			e := echo.New()
			e.Validator = validator.NewValidator()

			req := httptest.NewRequest(http.MethodPost, "/tabung", strings.NewReader(tt.args.reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			handler := &handler.Handler{
				Logger:  zap.NewNop(),
				Service: tt.fields.Service,
			}

			err := handler.TransactionCreditPost(c)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected.statusCode, rec.Code)

			bytes, _ := io.ReadAll(rec.Result().Body)
			if rec.Code == http.StatusOK {
				var response api.TransactionResponse
				json.Unmarshal(bytes, &response)
				assert.Equal(t, tt.expected.body, response)
			} else {
				var response api.DefaultErrorResponse
				json.Unmarshal(bytes, &response)
				assert.Equal(t, tt.expected.body, response)
			}
		})
	}
}
