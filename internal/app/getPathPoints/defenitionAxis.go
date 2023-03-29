package getPathPoints

// определение оси
func (d *data) defenitionAxis(width, height int) int {
	if width == 1 {
		return AxisY
	} else if height == 1 {
		return AxisX
	} else {
		return 0
	}
}

// логика требует в некоторых местах менять ось на противоположную.
func (d *data) changeAxis(axis int) int {
	if axis == AxisX {
		axis = AxisY
	} else {
		axis = AxisX
	}
	return axis
} 