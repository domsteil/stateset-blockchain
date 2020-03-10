package accountimport (
	"github.com/stateset/stateset-blockchain/x/bank/exported"
	"github.com/stateset/stateset-blockchain/x/distribution"
)

const (
	TransactionGift    = exported.TransactionGift
	TransactionBacking = exported.TransactionBacking

	UserGrowthPoolName = distribution.UserGrowthPoolName
)

var (
	FromModuleAccount = exported.FromModuleAccount
)