package standardtype

//StandardType is used for VIN identification
type StandardType int

const (
	ISO3779 StandardType = iota
	EU500More
	EU500Less
	NA2000More
	NA2000Less
	Pre1980
)

var standardTypes = [...]string{
	"ISO3779",
	"EU500More",
	"EU500Less",
	"NA2000More",
	"NA2000Less",
	"Pre1980"}

func (s StandardType) String() string {
	return standardTypes[s]
}
