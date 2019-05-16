package shell

import "fmt"
import "bufio"
import "os"
import "strings"

import "./twil"
import "./stats"
import "./ldaps"
import "gopkg.in/ldap.v3"

var ld *ldap.Conn

func Shell() string {

	ldaps.InitLDAP()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		words := strings.Fields(text)
		switch {
		case words[0] == "mem":
			stats.PrintMemUsage()
		case (words[0] == "quit" || words[0] == "exit"):
			os.Exit(3)
		case words[0] == "twil":
			fmt.Println(twil.SendText(words[1]))
		case words[0] == "clear":
			fmt.Println("\033[2J")
		case words[0] == "ldap":
			switch {
			case words[1] == "all":
				ldaps.CmdGetAllDNs()
			case words[1] == "search":
				ldaps.CmdSearch()
			default:
				fmt.Println("No Valid ldap action")
			}
		default:
			fmt.Println(words)
		}
	}

	return "exit"
}
