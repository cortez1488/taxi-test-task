package models

type TaxiData struct {
	Name        string  `redis:"name"`
	AdmArea     string  `redis:"admArea"`
	District    string  `redis:"district"`
	Address     string  `redis:"address"`
	CarCapacity int     `redis:"carCapacity"`
	Mode        string  `redis:"mode"`
	GlobalId    int64   `json:"global_id" redis:"global_id"`
	CoordX      float32 `redis:"latitude"`
	CoordY      float32 `redis:"longitude"`
}
