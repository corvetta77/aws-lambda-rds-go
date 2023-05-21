package model

type Sensor struct {
	Id           int     `json:"id,omitempty"`
	Type         *string `json:"type,omitempty"`
	Value        int     `json:"value,omitempty"`
	State        int     `json:"state,omitempty"`
	Trend        int     `json:"trend,omitempty"`
	ElapsedTimeS int     `json:"elapseTimeS,omitempty"`
}

type ResponseData struct {
	MultiSensor *MultiSensor `json:"multiSensor,omitempty"`
}

type MultiSensor struct {
	Sensors []Sensor `json:"sensors,omitempty"`
}
