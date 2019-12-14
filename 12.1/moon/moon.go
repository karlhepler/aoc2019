package moon

type Moons []Moon

func (moons Moons) Simulate(steps int) {
	for step := 0; step < steps; step++ {
		moons.ApplyGravity()
		moons.Move()
	}
}

func (moons Moons) Pairs() [][2]*Moon {
	pairs := make([][2]*Moon, 0)

	count := len(moons)
	for i := range moons {
		for j := i + 1; j < count; j++ {
			pairs = append(pairs, [2]*Moon{&(moons[i]), &(moons[j])})
		}
	}

	return pairs
}

func (moons Moons) Energy() (energy int) {
	for _, moon := range moons {
		energy += moon.Energy()
	}
	return
}

func (moons Moons) ApplyGravity() {
	for _, pair := range moons.Pairs() {
		ApplyGravity(pair)
	}
}

func (moons Moons) Move() {
	for i := range moons {
		moons[i].Velocity = moons[i].Velocity.Add(moons[i].DeltaVelocity)
		moons[i].Position = moons[i].Position.Add(moons[i].Velocity)
		moons[i].DeltaVelocity = Vector{}
	}
}

func NewMoon(pos Vector) Moon {
	return Moon{Position: pos}
}

type Moon struct {
	Position      Vector
	Velocity      Vector
	DeltaVelocity Vector
}

func (moon Moon) PotentialEnergy() (energy int) {
	for _, p := range moon.Position {
		energy += abs(p)
	}
	return
}

func (moon Moon) KineticEnergy() (energy int) {
	for _, v := range moon.Velocity {
		energy += abs(v)
	}
	return
}

func (moon Moon) Energy() int {
	return moon.PotentialEnergy() * moon.KineticEnergy()
}

func ApplyGravity(moons [2]*Moon) {
	for i := 0; i < 3; i++ {
		if moons[0].Position[i] < moons[1].Position[i] {
			moons[0].DeltaVelocity[i] += 1
			moons[1].DeltaVelocity[i] -= 1
		} else if moons[0].Position[i] > moons[1].Position[i] {
			moons[0].DeltaVelocity[i] -= 1
			moons[1].DeltaVelocity[i] += 1
		}
	}
	return
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
