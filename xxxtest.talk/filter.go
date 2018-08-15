func Filter(ctx Context, src <-chan Item, allow func(item Item) bool) <-chan Item {
	dst := make(chan Item)
	go runFilter(ctx, src, dst, allow)
}


	replyc := make(chan Answer)
	for item := range src {
		if item.Dir != nil {
			if !allow(item) {
				item.Reply <- Next
				continue
			}
		}
		dst <- item.WithReply(replyc)
		r := <-replyc
		item.Reply <- r
		switch r {
		case Quit:
			return
		case Next:
		case Skip:
		}
	}
}
