package fact

import (
	"math/big"
	"testing"
)

func TestFactRecursive(t *testing.T) {
	cases := []struct {
		input    int64
		expected string
	}{
		{1, "1"},
		{2, "2"},
		{3, "6"},
		{6, "720"},
		{9, "362880"},
		{10, "3628800"},
		{100, "93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000"},
		{200, "788657867364790503552363213932185062295135977687173263294742533244359449963403342920304284011984623904177212138919638830257642790242637105061926624952829931113462857270763317237396988943922445621451664240254033291864131227428294853277524242407573903240321257405579568660226031904170324062351700858796178922222789623703897374720000000000000000000000000000000000000000000000000"},
	}

	for _, c := range cases {
		input := big.NewInt(c.input)
		result := FactRecursive(input)
		if result.String() != c.expected {
			t.Errorf("FactRecursive(%d) is %s, want %s",
				input, result, c.expected)
		}
		result = FactIterative(input)
		if result.String() != c.expected {
			t.Errorf("FactIterative(%d) is %s, want %s",
				input, result, c.expected)
		}
	}
}

func BenchmarkFactRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactRecursive(big.NewInt(200))
	}
}

func BenchmarkFactIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactIterative(big.NewInt(200))
	}
}
