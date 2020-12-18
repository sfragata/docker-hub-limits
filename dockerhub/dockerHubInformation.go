package dockerhub

//Info struct
type Info struct {
	Repository   string
	Username     string
	Password     string
	token        string
	Verbose      bool
	authURL      string
	rateLimitURL string
}

type tokenResponse struct {
	Token string `json:"token"`
	// AccessToken string `json:"access_token"`
	// ExpiresIn   int    `json:"expires_in"`
	// IssuedAt    string `json:"issued_at"`
}

//RateLimitsInfo struct
type RateLimitsInfo struct {
	ImageName string
	Limit     int
	Remaining int
}

const tokenTemplateURL = "https://auth.docker.io/token?service=registry.docker.io&scope=repository:%s:pull"
const rateLimitsTemplateURL = "https://registry-1.docker.io/v2/%s/manifests/latest"
const bearerToken = "Bearer %s"
const httpTimeout = 5
