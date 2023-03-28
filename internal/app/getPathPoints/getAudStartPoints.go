package getPathPoints

func (d *data) getAudStartPoints() error {
	var err error

	axis := d.defenitionAxis(d.AudienceBorderPoint.Widht, d.AudienceBorderPoint.Height)

	err = d.getPoints(axis)
	if err != nil {
		return err
	}

	return nil
}