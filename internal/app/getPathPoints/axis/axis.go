package axes

func DefenitionAxis(width, height, AxisX, AxisY int) int {
	if width == 1 {
		return AxisY
	} else if height == 1 {
		return AxisX
	} else {
		return 0
	}
}

func ChangeAxis(axis, AxisX, AxisY int) int {
	if axis == AxisX {
		axis = AxisY
	} else {
		axis = AxisX
	}
	return axis
} 