package enums

type Status int32

const (
	StatusUnknow Status = iota
	StatusNormal
	StatusDeleted
)

func (s Status) String() string {
	return "Normal"
}
