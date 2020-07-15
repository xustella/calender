package calender

type Date struct {
	year, month, day int
}

//the setter methods:
func (d *Date) SetYear(y int) {
	d.year = y
	//or
	//(*d).year = y
}
func (d *Date) SetMonth(m int) {
	d.month = m
}
func (d *Date) SetDay(day int) {
	d.day = day
}
