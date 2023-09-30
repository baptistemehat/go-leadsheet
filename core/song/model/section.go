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

func (section *Section) TrimEmptyLines() Section {
	trimedSection := NewSection()
	trimedSection.SetName(section.Name)

	isPreviousLineEmpty := false
	isFirstNonEmptyLine := true

	for _, line := range section.Lines {

		if !line.IsEmpty() {
			if isPreviousLineEmpty && !isFirstNonEmptyLine {
				trimedSection.AddLine(NewLine())
			}
			trimedSection.AddLine(line)
			isFirstNonEmptyLine = false
		}

		isPreviousLineEmpty = line.IsEmpty()
	}
	return trimedSection
}

// Format
func (section *Section) Format(f Formatter) (string, error) {
	return f.FormatSection(section)
}
