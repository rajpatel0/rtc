package showcase

import (
	"fmt"

	tuple "github.com/rajpatel0/rtc/tools"
)

type projectile struct {
	position tuple.Tuple //a point
	velocity tuple.Tuple //a vector
}

type environment struct {
	gravity tuple.Tuple //a vector
	wind    tuple.Tuple //a vector
}

func tick(env environment, proj projectile) projectile {
	pos := tuple.AddTuple(proj.position, proj.velocity)
	vel := tuple.AddTuple(tuple.AddTuple(proj.velocity, env.gravity), env.wind)
	return projectile{pos, vel}
}

func Cannon() {

	p := projectile{tuple.Point(0, 1, 0), tuple.MultiplyScaler(tuple.Normalize(tuple.Vector(1, 1, 0)), 12.5)} //change the float at the end to change how far cannon launches
	e := environment{tuple.Vector(0, -0.1, 0), tuple.Vector(-0.01, 0, 0)}

	for i := 0; p.position.Y >= 0; i++ {
		fmt.Printf("Position after tick %d: x: %6.2f , y: %6.2f, z: %6.2f \n", i, p.position.X, p.position.Y, p.position.Z)
		p = tick(e, p)
	}
}
