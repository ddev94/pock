package helpers

import "fmt"

const (
	SettingListLayout = "listLayout"
	SettingDateFormat = "dateFormat"

	ListLayoutTable  = "table"
	ListLayoutSimple = "simple"

	DateFormatRelative = "relative"
	DateFormatLocale   = "locale"
	DateFormatISO      = "iso"
)

func ValidSettingKeys() []string {
	return []string{SettingListLayout, SettingDateFormat}
}

func ValidateSettingUpdate(key, value string) error {
	switch key {
	case SettingListLayout:
		if value == ListLayoutTable || value == ListLayoutSimple {
			return nil
		}
		return fmt.Errorf("invalid value for listLayout. Must be 'table' or 'simple'")
	case SettingDateFormat:
		if value == DateFormatRelative || value == DateFormatLocale || value == DateFormatISO {
			return nil
		}
		return fmt.Errorf("invalid value for dateFormat. Must be 'relative', 'locale', or 'iso'")
	default:
		return fmt.Errorf("unknown setting key")
	}
}
