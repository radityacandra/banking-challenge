package handler_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/radityacandra/banking-challenge/api"
	"github.com/radityacandra/banking-challenge/internal/application/user/handler"
	"github.com/radityacandra/banking-challenge/internal/application/user/service"
	"github.com/radityacandra/banking-challenge/internal/application/user/types"
	mockService "github.com/radityacandra/banking-challenge/mocks/internal_/application/user/service"
	"github.com/radityacandra/banking-challenge/pkg/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestUserPost(t *testing.T) {
	type fields struct {
		Service service.IService
	}

	type args struct {
		reqBodyStruct api.UserPostRequest
		reqBody       string
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
			name: "should return bad request if failed to bind req body",
			args: args{
				reqBody: `{"nama":{"nickName":"Will"},"nik":"2764920938679992","no_hp":"081123456789"}`,
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
				body: api.DefaultErrorResponse{
					Remarks: "code=400, message=Unmarshal type error: expected=string, got=object, field=nama, offset=9, internal=json: cannot unmarshal object into Go struct field UserPostRequest.nama of type string",
				},
			},
			mock: func(tt test) test {
				return tt
			},
		},
		{
			name: "should return bad request if failed to validate request",
			args: args{
				reqBodyStruct: api.UserPostRequest{
					Nama: "John Doe",
					Nik:  "123",
					NoHp: "081123456789",
				},
			},
			expected: expected{
				statusCode: http.StatusBadRequest,
				body: api.DefaultErrorResponse{
					Remarks: "nik must be greater than 16 length",
				},
			},
			mock: func(tt test) test {
				return tt
			},
		},
		{
			name: "should return 500 if service return unknown error",
			args: args{
				reqBodyStruct: api.UserPostRequest{
					Nama: "John Doe",
					Nik:  "1122334455667788",
					NoHp: "081123456789",
				},
			},
			expected: expected{
				statusCode: http.StatusInternalServerError,
				body: api.DefaultErrorResponse{
					Remarks: "unknown error",
				},
			},
			mock: func(tt test) test {
				service := mockService.NewMockIService(t)
				tt.fields.Service = service

				service.EXPECT().RegisterUser(mock.Anything, types.RegisterUserInput{
					Name:       tt.args.reqBodyStruct.Nama,
					PhoneNo:    tt.args.reqBodyStruct.NoHp,
					IdentityNo: tt.args.reqBodyStruct.Nik,
				}).Return(types.RegisterUserOutput{}, errors.New("some error")).Once()

				return tt
			},
		},
		{
			name: "should return status ok when service executed successfully",
			args: args{
				reqBodyStruct: api.UserPostRequest{
					Nama: "John Doe",
					Nik:  "1122334455667788",
					NoHp: "081123456789",
				},
			},
			expected: expected{
				statusCode: http.StatusOK,
				body: api.UserPostResponse{
					NoRekening: uuid.NewString(),
				},
			},
			mock: func(tt test) test {
				service := mockService.NewMockIService(t)
				tt.fields.Service = service

				service.EXPECT().RegisterUser(mock.Anything, types.RegisterUserInput{
					Name:       tt.args.reqBodyStruct.Nama,
					PhoneNo:    tt.args.reqBodyStruct.NoHp,
					IdentityNo: tt.args.reqBodyStruct.Nik,
				}).Return(types.RegisterUserOutput{
					UserId:        uuid.NewString(),
					BankAccountNo: tt.expected.body.(api.UserPostResponse).NoRekening,
				}, nil).Once()

				return tt
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt = tt.mock(tt)

			e := echo.New()
			e.Validator = validator.NewValidator()

			var reqBody string
			if tt.args.reqBody != "" {
				reqBody = tt.args.reqBody
			} else {
				bytes, _ := json.Marshal(tt.args.reqBodyStruct)
				reqBody = string(bytes)
			}

			req := httptest.NewRequest(http.MethodPost, "/daftar", strings.NewReader(reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			handler := &handler.Handler{
				Logger:  zap.NewNop(),
				Service: tt.fields.Service,
			}

			err := handler.UserPost(c)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected.statusCode, rec.Code)

			bytes, _ := io.ReadAll(rec.Result().Body)
			if tt.expected.statusCode != http.StatusOK {
				var response api.DefaultErrorResponse
				json.Unmarshal(bytes, &response)

				assert.Equal(t, tt.expected.body, response)
			} else {
				var response api.UserPostResponse
				json.Unmarshal(bytes, &response)
				assert.NotEmpty(t, response.NoRekening)
				assert.Equal(t, tt.expected.body, response)
			}
		})
	}
}
