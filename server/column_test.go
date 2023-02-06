// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"testing"

	"github.com/pingcap/tidb/parser/mysql"
	"github.com/stretchr/testify/require"
)

func TestDumpColumn(t *testing.T) {
	info := ColumnInfo{
		Schema:       "testSchema",
		Table:        "testTable",
		OrgTable:     "testOrgTable",
		Name:         "testName",
		OrgName:      "testOrgName",
		ColumnLength: 1,
		Charset:      106,
		Flag:         0,
		Decimal:      1,
		Type:         14,
		DefaultValue: []byte{5, 2},
	}
	r := info.Dump(nil, nil)
	exp := []byte{0x3, 0x64, 0x65, 0x66, 0xa, 0x74, 0x65, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x9, 0x74, 0x65, 0x73, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0xc, 0x74, 0x65, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x8, 0x74, 0x65, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xb, 0x74, 0x65, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xc, 0x6a, 0x0, 0x1, 0x0, 0x0, 0x0, 0xe, 0x0, 0x0, 0x1, 0x0, 0x0}
	require.Equal(t, exp, r)

	require.Equal(t, uint16(mysql.SetFlag), dumpFlag(mysql.TypeSet, 0))
	require.Equal(t, uint16(mysql.EnumFlag), dumpFlag(mysql.TypeEnum, 0))
	require.Equal(t, uint16(0), dumpFlag(mysql.TypeString, 0))

	require.Equal(t, mysql.TypeString, dumpType(mysql.TypeSet))
	require.Equal(t, mysql.TypeString, dumpType(mysql.TypeEnum))
	require.Equal(t, mysql.TypeBit, dumpType(mysql.TypeBit))
}

func TestDumpColumnWithDefault(t *testing.T) {
	info := ColumnInfo{
		Schema:       "testSchema",
		Table:        "testTable",
		OrgTable:     "testOrgTable",
		Name:         "testName",
		OrgName:      "testOrgName",
		ColumnLength: 1,
		Charset:      106,
		Flag:         0,
		Decimal:      1,
		Type:         14,
		DefaultValue: "test",
	}
	r := info.DumpWithDefault(nil, nil)
	exp := []byte{0x3, 0x64, 0x65, 0x66, 0xa, 0x74, 0x65, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x9, 0x74, 0x65, 0x73, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0xc, 0x74, 0x65, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x8, 0x74, 0x65, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xb, 0x74, 0x65, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xc, 0x6a, 0x0, 0x1, 0x0, 0x0, 0x0, 0xe, 0x0, 0x0, 0x1, 0x0, 0x0, 0x4, 0x74, 0x65, 0x73, 0x74}
	require.Equal(t, exp, r)

	require.Equal(t, uint16(mysql.SetFlag), dumpFlag(mysql.TypeSet, 0))
	require.Equal(t, uint16(mysql.EnumFlag), dumpFlag(mysql.TypeEnum, 0))
	require.Equal(t, uint16(0), dumpFlag(mysql.TypeString, 0))

	require.Equal(t, mysql.TypeString, dumpType(mysql.TypeSet))
	require.Equal(t, mysql.TypeString, dumpType(mysql.TypeEnum))
	require.Equal(t, mysql.TypeBit, dumpType(mysql.TypeBit))
}

func TestColumnNameLimit(t *testing.T) {
	aLongName := make([]byte, 0, 300)
	for i := 0; i < 300; i++ {
		aLongName = append(aLongName, 'a')
	}
	info := ColumnInfo{
		Schema:       "testSchema",
		Table:        "testTable",
		OrgTable:     "testOrgTable",
		Name:         string(aLongName),
		OrgName:      "testOrgName",
		ColumnLength: 1,
		Charset:      106,
		Flag:         0,
		Decimal:      1,
		Type:         14,
		DefaultValue: []byte{5, 2},
	}
	r := info.Dump(nil, nil)
	exp := []byte{0x3, 0x64, 0x65, 0x66, 0xa, 0x74, 0x65, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x9, 0x74, 0x65, 0x73, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0xc, 0x74, 0x65, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x54, 0x61, 0x62, 0x6c, 0x65, 0xfc, 0x0, 0x1, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0xb, 0x74, 0x65, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xc, 0x6a, 0x0, 0x1, 0x0, 0x0, 0x0, 0xe, 0x0, 0x0, 0x1, 0x0, 0x0}
	require.Equal(t, exp, r)
}
