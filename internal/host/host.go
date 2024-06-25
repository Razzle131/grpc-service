package host

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/Razzle131/grpc-service/internal/dns"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetHost() (string, error) {
	cmd := exec.Command("hostname")
	output, err := cmd.Output()
	if err != nil {
		return "", status.Errorf(codes.Unknown, fmt.Sprintf("cmd error: %s", err))
	}

	return strings.TrimSpace(string(output)), nil
}

func SetHost(newHostname, sudoPassword string) (string, error) {
	if newHostname == "" {
		return "", status.Errorf(codes.InvalidArgument, "the length of the new hostname must be greater than 0")
	}

	command := fmt.Sprintf("hostnamectl set-hostname %s", newHostname)
	// cmd := exec.Command("sudo", "-S", "hostnamectl", "set-hostname", newHostname)
	// cmd.Stdin = strings.NewReader(sudoPassword)

	if err := dns.ExecCommandByRoot(command, sudoPassword); err != nil {
		return "", status.Errorf(codes.PermissionDenied, "given sudo password is incorrect")
	}

	return GetHost()
}
