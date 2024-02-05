package masker

import "testing"

const (
	testCase1 = "AA11111BB1111111111111111CC"
	testCase2 = "A111AA11"
)

func TestCustomSingle(t *testing.T) {
	masker := New()
	out := func(i string) string {
		l := len([]rune(i))
		if l == 0 {
			return ""
		}
		return masker.fixedOverlay(i, 5, 16)
	}
	t.Log(masker.OTMaskerString(testCase1, out))
}

func TestCustomList(t *testing.T) {
	customKey1 := "RBS"
	customKey2 := "GRZ"
	masker := New()

	out1 := func(i string) string {
		l := len([]rune(i))
		if l == 0 {
			return ""
		}
		return masker.fixedOverlay(i, 5, 16)
	}
	masker.AddCustomMasker(customKey1, out1)

	out2 := func(i string) string {
		l := len([]rune(i))
		if l == 0 {
			return ""
		}
		return masker.fixedOverlay(i, 2, 3)
	}
	masker.AddCustomMasker(customKey2, out2)

	t.Log(masker.MaskerString(customKey1, testCase1))
	t.Log(masker.MaskerString(customKey2, testCase2))
}
