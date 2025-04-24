package main

import (
	"fmt"
	"os"
	"slices"
	"sort"

	"github.com/charmbracelet/bubbles/list"
	"github.com/revanite-io/sci/pkg/layer2"
)

type availableCapability struct {
	Data    layer2.Capability
	Threats []availableThreat
}

type availableThreat struct {
	Data     layer2.Threat
	Controls []availableControl
}

type availableControl struct {
	Data              layer2.Control
	FamilyTitle       string
	FamilyDescription string
}

func loadData() (output []availableCapability) {
	var catalog layer2.Catalog
	err := catalog.LoadFiles([]string{
		"https://raw.githubusercontent.com/finos/common-cloud-controls/refs/heads/dev/common/controls.yaml",
		"https://raw.githubusercontent.com/finos/common-cloud-controls/refs/heads/dev/common/threats.yaml",
		"https://raw.githubusercontent.com/finos/common-cloud-controls/refs/heads/dev/common/capabilities.yaml",
	})
	if err != nil {
		fmt.Printf("Error loading catalog: %v\n", err)
		os.Exit(1)
	}

	for _, cap := range catalog.Capabilities {
		if cap.Id == "" || cap.Title == "" {
			continue
		}
		sortedCapability := availableCapability{
			Data: cap,
		}
		for _, threat := range catalog.Threats {
			if threat.Id == "" || threat.Title == "" || len(threat.Capabilities) == 0 {
				continue
			}
			for _, tc := range threat.Capabilities {
				if tc.ReferenceId != "CCC" {
					continue
				}
				for _, mappedCapabilityId := range tc.Identifiers {
					if cap.Id == mappedCapabilityId {
						sortedCapability.Threats = append(sortedCapability.Threats, availableThreat{
							Data: threat,
						})
					}
				}
			}
		}
		for _, family := range catalog.ControlFamilies {
			for _, control := range family.Controls {
				if control.Id == "" {
					continue
				}
				for _, ct := range control.ThreatMappings {
					if ct.ReferenceId != "CCC" {
						continue
					}
					for _, threatId := range ct.Identifiers {
						for i, threat := range sortedCapability.Threats {
							if threat.Data.Id == threatId {
								sortedCapability.Threats[i].Controls = append(threat.Controls, availableControl{
									Data:              control,
									FamilyTitle:       family.Title,
									FamilyDescription: family.Description,
								})
							}
						}
					}
				}
			}
		}
		output = append(output, sortedCapability)
	}

	return output
}

func loadChoices() (choices []list.Item) {
	data := loadData()

	for _, capability := range data {
		var threatList []string
		var controlList []string
		for _, threat := range capability.Threats {
			threatList = append(threatList, threat.Data.Id)
			for _, control := range threat.Controls {
				if !slices.Contains(controlList, control.Data.Id) {
					controlList = append(controlList, control.Data.Id)
				}
			}
		}

		description := fmt.Sprintf("Threats: %v | Controls: %v", len(threatList), len(controlList))

		choice := item{
			id:          capability.Data.Id,
			title:       capability.Data.Title,
			capability:  capability,
			description: description,
		}
		choices = append(choices, choice)
	}

	// Sort by title
	sort.Slice(choices, func(i, j int) bool {
		return choices[i].(item).capability.Data.Id < choices[j].(item).capability.Data.Id
	})

	return choices
}
