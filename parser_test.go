package cronparser

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

func TestV(t *testing.T)  {
	result, err := ParseCron("*/5 1-3 * JAN SAT", true)
	if err != nil {
		log.Println("parse error:", err)
		return
	}
	fmt.Println(result)
}
func TestParseJson(t *testing.T) {
	convey.Convey("case minutes", t, func() {
		convey.Convey("case 20", func() {
			v, _ := ParseCron("20 1 * * *", true)
			convey.So(v.Minutes, convey.ShouldEqual, "20")
		})
		convey.Convey("case -1", func() {
			_, err := ParseCron("-1 1 * * *", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("case 60", func() {
			_, err := ParseCron("60 1 * * *", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("case *", func() {
			v, _ := ParseCron("* 1 * * *", true)
			convey.So(v.Minutes, convey.ShouldEqual, "*")
		})

		convey.Convey("case */number", func() {
			v, _ := ParseCron("*/12 1 * * *", true)
			convey.So(v.Minutes, convey.ShouldEqual, "*/12")
		})
		convey.Convey("case */number out of range", func() {
			_, err := ParseCron("*/64 1 * * *", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("case 3-5", func() {
			v, _ := ParseCron("3-5 1 * * *", true)
			convey.So(v.Minutes, convey.ShouldEqual, "3-5")
		})
		convey.Convey("case 1-70", func() {
			_, err := ParseCron("1-70 1 * * *", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("case 3,5", func() {
			v, _ := ParseCron("3,5 1 * * *", true)
			convey.So(v.Minutes, convey.ShouldEqual, "3,5")
		})
	})


	convey.Convey("case hours", t, func() {
		convey.Convey("case 20", func() {
			v, _ := ParseCron("* 20 * * *", true)
			convey.So(v.Hours, convey.ShouldEqual, "20")
		})
		convey.Convey("case -1", func() {
			_, err := ParseCron("* -1 * * *", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("case 24", func() {
			_, err := ParseCron("* 60 * * *", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("case *", func() {
			v, _ := ParseCron("* * * * *", true)
			convey.So(v.Hours, convey.ShouldEqual, "*")
		})

		convey.Convey("case */number", func() {
			v, _ := ParseCron("* */12 * * *", true)
			convey.So(v.Hours, convey.ShouldEqual, "*/12")
		})
		convey.Convey("case */number out of range", func() {
			_, err := ParseCron("* */24  * * *", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("case 3-5", func() {
			v, _ := ParseCron("* 3-5  * * *", true)
			convey.So(v.Hours, convey.ShouldEqual, "3-5")
		})
		convey.Convey("case 1-24", func() {
			_, err := ParseCron("* 1-24 * * *", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("case 3,5", func() {
			v, _ := ParseCron("1 3,5 * * *", true)
			convey.So(v.Hours, convey.ShouldEqual, "3,5")
		})
		convey.Convey("case 1,24", func() {
			_,  err := ParseCron("1 1,24 * * *", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
	})


	convey.Convey("case days", t, func() {
		convey.Convey("case 20", func() {
			v, _ := ParseCron("* * 20  * *", true)
			convey.So(v.Days, convey.ShouldEqual, "20")
		})
		convey.Convey("case -1", func() {
			_, err := ParseCron("* * -1 * *", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("case 32", func() {
			_, err := ParseCron("* * 32 * *", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("case *", func() {
			v, _ := ParseCron("* * * * *", true)
			convey.So(v.Days, convey.ShouldEqual, "*")
		})

		convey.Convey("case */number", func() {
			v, _ := ParseCron("* * */12 * *", true)
			convey.So(v.Days, convey.ShouldEqual, "*/12")
		})
		convey.Convey("case */number out of range", func() {
			_, err := ParseCron("* * */32  * *", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("case 3-5", func() {
			v, _ := ParseCron("* * 3-5  * *", true)
			convey.So(v.Days, convey.ShouldEqual, "3-5")
		})
		convey.Convey("case 1-31", func() {
			v, _ := ParseCron("* * 1-31 * *", true)
			convey.So(v.Days, convey.ShouldEqual, "1-31")
		})
		convey.Convey("case 3,5", func() {
			v, _ := ParseCron("1 * 3,5  * *", true)
			convey.So(v.Days, convey.ShouldEqual, "3,5")
		})
		convey.Convey("case 1,32", func() {
			_,  err := ParseCron("1 * 1,32 * *", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
	})


	convey.Convey("case month", t, func() {
		convey.Convey("case 12", func() {
			v, _ := ParseCron("* * * 12 *", true)
			convey.So(v.Months, convey.ShouldEqual, "12")
		})
		convey.Convey("case -1", func() {
			_, err := ParseCron("* * * -1 *", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("case 13", func() {
			_, err := ParseCron("* * * 13 *", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("case *", func() {
			v, _ := ParseCron("* * * * *", true)
			convey.So(v.Months, convey.ShouldEqual, "*")
		})

		convey.Convey("case */number", func() {
			v, _ := ParseCron("10 10 10 */3  *", true)
			convey.So(v.Months, convey.ShouldEqual, "*/3")
		})
		convey.Convey("case */number out of range", func() {
			_, err := ParseCron("* * 10 10 10 */13  *", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("case JAN", func() {
			v, _ := ParseCron("* * * JAN *", true)
			convey.So(v.Months, convey.ShouldEqual, "JAN")
		})
		convey.Convey("case JAN-DEC", func() {
			v, _ := ParseCron("* * * JAN-DEC *", true)
			convey.So(v.Months, convey.ShouldEqual, "JAN-DEC")
		})
		convey.Convey("case a-b", func() {
			_, err := ParseCron("* * * a-b*", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("case 3-5", func() {
			v, _ := ParseCron("* * * 3-5 *", true)
			convey.So(v.Months, convey.ShouldEqual, "3-5")
		})
		convey.Convey("case 3,5", func() {
			v, _ := ParseCron("1 * * 3,5 *", true)
			convey.So(v.Months, convey.ShouldEqual, "3,5")
		})
		convey.Convey("case JAN,DEC", func() {
			v, _ := ParseCron("1 * * JAN,DEC *", true)
			convey.So(v.Months, convey.ShouldEqual, "JAN,DEC")
		})
		convey.Convey("case 1,32", func() {
			_,  err := ParseCron("1 * * 1,32 *", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
	})


	convey.Convey("case week", t, func() {
		convey.Convey("case 5", func() {
			v, _ := ParseCron("* * * * 5", true)
			convey.So(v.Weeks, convey.ShouldEqual, "5")
		})
		convey.Convey("case -1", func() {
			_, err := ParseCron("* * * * -1", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("case 8", func() {
			_, err := ParseCron("* * * * 8", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("case *", func() {
			v, _ := ParseCron("* * * * *", true)
			convey.So(v.Weeks, convey.ShouldEqual, "*")
		})

		convey.Convey("case */number", func() {
			v, _ := ParseCron("10 10 10 10  */3", true)
			convey.So(v.Weeks, convey.ShouldEqual, "*/3")
		})
		convey.Convey("case */number out of range", func() {
			_, err := ParseCron("* * 10 10 10 * */13", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("case SUN", func() {
			v, _ := ParseCron("* * * * SUN", true)
			convey.So(v.Weeks, convey.ShouldEqual, "SUN")
		})
		convey.Convey("case SUN-SAT", func() {
			v, _ := ParseCron("* * * * SUN-SAT", true)
			convey.So(v.Weeks, convey.ShouldEqual, "SUN-SAT")
		})
		convey.Convey("case a-b", func() {
			_, err := ParseCron("* * * * a-b", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("case 3-5", func() {
			v, _ := ParseCron("* * * * 3-5", true)
			convey.So(v.Weeks, convey.ShouldEqual, "3-5")
		})
		convey.Convey("case 3,5", func() {
			v, _ := ParseCron("1 * * * 3,5", true)
			convey.So(v.Weeks, convey.ShouldEqual, "3,5")
		})
		convey.Convey("case JAN,DEC", func() {
			v, _ := ParseCron("1 * * * SUN,SAT", true)
			convey.So(v.Weeks, convey.ShouldEqual, "SUN,SAT")
		})
		convey.Convey("case 1,10", func() {
			_,  err := ParseCron("1 * * * 1,32", true)
			convey.So(err, convey.ShouldNotBeNil)
		})
	})
}
