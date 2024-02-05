package masker

type MaskerFunc func(i string) string

func (m *Masker) AddCustomMasker(key string, maskerFunc MaskerFunc) {
	m.customMaskerList.Store(key, maskerFunc)
}

func (m *Masker) findCustomMasker(key string) (MaskerFunc, bool) {
	res, ok := m.customMaskerList.Load(key)
	return res.(MaskerFunc), ok
}

// OTMaskerString one-time mask the string
func (m *Masker) OTMaskerString(in string, out func(i string) string) string {
	return out(in)
}

// MaskerString mask the string
func (m *Masker) MaskerString(key, in string) string {
	out, ok := m.findCustomMasker(key)
	if !ok {
		return ""
	}
	return m.OTMaskerString(in, out)
}

func (m *Masker) fixedOverlay(str string, start int, length int) (overlayed string) {
	r := []rune(str)
	l := len([]rune(r))
	if l == 0 {
		return ""
	}
	for i := start; i <= length+start; i++ {
		r[i] = []rune(m.mask)[0]
	}
	overlayed = string(r)
	return overlayed
}
