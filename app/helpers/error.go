package helpers

type ConstantError string

func (e ConstantError) Error() string {
	return string(e)
}

var (
	ErrInternalServer  = ConstantError("internal server error")
	ErrIdNotFound      = ConstantError("id not found")
	ErrBadRequest      = ConstantError("bad request")
	ErrForbiddenAccess = ConstantError("forbidden access")
	ErrDataNotFound    = ConstantError("data not found")
	ErrBindData        = ConstantError("bind data error")
	ErrReadData        = ConstantError("read data error")
	ErrLogin           = ConstantError("invalid username/password")
)
