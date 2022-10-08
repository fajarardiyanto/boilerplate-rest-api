package generate

import (
	"fmt"
	"github.com/fajarardiyanto/boilerplate-rest-api/config"
	"github.com/fajarardiyanto/boilerplate-rest-api/internal/dto"
	"github.com/fajarardiyanto/flt-go-utils/caller"
	"github.com/fajarardiyanto/flt-go-utils/flags"
	"github.com/spf13/cobra"
	"path/filepath"
	"strings"
)

var CmdConfig = &cobra.Command{
	Use:   "config",
	Short: "Generate Config",
	RunE:  Config,
}

func Config(cmd *cobra.Command, args []string) error {
	cfg := new(dto.Config)
	flags.ParseStruct(cfg)
	fmt.Println(cfg)

	call, ok := caller.GetCaller(2)
	if ok {
		basedir := strings.Split(filepath.Dir(call.File), "/")
		if len(basedir) > 0 {
			basedir = basedir[:len(basedir)-1]
		}

		pss := filepath.Join(strings.Join(basedir, "/"), "../config", "config.yaml")
		if bs, err := filepath.Abs(pss); err == nil {
			if err = flags.GenerateConfig(cfg, bs); err != nil {
				config.GetLogger().Error(err).Quit()
			}
		} else {
			return err
		}

		config.GetLogger().Info("Success Generate File [%s]", basedir)

	}

	return nil
}
