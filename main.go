package main

import (
	c "github.com/samekigor/quill-cli/cmd/clicommands"
	"github.com/samekigor/quill-cli/internal/utils"
)

func main() {
	utils.InitEnviromentVariables()
	utils.InitLogger()
	c.Execute()

}
