package api

// Stop processing images and return any results accumulated so far.
func (a *AAPI) Interrupt() error {
	_, err := a.post(a.Config.Path.Interrupt, nil)
	return err
}
