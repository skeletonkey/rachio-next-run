package pushover

type token struct {
	Account     string `json:"account"`
	Application string `json:"application"`
}
type pushover struct {
	Enabled bool   `json:"enabled"`
	URL     string `json:"url"`
	Token   token  `json:"token"`
}
