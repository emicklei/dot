package dot

// AttributesMap holds attribute=value pairs.
type AttributesMap struct {
	attributes map[string]interface{}
}

// Attr sets the value for an attribute (unless empty).
func (a AttributesMap) Attr(label string, value interface{}) {
	if len(label) == 0 || value == nil {
		return
	}
	if s, ok := value.(string); ok {
		if len(s) > 0 {
			a.attributes[label] = s
			return
		}
	}
	a.attributes[label] = value
}

// Value return the value added for this label.
func (a AttributesMap) Value(label string) interface{} {
	return a.attributes[label]
}
