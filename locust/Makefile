# Copyright 2023 Google LLC All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

build:
	echo "Nothing to build. Please run build-image"

build-image:
	echo "Building locust docker container"
	docker build . -t locust:latest

clean:
	echo "Running cleanup"
	rm -rf ${BUILD_DIR}/
	docker rmi -f locust

.PHONY: build build-image clean
