package todoist

import "strings"

type Section struct {
	HaveID
	HaveProjectID
	Collapsed    bool   `json:"collapsed"`
	Name         string `json:"name"`
	IsArchived   bool   `json:"is_archived"`
	IsDeleted    bool   `json:"is_deleted"`
	SectionOrder int    `json:"section_order"`
}

type Sections []Section

// GetIDByNameAndProject finds a section by name within a specific project.
// Returns empty string if not found.
func (a Sections) GetIDByNameAndProject(name string, projectID string) string {
	name = strings.TrimSpace(name)
	for _, sec := range a {
		if sec.ProjectID == projectID && strings.EqualFold(sec.Name, name) {
			return sec.GetID()
		}
	}
	return ""
}

// GetIDByName finds a section by name (across all projects).
// Returns empty string if not found. Use GetIDByNameAndProject for precision.
func (a Sections) GetIDByName(name string) string {
	name = strings.TrimSpace(name)
	for _, sec := range a {
		if strings.EqualFold(sec.Name, name) {
			return sec.GetID()
		}
	}
	return ""
}
