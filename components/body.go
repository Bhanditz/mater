package components

import (
	"encoding/json"
	"errors"
	"github.com/teomat/mater/collision"
	"github.com/teomat/mater/engine"
	"log"
)

type Body struct {
	Empty
	*collision.Body
}

func (body *Body) Name() string {
	return "Body"
}

func (body *Body) Init(owner *engine.Entity) {
	if body.Body == nil {
		log.Printf("Error: Body component is not initialized correctly!")
		return
	}

	if owner.Transform != nil {
		body.Transform = *owner.Transform
	}

	owner.Transform = &body.Transform

	body.Body.UserData = owner
}

func (body *Body) Destroy(owner *engine.Entity) {
	if body.Body == nil {
		log.Printf("Error: Body component is not initialized correctly!")
		return
	}
	body.Body.Enabled = false
	owner.Scene.Space.RemoveBody(body.Body)
	body.Body.UserData = nil
}

func (body *Body) Marshal(owner *engine.Entity) ([]byte, error) {
	return json.Marshal(body.Body)
}

func (body *Body) Unmarshal(owner *engine.Entity, data []byte) error {
	if body.Body == nil {
		body.Body = collision.NewBody(collision.BodyType_Static)
	}
	err := json.Unmarshal(data, body.Body)

	if err != nil {
		return err
	}

	if body.Body == nil {
		return errors.New("nil Body")
	}

	owner.Scene.Space.AddBody(body.Body)

	return nil
}

func init() {
	engine.RegisterComponent(&Body{})
}
