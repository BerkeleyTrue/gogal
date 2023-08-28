package cmd

import (
	"fmt"
	"path/filepath"

	"berkeleytrue/gogal/config"
	"berkeleytrue/gogal/internal/infra"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var CmdName string = "gogal"

var rootCmd = &cobra.Command{
	Use:   CmdName,
	Short: CmdName + " is cli that creates a gallery of your images",
	Long: `A fast and flexible CLI tool to create a gallery of your images.
   Just point it at a directory of images and it will create a gallery for you.`,
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]

		if dir == "" {
			cmd.Help()
			return
		}

		dirpath, err := filepath.Abs(dir)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Serving images from " + dirpath)

    infra.GetDirectores(dirpath)

    app := fx.New(
			config.Module,
			fx.Invoke(func(cfg *config.Config) {
	      cfg.Directory = dirpath
	    }),
			infra.Module,
		)

		app.Run()
	},
}

func Execute() error {

	rootCmd.Long = `Start serving a gallery of your images.
  ` + CmdName + ` [dir] => starts serving a gallery of images in the 'dir' directory`

	err := rootCmd.Execute()

	if err != nil {
		return err
	}
	return nil
}
