package main

import (
	"flag"
	"net/url"
	"os"

	"github.com/NickPresta/GoURLShortener"
	"github.com/itsabot/abot/shared/datatypes"
	"github.com/itsabot/abot/shared/log"
	"github.com/itsabot/abot/shared/nlp"
	"github.com/itsabot/abot/shared/pkg"
)

var p *pkg.Pkg

type Onboard string

const pkgName string = "onboard"

func main() {
	var coreaddr string
	flag.StringVar(&coreaddr, "coreaddr", "",
		"Port used to communicate with Abot.")
	flag.Parse()
	l := log.New(pkgName)
	trigger := &nlp.StructuredInput{
		Commands: []string{"onboard"},
		Objects:  []string{"onboard"},
	}
	var err error
	p, err = pkg.NewPackage(pkgName, coreaddr, trigger)
	if err != nil {
		l.Fatal("building", err)
	}
	onboard := new(Onboard)
	if err := p.Register(onboard); err != nil {
		l.Fatal("registering", err)
	}
}

func (t *Onboard) Run(m *dt.Msg, resp *string) error {
	u, err := getURL(m)
	if err != nil {
		return err
	}
	*resp = "Hi, I'm Abot, your new personal assistant. To get started, please sign up here: " + u
	return nil
}

func (t *Onboard) FollowUp(m *dt.Msg, resp *string) error {
	u, err := getURL(m)
	if err != nil {
		return err
	}
	*resp = "Hi, I'm Abot. To get started, you can sign up here: " + u
	return nil
}

// TODO fix, passing in flexid to msg
func getURL(m *dt.Msg) (string, error) {
	fid := m.FlexID
	v := url.Values{
		"fid": {fid},
	}
	u := os.Getenv("ABOT_URL") + "/signup?" + v.Encode()
	u, err := goisgd.Shorten(u)
	if err != nil {
		return "", err
	}
	return u, nil
}
