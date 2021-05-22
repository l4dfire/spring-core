/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package conf

import (
	"strings"

	"github.com/go-spring/spring-core/conf/prop"
	"github.com/go-spring/spring-core/conf/toml"
	"github.com/go-spring/spring-core/conf/yaml"
)

func init() {
	NewReader(prop.Read, "properties", "prop")
	NewReader(yaml.Read, "yaml", "yml")
	NewReader(toml.Read, "toml")
}

type Reader func([]byte) (map[string]interface{}, error)

// readers 属性读取器列表。
var readers = make(map[string]Reader)

// NewReader 注册属性读取器。
func NewReader(fn Reader, configTypes ...string) {
	for _, configType := range configTypes {
		readers[strings.ToLower(configType)] = fn
	}
}
