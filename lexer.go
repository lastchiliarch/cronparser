package cronparser

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"strings"
)

type CronResult struct {
	Minutes string
	Hours   string
	Days    string
	Months  string
	Weeks   string
}

// Parse parses the input and returs the result.
type Scanner struct {
	buf         *bufio.Reader
	s           string
	data        CronResult
	err         error
	debug       bool
	lastLiteral string
	pos         int
}

func NewScanner(s string, debug bool) *Scanner {
	return &Scanner{
		buf:   bufio.NewReader(strings.NewReader(s)),
		s:     s,
		debug: debug,
	}
}

func (sc *Scanner) Error(s string) {
	sc.err = fmt.Errorf("%v", s)
	if sc.debug {
		log.Println(sc.err)
		log.Println(sc.s)
		log.Println(strings.Repeat(" ", sc.pos) + "^")
	}
}

func (sc *Scanner) Reduced(rule, state int, lval *yySymType) bool {
	if sc.debug {
		fmt.Printf("rule: %v; state %v; lval: %v\n", rule, state, lval)
	}

	return false
}

func (s *Scanner) Lex(lval *yySymType) int {
	s.pos++
	return s.lex(lval)
}

func (s *Scanner) lex(lval *yySymType) int {
	for {
		r := s.read()
		if r == 0 {
			return 0
		}
		if isWhitespace(r) {
			continue
		}

		if isDigit(r) {
			s.unread()
			lval.i = s.scanNumber()
			return NUMBER
		}

		if isLetter(r) {
			s.unread()
			lval.s = s.scanString()
			return STRING
		}
		switch r {
		case ',':
			return COMMA
		case ':':
			return COLON
		case '*':
			return ASTERISK
		case '/':
			return SLASH
		case '-':
			return DASH
		default:
			s.err = errors.New("error")
			return 0
		}
	}
}


func (s *Scanner) scanString() string {
	var str []rune
	//JAN
	for {
		r := s.read()
		if !isLetter(r) {
			s.unread()
			break
		}
		str = append(str, r)
	}
	return string(str)
}

func (s *Scanner) scanNumber() int {
	var number []rune
	for {
		r := s.read()
		if !isDigit(r) {
			s.unread()
			break
		}
		number = append(number, r)
	}
	i, _ := strconv.Atoi(string(number))
	return i
}

func (s *Scanner) read() rune {
	ch, _, _ := s.buf.ReadRune()
	return ch
}

func (s *Scanner) unread() { _ = s.buf.UnreadRune() }

func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' || ch == '\n' }

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isLetter(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}

func isNumValid(i int,start int, end int) bool {
	if i >= start && i <= end {
		return true
	}
	return false
}

var MONTH = []string{"JAN", "FEB", "MAR", "APR", "MAY", "JUNE", "JULY", "AUG", "SEPT", "OCT", "NOV", "DEC", }

func isMonthValid(month interface{}) bool {
	switch month := month.(type) {
	case string:
		for _, v := range MONTH {
			if v == month {
				return true
			}
		}
		return false

	case int:
		return isNumValid(month, 1, 12)
	default:
		return false
	}
}


var WEEK = []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT", }
func isWeekValid(week interface{}) bool {
	switch week := week.(type) {
	case string:
		for _, v := range WEEK {
			if v == week {
				return true
			}
		}
		return false
	case int:
		return isNumValid(week, 0, 6)
	default:
		return false
	}
}
