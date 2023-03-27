package pushover

type token struct {
	Account     string `json:"account"`
	Application string `json:"application"`
}
type pushover struct {
	Url   string `json:"url"`
	Token token  `json:"token"`
}
