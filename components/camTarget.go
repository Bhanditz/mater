package components

import (
	. "github.com/teomat/mater"
)

type CamTarget struct{
	Empty
}

func (ct *CamTarget) Name() string {
	return "CamTarget"
}

func (ct *CamTarget) Init(owner *Entity) {
	owner.Scene.Camera.Position = owner.Transform.Position
}

func (ct *CamTarget) Update(owner *Entity, dt float64) {
	owner.Scene.Camera.Position = owner.Transform.Position
}

func init() {
	RegisterComponent(&CamTarget{})
}
