# Golang Masker (forked by @xamust)

Add custom masker

```
masker := New()
	out := func(i string) string {
		l := len([]rune(i))
		if l == 0 {
			return ""
		}
// set start index and length
		return masker.fixedOverlay(i, 5, 16)
	}
masker.OTMaskerString("AA11111BB1111111111111111CC", out)
---------
AA111*****************111CC
```

```
	out1 := func(i string) string {
		l := len([]rune(i))
		if l == 0 {
			return ""
		}
// set start index and length
		return masker.fixedOverlay(i, 5, 16)
	}
 // add to sync.Map func and key
	masker.AddCustomMasker("customKey1", out1)

	out2 := func(i string) string {
		l := len([]rune(i))
		if l == 0 {
			return ""
		}
// set start index and length
		return masker.fixedOverlay(i, 2, 3)
	}
// add to sync.Map func and key
	masker.AddCustomMasker("customKey2", out2)

masker.MaskerString("customKey1", "AA11111BB1111111111111111CC")
masker.MaskerString("customKey2", "A111AA11")
---------
AA111*****************111CC
A1****11
```
