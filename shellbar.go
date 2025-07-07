package shellbar

type Shellbar struct{}

func (s *Shellbar) Run() error {
	cmd, err := parseArgs()
	if err != nil {
		return err
	}
	return cmd.Run()
}