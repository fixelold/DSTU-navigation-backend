package axes

func DefenitionAxis(width, height int) int {
	if width == 1 {
		return AxisY
	} else if height == 1 {
		return AxisX
	} else {
		return 0
	}
}

func ChangeAxis(axis int) int {
	if axis == AxisX {
		axis = AxisY
	} else {
		axis = AxisX
	}
	return axis
} 