package ffmpeg

import (
	"fmt"
	"github.com/Paxx-RnD/go-helper/helpers/slice_helper"
	"os/exec"
	"strings"
)

type IBuilder interface {
	Build() []string
}

type Builder ffmpeg

func (f *ffmpeg) Output(path string) IBuilder {
	f.arguments.Output.Append(path)
	return (*Builder)(f)
}

func (f *Builder) Build() []string {
	var args []string
	args = append(args, f.arguments.Inputs.Build()...)
	if len(f.arguments.FilterGraph.FilterChain) == 0 {
		args = append(args, f.arguments.VideoFilters.Build()...)
		args = append(args, f.arguments.AudioFilters.Build()...)
	} else {
		args = append(args, f.arguments.FilterGraph.Build()...)
	}
	args = append(args, f.arguments.Options.Build()...)
	args = append(args, f.arguments.Output.Build()...)

	return args
}

func (f *ffmpeg) Run(args []string) error {
	command := exec.Command(f.configuration.FfmpegPath)

	command.Args = append(command.Args, f.headers...)
	command.Args = append(command.Args, args...)

	command.Args = slice_helper.RemoveEmptyEntries(command.Args)

	if f.configuration.IsDebug {
		fmt.Println(strings.Join(command.Args, " "))
	}

	result, err := command.CombinedOutput()
	if err != nil {
		err = fmt.Errorf(fmt.Sprint(err) + ": " + string(result))
		return err
	}

	return nil
}
