package racer

import (
	"github.com/spf13/cobra"
)

// Topic represents a command help topic
type Topic struct {
	// Name is the name of the topic
	Name string

	// Desc is the description of the topic
	Desc string

	// Primary indicates if the topic a primary topic
	Primary bool

	// Commands are commands associated with the topic
	Commands []*cobra.Command

	// RootCmdPath is the path of the root command
	RootCmdPath string
}

// Topics represnet a list of help topics
type Topics []*Topic

// Len returns the number of topics in Topics list
func (t Topics) Len() int { return len(t) }

// Swap swaps topic i with topic j
func (t Topics) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

// Primary returns the Primary help topics from the topic list
func (t Topics) Primary() Topics {
	res := make(Topics, 0)
	for _, pt := range t {
		if pt.Primary {
			res = append(res, pt)
		}
	}
	return res
}

// Additional returns the additional topics in the list
func (t Topics) Additional() Topics {
	res := make(Topics, 0)
	for _, at := range t {
		if !at.Primary {
			res = append(res, at)
		}
	}
	return res
}

// ByName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Topic value.
type ByName struct{ Topics }

func (t ByName) Less(i, j int) bool { return t.Topics[i].Name < t.Topics[j].Name }

// Returns the padding required for the name when rending topic usage
func (t *Topic) NamePadding() int {
	padding := 1
	if len(t.Commands) > 0 {
		for _, cmd := range t.Commands {
			if len(cmd.Use) > padding {
				padding = len(cmd.Use)
			}
		}
	}
	return padding
}

// Returns the padding required for the name when rendering main help
func (t Topics) TopicNamePadding() int {
	padding := 1
	for _, topic := range t {
		if len(topic.Name) > padding {
			padding = len(topic.Name)
		}
	}
	return padding
}
