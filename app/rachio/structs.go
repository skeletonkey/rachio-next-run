package rachio

type device struct {
	Name  string         `json:"name"`
	ID    string         `json:"device_id"`
	Hours map[string]int `json:"hours"`
}
type rachio struct {
	Enabled     bool     `json:"enabled"`
	BearerToken string   `json:"bearer_token"`
	Devices     []device `json:"devices"`
}
type deviceStateType struct {
	DeviceID string `json:"deviceId"`
	NextRun  string `json:"nextRun"`
}
type nextRun struct {
	State deviceStateType `json:"state"`
}

// NextScheduleData hold general information about the next scheduled Rachio run, what type of alter it represents,
// and if the alert should be acted on.
type NextScheduleData struct {
	DeviceName string
	HoursUntil int
	AlertType  string
	Alert      bool
}
