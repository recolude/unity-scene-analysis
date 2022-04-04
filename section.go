package unitysceneanalysis

type Section struct {
	Name         string
	Size         uint64
	Lines        int
	StartingLine int
	Info         string
}

func (section Section) EndingLine() int {
	return section.StartingLine + section.Lines
}

type SectionBySize []Section

func (a SectionBySize) Len() int           { return len(a) }
func (a SectionBySize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SectionBySize) Less(i, j int) bool { return a[i].Size < a[j].Size }

type SectionBySizeDescending []Section

func (a SectionBySizeDescending) Len() int           { return len(a) }
func (a SectionBySizeDescending) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SectionBySizeDescending) Less(i, j int) bool { return a[i].Size > a[j].Size }
