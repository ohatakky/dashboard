package twitter

// ## get bearer  token
// curl -u '< ID >:<Secret>' \
//   --data 'grant_type=client_credentials' \
//   'https://api.twitter.com/oauth2/token'

// ## get tweets
// curl --request GET \
//   --url 'https://api.twitter.com/1.1/statuses/user_timeline.json?screen_name=<Account Name>' \
//   --header 'authorization: Bearer <Bearer Token>' \
//   --header 'content-type: application/json'

const (
	AuthorizationEndpoint = "https://api.twitter.com/oauth2/token"
	GrantType             = "client_credentials"
	Endpoint              = "https://api.twitter.com/1.1"
	TimelineFmt           = Endpoint + "/statuses/user_timeline.json?screen_name=%s"
)
