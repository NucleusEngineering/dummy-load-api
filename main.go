// Copyright 2023 Google Inc. All Rights Reserved.
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

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResponseStruct struct {
	Status string
}

var (
	LISTEN_PORT = 8080
)

func main() {
	r := setupRouter()

	log.Printf("Dummie workload running on :%d\n", LISTEN_PORT)

	if err := r.Run(fmt.Sprintf(":%d", LISTEN_PORT)); err != nil {
		log.Fatalf("could not run gin router: %s", err)
		return
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	if err := r.SetTrustedProxies(nil); err != nil {
		log.Fatalf("Could not set trusted proxies: %s", err)
	}

	r.GET("/", handleRoot)
	r.GET("/load", dummyLoad)

	return r
}

func handleRoot(c *gin.Context) {
	c.JSON(http.StatusOK, ResponseStruct{Status: "OK"})
}

func dummyLoad(c *gin.Context) {
	// 1 core default
	cores, _ := strconv.Atoi(c.DefaultQuery("cores", "1"))

	// 100ms default
	time, _ := strconv.Atoi(c.DefaultQuery("time", "100"))

	// 5% cpu usage
	percentage, _ := strconv.Atoi(c.DefaultQuery("percentage", "5"))

	// 128mb of ram
	ram, _ := strconv.Atoi(c.DefaultQuery("ram", "128"))

	GenerateCPUAndMemoryLoad(cores, time, percentage, ram)

	c.JSON(http.StatusOK, ResponseStruct{Status: "OK"})
}
