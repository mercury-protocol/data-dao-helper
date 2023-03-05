package pkg

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mercury-protocol/data-dao-helper/internal"
)

// Retrieves a file from the filecoin network
func Retrieve(cid string, filename string) error {
	command := &exec.Cmd{
		Path:   internal.DaoConfig.LotusPath,
		Args:   []string{internal.DaoConfig.LotusPath, "client", "retrieve", cid, fmt.Sprintf("%s/%s", internal.DaoConfig.FileSavePath, filename)},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	if err := command.Run(); err != nil {
		return err
	}
	return nil
}
