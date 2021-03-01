# Cronparser
Crontab parser for golang(using yacc), you can use it to validate linux crontab etc.


# Usage

<pre>

	result, err := ParseCron("*/5 1-3 * JAN SAT", true)
	if err != nil {
		log.Println("parse error:", err)
		return
	}
	fmt.Println(result)

</pre>
