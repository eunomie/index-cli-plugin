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

package types

import v1 "github.com/google/go-containerregistry/pkg/v1"

type Score struct {
	Type  string `edn:"vulnerability.reference.score/type" json:"type"`
	Value string `edn:"vulnerability.reference.score/value" json:"value"`
}

type Reference struct {
	Source string  `edn:"vulnerability.reference/source" json:"source"`
	Scores []Score `edn:"vulnerability.reference/scores" json:"scores"`
}

type Cwe struct {
	SourceID string `edn:"vulnerability.cwe/source-id" json:"source_id"`
	Name     string `edn:"vulnerability.cwe/name" json:"name,omitempty"`
}

type URL struct {
	Name  string `edn:"vulnerability.url/name" json:"name"`
	Value string `edn:"vulnerability.url/value" json:"value,omitempty"`
}

type Advisory struct {
	Source      string      `edn:"vulnerability/source" json:"source"`
	SourceID    string      `edn:"vulnerability/source-id" json:"source_id"`
	References  []Reference `edn:"vulnerability/references" json:"references"`
	Description string      `edn:"vulnerability/description" json:"description,omitempty"`
	Cwes        []Cwe       `edn:"vulnerability/cwes" json:"cwes,omitempty"`
	Urls        []URL       `edn:"vulnerability/urls" json:"urls,omitempty"`
}

type Cve struct {
	Purl            string    `edn:"purl" json:"purl"`
	Source          string    `edn:"source" json:"source"`
	SourceID        string    `edn:"source-id" json:"source_id"`
	VulnerableRange string    `edn:"vulnerable-range" json:"vulnerable_range"`
	AdvisoryURL     string    `edn:"url" json:"-"`
	FixedBy         string    `edn:"fixed-by" json:"fixed_by,omitempty"`
	Advisory        *Advisory `edn:"v" json:"vendor_advisory,omitempty"`
	Cve             *Advisory `edn:"cve" json:"nist_cve,omitempty"`
}

type LayerMapping struct {
	ByDiffID        map[string]string
	ByDigest        map[string]string
	OrdinalByDiffID map[string]int
	DiffIDByOrdinal map[int]string
	DigestByOrdinal map[int]string
}

type IndexResult struct {
	Name     string
	Packages []Package
	Status   string
	Error    error
	Distro   Distro
}

type BaseImagesResult struct {
	BaseImages []BaseImage
	Status     string
	Error      error
}

const (
	Success string = "success"
	Failed  string = "failed"
)

type Distro struct {
	OsName    string `json:"os_name,omitempty"`
	OsVersion string `json:"os_version,omitempty"`
	OsDistro  string `json:"os_distro,omitempty"`
}

type Platform struct {
	Os           string `json:"os"`
	Architecture string `json:"architecture"`
	Variant      string `json:"variant,omitempty"`
}

type Location struct {
	Path   string `json:"path"`
	Digest string `json:"digest"`
	DiffID string `json:"diff_id"`
}

type ImageSource struct {
	Name        string         `json:"name"`
	Digest      string         `json:"digest"`
	Tags        *[]string      `json:"tags,omitempty"`
	Manifest    *v1.Manifest   `json:"manifest"`
	Config      *v1.ConfigFile `json:"config"`
	RawManifest string         `json:"raw_manifest"`
	RawConfig   string         `json:"raw_config"`
	Distro      Distro         `json:"distro"`
	Platform    Platform       `json:"platform"`
	Size        int64          `json:"size"`
	Details     *BaseImage     `json:"details,omitempty""`
}

type Descriptor struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	SbomVersion string `json:"sbom_version"`
}

type Source struct {
	Type       string           `json:"type"`
	Image      ImageSource      `json:"image"`
	BaseImages []BaseImageMatch `json:"base_images,omitempty"`
}

type Sbom struct {
	Source          Source                  `json:"source"`
	Artifacts       []Package               `json:"artifacts"`
	Vulnerabilities []VulnerabilitiesByPurl `json:"vulnerabilities,omitempty"`
	Descriptor      Descriptor              `json:"descriptor"`
}

type Package struct {
	Type          string     `json:"type"`
	Namespace     string     `json:"namespace,omitempty"`
	Name          string     `json:"name"`
	Version       string     `json:"version"`
	Purl          string     `json:"purl"`
	Author        string     `json:"author,omitempty"`
	Description   string     `json:"description,omitempty"`
	Licenses      []string   `json:"licenses,omitempty"`
	URL           string     `json:"url,omitempty"`
	Size          int        `json:"size,omitempty"`
	InstalledSize int        `json:"installed_size,omitempty"`
	Locations     []Location `json:"locations"`
	Files         []Location `json:"files,omitempty"`
	Parent        string     `json:"parent,omitempty"`
}

var NamespaceMapping = map[string]string{
	"oracle": "oraclelinux",
	"ol":     "oraclelinux",
	"amazon": "amazonlinux",
	"amzn":   "amazonlinux",
	"rhel":   "redhatlinux",
}

var PackageTypeMapping = map[string]string{
	"apk":            "alpine",
	"debian":         "deb",
	"ubuntu":         "deb",
	"node-pkg":       "npm",
	"java":           "maven",
	"gobinary":       "golang",
	"go":             "golang",
	"go-module":      "golang",
	"java-archive":   "maven",
	"jenkins-plugin": "maven",
	"python":         "pypi",
	"python-pkg":     "pypi",
	"jar":            "maven",
	"gemspec":        "gem",
	"centos":         "rpm",
	"oracle":         "rpm",
	"ol":             "rpm",
	"amzn":           "rpm",
	"amazon":         "rpm",
	"redhat":         "rpm",
	"photon":         "rpm",
	"sles":           "rpm",
	"rhel":           "rpm",
	"pip":            "pypi",
	"rubygems":       "gem",
	"rust":           "cargo",
	"crates.io":      "cargo",
	"packagist":      "composer",
}
