package shell

import (
	"./commands"
	"./general"
	"./ldaps"
	"./stats"
	"./twil"
	"fmt"
	"github.com/gobs/readline"
	"gopkg.in/ldap.v3"
	"strings"
)

var ld *ldap.Conn

var found string = "no"
var list []string

var matches = make([]string, 0, len(list))

var coms = commands.Commands{
	{
		Name:        "LDAP",
		ShortName:   "ldap",
		Usage:       "ldap commands",
		Action:      NoAction,
		SubCommands: ldaps.LdapSubs,
		Category:    "ldaps",
	},
	{
		Name:        "Twil",
		ShortName:   "twil",
		Usage:       "use the twilio api through the shell",
		Action:      NoAction,
		SubCommands: twil.TwilSubs,
		Category:    "twil",
	},
	{
		Name:        "stats",
		ShortName:   "stats",
		Usage:       "stats commands",
		Action:      NoAction,
		SubCommands: stats.StatsSubs,
		Category:    "stats",
	},
	{
		Name:      "Quit",
		ShortName: "quit",
		Usage:     "Exit the shell",
		Action:    general.End,
		Category:  "general",
	},
	{
		Name:      "Clear",
		ShortName: "clear",
		Usage:     "Clear the screen",
		Action:    general.Clear,
		Category:  "general",
	},
}

func AttemptedCompletion(text string, start, end int) []string {
	if start == 0 {
		return readline.CompletionMatches(text, CompletionEntry)
	} else {
		return nil
	}
}

func CompletionEntry(prefix string, index int) string {
	if index == 0 {
		matches = matches[:0]

		for _, w := range list {
			if strings.HasPrefix(w, prefix) {
				matches = append(matches, w)
			}
		}
	}

	if index < len(matches) {
		return matches[index]
	} else {
		return ""
	}
}
func NoAction() {
	fmt.Println("No action supplied with command")

}
func Shell() {

	ldaps.InitLDAP()
	defer ld.Close()

	for _, c := range coms {
		list = append(list, c.Name)
		list = append(list, c.ShortName)
		//		for _, s := range c.SubCommands {
		//			list = append(list, s.Name)
		//			list = append(list, s.ShortName)
		//		}
	}

	prompt := "> "
	matches = make([]string, 0, len(list))
L:
	for {

		found = "no"
		readline.SetCompletionEntryFunction(CompletionEntry)
		readline.SetAttemptedCompletionFunction(nil)
		result := readline.ReadLine(&prompt)
		if result == nil {
			break L
		}

		input := *result
		words := strings.Fields(input)
		if len(words) > 0 {
			readline.AddHistory(input)
			if coms.HasCommand(words[0]) && len(words) < 2 {
				cmd := coms.NameIs(words[0])
				cmd.Action()
			} else if coms.HasCommand(words[0]) && len(words) == 2 {
				for _, i := range coms {
					if i.SubCommands.HasCommand(words[1]) {
						cmd := i.SubCommands.NameIs(words[1])
						cmd.Action()
					}
				}
			}
		}
	}

	return
}
