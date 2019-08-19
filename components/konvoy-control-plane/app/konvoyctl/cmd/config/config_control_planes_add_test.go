package config_test

import (
	"bytes"
	"github.com/Kong/konvoy/components/konvoy-control-plane/app/konvoyctl/cmd"
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/spf13/cobra"
)

var _ = Describe("konvoy config control-planes add", func() {

	var configFile *os.File

	BeforeEach(func() {
		var err error
		configFile, err = ioutil.TempFile("", "")
		Expect(err).ToNot(HaveOccurred())
	})
	AfterEach(func() {
		if configFile != nil {
			os.Remove(configFile.Name())
		}
	})

	var rootCmd *cobra.Command
	var buf *bytes.Buffer

	BeforeEach(func() {
		rootCmd = cmd.DefaultRootCmd()
		buf = &bytes.Buffer{}
		rootCmd.SetOut(buf)
	})

	Describe("error cases", func() {

		It("should require name", func() {
			// given
			rootCmd.SetArgs([]string{"--config-file", configFile.Name(),
				"config", "control-planes", "add"})
			// when
			err := rootCmd.Execute()
			// then
			Expect(err.Error()).To(MatchRegexp(requiredFlagNotSet("name")))
		})

		It("should require API Server URL", func() {
			// given
			rootCmd.SetArgs([]string{"--config-file", configFile.Name(),
				"config", "control-planes", "add",
				"--name", "example"})
			// when
			err := rootCmd.Execute()
			// then
			Expect(err.Error()).To(MatchRegexp(requiredFlagNotSet("api-server-url")))
		})

		It("should fail to add a new Control Plane with duplicate name", func() {
			// given
			rootCmd.SetArgs([]string{"--config-file", filepath.Join("testdata", "config-control-planes-add.01.golden.yaml"),
				"config", "control-planes", "add",
				"--name", "example",
				"--api-server-url", "https://konvoy-control-plane.internal:5681"})
			// when
			err := rootCmd.Execute()
			// then
			Expect(err).To(MatchError(`Control Plane with name "example" already exists`))
		})
	})

	Describe("happy path", func() {

		type testCase struct {
			configFile string
			goldenFile string
		}

		DescribeTable("should add a new Control Plane by name and address",
			func(given testCase) {
				// setup
				initial, err := ioutil.ReadFile(filepath.Join("testdata", given.configFile))
				Expect(err).ToNot(HaveOccurred())
				err = ioutil.WriteFile(configFile.Name(), initial, 0600)
				Expect(err).ToNot(HaveOccurred())

				// given
				rootCmd.SetArgs([]string{"--config-file", configFile.Name(),
					"config", "control-planes", "add",
					"--name", "example",
					"--api-server-url", "https://konvoy-control-plane.internal:5681"})
				// when
				err = rootCmd.Execute()
				// then
				Expect(err).ToNot(HaveOccurred())

				// when
				expected, err := ioutil.ReadFile(filepath.Join("testdata", given.goldenFile))
				// then
				Expect(err).ToNot(HaveOccurred())

				// when
				actual, err := ioutil.ReadFile(configFile.Name())
				// then
				Expect(err).ToNot(HaveOccurred())

				// and
				Expect(actual).To(MatchYAML(expected))
			},
			Entry("should add a first Control Plane", testCase{
				configFile: "config-control-planes-add.01.initial.yaml",
				goldenFile: "config-control-planes-add.01.golden.yaml",
			}),
			Entry("should add a second Control Plane", testCase{
				configFile: "config-control-planes-add.02.initial.yaml",
				goldenFile: "config-control-planes-add.02.golden.yaml",
			}),
		)
	})
})