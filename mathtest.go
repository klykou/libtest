package libtest

// if center is 0,0
type Init struct {
	Gstring string
}

// Get Eccentricity of Ellipse
func (e *Init) GetString() string {
	return "hello lib world " + e.Gstring
}
