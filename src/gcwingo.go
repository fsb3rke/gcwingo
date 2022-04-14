package gcwingo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aiocat/cystem"
)

type Integrations struct {
	CommandsFileName string
}

func main() {

	cliCommand := os.Args[1]
	commandName := os.Args[2]

	integrationsFileContentBytes, err := ioutil.ReadFile("gcwinGO.integrations.json")
	if err != nil {
		log.Fatal(err)
	}

	var integrations Integrations
	json.Unmarshal(integrationsFileContentBytes, &integrations)

	byt, err := ioutil.ReadFile(integrations.CommandsFileName)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	if cliCommand == "use" {
		for key, value := range dat {
			if key == commandName {
				for _, v := range value.([]interface{}) {
					cystem.RunString(v.(string))
				}
			}
		}
	}

	if cliCommand == "show" {
		if commandName == "all" {
			for key, value := range dat {
				fmt.Println(key)
				for _, v := range value.([]interface{}) {
					fmt.Println(v)
				}
				fmt.Println("")
			}
		} else {
			for key, value := range dat {
				if key == commandName {
					for _, v := range value.([]interface{}) {
						fmt.Println(v)
					}
				}
			}
		}
	}

	if cliCommand == "edit" {
		cystem.RunString("notepad " + integrations.CommandsFileName)
	}

	if cliCommand == "help" {
		fmt.Printf("\nuse  -> starts the command\nshow -> shows command inside\nedit -> opens the notepad\n")
	}
}
