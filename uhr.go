package uhr

import (
	"fmt"
	"sort"
	"time"
)

func Uhr(t time.Time) []string {
	result := []string{}

	if t.Minute() != 0 {
		result = append(
			result,
			// t.Format("15 Uhr 04"),
			number(t.Hour())+" Uhr "+number(t.Minute()),
		)
	}

	switch t.Minute() {
	case 0:
		result = append(
			result,
			// t.Format("15 Uhr"),
			number(t.Hour())+" Uhr",
			// t.Format("punkt 15"),
			"punkt "+number(t.Hour()),
		)
	case 15:
		result = append(
			result,
			// t.Format("viertel nach 15"),
			"Viertel nach "+number(t.Hour()),
		)
	case 30:
		result = append(
			result,
			// t.Add(time.Hour).Format("halb 15"),
			"halb "+number(t.Hour()+1),
		)
	case 45:
		result = append(
			result,
			// t.Add(time.Hour).Format("viertel vor 15"),
			"viertel vor "+number(t.Hour()+1),
		)
	}

	if t.Minute() > 0 && t.Minute() <= 20 {
		result = append(
			result,
			// fmt.Sprintf("%d nach %d", t.Minute(), t.Hour()),
			fmt.Sprintf("%s nach %s", number(t.Minute()), number(t.Hour())),
		)
		if t.Minute() < 5 {
			result = append(
				result,
				fmt.Sprintf("kurz nach %s", number(t.Hour())),
			)
		}
	}
	if t.Minute() >= 40 && t.Minute() <= 59 {
		result = append(
			result,
			// fmt.Sprintf("%d vor %d", 60-t.Minute(), t.Hour()+1),
			fmt.Sprintf("%s vor %s", number(60-t.Minute()), number(t.Hour()+1)),
		)
		if t.Minute() > 55 {
			result = append(
				result,
				fmt.Sprintf("kurz vor %s", number(t.Hour()+1)),
			)
		}
	}

	if t.Minute() >= 20 && t.Minute() < 30 {
		result = append(
			result,
			// fmt.Sprintf("%d vor halb %d", 30-t.Minute(), t.Hour()),
			fmt.Sprintf("%s vor halb %s", number(30-t.Minute()), number(t.Hour())),
		)
	}

	if t.Minute() > 30 && t.Minute() <= 40 {
		result = append(
			result,
			// fmt.Sprintf("%d nach halb %d", t.Minute()-30, t.Hour()),
			fmt.Sprintf("%s nach halb %s", number(t.Minute()-30), number(t.Hour())),
		)
	}

	if t.Hour() > 12 {
		result = append(result, Uhr(t.Add(-12*time.Hour))...)
	}

	sort.Strings(result)
	return result
}

type Weekdayer interface {
	Weekday() time.Weekday
}

func Weekday(t Weekdayer) string {
	switch t.Weekday() {
	case time.Sunday:
		return "Sonntag"
	case time.Monday:
		return "Montag"
	case time.Tuesday:
		return "Dienstag"
	case time.Wednesday:
		return "Mittwoch"
	case time.Thursday:
		return "Donnerstag"
	case time.Friday:
		return "Freitag"
	case time.Saturday:
		return "Samstag"
	default:
		return ""
	}
}

func number(i int) string {
	switch abs(i) {
	case 0:
		return "null"
	case 1:
		return "eins"
	case 2:
		return "zwei"
	case 3:
		return "drei"
	case 4:
		return "vier"
	case 5:
		return "fünf"
	case 6:
		return "sechs"
	case 7:
		return "sieben"
	case 8:
		return "acht"
	case 9:
		return "neun"
	case 10:
		return "zehn"
	case 11:
		return "elf"
	case 12:
		return "zwölf"
	case 13:
		return "dreizehn"
	case 14:
		return "vierzehn"
	case 15:
		return "fünfzehn"
	case 16:
		return "sechzehn"
	case 17:
		return "siebzehn"
	case 18:
		return "achtzehn"
	case 19:
		return "neunzehn"
	case 20:
		return "zwanzig"
	case 30:
		return "dreißig"
	case 40:
		return "vierzig"
	case 50:
		return "fünfzig"
	}
	return handleEins(number(i%10)) + "und" + number((i/10)*10)
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return i * -1
}

func handleEins(s string) string {
	if s == "eins" {
		return "ein"
	}
	return s
}
