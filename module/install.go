package module

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func detectDistro() string {
	osReleaseFile, err := os.Open("/etc/os-release")
	if err != nil {
		return "Unknown"
	}
	defer osReleaseFile.Close()

	osReleaseBytes, err := io.ReadAll(osReleaseFile)
	if err != nil {
		return "Unknown"
	}

	osRelease := string(osReleaseBytes)
	for _, line := range strings.Split(osRelease, "\n") {
		if strings.HasPrefix(line, "ID=") {
			return strings.Trim(strings.Split(line, "=")[1], "\"")
		}
	}

	return "Unknown"
}

func InstallPackageWithFilePath(filePath string) error {
	distro := detectDistro()
	var cmd *exec.Cmd
	switch distro {
	case "debian":
		cmd = exec.Command("sudo", "dpkg", "-i", filePath)
	case "arch":
		cmd = exec.Command("sudo", "pacman", "-U", filePath)
	case "centos":
		cmd = exec.Command("sudo", "yum", "install", "-y", filePath)
	default:
		fmt.Println("Unknown distribution")
		return nil
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	return nil
}
