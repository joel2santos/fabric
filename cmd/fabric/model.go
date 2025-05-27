
package fabric

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Model struct {
	Fabricable
}

func NewModel() *Model {
	model := &Model{}
	model.Initialize(model)
	return model
}

var FabricModelCmd = &cobra.Command{
	Use:   "model [relative path] [lang] [modelName] [param1:type1:true] [param2:type2:false]... [paramN:typeN:optional]",
	Short: "Generate code based on model and model definitions",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		model := NewModel()

		creatingErr := model.CreateFromArgs(cmd, args)
		if creatingErr != nil {
			fmt.Printf("Error creating model: %v\n", creatingErr)
			return
		}
		
		tmplName, err := GetModelTemplate(model.Language)
		if err != nil {
			fmt.Printf("Error getting model template: %v\n", err)
			return
		}

		generatingErr := model.Generate(tmplName, args[0], "_model")
		if generatingErr != nil {
			fmt.Printf("Error generating model: %v\n", generatingErr)
			return
		}
	},
}
