package asteroid

// NewScout returns a copy of a new Scout instance
func NewScout(m Map, origin Vector, messages chan<- Message) Scout {
	return Scout{
		Map:          m,
		VisibleMap:   make(Map),
		InvisibleMap: make(Map),
		Origin:       origin,
		Messages:     messages,
	}
}

// Scout represents a person who is sent to an asteroid with a COPY of the map
// (Map). He knows where he is on the map (his Origin). He also has the ability
// to send a message back to the dispatcher via the Messages channel.
type Scout struct {
	Map
	VisibleMap   Map
	InvisibleMap Map
	Origin       Vector
	Messages     chan<- Message
}

// Search searches the map for asteroids, MOVING "visible" asteroids to
// VisibleMap and "invisible" asteroids to InvisibleMap.
func (s *Scout) Search() {
	// Loop until the map is "empty" (except for the origin)
	for len(s.Map) > 1 {
		visible := Vector{}
		foundVisible := false

		for check := range s.Map {
			// Don't check the origin, duh!
			if check == s.Origin {
				continue
			}

			// Grab the asteroid to check against
			if foundVisible == false {
				visible = check
				foundVisible = true
				continue
			}

			// If check "blocks" the view to visible, then add visible to InvisibleMap,
			// remove visible from Map, and set visible to check.
			ab, ac := check.Sub(s.Origin), visible.Sub(s.Origin)
			if Colinear(ab, ac) && ab.Dot(ac) > 0 && ab.Dot(ac) < ab.Dot(ab) {
				s.InvisibleMap[visible] = s.Map[visible]
				delete(s.Map, visible)
				visible = check
			}
		}

		// Add visible to VisibleMap and delete it from Map.
		s.VisibleMap[visible] = s.Map[visible]
		delete(s.Map, visible)
	}
}
