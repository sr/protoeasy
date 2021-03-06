package main

import (
	"fmt"
	"os"
	"strings"

	"google.golang.org/grpc"

	"go.pedge.io/env"
	"go.pedge.io/protoeasy"
	"go.pedge.io/protolog"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type appEnv struct {
	Address string `env:"PROTOEASY_ADDRESS"`
}

type options struct {
	GoModifiers  []string
	GoPluginType string
	OutDirPath   string
}

func main() {
	env.Main(do, &appEnv{})
}

func do(appEnvObj interface{}) error {
	appEnv := appEnvObj.(*appEnv)
	compileOptions := &protoeasy.CompileOptions{}
	options := &options{}

	rootCmd := &cobra.Command{
		Use: fmt.Sprintf("%s directory", os.Args[0]),
		RunE: func(_ *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("must pass one argument, the directory, but passed %d arguments", len(args))
			}
			if err := optionsToCompileOptions(options, compileOptions); err != nil {
				return err
			}
			dirPath := args[0]
			outDirPath := args[0]
			if options.OutDirPath != "" {
				outDirPath = options.OutDirPath
			}

			compiler := protoeasy.DefaultServerCompiler
			if appEnv.Address != "" {
				clientConn, err := grpc.Dial(appEnv.Address, grpc.WithInsecure())
				if err != nil {
					return err
				}
				compiler = protoeasy.NewClientCompiler(
					protoeasy.NewAPIClient(
						clientConn,
					),
					protoeasy.CompilerOptions{},
				)
			}

			commands, err := compiler.Compile(dirPath, outDirPath, compileOptions)
			if err != nil {
				return err
			}
			logCommands(commands)
			return nil
		},
	}
	bindCompileOptions(rootCmd.Flags(), compileOptions)
	bindOptions(rootCmd.Flags(), options)

	return rootCmd.Execute()
}

func bindCompileOptions(flagSet *pflag.FlagSet, compileOptions *protoeasy.CompileOptions) {
	flagSet.BoolVar(&compileOptions.Grpc, "grpc", false, "Output grpc files.")
	flagSet.BoolVar(&compileOptions.GrpcGateway, "grpc-gateway", false, "Output grpc-gateway .gw.go files.")
	flagSet.BoolVar(&compileOptions.NoDefaultIncludes, "no-default-includes", false, "Do not import the default include directories, implies --go-no-default-modifiers.")
	flagSet.StringSliceVar(&compileOptions.ExcludePattern, "exclude", []string{}, "Exclude file patterns.")

	flagSet.BoolVar(&compileOptions.Cpp, "cpp", false, "Output cpp files.")
	flagSet.StringVar(&compileOptions.CppRelOut, "cpp-rel-out", "", "The directory, relative to the output directory, to output cpp files.")

	flagSet.BoolVar(&compileOptions.Csharp, "csharp", false, "Output csharp files.")
	flagSet.StringVar(&compileOptions.CsharpRelOut, "csharp-rel-out", "", "The directory, relative to the output directory, to output csharp files.")

	flagSet.BoolVar(&compileOptions.Go, "go", false, "Output go files.")
	flagSet.StringVar(&compileOptions.GoRelOut, "go-rel-out", "", "The directory, relative to the output directory, to output go files.")
	flagSet.StringVar(&compileOptions.GoImportPath, "go-import-path", "", "Go package.")
	flagSet.BoolVar(&compileOptions.GoNoDefaultModifiers, "go-no-default-modifiers", false, "Do not set the default Mfile=package modifiers for --go_out.")

	flagSet.BoolVar(&compileOptions.Objc, "objc", false, "Output objc files.")
	flagSet.StringVar(&compileOptions.ObjcRelOut, "objc-rel-out", "", "The directory, relative to the output directory, to output objc files.")

	flagSet.BoolVar(&compileOptions.Python, "python", false, "Output python files.")
	flagSet.StringVar(&compileOptions.PythonRelOut, "python-rel-out", "", "The directory, relative to the output directory, to output python files.")

	flagSet.BoolVar(&compileOptions.Ruby, "ruby", false, "Output ruby files.")
	flagSet.StringVar(&compileOptions.RubyRelOut, "ruby-rel-out", "", "The directory, relative to the output directory, to output ruby files.")
}

func bindOptions(flagSet *pflag.FlagSet, options *options) {
	flagSet.StringSliceVar(&options.GoModifiers, "go-modifier", []string{}, "Extra Mfile=package modifiers for --go_out, specify just as file=package to this flag.")
	flagSet.StringVar(&options.GoPluginType, "go-protoc-plugin", "go", fmt.Sprintf("The go protoc plugin to use, allowed values are %s, if not go, --go-no-default-modifiers is implied.", strings.Join(protoeasy.AllGoPluginTypeSimpleStrings(), ",")))
	flagSet.StringVar(&options.OutDirPath, "out", "", "Customize out directory path.")
}

func optionsToCompileOptions(options *options, compileOptions *protoeasy.CompileOptions) error {
	if strings.ToLower(options.GoPluginType) == "none" {
		return fmt.Errorf("invalid value for --go-protoc-plugin: %s", options.GoPluginType)
	}
	goPluginType, err := protoeasy.GoPluginTypeSimpleValueOf(options.GoPluginType)
	if err != nil {
		return err
	}
	compileOptions.GoPluginType = goPluginType
	// TODO(pedge): duplicated logic in goPlugin struct
	if goPluginType != protoeasy.GoPluginType_GO_PLUGIN_TYPE_GO {
		compileOptions.GoNoDefaultModifiers = true
	}
	modifiers := make(map[string]string)
	for _, modifierString := range options.GoModifiers {
		split := strings.SplitN(modifierString, "=", 2)
		if len(split) != 2 {
			return fmt.Errorf("invalid valid for --go-modifier: %s", modifierString)
		}
		modifiers[split[0]] = split[1]
	}
	compileOptions.GoModifiers = modifiers
	// TODO(pedge): this should not be in this function
	// TODO(pedge): duplicated logic in goPlugin struct
	if compileOptions.NoDefaultIncludes {
		compileOptions.GoNoDefaultModifiers = true
	}
	return nil
}

func logCommands(commands []*protoeasy.Command) {
	for _, command := range commands {
		if len(command.Arg) > 0 {
			protolog.Infof("\n%s\n", strings.Join(command.Arg, " \\\n\t"))
		}
	}
}
