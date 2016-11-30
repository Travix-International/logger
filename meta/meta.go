package meta

type Meta struct {
	Fields map[string]string
}

func (m *Meta) Set(key string, value string) *Meta {
	m.Fields[key] = value

	return m
}

func (m *Meta) Get(key string) string {
	return m.Fields[key]
}

func (m *Meta) Remove(key string) *Meta {
	delete(m.Fields, key)

	return m
}

func (m *Meta) GetFields() map[string]string {
	return m.Fields
}

func New() Meta {
	m := Meta{}
	m.Fields = make(map[string]string)

	return m
}
