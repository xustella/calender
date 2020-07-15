package calender

import "errors"

type Date struct {
	year, month, day int
}

//the setter methods:
func (d *Date) SetYear(y int) error {
	if y < 2000 || y > 3000 {
		return errors.New("Invalid year")
	}
	d.year = y
	return nil
	//or
	//(*d).year = y
}
func (d *Date) SetMonth(m int) error {
	if m < 1 || m > 12 {
		return errors.New("Invalid month")
	}
	d.month = m
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid day")
	}
	d.day = day
	return nil
}

//the getter methods:
func (d *Date) Year() int {
	return d.year
}
func (d *Date) Month() int {
	return d.month
}
func (d *Date) Day() int {
	return d.day
}
