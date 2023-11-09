package nft

//go:generate mockery --name NFTEventAdapter
type NFTEventAdapter interface {
	RunService() error
}
