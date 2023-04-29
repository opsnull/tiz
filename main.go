package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"syscall"
	"time"

	"github.com/getlantern/systray"
	"github.com/opsnull/tiz/icon"
	"github.com/mitchellh/go-homedir"
	"github.com/skratchdot/open-golang/open"
)

var logFileName string
var logFile *os.File

func main() {
	startTime := time.Now().UnixNano()
	logFileName = fmt.Sprintf(`/tmp/on_exit_%d.txt`, startTime)
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	onExit := func() {
		now := time.Now()
		logFile.WriteString(now.String())
	}

	systray.Run(onReady, onExit)
}

func onReady() {
	var cmd *exec.Cmd
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuitOrig.ClickedCh
		pgid, err := syscall.Getpgid(cmd.Process.Pid)
		if err == nil {
			syscall.Kill(-pgid, syscall.SIGKILL)
		}
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	go func() {
		systray.SetTemplateIcon(icon.Data, icon.Data)
		//systray.SetTitle("TiZ")
		enableCheckbox := systray.AddMenuItemCheckbox("Enable", "", false)
		systray.AddSeparator()
		logItem := systray.AddMenuItem("Log", "")
		logItem.Hide()
		systray.AddSeparator()
		infoItem := systray.AddMenuItem("Info", "")
		infoItem.Hide()


		for {
			select {
			case <-enableCheckbox.ClickedCh:
				if enableCheckbox.Checked() {
					enableCheckbox.Uncheck()
					logItem.Hide()
					infoItem.Hide()
					pgid, err := syscall.Getpgid(cmd.Process.Pid)
					if err == nil {
						syscall.Kill(-pgid, syscall.SIGKILL)
					}

				} else {
					enableCheckbox.Check()
					logItem.Show()
					infoItem.Show()
					cmd, _ = runGost()
				}
			case <-logItem.ClickedCh:
				if err := open.Run(logFileName); err != nil {
					panic(err)
				}
			}
		}
	}()
}

func runGost() (*exec.Cmd, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	// 将进程添加到新的进程组中
	attr := syscall.SysProcAttr{
		Setpgid: true,
	}
	cmd := exec.Command("/bin/bash", path.Join(home, ".ssh/tiz.sh"))
	cmd.Dir = "/tmp"
	cmd.Stderr = logFile
	cmd.Stdout = logFile
	cmd.SysProcAttr = &attr
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return cmd, nil
}
