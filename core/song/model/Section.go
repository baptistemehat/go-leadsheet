package model

// Section
type Section struct {
	Name  string `json:"name"`
	Lines []Line `json:"lines"`
}

// NewSection
func NewSection() Section {
	return Section{
		Name:  "",
		Lines: []Line{},
	}
}

// SetName
func (section *Section) SetName(name string) {
	section.Name = name
}

// AddLine
func (section *Section) AddLine(line Line) {
	section.Lines = append(section.Lines, line)
}

// Clear
func (section *Section) Clear() {
	section.Name = ""
	section.Lines = []Line{}
}

// Format
func (section *Section) Format(f Formatter) (string, error) {
	return f.FormatSection(section)
}
