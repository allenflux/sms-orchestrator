package model

type Data struct {
	Type    string `json:"type"`
	DevId   string `json:"dev_id"`
	Content string `json:"content"`
	AdminId string `json:"admin_id"`
}
type DeviceReceiveCmd struct {
	Event string `json:"event"`
	Data  Data   `json:"data"`
}
