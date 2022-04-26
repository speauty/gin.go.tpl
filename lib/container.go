package lib

import "gin.go.tpl/lib/inject"

type Container struct {
	graph inject.Graph
}

func (c Container) Register(obj ...*inject.Object) (err error) {
	if err = c.graph.Provide(obj...); err != nil {
		return err
	}
	if err = c.graph.Populate(); err != nil {
		return err
	}
	return nil
}
