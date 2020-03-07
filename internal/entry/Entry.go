package entry

// Entry type
type Entry (map[string]string)

func (e Entry) Copy() Entry {
	out := make(Entry)
	for k, v := range e {
		out[k] = v
	}
	return out
}
