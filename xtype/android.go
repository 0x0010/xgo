package xtype

type Android struct {
	Person
	model string
}

func NewAndroid(name, model string) *Android {
	a := &Android{model: model}
	a.name = name
	return a
}
