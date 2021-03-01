package cronparser

func ParseCron(d string, debug bool) (CronResult, error) {
	s := NewScanner(d, debug)
	yyParse(s)
	if s.err != nil {
		return CronResult{}, s.err
	}
	return s.data, nil
}