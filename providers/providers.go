// Copyright 2019 Netflix, Inc.
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

package providers

import (
	"path/filepath"

	"github.com/Netflix/p2plab"
	"github.com/Netflix/p2plab/errdefs"
	"github.com/Netflix/p2plab/providers/terraform"
	"github.com/pkg/errors"
)

type ProviderSettings struct {
}

func GetNodeProvider(root, providerType string, settings ProviderSettings) (p2plab.NodeProvider, error) {
	root = filepath.Join(root, providerType)
	switch providerType {
	case "terraform":
		return terraform.New(root)
	default:
		return nil, errors.Wrapf(errdefs.ErrInvalidArgument, "unrecognized node provider type %q", providerType)
	}
}
