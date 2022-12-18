/*
 * Copyright © 2022 Docker, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package registry

import (
	"github.com/google/go-containerregistry/pkg/v1/layout"
	"github.com/pkg/errors"

	"github.com/atomist-skills/go-skill"
)

func ReadImage(name string, path string) (*ImageCache, error) {
	skill.Log.Infof("Loading image from %s", path)
	index, err := layout.ImageIndexFromPath(path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read manifest index at %s", path)
	}
	mani, err := index.IndexManifest()
	hash := mani.Manifests[0].Digest
	img, _ := index.Image(hash)
	skill.Log.Infof("Loaded image")
	return &ImageCache{
		ID:        hash.String(),
		Digest:    hash.String(),
		Tags:      []string{},
		Name:      name,
		Image:     &img,
		ImagePath: path,
		Ref:       nil,
		copy:      false,
	}, nil
}
