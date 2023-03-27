package rachio

type url struct {
	Public   string `json:"public"`
	Internal string `json:"internal"`
}
type device struct {
	Name  string         `json:"name"`
	Id    string         `json:"device_id"`
	Hours map[string]int `json:"hours"`
}
type rachio struct {
	Url         url      `json:"url"`
	BearerToken string   `json:"bearer_token"`
	Devices     []device `json:"devices"`
}
type deviceStateType struct {
	DeviceId string `json:"deviceId"`
	NextRun  string `json:"nextRun"`
}
type nextRun struct {
	State deviceStateType `json:"state"`
}
