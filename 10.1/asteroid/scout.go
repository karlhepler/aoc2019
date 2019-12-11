package asteroid

// NewScout returns a copy of a new Scout instance
func NewScout(m Map, origin Vector, messages chan<- Message) *Scout {
	return &Scout{
		Map:      m,
		Origin:   origin,
		Messages: messages,
	}
}

// Scout represents a person who is sent to an asteroid with a COPY of the map
// (Map). He knows where he is on the map (his Origin). He also has the ability
// to send a message back to the dispatcher via the Messages channel.
type Scout struct {
	Map
	Origin   Vector
	Messages chan<- Message
}

// Search searches the map for asteroids, appending "visible" asteroids to
// VisibleMap. Once this is done, it reports its findings to the dispatcher via
// the Messages channel.
func (s Scout) SearchAndReport() {
	visible := 0

	for _, check := range s.Map {
		// Don't count the origin!
		if check.Pos == s.Origin {
			continue
		}

		// Is anything blocking the scout's view of check?
		// From s.Origin to check, is there anything in between?
		// If so, add it to the Visible map.

		isBlocked := false
		for _, blocker := range s.Map {
			if blocker.Pos.OnSegment(s.Origin, check.Pos) {
				isBlocked = true
				break
			}
		}

		if !isBlocked {
			visible++
		}
	}

	// Report this to the dispatcher!
	s.Messages <- Message{visible}
}
