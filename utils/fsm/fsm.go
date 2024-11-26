package fsm

import (
	"github.com/sirikothe/gotextfsm"
	"log"
)

func ParseFsm(data, template string) ([]map[string]interface{}, error) {
	fsm := gotextfsm.TextFSM{}
	err := fsm.ParseString(template)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	parser := gotextfsm.ParserOutput{}
	err = parser.ParseTextString(data, fsm, true)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return parser.Dict, nil
}
