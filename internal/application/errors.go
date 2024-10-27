package application

import "errors"

var ErrDecodingPayload = errors.New("error decoding payload")
var ErrEncodingPayload = errors.New("error encoding payload")
var ErrValidatingPayload = errors.New("error validating payload")

var ErrCreatingUser = errors.New("error creating user")
var ErrExistingUser = errors.New("error existing user")

var ErrLoadingEnv = errors.New("error creating user")
var ErrOpeningDb = errors.New("error opening database")
var ErrListeningOnServer = errors.New("error listening on user")

var ErrCreatingMockDb = errors.New("error creating mock database")
