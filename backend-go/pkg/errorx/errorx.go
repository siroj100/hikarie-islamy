package errorx

import "errors"

var (
	ErrMissingParam    = errors.New("missing param")
	ErrInvalidParam    = errors.New("invalid param")
	ErrNotFound        = errors.New("not found")
	ErrInvalidResponse = errors.New("invalid response")
	ErrNoDbTx          = errors.New("no db transaction")

	ErrMultipleRows         = errors.New("multiple_rows")
	ErrNoRowsUpdated        = errors.New("no rows updated")
	ErrInvalidLevelID       = errors.New("invalid level id")
	ErrInvalidUserID        = errors.New("invalid user id")
	ErrInvalidHafalanBucket = errors.New("invalid hafalan bucket")
	ErrNotEnoughInventory   = errors.New("not enough inventory")
	ErrEmptyWallet          = errors.New("empty wallet")
	ErrNotEnoughBalance     = errors.New("not enough balance")
	ErrInactiveRecipient    = errors.New("inactive recipient")

	ErrServerError       = errors.New("server error")
	ErrForbidden         = errors.New("forbidden")
	ErrUnverifiedUser    = errors.New("unverified user")
	ErrInvalidStatusUser = errors.New("invalid status user")
	ErrUnauthenticated   = errors.New("unauthenticated")
)
