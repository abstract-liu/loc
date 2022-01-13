package utils
import (
	"os/exec"
)

func RunGitCommand(name string, arg ...string) (string, error) {
      cmd := exec.Command(name, arg...)
      msg, err := cmd.CombinedOutput()
      cmd.Run()

      return string(msg), err
  }

