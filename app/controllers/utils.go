package controllers

func defaultTo(value, default_value string) string {
	if value == "" {
		return default_value
	}
	return value
}
