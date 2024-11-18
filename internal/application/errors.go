package application

import "errors"

var ErrDecodingPayload = errors.New("error decoding payload")
var ErrEncodingPayload = errors.New("error encoding payload")
var ErrValidatingPayload = errors.New("error validating payload")

var ErrCreatingUser = errors.New("error creating user")
var ErrExistingUser = errors.New("error existing user")
var ErrGettingUserByUsername = errors.New("error getting user by username")
var ErrGettingUserById = errors.New("error getting user by id")
var ErrDeletingUserById = errors.New("error deleting user by id")
var ErrUpdatingUsernameById = errors.New("error updating user by id")
var ErrParsingUserIdParamToUUID = errors.New("error parsing user id parameter to uuid")

var ErrHashingPassword = errors.New("error hashing password")
var ErrComparingPasswords = errors.New("error comparing passwords")

var ErrCreatingAccessToken = errors.New("error creating access token")
var ErrCreatingRefreshToken = errors.New("error creating refresh token")

var ErrLoadingEnv = errors.New("error creating user")
var ErrOpeningDb = errors.New("error opening database")
var ErrListeningOnServer = errors.New("error listening on user")

var ErrCreatingMockDb = errors.New("error creating mock database")

var ErrValidatingBearerToken = errors.New("error validating bearer token")
var ErrValidatingJwt = errors.New("error validating jwt")
var ErrGettingJwtCookie = errors.New("error getting jwt cookie")
var ErrParsingJwtSubToUUID = errors.New("error parsing jwt subject to uuid")
