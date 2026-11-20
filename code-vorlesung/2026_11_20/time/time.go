package time

type Time int

func FromSeconds(m int) Time {
	return Time(m)
}

func FromMinutes(c int) Time {
	//return FromSeconds(c * 60)
	return Time(c * 60)
}

func FromHours(k int) Time {
	return FromSeconds(k * 3600)
}

func (l Time) Seconds() int {
	return int(l)
}

func (l Time) Minutes() int {
	return l.Seconds() / 60
}

func (l Time) Hours() int {
	return l.Minutes() / 60
}
