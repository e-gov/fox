package fox

type Statistics struct {
	TimeSinceLastOK 		int64	`json:"timeSinceLastOK"`
	TimeSinceLastNOK 		int64	`json:"timeSinceLastNOK"`
	ParallelRequestCount	int		`json:"currentRequestCount"`
	NodeName				string	`json:"nodeName"`
}
