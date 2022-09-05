package view

type ConfigView struct {
	OnConfigTap, OnStopTap, OnNextTap func()
	DefaultScreen                     int
}
