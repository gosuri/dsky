package dsky

import (
	"github.com/spf13/cobra"
	"sort"
	"testing"
)

func testTopics() []*Topic {
	return []*Topic{
		{"b", "", true, nil, "foo"},
		{"a", "", true, nil, "foo"},
		{"c1", "", false, nil, "foo"},
	}
}

func testCmds() []*cobra.Command {
	return []*cobra.Command{
		&cobra.Command{
			Use: "a:1",
			Run: runHelp,
		},
		&cobra.Command{
			Use: "a:10",
			Run: runHelp,
		},
		&cobra.Command{
			Use: "b:2",
			Run: runHelp,
		},
	}
}

func TestSortByName(t *testing.T) {
	topics := testTopics()
	sort.Sort(ByName{topics})
	if topics[0].Name != "a" {
		t.Fatalf("expected: %s, got: %s", "a", topics[0].Name)
	}
}

func TestNamePadding(t *testing.T) {
	top := &Topic{"a", "", true, testCmds(), "foo"}
	if top.NamePadding() != 4 {
		t.Fatalf("expected: %d, got: %d", 4, top.NamePadding())
	}
}

func TestTopicNamePadding(t *testing.T) {
	topics := Topics(testTopics())
	if topics.TopicNamePadding() != 2 {
		t.Fatalf("expected: %d, got: %d", 2, topics.TopicNamePadding())
	}
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
