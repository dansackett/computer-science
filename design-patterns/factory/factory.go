package factory

type Appliance interface {
	Start()
	Stop()
	GetName() string
}

type ApplianceFactory struct{}

type Blender struct {
	name   string
	active bool
}

func (b *Blender) Start() {
	b.active = true
}

func (b *Blender) Stop() {
	b.active = false
}

func (b *Blender) GetName() string {
	return b.name
}

type Microwave struct {
	name   string
	active bool
}

func (m *Microwave) Start() {
	m.active = true
}

func (m *Microwave) Stop() {
	m.active = false
}

func (m *Microwave) GetName() string {
	return m.name
}

func NewApplianceFactory() *ApplianceFactory {
	return &ApplianceFactory{}
}

func (f *ApplianceFactory) MakeBlender() *Blender {
	return &Blender{name: "Blender", active: false}
}

func (f *ApplianceFactory) MakeMicrowave() *Microwave {
	return &Microwave{name: "Microwave", active: false}
}
