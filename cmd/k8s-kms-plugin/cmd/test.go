/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/thalescpl-io/k8s-kms-plugin/apis/istio/v1"
	"golang.org/x/sync/errgroup"
	"os"
	"path/filepath"
	"time"
)

var loop bool
var maxLoops int
var loopTime, timeout time.Duration

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test connectivety to the socket for some encrypt/decrypt",

	RunE: func(cmd *cobra.Command, args []string) error {
		time.Sleep(2 * time.Second)

		g := &errgroup.Group{}
		if loop {
			g.Go(loopTestRun)
		} else {
			g.Go(runTest)
		}
		return g.Wait()
	},
}

func loopTestRun() error {
	count := 0
	for {
		logrus.Info("Running Tests")
		_ = runTest()
		time.Sleep(10 * time.Second)
		count++
		if count > maxLoops {
			break
		}
	}
	return nil
}

func runTest() error {
	// Run Istio e2e tests against the socket

	ctx, cancel, c, err := istio.GetClientSocket(socketPath, timeout)
	defer cancel()
	if err != nil {
		logrus.Fatal(err)
		return err
	}

	// Generate a random UUID for request

	testuuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	var testKid []byte
	testKid, err = testuuid.MarshalText()
	if err != nil {
		return err
	}

	/*
		GenerateDEK
	*/
	logrus.Info("Test 1 GenerateKEK")
	var genKEKResp *istio.GenerateKEKResponse
	genKEKResp, err = c.GenerateKEK(ctx, &istio.GenerateKEKRequest{
		KekKid: testKid,
	})
	if err != nil {
		logrus.Errorf("Test 1 Failed: %v", err)
		return err
	}
	logrus.Infof("Test 1 Returned KEK ID: %s", string(genKEKResp.KekKid))
	/*
		GenerateDEK
	*/
	logrus.Info("Test 2 GenerateDEK")
	var genDEKResp *istio.GenerateDEKResponse
	if genDEKResp, err = c.GenerateDEK(ctx, &istio.GenerateDEKRequest{
		Size:   256,
		Kind:   istio.KeyKind_AES,
		KekKid: genKEKResp.KekKid,
	}); err != nil {
		logrus.Fatal(err)

		return err
	}

	logrus.Infof("Test 2 Returned WrappedDEK: %s", genDEKResp.EncryptedDekBlob)

	/*
		GenerateSEK
	*/

	logrus.Info("Test 3 GenerateSEK RSA")
	var genSEKResp *istio.GenerateSEKResponse
	if genSEKResp, err = c.GenerateSEK(ctx, &istio.GenerateSEKRequest{
		Size:             4096,
		Kind:             istio.KeyKind_RSA,
		KekKid:           genKEKResp.KekKid,
		EncryptedDekBlob: genDEKResp.EncryptedDekBlob,
	}); err != nil {
		logrus.Fatal(err)
		return err
	}
	logrus.Infof("Test 3 Returned WrappedSEK: %s", genSEKResp.EncryptedSekBlob)

	/*
		LoadSEK
	*/
	logrus.Info("Test 4 LoadSEK RSA")
	var loadSEKResp *istio.LoadSEKResponse
	if loadSEKResp, err = c.LoadSEK(ctx, &istio.LoadSEKRequest{

		KekKid:           genKEKResp.KekKid,
		EncryptedDekBlob: genDEKResp.EncryptedDekBlob,
		EncryptedSekBlob: genSEKResp.EncryptedSekBlob,
	}); err != nil {
		logrus.Fatal(err)
		return err
	}
	var out string
	if debug {
		out = string(loadSEKResp.ClearSek)
	} else {

		out = "Success"
	}
	logrus.Infof("Test 4 Returned LoadedSEK in PEM Format: %v", out)
	logrus.Infof("------------------------------------------------------------")

	/*
		GenerateRootCAK
	*/
	logrus.Info("Test 5 GenerateRootCAK RSA")
	var genCAKResp *istio.GenerateRootCAKResponse
	if genCAKResp, err = c.GenerateRootCAK(ctx, &istio.GenerateRootCAKRequest{

	}); err != nil {
		logrus.Fatal(err)
		return err
	}

	logrus.Infof("Test GenerateRootCAK KID Returned: %v", genCAKResp.RootCaKid)
	logrus.Infof("------------------------------------------------------------")
	return err

}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.PersistentFlags().StringVar(&socketPath, "socket", filepath.Join(os.TempDir(), "run", ".sock"), "Unix Socket")
	testCmd.Flags().BoolVar(&loop, "loop", false, "Should we run the test in a loop?")
	testCmd.Flags().DurationVar(&loopTime, "loop-sleep", 10, "How many seconds to sleep between test runs ")
	testCmd.Flags().IntVar(&maxLoops, "max-loops", 100, "How many seconds to sleep between test runs ")

	testCmd.Flags().DurationVar(&timeout, "timeout", 10*time.Second, "Timeout Duration")
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}