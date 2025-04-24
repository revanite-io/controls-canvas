package main

import (
	"os"

	"github.com/revanite-io/sci/pkg/layer2"
	"gopkg.in/yaml.v3"
)

func writeOutputCatalog(catalog layer2.Catalog, path string) error {
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
		sharedCapabilities = append(sharedCapabilities, item.id)
		for _, threat := range item.capability.Threats {
			sharedThreats = append(sharedThreats, threat.Data.Id)
			for _, control := range threat.Controls {
				sharedControls = append(sharedControls, control.Data.Id)
			}
		}
	}

	outputCatalog = layer2.Catalog{
		SharedControls: []layer2.Mapping{
			{
				ReferenceId: "CCC",
				Identifiers: sharedControls,
			},
		},
		SharedThreats: []layer2.Mapping{
			{
				ReferenceId: "CCC",
				Identifiers: sharedThreats,
			},
		},
		SharedCapabilities: []layer2.Mapping{
			{
				ReferenceId: "CCC",
				Identifiers: sharedCapabilities,
			},
		},
	}
	return outputCatalog
}
