package unitysceneanalysis

import (
	"bufio"
	"io"
	"sort"
	"strings"
)

type Scene struct {
	Sections        []Section
	TotalMemory     uint64
	ObjectSummaries []ObjectSummary
}

type ObjectSummary struct {
	Name   string
	Memory uint64
	Count  int
}

type ObjectSummarySizeDescending []ObjectSummary

func (a ObjectSummarySizeDescending) Len() int           { return len(a) }
func (a ObjectSummarySizeDescending) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ObjectSummarySizeDescending) Less(i, j int) bool { return a[i].Memory > a[j].Memory }

func eatHeader(scanner *bufio.Scanner) {
	scanner.Scan()
	scanner.Scan()
}

func ReadScene(reader io.Reader) (*Scene, error) {
	scene := &Scene{}
	workingSummaries := map[string]*ObjectSummary{}

	scanner := bufio.NewScanner(reader)
	buf := make([]byte, 1024*1024*100)
	scanner.Buffer(buf, 1024*1024*100)
	eatHeader(scanner)

	var section *Section = nil
	curLine := 2
	for scanner.Scan() {
		line := scanner.Text()
		curLine++

		if strings.HasPrefix(line, "--- !u!") {
			if section != nil {
				scene.Sections = append(scene.Sections, *section)
			}
			section = &Section{}
			scanner.Scan()
			curLine++
			section.Name = scanner.Text()
			section.Name = section.Name[:len(section.Name)-1]

			if _, ok := workingSummaries[section.Name]; !ok {
				workingSummaries[section.Name] = &ObjectSummary{
					Name:   section.Name,
					Memory: 0,
					Count:  0,
				}
			}
			workingSummaries[section.Name].Count += 1
			section.StartingLine = curLine
			continue
		}

		if strings.HasPrefix(line, "  m_Script: ") {
			section.Info = "Script: " + strings.TrimPrefix(line, "  m_Script: ")
		}

		if strings.HasPrefix(line, "  m_SourcePrefab: ") {
			section.Info = "Source Prefab: " + strings.TrimPrefix(line, "  m_SourcePrefab: ")
		}

		section.Lines += 1
		section.Size += uint64(len(line))
		workingSummaries[section.Name].Memory += uint64(len(line))
		scene.TotalMemory += uint64(len(line))
	}

	if section != nil {
		scene.Sections = append(scene.Sections, *section)
	}

	for _, summary := range workingSummaries {
		scene.ObjectSummaries = append(scene.ObjectSummaries, *summary)
	}

	sort.Sort(SectionBySizeDescending(scene.Sections))
	sort.Sort(ObjectSummarySizeDescending(scene.ObjectSummaries))

	return scene, scanner.Err()
}
