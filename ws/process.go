package ws

func (b *ByBitWS) processOrderBook(ob *OrderBookV5, isSnapshot bool) { // ob []*OrderBookL2
	b.Emit(WSOrderBook50, ob, isSnapshot)
}

func (b *ByBitWS) processTrade(symbol string, data ...*Trade) {
	b.Emit(WSTrade, symbol, data)
}

func (b *ByBitWS) processKLine(symbol string, data KLine) {
	b.Emit(WSKLine, symbol, data)
}

func (b *ByBitWS) processInsurance(currency string, data ...*Insurance) {
	b.Emit(WSInsurance, currency, data)
}

func (b *ByBitWS) processInstrument(symbol string, data ...*Instrument) {
	b.Emit(WSInstrument, symbol, data)
}

func (b *ByBitWS) processPosition(data ...*Position) {
	b.Emit(WSPosition, data)
}

func (b *ByBitWS) processExecution(data ...*Execution) {
	b.Emit(WSExecution, data)
}

func (b *ByBitWS) processOrder(data ...*Order) {
	b.Emit(WSOrder, data)
}
