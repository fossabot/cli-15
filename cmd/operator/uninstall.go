package operator

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/calyptia/cli/k8s"
)

func NewCmdUninstall() *cobra.Command {
	// Create a new default kubectl command and retrieve its flags
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}

	cmd := &cobra.Command{
		Use:     "operator",
		Aliases: []string{"opr"},
		Short:   "Uninstall operator components",
		RunE: func(cmd *cobra.Command, args []string) error {
			kctl := newKubectlCmd()
			namespace := cmd.Flag("kube-namespace").Value.String()
			kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
			kubeClientConfig, err := kubeConfig.ClientConfig()
			if err != nil {
				return err
			}

			clientSet, err := kubernetes.NewForConfig(kubeClientConfig)
			if err != nil {
				return err
			}
			k := &k8s.Client{
				Interface: clientSet,
			}

			version, err := k.CheckOperatorVersion(context.Background())
			if err != nil {
				return err
			}

			yaml, err := prepareUninstallManifest(version, namespace)
			if err != nil {
				return err
			}

			kctl.SetArgs([]string{"delete", "-f", yaml})

			err = kctl.Execute()
			if err != nil {
				return err
			}
			defer os.RemoveAll(yaml)

			cmd.Printf("Calyptia Operator uninstalled successfully.\n")
			return nil
		},
	}
	fs := cmd.Flags()
	clientcmd.BindOverrideFlags(configOverrides, fs, clientcmd.RecommendedConfigOverrideFlags("kube-"))
	return cmd
}

func prepareUninstallManifest(version string, namespace string) (string, error) {
	file, err := f.ReadFile(manifestFile)
	if err != nil {
		return "", err
	}

	fullFile := string(file)

	solveNamespace := solveNamespaceCreation(false, fullFile, namespace)
	withNamespace := injectNamespace(solveNamespace, namespace)

	dir, err := os.MkdirTemp("", "calyptia-operator")
	if err != nil {
		return "", err
	}

	sysFile, err := os.CreateTemp(dir, "operator_*.yaml")
	if err != nil {
		return "", err
	}
	defer sysFile.Close()

	_, err = sysFile.WriteString(withNamespace)
	if err != nil {
		return "", err
	}

	return sysFile.Name(), nil
}
