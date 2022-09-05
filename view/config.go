package view

type ConfigView struct {
	OnConfigTap, OnStopTap, OnNextTap func()
	OnUserNameChange                  func(string)
	DefaultScreen                     int
}
