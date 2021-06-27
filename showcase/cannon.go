package showcase

import (
	"fmt"

	tools "github.com/rajpatel0/rtc/tools"
)

type projectile struct {
	position tools.Tuple //a point
	velocity tools.Tuple //a vector
}

type environment struct {
	gravity tools.Tuple //a vector
	wind    tools.Tuple //a vector
}

func tick(env environment, proj projectile) projectile {
	pos := tools.AddTuple(proj.position, proj.velocity)
	vel := tools.AddTuple(tools.AddTuple(proj.velocity, env.gravity), env.wind)
	return projectile{pos, vel}
}

func Cannon() {
	c := tools.MakeCanvas(900, 550)

	p := projectile{tools.Point(0, 1, 0), tools.MultiplyScaler(tools.Normalize(tools.Vector(1, 1.8, 0)), 11.25)} //change the float at the end to change how far cannon launches
	e := environment{tools.Vector(0, -0.1, 0), tools.Vector(-0.01, 0, 0)}
	red := tools.Color(1, 0, 0)

	for i := 0; p.position.Y >= 0; i++ {
		fmt.Printf("Position after tick %d: x: %6.2f , y: %6.2f, z: %6.2f \n", i, p.position.X, p.position.Y, p.position.Z)
		y := int(float64(c.Height) - p.position.Y)
		x := int(p.position.X)
		if x < c.Width && y < c.Height {
			c = tools.WritePixel(c, x, y, red)
		}
		p = tick(e, p)
	}
	tools.PrintCanvas(tools.CanvasToPPM(c))
}
