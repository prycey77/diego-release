package auctionmetricemitterdelegate

import (
	"time"

	"code.cloudfoundry.org/auction/auctiontypes"
	loggingclient "code.cloudfoundry.org/diego-logging-client"
)

const (
	LRPAuctionsStartedCounter     = "AuctioneerLRPAuctionsStarted"
	LRPAuctionsFailedCounter      = "AuctioneerLRPAuctionsFailed"
	TaskAuctionStartedCounter     = "AuctioneerTaskAuctionsStarted"
	TaskAuctionsFailedCounter     = "AuctioneerTaskAuctionsFailed"
	FetchStatesDuration           = "AuctioneerFetchStatesDuration"
	FailedCellStateRequestCounter = "AuctioneerFailedCellStateRequests"
)

type auctionMetricEmitterDelegate struct {
	metronClient loggingclient.IngressClient
}

func New(metronClient loggingclient.IngressClient) auctionMetricEmitterDelegate {
	return auctionMetricEmitterDelegate{
		metronClient: metronClient,
	}
}

func (d auctionMetricEmitterDelegate) FetchStatesCompleted(fetchStatesDuration time.Duration) error {
	return d.metronClient.SendDuration(FetchStatesDuration, fetchStatesDuration)
}

func (d auctionMetricEmitterDelegate) FailedCellStateRequest() {
	d.metronClient.IncrementCounter(FailedCellStateRequestCounter)
}

func (d auctionMetricEmitterDelegate) AuctionCompleted(results auctiontypes.AuctionResults) {
	d.metronClient.IncrementCounterWithDelta(LRPAuctionsStartedCounter, uint64(len(results.SuccessfulLRPs)))
	d.metronClient.IncrementCounterWithDelta(TaskAuctionStartedCounter, uint64(len(results.SuccessfulTasks)))

	d.metronClient.IncrementCounterWithDelta(LRPAuctionsFailedCounter, uint64(len(results.FailedLRPs)))
	d.metronClient.IncrementCounterWithDelta(TaskAuctionsFailedCounter, uint64(len(results.FailedTasks)))
}
