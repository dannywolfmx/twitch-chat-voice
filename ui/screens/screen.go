package screens

type Screen interface {
	Render(gtx Context) Dimensions
}
