package job

import (
	"fmt"
	"syscall"
	"testing"
)

const sleepCmd = "timeout 1"

func testSysProcAttr(t *testing.T, job *Job) {
	if job.Cmd.SysProcAttr.CreationFlags != syscall.CREATE_NEW_PROCESS_GROUP {
		t.Fatal("processes are not started in a new group")
	}

	if job.Cmd.SysProcAttr.CmdLine != fmt.Sprintf(`/C "%s"`, sleepCmd) {
		t.Fatal("CmdLine is not set correctly")
	}
}

func testStopErr(t *testing.T, err error) {
	// the error seems to vary on windows so we cannot test it meaningfully
}
