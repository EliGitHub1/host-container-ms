package model

type HostRec struct {
	ID        int    `json:"id"`
	Uuid      string `json:"uuid"`
	Name      string `json:"artist"`
	IpAddress string `json:"ip_address"`
}
