package fabric

import (
	"fmt"

	"github.com/joel2santos/fabric/cmd/fabric/templates"
)

func GetEntityTemplate(name string) (string, error) {
	switch (name) {
	case "ts":
		return templates.TS_ENTITY_TEMPLATE, nil
	default:
		return "", fmt.Errorf("unsupported template type: %s", name)
	}
}

func GetModelTemplate(name string) (string, error) {
	switch (name) {
	case "ts":
		return templates.TS_MODEL_TEMPLATE, nil
	default:
		return "", fmt.Errorf("unsupported template type: %s", name)
	}
}

func GetFileType(name string) (string, error) {
	switch (name) {
	case "ts":
		return "ts", nil
	default:
		return "", fmt.Errorf("unsupported template type: %s", name)
	}
}
