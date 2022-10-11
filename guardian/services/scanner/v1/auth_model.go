package v1

type (
	GenerateUserAccessTokenRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	GenerateUserAccessTokenResponse struct {
		AccessToken string `json:"access_token"`
	}
)
