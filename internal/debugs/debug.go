package debugs

import "fmt"

var (
	NumTransferRequestSend          int
	NumTransferRequestReceived      int
	NumTransferResponseSend         int
	NumTransferResponseReceived     int
	NumTransferResponseAwaitTimeout int
)

func LogNums() {
	fmt.Printf("NumTransferRequestSend: %d\n", NumTransferRequestSend)
	fmt.Printf("NumTransferRequestReceived: %d\n", NumTransferRequestReceived)
	fmt.Printf("NumTransferResponseSend: %d\n", NumTransferResponseSend)
	fmt.Printf("NumTransferResponseReceived: %d\n\n", NumTransferResponseReceived)

	fmt.Printf("NumTransferResponseAwaitTimeout: %d\n\n", NumTransferResponseAwaitTimeout)
}
