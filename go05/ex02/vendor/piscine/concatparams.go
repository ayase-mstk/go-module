package piscine

func ConcatParams(args []string) string {
	var ret string
	for i, arg := range args {
		if i != 0 {
			ret += "\n"
		}
		ret += arg
	}
	return ret
}
