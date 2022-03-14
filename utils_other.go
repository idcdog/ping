package ping

// +build !linux,!windows

// Returns the length of an ICMP message.
func (p *Pinger) getMessageLength() int {
	switch p.PingType {
	case "echo":
		// 24+8
		return p.Size+8
	case "timestamp":
		return 20
	default:
		return p.Size+8
	}
}

// Attempts to match the ID of an ICMP packet.
func (p *Pinger) matchID(ID int) bool {
	if p.protocol == "icmp" {
		if ID != p.id {
			return false
		}
	}
	return true
}