package jobworker

import (
	"encoding/json"
	"errors"
	"iot/conf"
	"iot/lib/formatter"
	"iot/lib/sender"
	"iot/lib/sgu_utils"

	"github.com/StabbyCutyou/buffstreams"
	"github.com/revel/revel"
)

func ProcessPacket(task string, args ...interface{}) error {
	name, ok := args[0].(string)
	if !ok {
		return errors.New("Invalid Worker")
	}
	if name == "parse_sgu_packets" {
		client := buffstreams.Client{}
		err := formatter.GetStructFromInterface(args[1], &client)

		if err != nil {
			revel.ERROR.Println(err)
			return err
		}
		sgu_utils.ParseInputPackets(&client)
	} else if name == "send_response_packets" {
		incoming := conf.Incoming{}

		packet_type := args[1].(json.Number)

		err := formatter.GetStructFromInterface(args[2], &incoming)
		if err != nil {
			revel.ERROR.Println(err)
			return err
		}
		pack_type, _ := (packet_type.Int64())
		sender.SendResponsePacket(int(pack_type), incoming)
	} else if name == "send_3000" {
		sender.SendServerPacket(0x3000, args[1])
	}
	return nil
}