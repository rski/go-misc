package nbf

import (
	"testing"
)

func TestMessage_ParseFilename(t *testing.T) {
	const name = "0000186F3C52A89B0042201000500000004030000000000000000000000000000+336303132330000009F"
	var msg Message
	err := msg.ParseFilename(name)
	if err != nil {
		t.Fatal(err)
	}

	if msg.Seq != 0x186f {
		t.Errorf("bad sequence number: 0x%x", msg.Seq)
	}
	if msg.Timestamp != 0x3c52a89b {
		t.Errorf("bad timestamp: 0x%x", msg.Timestamp)
	}
	t.Logf("timestamp: %s", DosTime(msg.Timestamp))
	if msg.MultipartSeq != 0x42 {
		t.Errorf("bad multipart sequence number: %d", msg.MultipartSeq)
	}
	if msg.Flags != 0x2010 {
		t.Errorf("bad flags: 0x%x", msg.Flags)
	}
	if msg.PartNo != 3 || msg.PartTotal != 4 {
		t.Errorf("got part %d/%d, expected 3/4",
			msg.PartNo, msg.PartTotal)
	}
	if peer := string(msg.Peer[:]); peer != "+33630313233" {
		t.Errorf("wrong peer %s, expected +33630313233", peer)
	}
}

func TestDecode7bit(t *testing.T) {
	var data = []byte{0xd2, 0xf7, 0xfb, 0xfd, 0x7e, 0x83, 0xe8, 0x75, 0x90, 0xbd, 0x5c, 0xc7, 0x83,
		0xe2, 0xf5, 0x32, 0x48, 0x7d, 0x0a, 0xc3, 0xe1, 0x65, 0x36, 0xbb, 0xfc, 0x3}
	u := unpack7bit(data)
	t.Logf("in: %d bytes, out: %d septets", len(data), len(u))
	s := translateSMS(unpack7bit(data), &basicSMSset)
	const ref = "Rooooo tu veux que j'appelle?"
	if s != ref {
		t.Errorf("got %q, expected %q", s, ref)
	}
	t.Logf("%s", s)
}