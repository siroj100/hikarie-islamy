package ctxs

type ContextKey int

const (
	theCtx ContextKey = iota
	IPAddress
	UserAgent
	AuthorizationHeader
	Transaction
	//JwtUser
	//AppUser
)

const (
//EchoUserToken = "user"
//EchoDbAppUser = "appUser"
)
