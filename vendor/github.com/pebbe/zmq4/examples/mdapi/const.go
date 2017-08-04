// Majordomo Protocol Client and Worker API.
// Implements the MDP/Worker spec at http://rfc.zeromq.org/spec:7.
package mdapi

const (
	//  This is the version of MDP/Client we implement
	MDPC_CLIENT = "MDPC01"

	//  This is the version of MDP/Worker we implement
	MDPW_WORKER = "MDPW01"
)

const (
	//  MDP/Server commands, as strings
	MDPW_READY = string(iota + 1)
	MDPW_REQUEST
	MDPW_REPLY
	MDPW_HEARTBEAT
	MDPW_DISCONNECT
)

var (
	MDPS_COMMANDS = map[string]string{
		MDPW_READY:      "READY",
		MDPW_REQUEST:    "REQUEST",
		MDPW_REPLY:      "REPLY",
		MDPW_HEARTBEAT:  "HEARTBEAT",
		MDPW_DISCONNECT: "DISCONNECT",
	}
)
