package dns

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const ipRegExp = `^(\d+\.){3}\d+$`

func GetDNS() ([]string, error) {
	//cmd := "cat /etc/resolv.conf | grep nameserver"
	cmd := "cat /home/razzle/go_projects/tz_yadro/foo.conf | grep nameserver"

	output, err := exec.Command("bash", "-c", cmd).Output()
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
	return regexp.MatchString(ipRegExp, ip)
}

func AddDns(ip string) error {
	if match, err := isValidIp(ip); !match || err != nil {
		return status.Errorf(codes.Aborted, "given ip is not valid")
	}

	cmd := fmt.Sprintf("echo -n '\nnameserver %s' >> /home/razzle/go_projects/tz_yadro/foo.conf", ip)

	err := exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		return status.Errorf(codes.Unknown, fmt.Sprintf("cmd error: %s", err))
	}

	return nil
}

func RemoveDns(ip string) error {
	if match, err := isValidIp(ip); !match || err != nil {
		return status.Errorf(codes.Aborted, "given ip is not valid")
	}

	cmd := fmt.Sprintf(`awk -v input="%s" '!index($0, input)' foo.conf > $$.temp && cat $$.temp > foo.conf && rm $$.temp`, ip)

	err := exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		return status.Errorf(codes.Unknown, fmt.Sprintf("cmd error: %s", err))
	}

	return nil
}
