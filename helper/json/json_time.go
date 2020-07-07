package json

import "time"

type SimpleTime struct {
	time.Time
}

func (m *SimpleTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	// tt, err := time.Parse(`"`+time.RFC3339+`"`, string(data))
	tt, err := time.Parse(`"`+"2006-01-02 15:04:05"+`"`, string(data))
	*m = SimpleTime{tt}
	return err
}
