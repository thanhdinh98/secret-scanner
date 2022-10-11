package database

import (
	"time"

	"github.com/jackc/pgtype"
)

func Text(v string) pgtype.Text {
	return pgtype.Text{Status: pgtype.Present, String: v}
}

func Int2(v int16) pgtype.Int2 {
	return pgtype.Int2{Int: v, Status: pgtype.Present}
}

func Int4(v int32) pgtype.Int4 {
	return pgtype.Int4{Int: v, Status: pgtype.Present}
}

func Int8(v int64) pgtype.Int8 {
	return pgtype.Int8{Int: v, Status: pgtype.Present}
}

func Float4(v float32) pgtype.Float4 {
	return pgtype.Float4{Float: v, Status: pgtype.Present}
}

func Float8(v float64) pgtype.Float8 {
	return pgtype.Float8{Float: v, Status: pgtype.Present}
}

func Bool(v bool) pgtype.Bool {
	return pgtype.Bool{Bool: v, Status: pgtype.Present}
}

func Int4Array(v []int32) pgtype.Int4Array {
	if v == nil {
		return pgtype.Int4Array{Status: pgtype.Null}
	}
	if len(v) == 0 {
		return pgtype.Int4Array{Status: pgtype.Present}
	}
	elements := make([]pgtype.Int4, len(v))
	for idx := range v {
		elements[idx] = Int4(v[idx])
	}
	return pgtype.Int4Array{
		Elements:   elements,
		Dimensions: []pgtype.ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
		Status:     pgtype.Present,
	}
}

func TextArray(v []string) pgtype.TextArray {
	if v == nil {
		return pgtype.TextArray{Status: pgtype.Null}
	}
	if len(v) == 0 {
		return pgtype.TextArray{Status: pgtype.Present}
	}
	elements := make([]pgtype.Text, len(v))
	for i := range v {
		elements[i] = Text(v[i])
	}
	return pgtype.TextArray{
		Elements:   elements,
		Dimensions: []pgtype.ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
		Status:     pgtype.Present,
	}
}

func JSONB(v interface{}) pgtype.JSONB {
	j := pgtype.JSONB{}
	_ = j.Set(v)
	return j
}

func Timestamptz(v time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{Time: v, Status: pgtype.Present}
}
