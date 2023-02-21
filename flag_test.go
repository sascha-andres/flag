package flag

import (
	"flag"
	"os"
	"testing"

	"golang.org/x/exp/slices"
)

type verbTestCase struct {
	args          []string
	expectedVerbs []string
	name          string
}

var testCasesGetVerbs = []verbTestCase{
	{
		args:          []string{"single-verb", "-test", "1"},
		expectedVerbs: []string{"single-verb"},
		name:          "single verb",
	},
	{
		args:          []string{"-test", "1"},
		expectedVerbs: []string{},
		name:          "no verb",
	},
	{
		args:          []string{"two-verb", "two-second-verb", "-test", "1"},
		expectedVerbs: []string{"two-verb", "two-second-verb"},
		name:          "two verbs",
	},
}

func TestGetVerbs(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	for _, testCase := range testCasesGetVerbs {
		t.Run(testCase.name, func(t *testing.T) {
			os.Args = testCase.args
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			t.Logf("%#v", os.Args)
			Parse()
			result := GetVerbs()
			for _, foundVerb := range result {
				if !slices.Contains(testCase.expectedVerbs, foundVerb) {
					t.Fatalf("found %q which is not part of expected verbs", foundVerb)
				}
			}
			for _, verb := range testCase.expectedVerbs {
				if !slices.Contains(result, verb) {
					t.Fatalf("expected %q to be present", verb)
				}
			}
		})
	}
}
