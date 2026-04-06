package examples

import "time"

type ServerTestData struct {
	NeedUpdate      string         `json:"need_update,omitempty"`
	WebServerState  string         `json:"web_server_state,omitempty"`
	OperatingSystem string         `json:"operating_system,omitempty"`
	Uptime          *time.Duration `json:"uptime,omitempty"`
	Testing         string         `json:"testing,omitempty"`
	Server          string         `json:"server,omitempty"`
	LinkUp          bool           `json:"link_up,omitempty"`
	PingMS          *int           `json:"ping_ms,omitempty"`
	SSDUsedPercent  *int           `json:"ssd_used_percent,omitempty"`
	RAMUsedPercent  *int           `json:"ram_used_percent,omitempty"`
}
