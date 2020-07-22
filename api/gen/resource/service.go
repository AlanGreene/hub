// Code generated by goa v3.2.0, DO NOT EDIT.
//
// resource service
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package resource

import (
	"context"

	resourceviews "github.com/tektoncd/hub/api/gen/resource/views"
	goa "goa.design/goa/v3/pkg"
)

// The resource service provides details about all type of resources
type Service interface {
	// Find resources by a combination of name, type
	Query(context.Context, *QueryPayload) (res ResourceCollection, err error)
	// List all resources sorted by rating and name
	List(context.Context, *ListPayload) (res ResourceCollection, err error)
	// Find all versions of a resource by its id
	VersionsByID(context.Context, *VersionsByIDPayload) (res *Versions, err error)
	// Find resource using name, type and version of resource
	ByTypeNameVersion(context.Context, *ByTypeNameVersionPayload) (res *Version, err error)
	// Find a resource using its version's id
	ByVersionID(context.Context, *ByVersionIDPayload) (res *Version, err error)
	// Find resources using name and type
	ByTypeName(context.Context, *ByTypeNamePayload) (res ResourceCollection, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "resource"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [6]string{"Query", "List", "VersionsByID", "ByTypeNameVersion", "ByVersionId", "ByTypeName"}

// QueryPayload is the payload type of the resource service Query method.
type QueryPayload struct {
	// Name of resource
	Name string
	// Type of resource
	Type string
	// Maximum number of resources to be returned
	Limit uint
}

// ResourceCollection is the result type of the resource service Query method.
type ResourceCollection []*Resource

// ListPayload is the payload type of the resource service List method.
type ListPayload struct {
	// Maximum number of resources to be returned
	Limit uint
}

// VersionsByIDPayload is the payload type of the resource service VersionsByID
// method.
type VersionsByIDPayload struct {
	// ID of a resource
	ID uint
}

// Versions is the result type of the resource service VersionsByID method.
type Versions struct {
	// Latest Version of resource
	Latest *Version
	// List of all versions of resource
	Versions []*Version
}

// ByTypeNameVersionPayload is the payload type of the resource service
// ByTypeNameVersion method.
type ByTypeNameVersionPayload struct {
	// type of resource
	Type string
	// name of resource
	Name string
	// version of resource
	Version string
}

// Version is the result type of the resource service ByTypeNameVersion method.
type Version struct {
	// ID is the unique id of resource's version
	ID uint
	// Version of resource
	Version string
	// Display name of version
	DisplayName string
	// Description of version
	Description string
	// Minimum pipelines version the resource's version is compatible with
	MinPipelinesVersion string
	// Raw URL of resource's yaml file of the version
	RawURL string
	// Web URL of resource's yaml file of the version
	WebURL string
	// Timestamp when version was last updated
	UpdatedAt string
	// Resource to which the version belongs
	Resource *Resource
}

// ByVersionIDPayload is the payload type of the resource service ByVersionId
// method.
type ByVersionIDPayload struct {
	// Version ID of a resource's version
	VersionID uint
}

// ByTypeNamePayload is the payload type of the resource service ByTypeName
// method.
type ByTypeNamePayload struct {
	// Type of resource
	Type string
	// Name of resource
	Name string
}

// The resource type describes resource information.
type Resource struct {
	// ID is the unique id of the resource
	ID uint
	// Name of resource
	Name string
	// Type of catalog to which resource belongs
	Catalog *Catalog
	// Type of resource
	Type string
	// Latest version of resource
	LatestVersion *LatestVersion
	// Tags related to resource
	Tags []*Tag
	// Rating of resource
	Rating float64
}

type Catalog struct {
	// ID is the unique id of the catalog
	ID uint
	// Type of catalog
	Type string
}

type LatestVersion struct {
	// ID is the unique id of resource's version
	ID uint
	// Version of resource
	Version string
	// Display name of version
	DisplayName string
	// Description of version
	Description string
	// Minimum pipelines version the resource's version is compatible with
	MinPipelinesVersion string
	// Raw URL of resource's yaml file of the version
	RawURL string
	// Web URL of resource's yaml file of the version
	WebURL string
	// Timestamp when version was last updated
	UpdatedAt string
}

type Tag struct {
	// ID is the unique id of tag
	ID uint
	// Name of tag
	Name string
}

// MakeInternalError builds a goa.ServiceError from an error.
func MakeInternalError(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "internal-error",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not-found",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewResourceCollection initializes result type ResourceCollection from viewed
// result type ResourceCollection.
func NewResourceCollection(vres resourceviews.ResourceCollection) ResourceCollection {
	var res ResourceCollection
	switch vres.View {
	case "info":
		res = newResourceCollectionInfo(vres.Projected)
	case "default", "":
		res = newResourceCollection(vres.Projected)
	}
	return res
}

// NewViewedResourceCollection initializes viewed result type
// ResourceCollection from result type ResourceCollection using the given view.
func NewViewedResourceCollection(res ResourceCollection, view string) resourceviews.ResourceCollection {
	var vres resourceviews.ResourceCollection
	switch view {
	case "info":
		p := newResourceCollectionViewInfo(res)
		vres = resourceviews.ResourceCollection{Projected: p, View: "info"}
	case "default", "":
		p := newResourceCollectionView(res)
		vres = resourceviews.ResourceCollection{Projected: p, View: "default"}
	}
	return vres
}

// NewVersions initializes result type Versions from viewed result type
// Versions.
func NewVersions(vres *resourceviews.Versions) *Versions {
	return newVersions(vres.Projected)
}

// NewViewedVersions initializes viewed result type Versions from result type
// Versions using the given view.
func NewViewedVersions(res *Versions, view string) *resourceviews.Versions {
	p := newVersionsView(res)
	return &resourceviews.Versions{Projected: p, View: "default"}
}

// NewVersion initializes result type Version from viewed result type Version.
func NewVersion(vres *resourceviews.Version) *Version {
	var res *Version
	switch vres.View {
	case "urls":
		res = newVersionUrls(vres.Projected)
	case "default", "":
		res = newVersion(vres.Projected)
	}
	return res
}

// NewViewedVersion initializes viewed result type Version from result type
// Version using the given view.
func NewViewedVersion(res *Version, view string) *resourceviews.Version {
	var vres *resourceviews.Version
	switch view {
	case "urls":
		p := newVersionViewUrls(res)
		vres = &resourceviews.Version{Projected: p, View: "urls"}
	case "default", "":
		p := newVersionView(res)
		vres = &resourceviews.Version{Projected: p, View: "default"}
	}
	return vres
}

// newResourceCollectionInfo converts projected type ResourceCollection to
// service type ResourceCollection.
func newResourceCollectionInfo(vres resourceviews.ResourceCollectionView) ResourceCollection {
	res := make(ResourceCollection, len(vres))
	for i, n := range vres {
		res[i] = newResourceInfo(n)
	}
	return res
}

// newResourceCollection converts projected type ResourceCollection to service
// type ResourceCollection.
func newResourceCollection(vres resourceviews.ResourceCollectionView) ResourceCollection {
	res := make(ResourceCollection, len(vres))
	for i, n := range vres {
		res[i] = newResource(n)
	}
	return res
}

// newResourceCollectionViewInfo projects result type ResourceCollection to
// projected type ResourceCollectionView using the "info" view.
func newResourceCollectionViewInfo(res ResourceCollection) resourceviews.ResourceCollectionView {
	vres := make(resourceviews.ResourceCollectionView, len(res))
	for i, n := range res {
		vres[i] = newResourceViewInfo(n)
	}
	return vres
}

// newResourceCollectionView projects result type ResourceCollection to
// projected type ResourceCollectionView using the "default" view.
func newResourceCollectionView(res ResourceCollection) resourceviews.ResourceCollectionView {
	vres := make(resourceviews.ResourceCollectionView, len(res))
	for i, n := range res {
		vres[i] = newResourceView(n)
	}
	return vres
}

// newResourceInfo converts projected type Resource to service type Resource.
func newResourceInfo(vres *resourceviews.ResourceView) *Resource {
	res := &Resource{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Type != nil {
		res.Type = *vres.Type
	}
	if vres.Rating != nil {
		res.Rating = *vres.Rating
	}
	if vres.Catalog != nil {
		res.Catalog = transformResourceviewsCatalogViewToCatalog(vres.Catalog)
	}
	if vres.Tags != nil {
		res.Tags = make([]*Tag, len(vres.Tags))
		for i, val := range vres.Tags {
			res.Tags[i] = transformResourceviewsTagViewToTag(val)
		}
	}
	return res
}

// newResource converts projected type Resource to service type Resource.
func newResource(vres *resourceviews.ResourceView) *Resource {
	res := &Resource{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Type != nil {
		res.Type = *vres.Type
	}
	if vres.Rating != nil {
		res.Rating = *vres.Rating
	}
	if vres.Catalog != nil {
		res.Catalog = transformResourceviewsCatalogViewToCatalog(vres.Catalog)
	}
	if vres.LatestVersion != nil {
		res.LatestVersion = transformResourceviewsLatestVersionViewToLatestVersion(vres.LatestVersion)
	}
	if vres.Tags != nil {
		res.Tags = make([]*Tag, len(vres.Tags))
		for i, val := range vres.Tags {
			res.Tags[i] = transformResourceviewsTagViewToTag(val)
		}
	}
	return res
}

// newResourceViewInfo projects result type Resource to projected type
// ResourceView using the "info" view.
func newResourceViewInfo(res *Resource) *resourceviews.ResourceView {
	vres := &resourceviews.ResourceView{
		ID:     &res.ID,
		Name:   &res.Name,
		Type:   &res.Type,
		Rating: &res.Rating,
	}
	if res.Catalog != nil {
		vres.Catalog = transformCatalogToResourceviewsCatalogView(res.Catalog)
	}
	if res.Tags != nil {
		vres.Tags = make([]*resourceviews.TagView, len(res.Tags))
		for i, val := range res.Tags {
			vres.Tags[i] = transformTagToResourceviewsTagView(val)
		}
	}
	return vres
}

// newResourceView projects result type Resource to projected type ResourceView
// using the "default" view.
func newResourceView(res *Resource) *resourceviews.ResourceView {
	vres := &resourceviews.ResourceView{
		ID:     &res.ID,
		Name:   &res.Name,
		Type:   &res.Type,
		Rating: &res.Rating,
	}
	if res.Catalog != nil {
		vres.Catalog = transformCatalogToResourceviewsCatalogView(res.Catalog)
	}
	if res.LatestVersion != nil {
		vres.LatestVersion = transformLatestVersionToResourceviewsLatestVersionView(res.LatestVersion)
	}
	if res.Tags != nil {
		vres.Tags = make([]*resourceviews.TagView, len(res.Tags))
		for i, val := range res.Tags {
			vres.Tags[i] = transformTagToResourceviewsTagView(val)
		}
	}
	return vres
}

// newVersions converts projected type Versions to service type Versions.
func newVersions(vres *resourceviews.VersionsView) *Versions {
	res := &Versions{}
	if vres.Versions != nil {
		res.Versions = make([]*Version, len(vres.Versions))
		for i, val := range vres.Versions {
			res.Versions[i] = transformResourceviewsVersionViewToVersion(val)
		}
	}
	if vres.Latest != nil {
		res.Latest = newVersionUrls(vres.Latest)
	}
	return res
}

// newVersionsView projects result type Versions to projected type VersionsView
// using the "default" view.
func newVersionsView(res *Versions) *resourceviews.VersionsView {
	vres := &resourceviews.VersionsView{}
	if res.Versions != nil {
		vres.Versions = make([]*resourceviews.VersionView, len(res.Versions))
		for i, val := range res.Versions {
			vres.Versions[i] = transformVersionToResourceviewsVersionView(val)
		}
	}
	if res.Latest != nil {
		vres.Latest = newVersionViewUrls(res.Latest)
	}
	return vres
}

// newVersionUrls converts projected type Version to service type Version.
func newVersionUrls(vres *resourceviews.VersionView) *Version {
	res := &Version{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Version != nil {
		res.Version = *vres.Version
	}
	if vres.RawURL != nil {
		res.RawURL = *vres.RawURL
	}
	if vres.WebURL != nil {
		res.WebURL = *vres.WebURL
	}
	if vres.Resource != nil {
		res.Resource = newResource(vres.Resource)
	}
	return res
}

// newVersion converts projected type Version to service type Version.
func newVersion(vres *resourceviews.VersionView) *Version {
	res := &Version{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Version != nil {
		res.Version = *vres.Version
	}
	if vres.Description != nil {
		res.Description = *vres.Description
	}
	if vres.MinPipelinesVersion != nil {
		res.MinPipelinesVersion = *vres.MinPipelinesVersion
	}
	if vres.DisplayName != nil {
		res.DisplayName = *vres.DisplayName
	}
	if vres.RawURL != nil {
		res.RawURL = *vres.RawURL
	}
	if vres.WebURL != nil {
		res.WebURL = *vres.WebURL
	}
	if vres.UpdatedAt != nil {
		res.UpdatedAt = *vres.UpdatedAt
	}
	if vres.Resource != nil {
		res.Resource = newResourceInfo(vres.Resource)
	}
	return res
}

// newVersionViewUrls projects result type Version to projected type
// VersionView using the "urls" view.
func newVersionViewUrls(res *Version) *resourceviews.VersionView {
	vres := &resourceviews.VersionView{
		ID:      &res.ID,
		Version: &res.Version,
		RawURL:  &res.RawURL,
		WebURL:  &res.WebURL,
	}
	return vres
}

// newVersionView projects result type Version to projected type VersionView
// using the "default" view.
func newVersionView(res *Version) *resourceviews.VersionView {
	vres := &resourceviews.VersionView{
		ID:                  &res.ID,
		Version:             &res.Version,
		DisplayName:         &res.DisplayName,
		Description:         &res.Description,
		MinPipelinesVersion: &res.MinPipelinesVersion,
		RawURL:              &res.RawURL,
		WebURL:              &res.WebURL,
		UpdatedAt:           &res.UpdatedAt,
	}
	if res.Resource != nil {
		vres.Resource = newResourceViewInfo(res.Resource)
	}
	return vres
}

// transformResourceviewsCatalogViewToCatalog builds a value of type *Catalog
// from a value of type *resourceviews.CatalogView.
func transformResourceviewsCatalogViewToCatalog(v *resourceviews.CatalogView) *Catalog {
	if v == nil {
		return nil
	}
	res := &Catalog{
		ID:   *v.ID,
		Type: *v.Type,
	}

	return res
}

// transformResourceviewsTagViewToTag builds a value of type *Tag from a value
// of type *resourceviews.TagView.
func transformResourceviewsTagViewToTag(v *resourceviews.TagView) *Tag {
	if v == nil {
		return nil
	}
	res := &Tag{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}

// transformResourceviewsLatestVersionViewToLatestVersion builds a value of
// type *LatestVersion from a value of type *resourceviews.LatestVersionView.
func transformResourceviewsLatestVersionViewToLatestVersion(v *resourceviews.LatestVersionView) *LatestVersion {
	if v == nil {
		return nil
	}
	res := &LatestVersion{
		ID:                  *v.ID,
		Version:             *v.Version,
		DisplayName:         *v.DisplayName,
		Description:         *v.Description,
		MinPipelinesVersion: *v.MinPipelinesVersion,
		RawURL:              *v.RawURL,
		WebURL:              *v.WebURL,
		UpdatedAt:           *v.UpdatedAt,
	}

	return res
}

// transformCatalogToResourceviewsCatalogView builds a value of type
// *resourceviews.CatalogView from a value of type *Catalog.
func transformCatalogToResourceviewsCatalogView(v *Catalog) *resourceviews.CatalogView {
	res := &resourceviews.CatalogView{
		ID:   &v.ID,
		Type: &v.Type,
	}

	return res
}

// transformTagToResourceviewsTagView builds a value of type
// *resourceviews.TagView from a value of type *Tag.
func transformTagToResourceviewsTagView(v *Tag) *resourceviews.TagView {
	res := &resourceviews.TagView{
		ID:   &v.ID,
		Name: &v.Name,
	}

	return res
}

// transformLatestVersionToResourceviewsLatestVersionView builds a value of
// type *resourceviews.LatestVersionView from a value of type *LatestVersion.
func transformLatestVersionToResourceviewsLatestVersionView(v *LatestVersion) *resourceviews.LatestVersionView {
	res := &resourceviews.LatestVersionView{
		ID:                  &v.ID,
		Version:             &v.Version,
		DisplayName:         &v.DisplayName,
		Description:         &v.Description,
		MinPipelinesVersion: &v.MinPipelinesVersion,
		RawURL:              &v.RawURL,
		WebURL:              &v.WebURL,
		UpdatedAt:           &v.UpdatedAt,
	}

	return res
}

// transformResourceviewsVersionViewToVersion builds a value of type *Version
// from a value of type *resourceviews.VersionView.
func transformResourceviewsVersionViewToVersion(v *resourceviews.VersionView) *Version {
	if v == nil {
		return nil
	}
	res := &Version{
		ID:                  *v.ID,
		Version:             *v.Version,
		DisplayName:         *v.DisplayName,
		Description:         *v.Description,
		MinPipelinesVersion: *v.MinPipelinesVersion,
		RawURL:              *v.RawURL,
		WebURL:              *v.WebURL,
		UpdatedAt:           *v.UpdatedAt,
	}
	if v.Resource != nil {
		res.Resource = transformResourceviewsResourceViewToResource(v.Resource)
	}

	return res
}

// transformResourceviewsResourceViewToResource builds a value of type
// *Resource from a value of type *resourceviews.ResourceView.
func transformResourceviewsResourceViewToResource(v *resourceviews.ResourceView) *Resource {
	res := &Resource{
		ID:     *v.ID,
		Name:   *v.Name,
		Type:   *v.Type,
		Rating: *v.Rating,
	}
	if v.Catalog != nil {
		res.Catalog = transformResourceviewsCatalogViewToCatalog(v.Catalog)
	}
	if v.LatestVersion != nil {
		res.LatestVersion = transformResourceviewsLatestVersionViewToLatestVersion(v.LatestVersion)
	}
	if v.Tags != nil {
		res.Tags = make([]*Tag, len(v.Tags))
		for i, val := range v.Tags {
			res.Tags[i] = transformResourceviewsTagViewToTag(val)
		}
	}

	return res
}

// transformVersionToResourceviewsVersionView builds a value of type
// *resourceviews.VersionView from a value of type *Version.
func transformVersionToResourceviewsVersionView(v *Version) *resourceviews.VersionView {
	res := &resourceviews.VersionView{
		ID:                  &v.ID,
		Version:             &v.Version,
		DisplayName:         &v.DisplayName,
		Description:         &v.Description,
		MinPipelinesVersion: &v.MinPipelinesVersion,
		RawURL:              &v.RawURL,
		WebURL:              &v.WebURL,
		UpdatedAt:           &v.UpdatedAt,
	}
	if v.Resource != nil {
		res.Resource = transformResourceToResourceviewsResourceView(v.Resource)
	}

	return res
}

// transformResourceToResourceviewsResourceView builds a value of type
// *resourceviews.ResourceView from a value of type *Resource.
func transformResourceToResourceviewsResourceView(v *Resource) *resourceviews.ResourceView {
	res := &resourceviews.ResourceView{
		ID:     &v.ID,
		Name:   &v.Name,
		Type:   &v.Type,
		Rating: &v.Rating,
	}
	if v.Catalog != nil {
		res.Catalog = transformCatalogToResourceviewsCatalogView(v.Catalog)
	}
	if v.LatestVersion != nil {
		res.LatestVersion = transformLatestVersionToResourceviewsLatestVersionView(v.LatestVersion)
	}
	if v.Tags != nil {
		res.Tags = make([]*resourceviews.TagView, len(v.Tags))
		for i, val := range v.Tags {
			res.Tags[i] = transformTagToResourceviewsTagView(val)
		}
	}

	return res
}
