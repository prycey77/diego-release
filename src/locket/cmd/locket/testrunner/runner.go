package testrunner

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"code.cloudfoundry.org/durationjson"
	"code.cloudfoundry.org/lager/lagerflags"
	"code.cloudfoundry.org/diego-release/locket"
	"code.cloudfoundry.org/diego-release/locket/cmd/locket/config"
	"code.cloudfoundry.org/tlsconfig"
	. "github.com/onsi/gomega"
	"github.com/tedsuo/ifrit/ginkgomon"
)

var (
	fixturesPath = filepath.Join(os.Getenv("DIEGO_RELEASE_DIR"), "src/code.cloudfoundry.org/diego-release/locket/cmd/locket/fixtures")

	caCertFile = filepath.Join(fixturesPath, "ca.crt")
	certFile   = filepath.Join(fixturesPath, "cert.crt")
	keyFile    = filepath.Join(fixturesPath, "cert.key")
)

func NewLocketRunner(locketBinPath string, fs ...func(cfg *config.LocketConfig)) *ginkgomon.Runner {
	cfg := &config.LocketConfig{
		LagerConfig: lagerflags.LagerConfig{
			LogLevel:   lagerflags.INFO,
			TimeFormat: lagerflags.FormatUnixEpoch,
		},
		DatabaseDriver: "mysql",
		ReportInterval: durationjson.Duration(1 * time.Minute),
		CaFile:         caCertFile,
		CertFile:       certFile,
		KeyFile:        keyFile,
	}

	for _, f := range fs {
		f(cfg)
	}

	locketConfig, err := ioutil.TempFile("", "locket-config")
	Expect(err).NotTo(HaveOccurred())

	locketConfigFilePath := locketConfig.Name()

	encoder := json.NewEncoder(locketConfig)
	err = encoder.Encode(cfg)
	Expect(err).NotTo(HaveOccurred())
	Expect(locketConfig.Close()).To(Succeed())

	return ginkgomon.New(ginkgomon.Config{
		Name:              "locket",
		StartCheck:        "locket.started",
		StartCheckTimeout: 10 * time.Second,
		Command:           exec.Command(locketBinPath, "-config="+locketConfigFilePath),
		Cleanup: func() {
			os.RemoveAll(locketConfigFilePath)
		},
	})
}

func LocketClientTLSConfig() *tls.Config {
	tlsConfig, err := tlsconfig.Build(
		tlsconfig.WithInternalServiceDefaults(),
		tlsconfig.WithIdentityFromFile(certFile, keyFile),
	).Client(tlsconfig.WithAuthorityFromFile(caCertFile))
	Expect(err).NotTo(HaveOccurred())
	return tlsConfig
}

func ClientLocketConfig() locket.ClientLocketConfig {
	return locket.ClientLocketConfig{
		LocketCACertFile:     caCertFile,
		LocketClientCertFile: certFile,
		LocketClientKeyFile:  keyFile,
	}
}
