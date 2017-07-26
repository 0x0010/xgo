package xtype

type Android struct {
	Person
	model string
}

func NewAndroid(name, model string) (a Android) {
	a.model = model
	a.name = name
	return
}
