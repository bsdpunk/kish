package shell

import "fmt"
import "strings"

import "./commands"
import "./general"

import "./twil"
import "./stats"
import "./ldaps"
import "gopkg.in/ldap.v3"

import "github.com/gobs/readline"

var ld *ldap.Conn

var found string = "no"
var list []string

var matches = make([]string, 0, len(list))
var statsSubs = commands.Commands{
	{
		Name:      "PrintMemUsage",
		ShortName: "mem",
		Usage:     "Print Memory Usage of the Shell",
		Action:    stats.PrintMemUsage,
		Category:  "stats",
	},
}

var twilSubs = commands.Commands{
	{
		Name:      "send",
		ShortName: "send",
		Usage:     "Send Text Message",
		Action:    twil.SendText,
		Category:  "twil",
	},
}
var ldapSubs = commands.Commands{
	{
		Name:      "GetAllDNs",
		ShortName: "all",
		Usage:     "Get All DNs",
		Action:    ldaps.CmdGetAllDNs,
		Category:  "ldaps",
	},
	//	{
	//		Name:   "GetAllThirds",
	//		Usage:  "Get All DNs",
	//		Action: command.CmdGetAllThirds,
	//		Flags:  []cli.Flag{},
	//	},

	//	{
	//		Name:   "GetAllAttr",
	//		Usage:  "Get All Attributes",
	//		Action: command.CmdGetAllAttr,
	//		Flags:  []cli.Flag{},
	//	},
	{
		Name:      "Search",
		ShortName: "search",
		Usage:     "Search LDAP with LDAP filter object",
		Action:    ldaps.CmdSearch,
		Category:  "ldaps",
	},
}
var coms = commands.Commands{
	{
		Name:        "LDAP",
		ShortName:   "ldap",
		Usage:       "ldap commands",
		Action:      NoAction,
		SubCommands: ldapSubs,
		Category:    "ldaps",
	},
	{
		Name:        "Twil",
		ShortName:   "twil",
		Usage:       "use the twilio api through the shell",
		Action:      NoAction,
		SubCommands: twilSubs,
		Category:    "twil",
	},
	{
		Name:        "stats",
		ShortName:   "stats",
		Usage:       "stats commands",
		Action:      NoAction,
		SubCommands: statsSubs,
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
	if start == 0 { // this is the command to match
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
	fmt.Println("Command not found")
	//	fmt.Printf("%+v", c)

}
func Shell() string {

	ldaps.InitLDAP()

	for _, c := range coms {
		list = append(list, c.Name)
		list = append(list, c.ShortName)
		for _, s := range c.SubCommands {
			list = append(list, s.Name)
			list = append(list, s.ShortName)
		}
	}

	//	reader := bufio.NewReader(os.Stdin)
	prompt := "> "
	matches = make([]string, 0, len(list))
L:
	for {

		found = "no"
		readline.SetCompletionEntryFunction(CompletionEntry)
		readline.SetAttemptedCompletionFunction(nil)
		result := readline.ReadLine(&prompt)
		if result == nil { // exit loop
			break L
		}

		input := *result
		//input = strings.TrimSpace(input)
		//		fmt.Print("$ ")
		//		text, _ := reader.ReadString('\n')
		//		text = strings.Replace(text, "\n", "", -1)
		words := strings.Fields(input)
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
		//		switch {
		//		case words[0] == "mem":
		//			stats.PrintMemUsage()
		//		case (words[0] == "quit" || words[0] == "exit"):
		//			os.Exit(3)
		//		case words[0] == "twil":
		//			fmt.Println(twil.SendText(words[1]))
		//		case words[0] == "clear":
		//			fmt.Println("\033[2J")
		//		case words[0] == "ldap":
		//			if len(words) > 1 {
		//				switch {
		//				case words[1] == "all":
		//					ldaps.CmdGetAllDNs()
		//				case words[1] == "search":
		//					ldaps.CmdSearch()
		//				default:
		//					fmt.Println("No Valid ldap Action")
		//				}
		//			} else {
		//				fmt.Println("No ldap Action Given")
		//			}
		//		default:
		//			fmt.Println(words)
		//		}
	}

	return "exit"
}
