package data

type CV struct {
	FullName    string `yaml:"full-name"`
	ContactInfo []struct {
		Label string `yaml:"label"`
		Icon  string `yaml:"icon"`
		Value string `yaml:"value"`
		URI   string `yaml:"uri"`
	} `yaml:"contact-info"`
	Summary     string `yaml:"summary"`
	Experiences []struct {
		Title       string   `yaml:"title"`
		Company     string   `yaml:"company"`
		Location    string   `yaml:"location"`
		Start       string   `yaml:"start"`
		End         string   `yaml:"end,omitempty"`
		Current     bool     `yaml:"current,omitempty"`
		Description []string `yaml:"description"`
		TechStack   []string `yaml:"tech-stack"`
	} `yaml:"experiences"`
	Educations []struct {
		Degree      string `yaml:"degree"`
		Institution string `yaml:"institution"`
		Location    string `yaml:"location"`
		Start       string `yaml:"start"`
		End         string `yaml:"end"`
	} `yaml:"educations"`
}
