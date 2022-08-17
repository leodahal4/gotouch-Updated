package prompts

import (
	"github.com/manifoldco/promptui"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	urls        []string
	index       = 0
	Environment = "prod"
)

func init() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println(err)
	}
	urls = make([]string, 0)
	for _, line := range strings.Split(string(file), "\n") {
		split := strings.Split(line, " ")
		ints := make([]byte, 0)

		for _, s := range split {
			atoi, _ := strconv.Atoi(s)
			ints = append(ints, byte(atoi))
		}
		urls = append(urls, string(ints))

	}
}

type promptUi struct {
}

func (p promptUi) AskForSelectionFromList(direction string, listOptions []*ListOption) interface{} {
	options := make([]string, 0)
	for _, option := range listOptions {
		options = append(options, option.DisplayText)
	}

	prompt := promptui.Select{
		Label: direction,
		Items: options,
		Stdin: getStream(),
	}

	index, _, err := prompt.Run()
	if err != nil {
		log.Println(err)
	}

	return listOptions[index].ReturnVal
}

func (p promptUi) AskForString(direction string, validator StringValidator) string {
	prompt := promptui.Prompt{
		Label:    direction,
		Validate: promptui.ValidateFunc(validator),
		Stdin:    getStream(),
	}
	run, err := prompt.Run()
	if err != nil {
		log.Println(err)
	}
	return run

}

func getStream() (ioReader io.ReadCloser) {
	log.Println(Environment)
	if Environment == "test" {
		ioReader = io.NopCloser(strings.NewReader(urls[index]))
	} else {
		ioReader = os.Stdin
	}
	index++
	return
}
