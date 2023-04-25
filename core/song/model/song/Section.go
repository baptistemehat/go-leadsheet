package song

type Section struct {
	Name  string `json:"name"`
	Lines []Line `json:"lines"`
}

func NewSection() Section {
	return Section{
		Name:  "",
		Lines: []Line{},
	}
}

func (section *Section) SetName(name string) {
	section.Name = name
}

func (section *Section) AddLine(line Line) {
	section.Lines = append(section.Lines, line)
}

func (section *Section) Clear() {
	section.Name = ""
	section.Lines = []Line{}
}

// func (section *Section) Format(f songFormatter.SongFormatter) string {
// 	return f.FormatSection(section)
// }
