package gidit

import (
	"../sets"
	"fmt"
	"os"
	"strings"
	//	"../Explore"
	//"github.com/urfave/cli"
	"../commands"
	"bufio"
	"github.com/nfnt/resize"
	"log"
	//	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strconv"
)

var GiditSubs = commands.Commands{
	{
		Name:      "ResizePng",
		ShortName: "ResizePng",
		Usage:     "Get All DNs",
		Action:    ResizePng,
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
func ResizePng() {

	// open "test.jpg"
	var width uint
	var height uint
	//var ratio uint
	file, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		panic("No file selected")
	}

	// decode jpeg into image.Image
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	file.Close()
	b := img.Bounds()
	if w, err := strconv.Atoi(os.Args[1]); err == nil {
		width = uint(w)
	} else {
		panic("No resize given")
	}
	if h, err := strconv.Atoi(os.Args[2]); err == nil {
		height = uint(h)
	} else {
		var tmp float64
		//height = int(float64(b.Dy)/float64(b.Dx)) * width
		//ratio := uint(0)
		tmp = float64(b.Max.Y) / float64(b.Max.X)
		height = uint(tmp) * width
	}
	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(width, height, img, resize.Lanczos3)
	//fmt.Printf("%+v\n", b)
	//fmt.Println("Original", b)

	//fmt.Println(m.Bounds())
	out, err := os.Create("new" + os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)
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

func getTwoFiles(re bool) (one string, two string) {

	reader := bufio.NewReader(os.Stdin)
	//fmt.Println(urlStr)
	fmt.Print("ImageFileOne$ ")
	to, _ := reader.ReadString('\n')
	to = strings.Replace(to, "\n", "", -1)
	if re {
		fmt.Print("Width: ")
		w, _ := reader.ReadString('\n')
		w = strings.Replace(to, "\n", "", -1)

		fmt.Print("Height: ")
		h, _ := reader.ReadString('\n')
		h = strings.Replace(to, "\n", "", -1)

		return to, w + "X" + h
	}
	fmt.Print("ImageFileTwo$ ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return to, text
}
