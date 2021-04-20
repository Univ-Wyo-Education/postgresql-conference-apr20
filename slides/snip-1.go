package main

import "fmt"

func MainProcessing() {
	for {
		select {
		case tx <- TxChannel:
			if !ValidateSignature(tx) {
				fmt.Printf(logFile, "Invalid Signature %s\n", tx.Signature)
			} else {
				P2pSendTx(tx)
				InsertTx(tx)
				if IAmARunner() {
					if CountTxByHash() >= 9 {
						TxMidpoint <- 1
					}
				}
			}
		case start <- RanChanel:
			if IAmARunner() {
				StartBlock()
			}
		case runTx <- TxMidpoint:
			ClearTimeout()
			if IAmARunner() {
				PickTxForBlock()
				RunTx()
				MerkelHashTx()
				SignTx()
				P2pSendTxBlock()
			}
		}
	}
}
