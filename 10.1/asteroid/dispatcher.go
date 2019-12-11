package asteroid

// NewDispatcher returns a pointer to a newly configured Dispatcher.
func NewDispatcher(m Map) *Dispatcher {
	return &Dispatcher{
		Map:      m,
		Scouts:   make([]Scout, 0),
		Messages: make(chan Message),
	}
}

// Dispatcher represents a person in charge of dispatching scouts to
// each asteroid and listening for their responses (the number of
// visible asteroids from the scout's perspective). As he receives
// responses, he determines which scout is able to see the MOST
// asteroids.
type Dispatcher struct {
	Map
	Scouts   []Scout
	Messages chan Message
}

// DispatchScout dispatches a scout to a given position on the Map and returns
// a pointer to that Scout.
func (d *Dispatcher) DispatchScout(pos Vector) *Scout {
	d.Scouts = append(d.Scouts, NewScout(d.Map.Copy(), pos, d.Messages))
	return &d.Scouts[len(d.Scouts)-1]
}

// Message is sent by the scouts back to the dispatcher via the Messages
// channel.
type Message struct {
	//
}
