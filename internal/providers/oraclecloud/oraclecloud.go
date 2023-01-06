// Copyright 2023 CoreOS, Inc.
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

// The Oracle Cloud provider fetches a remote configuration from the
// instance metadata service (IMDS).
// For details see https://docs.oracle.com/en-us/iaas/Content/Compute/Tasks/gettingmetadata.htm

package oraclecloud

import (
	"net/http"
	"net/url"

	"github.com/coreos/ignition/v2/config/v3_4_experimental/types"
	"github.com/coreos/ignition/v2/internal/providers/util"
	"github.com/coreos/ignition/v2/internal/resource"

	"github.com/coreos/vcontext/report"
)

var (
	userdataURL = url.URL{
		Scheme: "http",
		Host:   "169.254.169.254",
		Path:   "opc/v2/instance/metadata/user_data",
	}
	fetchOptions = resource.FetchOptions{
		Headers: http.Header{
			"Authorization": {"Bearer Oracle"},
		},
	}
)

func FetchConfig(f *resource.Fetcher) (types.Config, report.Report, error) {
	data, err := f.FetchToBuffer(userdataURL, fetchOptions)
	if err != nil && err != resource.ErrNotFound {
		return types.Config{}, report.Report{}, err
	}

	return util.ParseConfig(f.Logger, data)
}
