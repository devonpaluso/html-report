// Copyright 2015 ThoughtWorks, Inc.

// This file is part of getgauge/html-report.

// getgauge/html-report is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// getgauge/html-report is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with getgauge/html-report.  If not, see <http://www.gnu.org/licenses/>.

package generator

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/getgauge/html-report/gauge_messages"
	"github.com/golang/protobuf/proto"
)

var passSpecRes1 = &gauge_messages.ProtoSpecResult{
	Failed:        proto.Bool(false),
	Skipped:       proto.Bool(false),
	ExecutionTime: proto.Int64(211316),
	ProtoSpec: &gauge_messages.ProtoSpec{
		SpecHeading: proto.String("Passing Specification 1"),
		Tags:        []string{"tag1", "tag2"},
		FileName:    proto.String("/tmp/gauge/specs/foobar.spec"),
	},
}

var passSpecRes2 = &gauge_messages.ProtoSpecResult{
	Failed:        proto.Bool(false),
	Skipped:       proto.Bool(false),
	ExecutionTime: proto.Int64(211316),
	ProtoSpec: &gauge_messages.ProtoSpec{
		SpecHeading: proto.String("Passing Specification 2"),
		Tags:        []string{},
	},
}

var passSpecRes3 = &gauge_messages.ProtoSpecResult{
	Failed:        proto.Bool(false),
	Skipped:       proto.Bool(false),
	ExecutionTime: proto.Int64(211316),
	ProtoSpec: &gauge_messages.ProtoSpec{
		SpecHeading: proto.String("Passing Specification 3"),
		Tags:        []string{"foo"},
	},
}

var failSpecRes1 = &gauge_messages.ProtoSpecResult{
	Failed:        proto.Bool(true),
	Skipped:       proto.Bool(false),
	ExecutionTime: proto.Int64(0),
	ProtoSpec: &gauge_messages.ProtoSpec{
		SpecHeading: proto.String("Failing Specification 1"),
		Tags:        []string{},
	},
}

var skipSpecRes1 = &gauge_messages.ProtoSpecResult{
	Failed:        proto.Bool(false),
	Skipped:       proto.Bool(true),
	ExecutionTime: proto.Int64(0),
	ProtoSpec: &gauge_messages.ProtoSpec{
		SpecHeading: proto.String("Skipped Specification 1"),
		Tags:        []string{"bar"},
	},
}

var suiteRes = &gauge_messages.ProtoSuiteResult{
	SpecResults:       []*gauge_messages.ProtoSpecResult{passSpecRes1, passSpecRes2, passSpecRes3, failSpecRes1, skipSpecRes1},
	Failed:            proto.Bool(false),
	SpecsFailedCount:  proto.Int32(1),
	ExecutionTime:     proto.Int64(122609),
	SuccessRate:       proto.Float32(60),
	Environment:       proto.String("default"),
	Tags:              proto.String(""),
	ProjectName:       proto.String("Gauge Project"),
	Timestamp:         proto.String("Jul 13, 2016 at 11:49am"),
	SpecsSkippedCount: proto.Int32(1),
}

func TestHTMLGeneration(t *testing.T) {
	content, err := ioutil.ReadFile("_testdata/expected.html")
	if err != nil {
		t.Errorf("Error reading expected HTML file: %s", err.Error())
	}

	buf := new(bytes.Buffer)
	generate(suiteRes, buf)

	want := removeNewline(string(content))
	got := removeNewline(buf.String())

	if got != want {
		t.Errorf("want:\n%q\ngot:\n%q\n", want, got)
	}
}