package onboard

import (
	"log"
	"net/url"
	"os"

	"github.com/NickPresta/GoURLShortener"
	"github.com/itsabot/abot/shared/datatypes"
	"github.com/itsabot/abot/shared/nlp"
	"github.com/itsabot/abot/shared/plugin"
)

var p *dt.Plugin

func init() {
	trigger := &nlp.StructuredInput{
		Commands: []string{"onboard"},
		Objects:  []string{"onboard"},
	}
	fns := &dt.PluginFns{Run: Run, FollowUp: FollowUp}
	var err error
	p, err = plugin.New("github.com/itsabot/plugin_onboard", trigger, fns)
	if err != nil {
		log.Fatal("building", err)
	}
}

func Run(in *dt.Msg) (string, error) {
	u, err := getURL(in)
	if err != nil {
		return "Something went wrong. I'll figure it out. :(", err
	}
	return "Hi, I'm Abot, your new personal assistant. To get started, please sign up here: " + u, nil
}

func FollowUp(in *dt.Msg) (string, error) {
	u, err := getURL(in)
	if err != nil {
		return "", err
	}
	return "Hi, I'm Abot. To get started, you can sign up here: " + u, nil
}

// TODO fix, passing in flexid to msg
func getURL(in *dt.Msg) (string, error) {
	fid := in.FlexID
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
