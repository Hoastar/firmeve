package firmeve

import (
	"github.com/firmeve/firmeve/event"
	"github.com/firmeve/firmeve/http"
	"github.com/firmeve/firmeve/kernel"
	"github.com/firmeve/firmeve/kernel/contract"
	logging "github.com/firmeve/firmeve/logger"
	"github.com/firmeve/firmeve/support"
	"github.com/spf13/cobra"
)

var (
	initProviders = []contract.Provider{
		new(event.Provider),
		new(logging.Provider),
	}

	defaultProviders = []contract.Provider{
		new(http.Provider),
	}

	defaultCommands = []contract.Command{
		new(http.HttpCommand),
	}
)

type (
	option struct {
		providers []contract.Provider
		commands  []contract.Command
	}

	//Running func(options ...support.Option) (contract.Application, *cobra.Command)
)

func WithProviders(providers []contract.Provider) support.Option {
	return func(object support.Object) {
		object.(*option).providers = providers
	}
}

func WithCommands(commands []contract.Command) support.Option {
	return func(object support.Object) {
		object.(*option).commands = commands
	}
}
func registerBaseProvider(f contract.Application) {
	f.RegisterMultiple([]contract.Provider{
		new(event.Provider),
		new(logging.Provider),
	}, false)
}

func RunDefault(options ...support.Option) (contract.Application, contract.BaseCommand) {
	option := parseOption(options)
	option.providers = append(defaultProviders, option.providers...)
	option.commands = append(defaultCommands, option.commands...)

	return Run(WithProviders(option.providers), WithCommands(option.commands))
}

func Run(options ...support.Option) (contract.Application, contract.BaseCommand) {
	option := parseOption(options)
	//app := kernel.New()
	//root := kernel.CommandRoot(app)
	option.providers = append(initProviders, option.providers...)
	command := kernel.NewCommand(option.providers, option.commands...)
	//command.AddCommand(option.commands...)

	err := command.Run()
	if err != nil && err != cobra.ErrSubCommandRequired {
		panic(err)
	}
	//for i, _ := range option.commands {
	//	option.commands[i].SetApplication(app)
	//	option.commands[i].SetProviders(option.providers)
	//	root.AddCommand(option.commands[i].Cmd())
	//}
	//
	//run(root)

	return command.Application(), command
}

//func run(cmd *cobra.Command) {
//	err := cmd.Execute()
//	if err != nil && err != cobra.ErrSubCommandRequired {
//		panic(err)
//	}
//}

func parseOption(options []support.Option) *option {
	return support.ApplyOption(&option{
		providers: make([]contract.Provider, 0),
	}, options...).(*option)
}
