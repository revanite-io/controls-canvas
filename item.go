package main

type item struct {
	id          string
	title       string
	description string
	capability  availableCapability
}

func (i item) Title() string       { return i.id + ": " + i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title }
