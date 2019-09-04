package bash

import "fmt"

func KillPid(pid int) string {
	return fmt.Sprintf("kill %d", pid)
}
