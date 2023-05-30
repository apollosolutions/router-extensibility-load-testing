package graph

// Query returns QueryResolver implementation.
// func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
