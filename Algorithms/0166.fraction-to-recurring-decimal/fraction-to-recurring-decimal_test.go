package Problem0166

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	numerator   int
	denominator int
	ans         string
}{

	{
		-7,
		-12,
		"0.58(3)",
	},

	{
		7,
		-12,
		"-0.58(3)",
	},

	{
		1,
		90,
		"0.0(1)",
	},

	{
		5425221,
		-45664,
		"-118.80739(75122634898388227049754730203223545900490539593552908199018920812894183601962158374211632796075683251576734407848633496846531184302733006306937631394533987386124737210932025227750525578135949544498948843728100911002102312543798177995795374912403644008409250175192711983181499649614576033637000700770847932725998598458304134548002803083391730903994393833216538192011212333566923615977575332866152768044849334267694463910301331464611072179397337070777855641205325858444288717589348283111422564821303433777154870357393132445690259285213735108619481429572529782761037140854940434477925718290119131044148563419761737911702873160476524176594253679046951646811492641906096706377014716187806587245970567624386825508058864)",
	},

	{
		5425221,
		45664,
		"118.80739(75122634898388227049754730203223545900490539593552908199018920812894183601962158374211632796075683251576734407848633496846531184302733006306937631394533987386124737210932025227750525578135949544498948843728100911002102312543798177995795374912403644008409250175192711983181499649614576033637000700770847932725998598458304134548002803083391730903994393833216538192011212333566923615977575332866152768044849334267694463910301331464611072179397337070777855641205325858444288717589348283111422564821303433777154870357393132445690259285213735108619481429572529782761037140854940434477925718290119131044148563419761737911702873160476524176594253679046951646811492641906096706377014716187806587245970567624386825508058864)",
	},

	{
		1,
		2,
		"0.5",
	},

	{
		2,
		1,
		"2",
	},

	{
		2,
		3,
		"0.(6)",
	},

	{
		4,
		7,
		"0.(571428)",
	},

	// 可以有多个 testcase
}

func Test_fractionToDecimal(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, fractionToDecimal(tc.numerator, tc.denominator), "输入:%v", tc)
	}
}

func Benchmark_fractionToDecimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			fractionToDecimal(tc.numerator, tc.denominator)
		}
	}
}
