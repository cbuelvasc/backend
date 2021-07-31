package enums

const (
	BasePath = "/api/yellow/v1"

	SignIn = "/signin"
	SignUp = "/signup"

	GetUsers       = "/users"
	GetUserById    = "/users/:id"
	UpdateUserById = "/users/:id"
	DeleteUserById = "/users/:id"

	GetTweets       = "/tweets"
	CreateTweets    = "/tweets"
	GetTweetById    = "/tweets/:id"
	UpdateTweetById = "/tweets/:id"
	DeleteTweetById = "/tweets/:id/user/:userId"
)
