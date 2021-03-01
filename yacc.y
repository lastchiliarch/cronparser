%{
package cronparser
import "fmt"

func setScannerData(yylex interface{}, data CronResult) {
	yylex.(*Scanner).data = data
}

%}

%union {
empty   struct{}
s       string
i       int
c    	CronResult
}


%token<empty> COMMA COLON ASTERISK SLASH DASH
%token<s> STRING
%token<i> NUMBER

%type<s> minutes hours days months weeks
%type<c> value

%left SLASH ASTERISK DASH COMMA NUMBER

%%

cron:
	value {
		setScannerData(yylex, $1)
	}

value:
	minutes hours days months weeks
	{
		$$ = CronResult{Minutes: $1, Hours: $2, Days: $3, Months: $4, Weeks: $5}
	}

minutes:
	ASTERISK 	{ $$ = "*" }
	| ASTERISK SLASH minutes {
		$$ = fmt.Sprintf("*/%s", $3)
	}
	| minutes DASH minutes {
		$$ = fmt.Sprintf("%s-%s", $1, $3)
	}
	| minutes COMMA minutes {
		$$ = fmt.Sprintf("%s,%s", $1, $3)
	}
	| NUMBER {
		if !isNumValid($1, 0, 59) {
			yylex.Error(fmt.Sprintf("minutes:%d should between 0-59", $1))
		}
		$$ = fmt.Sprintf("%d", $1)
	}

hours:
	ASTERISK 	{ $$ = "*" }
	| ASTERISK SLASH hours {
		$$ = fmt.Sprintf("*/%s", $3)
	}
	| hours DASH hours{
		$$ = fmt.Sprintf("%s-%s", $1, $3)
	}
	| hours COMMA hours{
		$$ = fmt.Sprintf("%s,%s", $1, $3)
	}
	| NUMBER {
		if !isNumValid($1, 0, 23) {
			yylex.Error(fmt.Sprintf("hours:%d should between 0-23", $1))
		}
		$$ = fmt.Sprintf("%d", $1)
	}

days:
	ASTERISK 	{ $$ = "*" }
	| ASTERISK SLASH days {
		$$ = fmt.Sprintf("*/%s", $3)
	}
	| days DASH days {
		$$ = fmt.Sprintf("%s-%s", $1, $3)
	}
	| days COMMA days {
		$$ = fmt.Sprintf("%s,%s", $1, $3)
	}
	| NUMBER {
		if !isNumValid($1, 1, 31) {
			yylex.Error(fmt.Sprintf("days:%d should between 1-31", $1))
		}
		$$ = fmt.Sprintf("%d", $1)
	}


months:
	ASTERISK 	{ $$ = "*" }
	| ASTERISK SLASH months {
		$$ = fmt.Sprintf("*/%s", $3)
	}
	| months DASH months {
		$$ = fmt.Sprintf("%s-%s", $1, $3)
	}
	| months COMMA months {
		$$ = fmt.Sprintf("%s,%s", $1, $3)
	}
	| NUMBER {
		if !isMonthValid($1) {
			yylex.Error(fmt.Sprintf("months:%d should between 1–12 or JAN–DEC", $1))
		}
		$$ = fmt.Sprintf("%d", $1)
	}
	| STRING {
		if !isMonthValid($1) {
			yylex.Error(fmt.Sprintf("months:%s should between 1–12 or JAN–DEC", $1))
		}
		$$ = fmt.Sprintf("%s", $1)
	}


weeks:
	ASTERISK 	{ $$ = "*" }
	| ASTERISK SLASH weeks {
		$$ = fmt.Sprintf("*/%s", $3)
	}
	| weeks DASH weeks {
		$$ = fmt.Sprintf("%s-%s", $1, $3)
	}
	| weeks COMMA weeks {
		$$ = fmt.Sprintf("%s,%s", $1, $3)
	}
	| NUMBER {
		if !isWeekValid($1) {
			yylex.Error(fmt.Sprintf("weeks:%d should between 0–6 or SUN–SAT", $1))
		}
		$$ = fmt.Sprintf("%d", $1)
	}
	| STRING {
		if !isWeekValid($1) {
			yylex.Error(fmt.Sprintf("months:%s should between 0-6 or SUN–SAT", $1))
		}
		$$ = fmt.Sprintf("%s", $1)
	}
