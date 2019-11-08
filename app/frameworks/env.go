package frameworks

import (
	"bufio"
	"os"
	"strings"

	"github.com/NasSilverBullet/twitter-clone-api/app/usecases"
)

func LoadEnv(logger usecases.Logger) error {
	logger.Info("Start loading environment variables")

	f, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer f.Close()

	lines := make([]string, 0, 100)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	for _, l := range lines {
		e := strings.Split(l, "=")
		os.Setenv(e[0], e[1])
	}

	logger.Infof("Finish loading environment variables")
	return nil
}
