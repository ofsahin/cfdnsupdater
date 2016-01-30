package core

import (
	"encoding/json"
	"errors"
	"fmt"
)

type RecordType int

type RecordTypeSlice []RecordType

const (
	A     RecordType = 0 << iota
	AAAA  RecordType = 1
	CNAME RecordType = 2
	MX    RecordType = 3
	LOC   RecordType = 4
	SRV   RecordType = 5
	SPF   RecordType = 6
	TXT   RecordType = 7
	NS    RecordType = 8
)

func (self *RecordType) UnmarshalJSON(b []byte) (err error) {
	var val string
	if err = json.Unmarshal(b, &val); err == nil {
		*self, err = FromString(val)
		return
	}
	return
}

func (self RecordType) String() (string, error) {
	var conversionArray = []string{"A", "AAAA", "CNAME", "MX", "LOC", "SRV", "SPF", "TXT", "NS"}
	if int(self) < len(conversionArray) {
		return conversionArray[self], nil
	} else {
		return string(""), errors.New("Invalid type of record")
	}
}

func (self RecordTypeSlice) Contains(valueToSearch RecordType) bool {
	for _, v := range self {
		if v == valueToSearch {
			return true
		}
	}
	return false
}

func FromString(valToConvert string) (res RecordType, err error) {
	var conversionArray = []string{"A", "AAAA", "CNAME", "MX", "LOC", "SRV", "SPF", "TXT", "NS"}
	for p, v := range conversionArray {
		if v == valToConvert {
			res = RecordType(p)
			err=nil
			return
		}
	}
	err = errors.New(fmt.Sprintf("'%s' is not a valid record type", valToConvert))
	return
}
