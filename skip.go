package api

// Stop processing current image and continue processing.
func (a *AAPI) Skip() error {
	_, err := a.post(a.Config.Path.Skip, nil)
	return err
}
