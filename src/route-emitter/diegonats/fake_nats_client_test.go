package diegonats_test

import (
	"testing"

	"code.cloudfoundry.org/diego-release/route-emitter/diegonats"
)

func FunctionTakingNATSClient(diegonats.NATSClient) {

}

func TestCanPassFakeYagnatsAsNATSClient(t *testing.T) {
	FunctionTakingNATSClient(diegonats.NewFakeClient())
}
