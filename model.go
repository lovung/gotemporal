package gotemporal

type TemporalModel interface {
	IDer
	TIDer

	Clean()
}

type TIDer interface {
	GetTID() string
	SetTID(string)
}

type IDer interface {
	GetID() uint64
}
