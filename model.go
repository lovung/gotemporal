package gotemporal

type TemporalModel interface {
	IDer
	TIDer

	Clean()
}

type TIDer interface {
	GetTID() interface{}
}

type IDer interface {
	GetID() interface{}
}
