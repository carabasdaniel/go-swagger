// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spec

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v2"
)

func TestIntegrationLicense(t *testing.T) {
	Convey("all fields of license should", t, func() {
		Convey("serialize to JSON", func() {
			b, err := json.Marshal(License{"the name", "the url"})
			So(err, ShouldBeNil)
			So(string(b), ShouldEqual, `{"name":"the name","url":"the url"}`)
		})

		Convey("serialize to YAML", func() {
			b, err := yaml.Marshal(License{"the name", "the url"})
			So(err, ShouldBeNil)
			So(string(b), ShouldEqual, "name: the name\nurl: the url\n")
		})

		Convey("deserialize from JSON", func() {
			actual := License{}
			err := json.Unmarshal([]byte(`{"name":"the name","url":"the url"}`), &actual)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, License{"the name", "the url"})
		})

		Convey("deserialize from YAML", func() {
			actual := License{}
			err := yaml.Unmarshal([]byte("name: the name\nurl: the url\n"), &actual)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, License{"the name", "the url"})
		})
	})
}
