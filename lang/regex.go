package lang

import (
	"regexp"
)

func RegexExamples() {
	pattern := regexp.MustCompile(`\d+`)
	pl(pattern.FindString("1234567890"))

	reStr := "The ape was at the apex"
	match, _ := regexp.MatchString(`(ape[^ ]?)`, reStr)
	pl("Match: ", match)

	reStr2 := "Cat rat mat fat pat"
	r, _ := regexp.Compile(`([crmfp]at)`)
	pl("MatchString: ", r.MatchString(reStr2))
	pl("FindString: ", r.FindString(reStr2))
	pl("Index: ", r.FindStringIndex(reStr2))
	pl("All: ", r.FindAllString(reStr2, -1))
	pl("FindAllStringSubmatchIndex: ", r.FindAllStringSubmatchIndex(reStr2, -1))
	pl("ReplaceAllString: ", r.ReplaceAllString(reStr2, "dog"))

}
