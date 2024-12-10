package service

import "errors"

var (
	ErrUserIsNil              = errors.New("user is nil")
	ErrUserLoginIncorected    = errors.New("login user incorected")
	ErrUserPasswordIncorected = errors.New("password user incorected")
	ErrRegistrationUser       = errors.New("user not registration")
	ErrUserNotFound           = errors.New("user not found")
	ErrUserExists             = errors.New("user alresdy exists")
	ErrAuthenticationUser     = errors.New("user not authentication")
	ErrNoAccess               = errors.New("there is no access to the file")

	ErrTokenNotFound     = errors.New("token not found")
	ErrDocumentNotFound  = errors.New("document not found")
	ErrDocumentsNotFound = errors.New("documents not found")
	ErrLogOutUser        = errors.New("user not finish the session")
)
