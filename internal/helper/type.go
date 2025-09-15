package helper

type ctxKey string

const (
	CtxUserID  ctxKey = "user_id"
	CtxPhone   ctxKey = "phone"
	CtxIsAdmin ctxKey = "is_admin"
)

type M map[string]any
