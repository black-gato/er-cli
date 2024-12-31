package cmd

type Add struct {
	ErrDescription string   `arg:"" help:"Give a description of the error"`
	ErrValue       string   `arg:"" help:"The value of the error"`
	Tags           []string `arg:"" help:"keywords to assoicate with this entry" placeholder:"tag1 tag2 tag3,"`
}

func hell() {}
