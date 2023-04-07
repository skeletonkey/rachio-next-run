package healthcheck

type visible struct {
	Dependencies bool `json:"dependencies"`
	Settings     bool `json:"settings"`
}
type healthcheck struct {
	Addr     string  `json:"address"`
	EndPoint string  `json:"end_point"`
	Visible  visible `json:"visible"`
}
