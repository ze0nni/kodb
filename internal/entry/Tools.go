package entry

import (
	"strconv"
)

func Int(key string, entry Entry) (int, bool) {
	if v, ok := entry[key]; ok {
		if i, err := strconv.Atoi(v); nil == err {
			return i, true
		}
	}
	return 0, false
}

func IntDef(key string, def int, entry Entry) int {
	if v, ok := entry[key]; ok {
		if i, err := strconv.Atoi(v); nil == err {
			return i
		}
	}
	return def
}

func SetInt(key string, value int, entry Entry) {
	entry[key] = strconv.Itoa(value)
}

func String(key string, entry Entry) (string, bool) {
	if v, ok := entry[key]; ok {
		return v, true
	}
	return "", false
}

func StringDef(key string, def string, entry Entry) string {
	if v, ok := entry[key]; ok {
		return v
	}
	return def
}

func SetString(key string, value string, entry Entry) {
	entry[key] = value
}
