package footsites

type SessionResponse struct {
	Data struct {
		CSRFToken string `json:"csrfToken"`
	} `json:"data"`
}

type footsites struct {
	// internal things
	Host string

	CSRFToken string
}
