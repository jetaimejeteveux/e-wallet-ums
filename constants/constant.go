package constants

const (
	SuccessMessage = "success"
	Token          = "token"
	RefreshToken   = "refresh_token"
)

const (
	ErrFailedParseRequest         = "failed to parse request"
	ErrFailedValidateRequest      = "failed to validate request"
	ErrFailedValidateToken        = "failed to validate token"
	ErrFailedRegisterUser         = "failed to register user"
	ErrFailedGetUserSession       = "failed to get user session"
	ErrFailedLoginUser            = "user failed to login"
	ErrFailedGetUser              = "failed to get user"
	ErrUserNotFound               = "user not found"
	ErrFailedGenerateToken        = "failed to generate token"
	ErrFailedGenerateRefreshToken = "failed to generate refresh token"
	ErrFailedInsertSession        = "failed to insert session token"
	ErrFailedLogout               = "user failed to logout"
	ErrUnauthorized               = "error unauthorized"
	ErrSessionNotFound            = "error session not found"
	ErrFailedUpdateToken          = "error update token in db"
	ErrFailedGetToken             = "failed to get token from context"
	ErrFailedParseToken           = "failed to parse token"
	ErrFailedRefreshToken         = "failed to refresh token"
)
