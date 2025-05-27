package fabric

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Entity struct {
	Fabricable
}

func NewEntity() *Entity {
	entity := &Entity{}
	entity.Initialize(entity)
	return entity
}

var FabricEntityCmd = &cobra.Command{
	Use:   "entity [relative path] [lang] [entityName] [param1:type1:true] [param2:type2:false]... [paramN:typeN:optional]",
	Short: "Generate code based on entity and model definitions",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		entity := NewEntity()

		creatingErr := entity.CreateFromArgs(cmd, args)
		if creatingErr != nil {
			fmt.Printf("Error creating entity: %v\n", creatingErr)
			return
		}
		
		tmplName, err := GetEntityTemplate(entity.Language)
		if err != nil {
			fmt.Printf("Error getting entity template: %v\n", err)
			return
		}

		generatingErr := entity.Generate(tmplName, args[0], "")
		if generatingErr != nil {
			fmt.Printf("Error generating entity: %v\n", generatingErr)
			return
		}
	},
}
