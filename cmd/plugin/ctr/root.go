package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"os"
	"strings"
)

var (
	KubernetesConfigFlags *genericclioptions.ConfigFlags
)

func newCtrOptions() *CtrOptions {
	return &CtrOptions{
		configFlags: genericclioptions.NewConfigFlags(true),
	}
}

func RootCmd() *cobra.Command {
	o := newCtrOptions()
	cmd := &cobra.Command{
		Use:           "ctr <pod-name>",
		Short:         "display all containers in the pod",
		Long:          `display all containers in the pod`,
		SilenceErrors: true,
		SilenceUsage:  true,
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if ver, err := cmd.Flags().GetBool("version"); err != nil {
				return err
			} else if ver {
				fmt.Println(FullVersion())
				return nil
			}

			if len(args) == 0 {
				_ = cmd.Help()
				return nil
			}

			if err := o.Complete(cmd, args); err != nil {
				return err
			}
			o.Run()
			return nil

			//log := logger.NewLogger()
			//log.Info("")
			//
			//s := spin.New()
			//finishedCh := make(chan bool, 1)
			//namespaceName := make(chan string, 1)
			//go func() {
			//	lastNamespaceName := ""
			//	for {
			//		select {
			//		case <-finishedCh:
			//			fmt.Printf("\r")
			//			return
			//		case n := <-namespaceName:
			//			lastNamespaceName = n
			//			fmt.Println(lastNamespaceName + "123")
			//		case <-time.After(time.Millisecond * 100):
			//			if lastNamespaceName == "" {
			//				fmt.Printf("\r  \033[36mSearching for namespaces\033[m %s", s.Next())
			//			} else {
			//				fmt.Printf("\r  \033[36mSearching for namespaces\033[m %s (%s)", s.Next(), lastNamespaceName)
			//			}
			//		}
			//	}
			//}()
			//defer func() {
			//	finishedCh <- true
			//}()
			//
			//if err := RunPlugin(KubernetesConfigFlags, namespaceName); err != nil {
			//	return errors.Unwrap(err)
			//}
			//
			//log.Info("")
			//
			//return nil
		},
	}

	cobra.OnInitialize(initConfig)

	//KubernetesConfigFlags = genericclioptions.NewConfigFlags(false)
	o.configFlags.AddFlags(cmd.Flags())

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	return cmd
}

func InitAndExecute() {
	if err := RootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	viper.AutomaticEnv()
}
