package null

import (
	"database/sql"
	"database/sql/driver"
	"reflect"
	"strconv"
	"time"
)

func String(v *string) ValueScanner   { return &NullString{String_: v} }
func Uint8(v *uint8) ValueScanner     { return &NullUint8{Uint8: v} }
func Byte(v *byte) ValueScanner       { return &NullByte{Byte: v} }
func Int8(v *int8) ValueScanner       { return &NullInt8{Int8: v} }
func Uint16(v *uint16) ValueScanner   { return &NullUint16{Uint16: v} }
func Int16(v *int16) ValueScanner     { return &NullInt16{Int16: v} }
func Uint32(v *uint32) ValueScanner   { return &NullUint32{Uint32: v} }
func Int32(v *int32) ValueScanner     { return &NullInt32{Int32: v} }
func Rune(v *rune) ValueScanner       { return &NullRune{Rune: v} }
func Int(v *int) ValueScanner         { return &NullInt{Int: v} }
func Uint64(v *uint64) ValueScanner   { return &NullUint64{Uint64: v} }
func Int64(v *int64) ValueScanner     { return &NullInt64{Int64: v} }
func Float32(v *float32) ValueScanner { return &NullFloat32{Float32: v} }
func Float64(v *float64) ValueScanner { return &NullFloat64{Float64: v} }
func Time(v *time.Time) ValueScanner  { return &NullTime{Time: v} }

type ValueScanner interface {
	driver.Valuer
	sql.Scanner
}

// NullString represents a string that may be null.
// NullString implements the Scanner interface so
// it can be used as a scan destination:
//
//  var plain string
//  err := db.QueryRow("SELECT name FROM foo WHERE id=?", id).Scan(&String{S:&s})
//  ...
//  use plain if database return null, plain is blank
type NullString struct {
	String_ *string
}

func (n *NullString) String() string {
	if n.String_ != nil {
		return *n.String_
	}
	return ""
}

// Scan implements the Scanner interface.
func (s *NullString) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case []byte:
		*s.String_ = string(v)
	case *[]byte:
		*s.String_ = string(*v)
	default:
		panic("unsupported type " + reflect.TypeOf(v).String())
	}
	return nil
}

// Value implements the driver Valuer interface.
func (s NullString) Value() (driver.Value, error) {
	if s.String_ == nil {
		return nil, nil
	}
	if *s.String_ == "" {
		return nil, nil
	}
	return *s.String_, nil
}

type NullInt struct {
	Int *int
}

func (n *NullInt) String() string {
	if n.Int != nil {
		return strconv.FormatInt(int64(*n.Int), 10)
	}
	return "0"
}
func (n *NullInt) Scan(value interface{}) error {
	if value != nil {
		*n.Int = int(scan(value))
	}
	return nil
}
func (n NullInt) Value() (driver.Value, error) {
	if n.Int == nil {
		return nil, nil
	}
	return int64(*n.Int), nil
}

type NullUint8 struct {
	Uint8 *uint8
}

func (n *NullUint8) String() string {
	if n.Uint8 != nil {
		return strconv.FormatInt(int64(*n.Uint8), 10)
	}
	return "0"
}
func (n *NullUint8) Scan(value interface{}) error {
	if value != nil {
		*n.Uint8 = uint8(scan(value))
	}
	return nil
}
func (n NullUint8) Value() (driver.Value, error) {
	if n.Uint8 == nil {
		return nil, nil
	}
	return int64(*n.Uint8), nil
}

type NullByte struct {
	Byte *byte
}

func (n *NullByte) String() string {
	if n.Byte != nil {
		return strconv.FormatInt(int64(*n.Byte), 10)
	}
	return "0"
}
func (n *NullByte) Scan(value interface{}) error {
	if value != nil {
		*n.Byte = byte(scan(value))
	}
	return nil
}
func (n NullByte) Value() (driver.Value, error) {
	if n.Byte == nil {
		return nil, nil
	}
	return int64(*n.Byte), nil
}

type NullInt8 struct {
	Int8 *int8
}

func (n *NullInt8) String() string {
	if n.Int8 != nil {
		return strconv.FormatInt(int64(*n.Int8), 10)
	}
	return "0"
}
func (n *NullInt8) Scan(value interface{}) error {
	if value != nil {
		*n.Int8 = int8(scan(value))
	}
	return nil
}
func (n NullInt8) Value() (driver.Value, error) {
	if n.Int8 == nil {
		return nil, nil
	}
	return int64(*n.Int8), nil
}

type NullUint16 struct {
	Uint16 *uint16
}

func (n *NullUint16) String() string {
	if n.Uint16 != nil {
		return strconv.FormatInt(int64(*n.Uint16), 10)
	}
	return "0"
}
func (n *NullUint16) Scan(value interface{}) error {
	if value != nil {
		*n.Uint16 = uint16(scan(value))
	}
	return nil
}
func (n NullUint16) Value() (driver.Value, error) {
	if n.Uint16 == nil {
		return nil, nil
	}
	return int64(*n.Uint16), nil
}

type NullInt16 struct {
	Int16 *int16
}

func (n *NullInt16) String() string {
	if n.Int16 != nil {
		return strconv.FormatInt(int64(*n.Int16), 10)
	}
	return "0"
}
func (n *NullInt16) Scan(value interface{}) error {
	if value != nil {
		*n.Int16 = int16(scan(value))
	}
	return nil
}
func (n NullInt16) Value() (driver.Value, error) {
	if n.Int16 == nil {
		return nil, nil
	}
	return int64(*n.Int16), nil
}

type NullUint32 struct {
	Uint32 *uint32
}

func (n *NullUint32) String() string {
	if n.Uint32 != nil {
		return strconv.FormatInt(int64(*n.Uint32), 10)
	}
	return "0"
}
func (n *NullUint32) Scan(value interface{}) error {
	if value != nil {
		*n.Uint32 = uint32(scan(value))
	}
	return nil
}
func (n NullUint32) Value() (driver.Value, error) {
	if n.Uint32 == nil {
		return nil, nil
	}
	return int64(*n.Uint32), nil
}

type NullInt32 struct {
	Int32 *int32
}

func (n *NullInt32) String() string {
	if n.Int32 != nil {
		return strconv.FormatInt(int64(*n.Int32), 10)
	}
	return "0"
}
func (n *NullInt32) Scan(value interface{}) error {
	if value != nil {
		*n.Int32 = int32(scan(value))
	}
	return nil
}
func (n NullInt32) Value() (driver.Value, error) {
	if n.Int32 == nil {
		return nil, nil
	}
	return int64(*n.Int32), nil
}

type NullRune struct {
	Rune *rune
}

func (n *NullRune) String() string {
	if n.Rune != nil {
		return strconv.FormatInt(int64(*n.Rune), 10)
	}
	return "0"
}
func (n *NullRune) Scan(value interface{}) error {
	if value != nil {
		*n.Rune = rune(scan(value))
	}
	return nil
}
func (n NullRune) Value() (driver.Value, error) {
	if n.Rune == nil {
		return nil, nil
	}
	return int64(*n.Rune), nil
}

type NullUint64 struct {
	Uint64 *uint64
}

func (n *NullUint64) String() string {
	if n.Uint64 != nil {
		return strconv.FormatInt(int64(*n.Uint64), 10)
	}
	return "0"
}
func (n *NullUint64) Scan(value interface{}) error {
	if value != nil {
		*n.Uint64 = uint64(scan(value))
	}
	return nil
}
func (n NullUint64) Value() (driver.Value, error) {
	if n.Uint64 == nil {
		return nil, nil
	}
	return int64(*n.Uint64), nil
}

type NullInt64 struct {
	Int64 *int64
}

func (n *NullInt64) String() string {
	if n.Int64 != nil {
		return strconv.FormatInt(int64(*n.Int64), 10)
	}
	return "0"
}
func (n *NullInt64) Scan(value interface{}) error {
	if value != nil {
		*n.Int64 = scan(value)
	}
	return nil
}
func (n NullInt64) Value() (driver.Value, error) {
	if n.Int64 == nil {
		return nil, nil
	}
	return *n.Int64, nil
}

type NullFloat32 struct {
	Float32 *float32
}

func (n *NullFloat32) String() string {
	if n.Float32 != nil {
		return strconv.FormatFloat(float64(*n.Float32), 'e', 2, 32)
	}
	return "0"
}
func (n *NullFloat32) Scan(value interface{}) error {
	if value != nil {
		switch v := value.(type) {
		case float32:
			*n.Float32 = v
		case *float32:
			*n.Float32 = *v
		default:
			panic("unsupported type " + reflect.TypeOf(v).String())
		}
	}
	return nil
}
func (n NullFloat32) Value() (driver.Value, error) {
	if n.Float32 == nil {
		return nil, nil
	}
	return float64(*n.Float32), nil
}

type NullFloat64 struct {
	Float64 *float64
}

func (n *NullFloat64) String() string {
	if n.Float64 != nil {
		return strconv.FormatFloat(*n.Float64, 'e', 2, 32)
	}
	return "0"
}
func (n *NullFloat64) Scan(value interface{}) error {
	if value != nil {
		switch v := value.(type) {
		case float64:
			*n.Float64 = v
		case *float64:
			*n.Float64 = *v
		default:
			panic("unsupported type " + reflect.TypeOf(v).String())
		}
	}
	return nil
}
func (n NullFloat64) Value() (driver.Value, error) {
	if n.Float64 == nil {
		return nil, nil
	}
	return *n.Float64, nil
}

func scan(value interface{}) int64 {
	switch v := value.(type) {
	case int64:
		return v
	case *int64:
		return *v
	default:
		panic("unsupported type " + reflect.TypeOf(v).String())
	}
}

type NullTime struct {
	Time *time.Time
}

func (n *NullTime) String() string {
	if n.Time != nil {
		return n.Time.Format("2006-01-02 15:04:05.999")
	}
	return "0000-00-00 00:00:00.000"
}
func (n *NullTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*n.Time = v

	case *time.Time:
		*n.Time = *v

	default:
		panic("unsupported type " + reflect.TypeOf(v).String())
	}
	return nil
}
func (n NullTime) Value() (driver.Value, error) {
	if n.Time == nil {
		return nil, nil
	}
	return n.Time, nil
}
