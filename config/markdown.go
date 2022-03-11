package config

type Markdown struct {
	Content     string `yaml:"content"`
	ContentFile File   `yaml:"content_file"`
}

func (m Markdown) Parse() (string, error) {
	if len(m.Content) != 0 {
		return m.Content, nil
	}
	b, err := m.ContentFile.Content()
	return string(b), err
}
