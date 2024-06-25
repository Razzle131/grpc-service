package host

import (
	"fmt"
	"os/exec"
	"strings"

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

func SetHost(hostname string) (string, error) {
	if hostname == "" {
		return "", status.Errorf(codes.InvalidArgument, "the length of the new hostname must be greater than 0")
	}

	// change hostname
	changeHostNameCmd := exec.Command("sudo", "hostnamectl", "set-hostname", hostname)
	if err := changeHostNameCmd.Run(); err != nil {
		return "", status.Errorf(codes.Unknown, fmt.Sprintf("cmd error: %s", err))
	}

	return GetHost()
}
