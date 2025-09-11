package main

import (
	"os"
	"sort"

	"github.com/ossf/gemara/layer2"
	"gopkg.in/yaml.v3"
)

func writeOutputCatalog(path string) error {
	catalog := generateOutputCatalog()
	data, err := yaml.Marshal(catalog)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func generateOutputCatalog() (outputCatalog layer2.Catalog) {
	var sharedControls []string
	var sharedThreats []string
	var sharedCapabilities []string

	for _, item := range selectedCapabilities {
		sharedCapabilities = appendIfMissing(sharedCapabilities, item.id)
		for _, threat := range item.capability.Threats {
			sharedThreats = appendIfMissing(sharedThreats, threat.Data.Id)
			for _, control := range threat.Controls {
				sharedControls = appendIfMissing(sharedControls, control.Data.Id)
			}
		}
	}

	sort.Sort(sort.StringSlice(sharedControls))
	sort.Sort(sort.StringSlice(sharedThreats))
	sort.Sort(sort.StringSlice(sharedCapabilities))

	// Convert string slices to MappingEntry slices
	controlEntries := make([]layer2.MappingEntry, len(sharedControls))
	for i, controlId := range sharedControls {
		controlEntries[i] = layer2.MappingEntry{
			ReferenceId: controlId,
			Strength:    1, // Default strength
		}
	}

	threatEntries := make([]layer2.MappingEntry, len(sharedThreats))
	for i, threatId := range sharedThreats {
		threatEntries[i] = layer2.MappingEntry{
			ReferenceId: threatId,
			Strength:    1, // Default strength
		}
	}

	capabilityEntries := make([]layer2.MappingEntry, len(sharedCapabilities))
	for i, capabilityId := range sharedCapabilities {
		capabilityEntries[i] = layer2.MappingEntry{
			ReferenceId: capabilityId,
			Strength:    1, // Default strength
		}
	}

	outputCatalog = layer2.Catalog{
		Metadata: layer2.Metadata{
			Title: catalogName,
		},
		ImportedControls: []layer2.Mapping{
			{
				ReferenceId: "CCC",
				Entries:     controlEntries,
			},
		},
		ImportedThreats: []layer2.Mapping{
			{
				ReferenceId: "CCC",
				Entries:     threatEntries,
			},
		},
		ImportedCapabilities: []layer2.Mapping{
			{
				ReferenceId: "CCC",
				Entries:     capabilityEntries,
			},
		},
	}
	return outputCatalog
}

func appendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}
