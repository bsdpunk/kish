package ldaps

import (
	"../sets"
	"fmt"
	"gopkg.in/ldap.v3"
	"os"
	"strings"
	//	"../Explore"
	//"github.com/urfave/cli"
	"../commands"
	"bufio"
	"log"
)

var (
	//ldapServer = "ds.trozlabs.local:389"
	ldapServer   = string(os.Getenv("LDAPServer"))
	ldapBind     = "CN=Administrator,CN=Users,DC=trozlabs,DC=local"
	ldapPassword = string(os.Getenv("LDAPPassword"))

	filterDN      = "(objectClass=*)"
	baseDN        = string(os.Getenv("LDAPBase"))
	loginUsername = string(os.Getenv("LDAPUser"))
	loginPassword = string(os.Getenv("LDAPPassword"))
)
var ld *ldap.Conn

var LdapSubs = commands.Commands{
	{
		Name:      "GetAllDNs",
		ShortName: "all",
		Usage:     "Get All DNs",
		Action:    CmdGetAllDNs,
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
		Action:    CmdSearch,
		Category:  "ldaps",
	},
}

//func init() (*ldap.Conn, error) {
//tlsConfig := &tls.Config{InsecureSkipVerify: true}
func InitLDAP() (*ldap.Conn, error) {
	//tlsConfig := &tls.Config{InsecureSkipVerify: true}

	conn, err := ldap.Dial("tcp", ldapServer)

	if err != nil {
		return nil, fmt.Errorf("Failed to connect. %s", err)
	}

	if err := conn.Bind(ldapBind, ldapPassword); err != nil {
		return nil, fmt.Errorf("Failed to bind. %s", err)
	}

	ld = conn
	return ld, nil
}

//	return ld
//}

//func CmdGetAllAttr(c *cli.Context) {
//	l := ldap.NewSearchRequest(
//		baseDN,
//		ldap.ScopeWholeSubtree,
//		ldap.NeverDerefAliases,
//		0,
//		0,
//		false,
//		"(objectClass=*)",
//		[]string{},
//		nil,
//	)
//	sr, err := ld.Search(l)
//	if err != nil {
//		log.Fatal(err)
//	}
//	ns := sets.NewSet()
//	for _, entry := range sr.Entries {
//		//entry.PrettyPrint(1)
//		//fmt.Println(entry.DN)
//		for _, attr := range entry.Attributes {
//			//fmt.Println(attr.Name)
//			ns.Add(attr.Name)
//		}
//
//	}
//
//	//var removeErrant = regexp.MustCompile(`[a-zA-Z: 0-9=]+`)
//
//	//theSplit := strings.Split(strings.Join(removeErrant.FindAllString(fil, -1), ""), "=")
//
//	//fmt.Println(strings.Join(theSplit, ", "))
//	ns.PrintAll()
//	//	return ns
//}
func CmdGetAllDNs() {
	//conn = ld
	l := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(objectClass=*)",
		[]string{},
		nil,
	)
	//fmt.Println(c.Ldap())
	//fmt.Println(c)
	//conn := c.Conn
	fmt.Println(ld)
	sr, err := ld.Search(l)
	if err != nil {
		fmt.Println(err)
	}
	ns := sets.NewSet()
	for _, entry := range sr.Entries {
		//entry.PrettyPrint(1)
		ns.Add(entry.DN)
		fmt.Println(entry.DN)
	}

	//var removeErrant = regexp.MustCompile(`[a-zA-Z: 0-9=]+`)

	//theSplit := strings.Split(strings.Join(removeErrant.FindAllString(fil, -1), ""), "=")

	//fmt.Println(strings.Join(theSplit, ", "))
	ns.PrintAll()
}

func CmdSearch() {
	//	fmt.Println(c.Args())
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("ex(objectClass=*) Search$ ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	l := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		text,
		[]string{},
		nil,
	)
	sr, err := ld.Search(l)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range sr.Entries {
		entry.Print()
		fmt.Println("")
	}
	return
}

//func CmdGetAllThirds(c *cli.Context) {
//	//conn = ld
//	l := ldap.NewSearchRequest(
//		baseDN,
//		ldap.ScopeWholeSubtree,
//		ldap.NeverDerefAliases,
//		0,
//		0,
//		false,
//		"(objectClass=*)",
//		[]string{},
//		nil,
//	)
//	//fmt.Println(c.Ldap())
//	//fmt.Println(c)
//	//conn := c.Conn
//	sr, err := ld.Search(l)
//	if err != nil {
//		log.Fatal(err)
//	}
//	ns := sets.NewSet()
//	for _, entry := range sr.Entries {
//		//entry.PrettyPrint(1)
//		ns.Add(entry.DN)
//		//	fmt.Println(entry.DN)
//	}
//
//	ns.PrintThird()
//}
//
func Explore() (ns *sets.Set) {
	//conn = ld
	l := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(objectClass=*)",
		[]string{},
		nil,
	)
	//fmt.Println(c.Ldap())
	//fmt.Println(c)
	//conn := c.Conn
	sr, err := ld.Search(l)
	if err != nil {
		log.Fatal(err)
	}
	s := sets.NewSet()
	for _, entry := range sr.Entries {
		//entry.PrettyPrint(1)
		s.Add(entry.DN)
		//	fmt.Println(entry.DN)
	}

	//ns.PrintThird()
	return s
}
