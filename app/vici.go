package app

import (
	"log"

	"github.com/strongswan/govici/vici"
)

func listSAs() ([]LoadedIKE, error) {
	s, err := vici.NewSession()
	if err != nil {
		log.Printf("Error Connecting to vici: %s", err)
		return nil, err
	}
	defer s.Close()

	var retVar []LoadedIKE
	msgs, err := s.StreamedCommandRequest("list-sas", "list-sa", nil)
	if err != nil {
		return retVar, err
	}

	for _, m := range msgs.Messages() { // <- Directly iterate over msgs
		if e := m.Err(); e != nil {
			// ignoring this error
			continue
		}

		for _, k := range m.Keys() {
			inbound, ok := m.Get(k).(*vici.Message)
			if !ok {
				continue
			}

			var ike LoadedIKE
			if e := vici.UnmarshalMessage(inbound, &ike); e != nil {
				// ignoring this marshal/unmarshal error!
				continue
			}
			ike.Name = k
			retVar = append(retVar, ike)
		}
	}

	return retVar, nil
}
