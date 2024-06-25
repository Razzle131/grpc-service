package dns

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/Razzle131/grpc-service/internal/consts"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetDNS() ([]string, error) {
	command := "cat /etc/resolv.conf | grep nameserver"

	output, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		return nil, status.Errorf(codes.Unknown, fmt.Sprintf("cmd error: %s", err))
	}

	rawData := strings.Split(string(output), "\n")
	var res []string

	for _, val := range rawData {
		// skipping wrong lines (empty, newlines, etc)
		if len(val) < 3 {
			continue
		}
		res = append(res, strings.Split(val, " ")[1])
	}

	return res, nil
}

func isValidIp(ip string) (bool, error) {
	return regexp.MatchString(consts.IpRegExp, ip)
}

func ExecCommandByRoot(command, sudoPassword string) error {
	cmd := exec.Command("sudo", "-S", "su", "root", "bash", "-c", command)
	cmd.Stdin = strings.NewReader(sudoPassword)

	if err := cmd.Run(); err != nil {
		return status.Errorf(codes.PermissionDenied, "given sudo password is incoreect")
	}

	return nil
}

func AddDns(ip, sudoPassword string) error {
	if match, err := isValidIp(ip); !match || err != nil {
		return status.Errorf(codes.Aborted, "given ip is not valid")
	}

	command := fmt.Sprintf("sudo echo -n '\nnameserver %s' >> /etc/resolv.conf", ip)

	if err := ExecCommandByRoot(command, sudoPassword); err != nil {
		return status.Errorf(codes.PermissionDenied, "cannot access file for dns configuration")
	}

	return nil
}

func RemoveDns(ip, sudoPassword string) error {
	if match, err := isValidIp(ip); !match || err != nil {
		return status.Errorf(codes.Aborted, "given ip is not valid")
	}

	command := fmt.Sprintf(`awk -v input="%s" '!index($0, input)' /etc/resolv.conf > $$.temp && cat $$.temp > /etc/resolv.conf && rm $$.temp`, ip)

	if err := ExecCommandByRoot(command, sudoPassword); err != nil {
		return status.Errorf(codes.PermissionDenied, "cannot access file for dns configuration")
	}

	return nil
}
