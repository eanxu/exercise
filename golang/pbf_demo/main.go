/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"pbf_demo/cmd"
	"time"
)

// @title mvt
// @version 1.0
// @description mvt

// @contact.name xuyi
// @contact.email xuyi@diit.cn
// @BasePath /test
func main() {
	cmd.Execute("pbf_demo", time.Now().String(), "v1.0.0")
}
