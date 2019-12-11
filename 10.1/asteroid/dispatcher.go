package asteroid

// NewDispatcher returns a pointer to a newly configured Dispatcher.
func NewDispatcher(m Map) *Dispatcher {
	return &Dispatcher{
		Map: m,
	}
}

// Dispatcher represents a person in charge of dispatching scouts to
// each asteroid and listening for their responses (the number of
// visible asteroids from the scout's perspective). As he receives
// responses, he determines which scout is able to see the MOST
// asteroids.
type Dispatcher struct {
	Map
}

func (d *Dispatcher) SurveyAndListen() <-chan Message {
	messages := make(chan Message)

	for _, ast := range d.Map {
		scout := NewScout(d.Map, ast.Pos, messages)
		go scout.SearchAndReport()
	}

	return messages
}

// Message is sent by the scouts back to the dispatcher via the Messages
// channel.
type Message struct {
	NumVisibleAsteroids int
}
