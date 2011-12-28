package components

import (
	. "mater"
	"box2d"
	"json"
	"os"
)

//wrapper around box2d.Transform
type Transform struct {
	*box2d.Transform
}

func (xf *Transform) Name () string {
	return "Transform"
}

//If the Transform is set and the owner has a body attached, we replace its transform
//else if the Transform is not set we try to take it from the owners body or create a new one
func (xf *Transform) Init (owner *Entity) {
	//check if we already have a transform
	if xf.Transform == nil {
		//we don't, see if our owner has a body attached
		if bcomp, ok := owner.Components["Body"]; ok {
			body := bcomp.(*Body)
			//take its transform
			xf.Transform = body.Transform()
		} else {
			//create a new one
			xf.Transform = new(box2d.Transform)
		}
	} else {
		//we do, check if our owner has a body attached
		if bcomp, ok := owner.Components["Body"]; ok {
			body := bcomp.(*Body)
			//set its transform
			body.SetTransform(&xf.Position, xf.Angle())
			//take the address of the bodies transform
			xf.Transform = body.Transform()
		}
	}
}

func (xf *Transform) Update (owner *Entity, dt float64) {
	
}

func (xf *Transform) Destroy (owner *Entity) {
	
}

func (xf *Transform) Marshal(owner *Entity) ([]byte, os.Error) {
	return json.Marshal(xf.Transform)
}

func (xf *Transform) Unmarshal(owner *Entity, data []byte) (os.Error) {
	return json.Unmarshal(data, xf.Transform)
}

func init() {
	RegisterComponent(&Transform{})
}