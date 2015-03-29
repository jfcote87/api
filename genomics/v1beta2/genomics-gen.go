// Package genomics provides access to the Genomics API.
//
// See https://developers.google.com/genomics/v1beta2/reference
//
// Usage example:
//
//   import "github.com/jfcote87/api/genomics/v1beta2"
//   ...
//   genomicsService, err := genomics.New(oauthHttpClient)
package genomics

import (
	"errors"
	"fmt"
	"github.com/jfcote87/api/googleapi"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Background

const apiId = "genomics:v1beta2"
const apiName = "genomics"
const apiVersion = "v1beta2"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/genomics/v1beta2/"}

// OAuth2 scopes used by this API.
const (
	// View and manage your data in Google BigQuery
	BigqueryScope = "https://www.googleapis.com/auth/bigquery"

	// Manage your data in Google Cloud Storage
	DevstorageRead_writeScope = "https://www.googleapis.com/auth/devstorage.read_write"

	// View and manage Genomics data
	GenomicsScope = "https://www.googleapis.com/auth/genomics"

	// View Genomics data
	GenomicsReadonlyScope = "https://www.googleapis.com/auth/genomics.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.AnnotationSets = NewAnnotationSetsService(s)
	s.Annotations = NewAnnotationsService(s)
	s.Callsets = NewCallsetsService(s)
	s.Datasets = NewDatasetsService(s)
	s.Experimental = NewExperimentalService(s)
	s.Jobs = NewJobsService(s)
	s.Readgroupsets = NewReadgroupsetsService(s)
	s.Reads = NewReadsService(s)
	s.References = NewReferencesService(s)
	s.Referencesets = NewReferencesetsService(s)
	s.StreamingVariantStore = NewStreamingVariantStoreService(s)
	s.Variants = NewVariantsService(s)
	s.Variantsets = NewVariantsetsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	AnnotationSets *AnnotationSetsService

	Annotations *AnnotationsService

	Callsets *CallsetsService

	Datasets *DatasetsService

	Experimental *ExperimentalService

	Jobs *JobsService

	Readgroupsets *ReadgroupsetsService

	Reads *ReadsService

	References *ReferencesService

	Referencesets *ReferencesetsService

	StreamingVariantStore *StreamingVariantStoreService

	Variants *VariantsService

	Variantsets *VariantsetsService
}

func NewAnnotationSetsService(s *Service) *AnnotationSetsService {
	rs := &AnnotationSetsService{s: s}
	return rs
}

type AnnotationSetsService struct {
	s *Service
}

func NewAnnotationsService(s *Service) *AnnotationsService {
	rs := &AnnotationsService{s: s}
	return rs
}

type AnnotationsService struct {
	s *Service
}

func NewCallsetsService(s *Service) *CallsetsService {
	rs := &CallsetsService{s: s}
	return rs
}

type CallsetsService struct {
	s *Service
}

func NewDatasetsService(s *Service) *DatasetsService {
	rs := &DatasetsService{s: s}
	return rs
}

type DatasetsService struct {
	s *Service
}

func NewExperimentalService(s *Service) *ExperimentalService {
	rs := &ExperimentalService{s: s}
	rs.Jobs = NewExperimentalJobsService(s)
	return rs
}

type ExperimentalService struct {
	s *Service

	Jobs *ExperimentalJobsService
}

func NewExperimentalJobsService(s *Service) *ExperimentalJobsService {
	rs := &ExperimentalJobsService{s: s}
	return rs
}

type ExperimentalJobsService struct {
	s *Service
}

func NewJobsService(s *Service) *JobsService {
	rs := &JobsService{s: s}
	return rs
}

type JobsService struct {
	s *Service
}

func NewReadgroupsetsService(s *Service) *ReadgroupsetsService {
	rs := &ReadgroupsetsService{s: s}
	rs.Coveragebuckets = NewReadgroupsetsCoveragebucketsService(s)
	return rs
}

type ReadgroupsetsService struct {
	s *Service

	Coveragebuckets *ReadgroupsetsCoveragebucketsService
}

func NewReadgroupsetsCoveragebucketsService(s *Service) *ReadgroupsetsCoveragebucketsService {
	rs := &ReadgroupsetsCoveragebucketsService{s: s}
	return rs
}

type ReadgroupsetsCoveragebucketsService struct {
	s *Service
}

func NewReadsService(s *Service) *ReadsService {
	rs := &ReadsService{s: s}
	return rs
}

type ReadsService struct {
	s *Service
}

func NewReferencesService(s *Service) *ReferencesService {
	rs := &ReferencesService{s: s}
	rs.Bases = NewReferencesBasesService(s)
	return rs
}

type ReferencesService struct {
	s *Service

	Bases *ReferencesBasesService
}

func NewReferencesBasesService(s *Service) *ReferencesBasesService {
	rs := &ReferencesBasesService{s: s}
	return rs
}

type ReferencesBasesService struct {
	s *Service
}

func NewReferencesetsService(s *Service) *ReferencesetsService {
	rs := &ReferencesetsService{s: s}
	return rs
}

type ReferencesetsService struct {
	s *Service
}

func NewStreamingVariantStoreService(s *Service) *StreamingVariantStoreService {
	rs := &StreamingVariantStoreService{s: s}
	return rs
}

type StreamingVariantStoreService struct {
	s *Service
}

func NewVariantsService(s *Service) *VariantsService {
	rs := &VariantsService{s: s}
	return rs
}

type VariantsService struct {
	s *Service
}

func NewVariantsetsService(s *Service) *VariantsetsService {
	rs := &VariantsetsService{s: s}
	return rs
}

type VariantsetsService struct {
	s *Service
}

type AlignReadGroupSetsRequest struct {
	// BamSourceUris: The BAM source files for alignment. Exactly one of
	// readGroupSetId, bamSourceUris, interleavedFastqSource or
	// pairedFastqSource must be provided. The caller must have READ
	// permissions for these files.
	BamSourceUris []string `json:"bamSourceUris,omitempty"`

	// DatasetId: Required. The ID of the dataset the newly aligned read
	// group sets will belong to. The caller must have WRITE permissions to
	// this dataset.
	DatasetId string `json:"datasetId,omitempty"`

	// InterleavedFastqSource: The interleaved FASTQ source files for
	// alignment, where both members of each pair of reads are found on
	// consecutive records within the same FASTQ file. Exactly one of
	// readGroupSetId, bamSourceUris, interleavedFastqSource or
	// pairedFastqSource must be provided.
	InterleavedFastqSource *InterleavedFastqSource `json:"interleavedFastqSource,omitempty"`

	// PairedFastqSource: The paired end FASTQ source files for alignment,
	// where each member of a pair of reads are found in separate files.
	// Exactly one of readGroupSetId, bamSourceUris, interleavedFastqSource
	// or pairedFastqSource must be provided.
	PairedFastqSource *PairedFastqSource `json:"pairedFastqSource,omitempty"`

	// ReadGroupSetId: The ID of the read group set which will be aligned. A
	// new read group set will be generated to hold the aligned data, the
	// originals will not be modified. The caller must have READ permissions
	// for this read group set. Exactly one of readGroupSetId,
	// bamSourceUris, interleavedFastqSource or pairedFastqSource must be
	// provided.
	ReadGroupSetId string `json:"readGroupSetId,omitempty"`
}

type AlignReadGroupSetsResponse struct {
	// JobId: A job ID that can be used to get status information.
	JobId string `json:"jobId,omitempty"`
}

type Annotation struct {
	// AnnotationSetId: The ID of the containing annotation set.
	AnnotationSetId string `json:"annotationSetId,omitempty"`

	// Id: The generated unique ID for this annotation.
	Id string `json:"id,omitempty"`

	// Info: A map of additional data for this annotation.
	Info map[string][]string `json:"info,omitempty"`

	// Name: The display name of this annotation.
	Name string `json:"name,omitempty"`

	// Position: The position of this annotation on the reference sequence.
	Position *RangePosition `json:"position,omitempty"`

	// Transcript: A transcript value represents the assertion that a
	// particular region of the reference genome may be transcribed as RNA.
	// An alternative splicing pattern would be represented as a separate
	// transcript object. This field is only set for annotations of type
	// TRANSCRIPT.
	Transcript *Transcript `json:"transcript,omitempty"`

	// Type: The data type for this annotation. Must match the containing
	// annotation set's type.
	Type string `json:"type,omitempty"`

	// Variant: A variant annotation, which describes the effect of a
	// variant on the genome, the coding sequence, and/or higher level
	// consequences at the organism level e.g. pathogenicity. This field is
	// only set for annotations of type VARIANT.
	Variant *VariantAnnotation `json:"variant,omitempty"`
}

type AnnotationSet struct {
	// DatasetId: The ID of the containing dataset.
	DatasetId string `json:"datasetId,omitempty"`

	// Id: The generated unique ID for this annotation set.
	Id string `json:"id,omitempty"`

	// Info: A map of additional data for this annotation set.
	Info map[string][]string `json:"info,omitempty"`

	// Name: The display name for this annotation set.
	Name string `json:"name,omitempty"`

	// ReferenceSetId: The ID of the reference set that defines the
	// coordinate space for this set's annotations.
	ReferenceSetId string `json:"referenceSetId,omitempty"`

	// SourceUri: The source URI describing the file from which this
	// annotation set was generated, if any.
	SourceUri string `json:"sourceUri,omitempty"`

	// Type: The type of annotations contained within this set.
	Type string `json:"type,omitempty"`
}

type BatchAnnotationsResponse struct {
	// Entries: The resulting per-annotation entries, ordered consistently
	// with the original request.
	Entries []*BatchAnnotationsResponseEntry `json:"entries,omitempty"`
}

type BatchAnnotationsResponseEntry struct {
	// Annotation: The annotation, if any.
	Annotation *Annotation `json:"annotation,omitempty"`

	// Status: The resulting status for this annotation operation.
	Status *BatchAnnotationsResponseEntryStatus `json:"status,omitempty"`
}

type BatchAnnotationsResponseEntryStatus struct {
	// Code: The HTTP status code for this operation.
	Code int64 `json:"code,omitempty"`

	// Message: Error message for this status, if any.
	Message string `json:"message,omitempty"`
}

type BatchCreateAnnotationsRequest struct {
	// Annotations: The annotations to be created. At most 4096 can be
	// specified in a single request.
	Annotations []*Annotation `json:"annotations,omitempty"`
}

type Call struct {
	// CallSetId: The ID of the call set this variant call belongs to.
	CallSetId string `json:"callSetId,omitempty"`

	// CallSetName: The name of the call set this variant call belongs to.
	CallSetName string `json:"callSetName,omitempty"`

	// Genotype: The genotype of this variant call. Each value represents
	// either the value of the referenceBases field or a 1-based index into
	// alternateBases. If a variant had a referenceBases value of T and an
	// alternateBases value of ["A", "C"], and the genotype was [2, 1], that
	// would mean the call represented the heterozygous value CA for this
	// variant. If the genotype was instead [0, 1], the represented value
	// would be TA. Ordering of the genotype values is important if the
	// phaseset is present. If a genotype is not called (that is, a . is
	// present in the GT string) -1 is returned.
	Genotype []int64 `json:"genotype,omitempty"`

	// GenotypeLikelihood: The genotype likelihoods for this variant call.
	// Each array entry represents how likely a specific genotype is for
	// this call. The value ordering is defined by the GL tag in the VCF
	// spec. If Phred-scaled genotype likelihood scores (PL) are available
	// and log10(P) genotype likelihood scores (GL) are not, PL scores are
	// converted to GL scores. If both are available, PL scores are stored
	// in info.
	GenotypeLikelihood []float64 `json:"genotypeLikelihood,omitempty"`

	// Info: A map of additional variant call information.
	Info map[string][]string `json:"info,omitempty"`

	// Phaseset: If this field is present, this variant call's genotype
	// ordering implies the phase of the bases and is consistent with any
	// other variant calls in the same reference sequence which have the
	// same phaseset value. When importing data from VCF, if the genotype
	// data was phased but no phase set was specified this field will be set
	// to *.
	Phaseset string `json:"phaseset,omitempty"`
}

type CallReadGroupSetsRequest struct {
	// DatasetId: Required. The ID of the dataset the called variants will
	// belong to. The caller must have WRITE permissions to this dataset.
	DatasetId string `json:"datasetId,omitempty"`

	// ReadGroupSetId: The IDs of the read group sets which will be called.
	// The caller must have READ permissions for these read group sets. One
	// of readGroupSetId or sourceUris must be provided.
	ReadGroupSetId string `json:"readGroupSetId,omitempty"`

	// SourceUris: A list of URIs pointing at BAM files in Google Cloud
	// Storage which will be called. FASTQ files are not allowed. The caller
	// must have READ permissions for these files. One of readGroupSetId or
	// sourceUris must be provided.
	SourceUris []string `json:"sourceUris,omitempty"`
}

type CallReadGroupSetsResponse struct {
	// JobId: A job ID that can be used to get status information.
	JobId string `json:"jobId,omitempty"`
}

type CallSet struct {
	// Created: The date this call set was created in milliseconds from the
	// epoch.
	Created int64 `json:"created,omitempty,string"`

	// Id: The Google generated ID of the call set, immutable.
	Id string `json:"id,omitempty"`

	// Info: A map of additional call set information.
	Info map[string][]string `json:"info,omitempty"`

	// Name: The call set name.
	Name string `json:"name,omitempty"`

	// SampleId: The sample ID this call set corresponds to.
	SampleId string `json:"sampleId,omitempty"`

	// VariantSetIds: The IDs of the variant sets this call set belongs to.
	VariantSetIds []string `json:"variantSetIds,omitempty"`
}

type CigarUnit struct {
	Operation string `json:"operation,omitempty"`

	// OperationLength: The number of bases that the operation runs for.
	// Required.
	OperationLength int64 `json:"operationLength,omitempty,string"`

	// ReferenceSequence: referenceSequence is only used at mismatches
	// (SEQUENCE_MISMATCH) and deletions (DELETE). Filling this field
	// replaces SAM's MD tag. If the relevant information is not available,
	// this field is unset.
	ReferenceSequence string `json:"referenceSequence,omitempty"`
}

type CoverageBucket struct {
	// MeanCoverage: The average number of reads which are aligned to each
	// individual reference base in this bucket.
	MeanCoverage float64 `json:"meanCoverage,omitempty"`

	// Range: The genomic coordinate range spanned by this bucket.
	Range *Range `json:"range,omitempty"`
}

type Dataset struct {
	// Id: The Google generated ID of the dataset, immutable.
	Id string `json:"id,omitempty"`

	// IsPublic: Flag indicating whether or not a dataset is publicly
	// viewable. If a dataset is not public, it inherits viewing permissions
	// from its project.
	IsPublic bool `json:"isPublic,omitempty"`

	// Name: The dataset name.
	Name string `json:"name,omitempty"`

	// ProjectNumber: The Google Developers Console project number that this
	// dataset belongs to.
	ProjectNumber int64 `json:"projectNumber,omitempty,string"`
}

type ExperimentalCreateJobRequest struct {
	// Align: Specifies whether or not to run the alignment pipeline. Either
	// align or callVariants must be set.
	Align bool `json:"align,omitempty"`

	// CallVariants: Specifies whether or not to run the variant calling
	// pipeline. Either align or callVariants must be set.
	CallVariants bool `json:"callVariants,omitempty"`

	// GcsOutputPath: Specifies where to copy the results of certain
	// pipelines. This should be in the form of gs://bucket/path.
	GcsOutputPath string `json:"gcsOutputPath,omitempty"`

	// PairedSourceUris: A list of Google Cloud Storage URIs of paired end
	// .fastq files to operate upon. If specified, this represents the
	// second file of each paired .fastq file. The first file of each pair
	// should be specified in sourceUris.
	PairedSourceUris []string `json:"pairedSourceUris,omitempty"`

	// ProjectNumber: Required. The Google Cloud Project ID with which to
	// associate the request.
	ProjectNumber int64 `json:"projectNumber,omitempty,string"`

	// SourceUris: A list of Google Cloud Storage URIs of data files to
	// operate upon. These can be .bam, interleaved .fastq, or paired
	// .fastq. If specifying paired .fastq files, the first of each pair of
	// files should be listed here, and the second of each pair should be
	// listed in pairedSourceUris.
	SourceUris []string `json:"sourceUris,omitempty"`
}

type ExperimentalCreateJobResponse struct {
	// JobId: A job ID that can be used to get status information.
	JobId string `json:"jobId,omitempty"`
}

type ExportReadGroupSetsRequest struct {
	// ExportUri: A Google Cloud Storage URI for the exported BAM file. The
	// currently authenticated user must have write access to the new file.
	// An error will be returned if the URI already contains data.
	ExportUri string `json:"exportUri,omitempty"`

	// ProjectNumber: The Google Developers Console project number that owns
	// this export.
	ProjectNumber int64 `json:"projectNumber,omitempty,string"`

	// ReadGroupSetIds: The IDs of the read group sets to export.
	ReadGroupSetIds []string `json:"readGroupSetIds,omitempty"`

	// ReferenceNames: The reference names to export. If this is not
	// specified, all reference sequences, including unmapped reads, are
	// exported. Use * to export only unmapped reads.
	ReferenceNames []string `json:"referenceNames,omitempty"`
}

type ExportReadGroupSetsResponse struct {
	// JobId: A job ID that can be used to get status information.
	JobId string `json:"jobId,omitempty"`
}

type ExportVariantSetRequest struct {
	// BigqueryDataset: The BigQuery dataset to export data to. Note that
	// this is distinct from the Genomics concept of "dataset".
	BigqueryDataset string `json:"bigqueryDataset,omitempty"`

	// BigqueryTable: The BigQuery table to export data to. If the table
	// doesn't exist, it will be created. If it already exists, it will be
	// overwritten.
	BigqueryTable string `json:"bigqueryTable,omitempty"`

	// CallSetIds: If provided, only variant call information from the
	// specified call sets will be exported. By default all variant calls
	// are exported.
	CallSetIds []string `json:"callSetIds,omitempty"`

	// Format: The format for the exported data.
	Format string `json:"format,omitempty"`

	// ProjectNumber: The Google Cloud project number that owns the
	// destination BigQuery dataset. The caller must have WRITE access to
	// this project. This project will also own the resulting export job.
	ProjectNumber int64 `json:"projectNumber,omitempty,string"`
}

type ExportVariantSetResponse struct {
	// JobId: A job ID that can be used to get status information.
	JobId string `json:"jobId,omitempty"`
}

type ExternalId struct {
	// Id: The id used by the source of this data.
	Id string `json:"id,omitempty"`

	// SourceName: The name of the source of this data.
	SourceName string `json:"sourceName,omitempty"`
}

type FastqMetadata struct {
	// LibraryName: Optionally specifies the library name for alignment from
	// FASTQ.
	LibraryName string `json:"libraryName,omitempty"`

	// PlatformName: Optionally specifies the platform name for alignment
	// from FASTQ. For example: CAPILLARY, LS454, ILLUMINA, SOLID, HELICOS,
	// IONTORRENT, PACBIO.
	PlatformName string `json:"platformName,omitempty"`

	// PlatformUnit: Optionally specifies the platform unit for alignment
	// from FASTQ. For example: flowcell-barcode.lane for Illumina or slide
	// for SOLID.
	PlatformUnit string `json:"platformUnit,omitempty"`

	// ReadGroupName: Optionally specifies the read group name for alignment
	// from FASTQ.
	ReadGroupName string `json:"readGroupName,omitempty"`

	// SampleName: Optionally specifies the sample name for alignment from
	// FASTQ.
	SampleName string `json:"sampleName,omitempty"`
}

type ImportReadGroupSetsRequest struct {
	// DatasetId: Required. The ID of the dataset these read group sets will
	// belong to. The caller must have WRITE permissions to this dataset.
	DatasetId string `json:"datasetId,omitempty"`

	// PartitionStrategy: The partition strategy describes how read groups
	// are partitioned into read group sets.
	PartitionStrategy string `json:"partitionStrategy,omitempty"`

	// ReferenceSetId: The reference set to which the imported read group
	// sets are aligned to, if any. The reference names of this reference
	// set must be a superset of those found in the imported file headers.
	// If no reference set id is provided, a best effort is made to
	// associate with a matching reference set.
	ReferenceSetId string `json:"referenceSetId,omitempty"`

	// SourceUris: A list of URIs pointing at BAM files in Google Cloud
	// Storage.
	SourceUris []string `json:"sourceUris,omitempty"`
}

type ImportReadGroupSetsResponse struct {
	// JobId: A job ID that can be used to get status information.
	JobId string `json:"jobId,omitempty"`
}

type ImportVariantsRequest struct {
	// Format: The format of the variant data being imported.
	Format string `json:"format,omitempty"`

	// SourceUris: A list of URIs pointing at VCF files in Google Cloud
	// Storage. See the VCF Specification for more details on the input
	// format.
	SourceUris []string `json:"sourceUris,omitempty"`
}

type ImportVariantsResponse struct {
	// JobId: A job ID that can be used to get status information.
	JobId string `json:"jobId,omitempty"`
}

type Int32Value struct {
	// Value: The int32 value.
	Value int64 `json:"value,omitempty"`
}

type InterleavedFastqSource struct {
	// Metadata: Optionally specifies the metadata to be associated with the
	// final aligned read group set.
	Metadata *FastqMetadata `json:"metadata,omitempty"`

	// SourceUris: A list of URIs pointing at interleaved FASTQ files in
	// Google Cloud Storage which will be aligned. The caller must have READ
	// permissions for these files.
	SourceUris []string `json:"sourceUris,omitempty"`
}

type Job struct {
	// Created: The date this job was created, in milliseconds from the
	// epoch.
	Created int64 `json:"created,omitempty,string"`

	// DetailedStatus: A more detailed description of this job's current
	// status.
	DetailedStatus string `json:"detailedStatus,omitempty"`

	// Errors: Any errors that occurred during processing.
	Errors []string `json:"errors,omitempty"`

	// Id: The job ID.
	Id string `json:"id,omitempty"`

	// ImportedIds: If this Job represents an import, this field will
	// contain the IDs of the objects that were successfully imported.
	ImportedIds []string `json:"importedIds,omitempty"`

	// ProjectNumber: The Google Developers Console project number to which
	// this job belongs.
	ProjectNumber int64 `json:"projectNumber,omitempty,string"`

	// Request: A summarized representation of the original service request.
	Request *JobRequest `json:"request,omitempty"`

	// Status: The status of this job.
	Status string `json:"status,omitempty"`

	// Warnings: Any warnings that occurred during processing.
	Warnings []string `json:"warnings,omitempty"`
}

type JobRequest struct {
	// Destination: The data destination of the request, for example, a
	// Google BigQuery Table or Dataset ID.
	Destination []string `json:"destination,omitempty"`

	// Source: The data source of the request, for example, a Google Cloud
	// Storage object path or Readset ID.
	Source []string `json:"source,omitempty"`

	// Type: The original request type.
	Type string `json:"type,omitempty"`
}

type LinearAlignment struct {
	// Cigar: Represents the local alignment of this sequence (alignment
	// matches, indels, etc) against the reference.
	Cigar []*CigarUnit `json:"cigar,omitempty"`

	// MappingQuality: The mapping quality of this alignment. Represents how
	// likely the read maps to this position as opposed to other locations.
	MappingQuality int64 `json:"mappingQuality,omitempty"`

	// Position: The position of this alignment.
	Position *Position `json:"position,omitempty"`
}

type ListBasesResponse struct {
	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Offset: The offset position (0-based) of the given sequence from the
	// start of this Reference. This value will differ for each page in a
	// paginated request.
	Offset int64 `json:"offset,omitempty,string"`

	// Sequence: A substring of the bases that make up this reference.
	Sequence string `json:"sequence,omitempty"`
}

type ListCoverageBucketsResponse struct {
	// BucketWidth: The length of each coverage bucket in base pairs. Note
	// that buckets at the end of a reference sequence may be shorter. This
	// value is omitted if the bucket width is infinity (the default
	// behaviour, with no range or targetBucketWidth).
	BucketWidth int64 `json:"bucketWidth,omitempty,string"`

	// CoverageBuckets: The coverage buckets. The list of buckets is sparse;
	// a bucket with 0 overlapping reads is not returned. A bucket never
	// crosses more than one reference sequence. Each bucket has width
	// bucketWidth, unless its end is the end of the reference sequence.
	CoverageBuckets []*CoverageBucket `json:"coverageBuckets,omitempty"`

	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListDatasetsResponse struct {
	// Datasets: The list of matching Datasets.
	Datasets []*Dataset `json:"datasets,omitempty"`

	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type MergeVariantsRequest struct {
	// Variants: The variants to be merged with existing variants.
	Variants []*Variant `json:"variants,omitempty"`
}

type Metadata struct {
	// Description: A textual description of this metadata.
	Description string `json:"description,omitempty"`

	// Id: User-provided ID field, not enforced by this API. Two or more
	// pieces of structured metadata with identical id and key fields are
	// considered equivalent.
	Id string `json:"id,omitempty"`

	// Info: Remaining structured metadata key-value pairs.
	Info map[string][]string `json:"info,omitempty"`

	// Key: The top-level key.
	Key string `json:"key,omitempty"`

	// Number: The number of values that can be included in a field
	// described by this metadata.
	Number string `json:"number,omitempty"`

	// Type: The type of data. Possible types include: Integer, Float, Flag,
	// Character, and String.
	Type string `json:"type,omitempty"`

	// Value: The value field for simple metadata
	Value string `json:"value,omitempty"`
}

type PairedFastqSource struct {
	// FirstSourceUris: A list of URIs pointing at paired end FASTQ files in
	// Google Cloud Storage which will be aligned. The first of each paired
	// file should be specified here, in an order that matches the second of
	// each paired file specified in secondSourceUris. For example:
	// firstSourceUris: [file1_1.fq, file2_1.fq], secondSourceUris:
	// [file1_2.fq, file2_2.fq]. The caller must have READ permissions for
	// these files.
	FirstSourceUris []string `json:"firstSourceUris,omitempty"`

	// Metadata: Optionally specifies the metadata to be associated with the
	// final aligned read group set.
	Metadata *FastqMetadata `json:"metadata,omitempty"`

	// SecondSourceUris: A list of URIs pointing at paired end FASTQ files
	// in Google Cloud Storage which will be aligned. The second of each
	// paired file should be specified here, in an order that matches the
	// first of each paired file specified in firstSourceUris. For example:
	// firstSourceUris: [file1_1.fq, file2_1.fq], secondSourceUris:
	// [file1_2.fq, file2_2.fq]. The caller must have READ permissions for
	// these files.
	SecondSourceUris []string `json:"secondSourceUris,omitempty"`
}

type Position struct {
	// Position: The 0-based offset from the start of the forward strand for
	// that reference.
	Position int64 `json:"position,omitempty,string"`

	// ReferenceName: The name of the reference in whatever reference set is
	// being used.
	ReferenceName string `json:"referenceName,omitempty"`

	// ReverseStrand: Whether this position is on the reverse strand, as
	// opposed to the forward strand.
	ReverseStrand bool `json:"reverseStrand,omitempty"`
}

type QueryRange struct {
	// End: The end position of the range on the reference, 0-based
	// exclusive. If specified, referenceId or referenceName must also be
	// specified. If unset or 0, defaults to the length of the reference.
	End int64 `json:"end,omitempty,string"`

	// ReferenceId: The ID of the reference to query. At most one of
	// referenceId and referenceName should be specified.
	ReferenceId string `json:"referenceId,omitempty"`

	// ReferenceName: The name of the reference to query, within the
	// reference set associated with this query. At most one of referenceId
	// and referenceName pshould be specified.
	ReferenceName string `json:"referenceName,omitempty"`

	// Start: The start position of the range on the reference, 0-based
	// inclusive. If specified, referenceId or referenceName must also be
	// specified. Defaults to 0.
	Start int64 `json:"start,omitempty,string"`
}

type Range struct {
	// End: The end position of the range on the reference, 0-based
	// exclusive. If specified, referenceName must also be specified.
	End int64 `json:"end,omitempty,string"`

	// ReferenceName: The reference sequence name, for example chr1, 1, or
	// chrX.
	ReferenceName string `json:"referenceName,omitempty"`

	// Start: The start position of the range on the reference, 0-based
	// inclusive. If specified, referenceName must also be specified.
	Start int64 `json:"start,omitempty,string"`
}

type RangePosition struct {
	// End: The end position of the range on the reference, 0-based
	// exclusive.
	End int64 `json:"end,omitempty,string"`

	// ReferenceId: The ID of the Google Genomics reference associated with
	// this range.
	ReferenceId string `json:"referenceId,omitempty"`

	// ReferenceName: The display name corresponding to the reference
	// specified by referenceId, for example chr1, 1, or chrX.
	ReferenceName string `json:"referenceName,omitempty"`

	// ReverseStrand: Whether this range refers to the reverse strand, as
	// opposed to the forward strand. Note that regardless of this field,
	// the start/end position of the range always refer to the forward
	// strand.
	ReverseStrand bool `json:"reverseStrand,omitempty"`

	// Start: The start position of the range on the reference, 0-based
	// inclusive.
	Start int64 `json:"start,omitempty,string"`
}

type Read struct {
	// AlignedQuality: The quality of the read sequence contained in this
	// alignment record. alignedSequence and alignedQuality may be shorter
	// than the full read sequence and quality. This will occur if the
	// alignment is part of a chimeric alignment, or if the read was
	// trimmed. When this occurs, the CIGAR for this read will begin/end
	// with a hard clip operator that will indicate the length of the
	// excised sequence.
	AlignedQuality []int64 `json:"alignedQuality,omitempty"`

	// AlignedSequence: The bases of the read sequence contained in this
	// alignment record, without CIGAR operations applied. alignedSequence
	// and alignedQuality may be shorter than the full read sequence and
	// quality. This will occur if the alignment is part of a chimeric
	// alignment, or if the read was trimmed. When this occurs, the CIGAR
	// for this read will begin/end with a hard clip operator that will
	// indicate the length of the excised sequence.
	AlignedSequence string `json:"alignedSequence,omitempty"`

	// Alignment: The linear alignment for this alignment record. This field
	// will be null if the read is unmapped.
	Alignment *LinearAlignment `json:"alignment,omitempty"`

	// DuplicateFragment: The fragment is a PCR or optical duplicate (SAM
	// flag 0x400)
	DuplicateFragment bool `json:"duplicateFragment,omitempty"`

	// FailedVendorQualityChecks: SAM flag 0x200
	FailedVendorQualityChecks bool `json:"failedVendorQualityChecks,omitempty"`

	// FragmentLength: The observed length of the fragment, equivalent to
	// TLEN in SAM.
	FragmentLength int64 `json:"fragmentLength,omitempty"`

	// FragmentName: The fragment name. Equivalent to QNAME (query template
	// name) in SAM.
	FragmentName string `json:"fragmentName,omitempty"`

	// Id: The unique ID for this read. This is a generated unique ID, not
	// to be confused with fragmentName.
	Id string `json:"id,omitempty"`

	// Info: A map of additional read alignment information.
	Info map[string][]string `json:"info,omitempty"`

	// NextMatePosition: The mapping of the primary alignment of the
	// (readNumber+1)%numberReads read in the fragment. It replaces mate
	// position and mate strand in SAM.
	NextMatePosition *Position `json:"nextMatePosition,omitempty"`

	// NumberReads: The number of reads in the fragment (extension to SAM
	// flag 0x1).
	NumberReads int64 `json:"numberReads,omitempty"`

	// ProperPlacement: The orientation and the distance between reads from
	// the fragment are consistent with the sequencing protocol (SAM flag
	// 0x2)
	ProperPlacement bool `json:"properPlacement,omitempty"`

	// ReadGroupId: The ID of the read group this read belongs to. (Every
	// read must belong to exactly one read group.)
	ReadGroupId string `json:"readGroupId,omitempty"`

	// ReadGroupSetId: The ID of the read group set this read belongs to.
	// (Every read must belong to exactly one read group set.)
	ReadGroupSetId string `json:"readGroupSetId,omitempty"`

	// ReadNumber: The read number in sequencing. 0-based and less than
	// numberReads. This field replaces SAM flag 0x40 and 0x80.
	ReadNumber int64 `json:"readNumber,omitempty"`

	// SecondaryAlignment: Whether this alignment is secondary. Equivalent
	// to SAM flag 0x100. A secondary alignment represents an alternative to
	// the primary alignment for this read. Aligners may return secondary
	// alignments if a read can map ambiguously to multiple coordinates in
	// the genome. By convention, each read has one and only one alignment
	// where both secondaryAlignment and supplementaryAlignment are false.
	SecondaryAlignment bool `json:"secondaryAlignment,omitempty"`

	// SupplementaryAlignment: Whether this alignment is supplementary.
	// Equivalent to SAM flag 0x800. Supplementary alignments are used in
	// the representation of a chimeric alignment. In a chimeric alignment,
	// a read is split into multiple linear alignments that map to different
	// reference contigs. The first linear alignment in the read will be
	// designated as the representative alignment; the remaining linear
	// alignments will be designated as supplementary alignments. These
	// alignments may have different mapping quality scores. In each linear
	// alignment in a chimeric alignment, the read will be hard clipped. The
	// alignedSequence and alignedQuality fields in the alignment record
	// will only represent the bases for its respective linear alignment.
	SupplementaryAlignment bool `json:"supplementaryAlignment,omitempty"`
}

type ReadGroup struct {
	// DatasetId: The ID of the dataset this read group belongs to.
	DatasetId string `json:"datasetId,omitempty"`

	// Description: A free-form text description of this read group.
	Description string `json:"description,omitempty"`

	// Experiment: The experiment used to generate this read group.
	Experiment *ReadGroupExperiment `json:"experiment,omitempty"`

	// Id: The generated unique read group ID. Note: This is different than
	// the @RG ID field in the SAM spec. For that value, see the name field.
	Id string `json:"id,omitempty"`

	// Info: A map of additional read group information.
	Info map[string][]string `json:"info,omitempty"`

	// Name: The read group name. This corresponds to the @RG ID field in
	// the SAM spec.
	Name string `json:"name,omitempty"`

	// PredictedInsertSize: The predicted insert size of this read group.
	// The insert size is the length the sequenced DNA fragment from
	// end-to-end, not including the adapters.
	PredictedInsertSize int64 `json:"predictedInsertSize,omitempty"`

	// Programs: The programs used to generate this read group. Programs are
	// always identical for all read groups within a read group set. For
	// this reason, only the first read group in a returned set will have
	// this field populated.
	Programs []*ReadGroupProgram `json:"programs,omitempty"`

	// ReferenceSetId: The reference set the reads in this read group are
	// aligned to. Required if there are any read alignments.
	ReferenceSetId string `json:"referenceSetId,omitempty"`

	// SampleId: The sample this read group's data was generated from. Note:
	// This is not an actual ID within this repository, but rather an
	// identifier for a sample which may be meaningful to some external
	// system.
	SampleId string `json:"sampleId,omitempty"`
}

type ReadGroupExperiment struct {
	// InstrumentModel: The instrument model used as part of this
	// experiment. This maps to sequencing technology in BAM.
	InstrumentModel string `json:"instrumentModel,omitempty"`

	// LibraryId: The library used as part of this experiment. Note: This is
	// not an actual ID within this repository, but rather an identifier for
	// a library which may be meaningful to some external system.
	LibraryId string `json:"libraryId,omitempty"`

	// PlatformUnit: The platform unit used as part of this experiment e.g.
	// flowcell-barcode.lane for Illumina or slide for SOLiD. Corresponds to
	// the
	PlatformUnit string `json:"platformUnit,omitempty"`

	// SequencingCenter: The sequencing center used as part of this
	// experiment.
	SequencingCenter string `json:"sequencingCenter,omitempty"`
}

type ReadGroupProgram struct {
	// CommandLine: The command line used to run this program.
	CommandLine string `json:"commandLine,omitempty"`

	// Id: The user specified locally unique ID of the program. Used along
	// with prevProgramId to define an ordering between programs.
	Id string `json:"id,omitempty"`

	// Name: The name of the program.
	Name string `json:"name,omitempty"`

	// PrevProgramId: The ID of the program run before this one.
	PrevProgramId string `json:"prevProgramId,omitempty"`

	// Version: The version of the program run.
	Version string `json:"version,omitempty"`
}

type ReadGroupSet struct {
	// DatasetId: The dataset ID.
	DatasetId string `json:"datasetId,omitempty"`

	// Filename: The filename of the original source file for this read
	// group set, if any.
	Filename string `json:"filename,omitempty"`

	// Id: The read group set ID.
	Id string `json:"id,omitempty"`

	// Info: A map of additional read group set information.
	Info map[string][]string `json:"info,omitempty"`

	// Name: The read group set name. By default this will be initialized to
	// the sample name of the sequenced data contained in this set.
	Name string `json:"name,omitempty"`

	// ReadGroups: The read groups in this set. There are typically 1-10
	// read groups in a read group set.
	ReadGroups []*ReadGroup `json:"readGroups,omitempty"`

	// ReferenceSetId: The reference set the reads in this read group set
	// are aligned to.
	ReferenceSetId string `json:"referenceSetId,omitempty"`
}

type Reference struct {
	// Id: The Google generated immutable ID of the reference.
	Id string `json:"id,omitempty"`

	// Length: The length of this reference's sequence.
	Length int64 `json:"length,omitempty,string"`

	// Md5checksum: MD5 of the upper-case sequence excluding all whitespace
	// characters (this is equivalent to SQ:M5 in SAM). This value is
	// represented in lower case hexadecimal format.
	Md5checksum string `json:"md5checksum,omitempty"`

	// Name: The name of this reference, for example 22.
	Name string `json:"name,omitempty"`

	// NcbiTaxonId: ID from http://www.ncbi.nlm.nih.gov/taxonomy (e.g.
	// 9606->human) if not specified by the containing reference set.
	NcbiTaxonId int64 `json:"ncbiTaxonId,omitempty"`

	// SourceAccessions: All known corresponding accession IDs in INSDC
	// (GenBank/ENA/DDBJ) ideally with a version number, for example
	// GCF_000001405.26.
	SourceAccessions []string `json:"sourceAccessions,omitempty"`

	// SourceURI: The URI from which the sequence was obtained. Specifies a
	// FASTA format file/string with one name, sequence pair.
	SourceURI string `json:"sourceURI,omitempty"`
}

type ReferenceBound struct {
	// ReferenceName: The reference the bound is associate with.
	ReferenceName string `json:"referenceName,omitempty"`

	// UpperBound: An upper bound (inclusive) on the starting coordinate of
	// any variant in the reference sequence.
	UpperBound int64 `json:"upperBound,omitempty,string"`
}

type ReferenceSet struct {
	// AssemblyId: Public id of this reference set, such as GRCh37.
	AssemblyId string `json:"assemblyId,omitempty"`

	// Description: Free text description of this reference set.
	Description string `json:"description,omitempty"`

	// Id: The Google generated immutable ID of the reference set.
	Id string `json:"id,omitempty"`

	// Md5checksum: Order-independent MD5 checksum which identifies this
	// reference set. The checksum is computed by sorting all lower case
	// hexidecimal string reference.md5checksum (for all reference in this
	// set) in ascending lexicographic order, concatenating, and taking the
	// MD5 of that value. The resulting value is represented in lower case
	// hexadecimal format.
	Md5checksum string `json:"md5checksum,omitempty"`

	// NcbiTaxonId: ID from http://www.ncbi.nlm.nih.gov/taxonomy (e.g.
	// 9606->human) indicating the species which this assembly is intended
	// to model. Note that contained references may specify a different
	// ncbiTaxonId, as assemblies may contain reference sequences which do
	// not belong to the modeled species, e.g. EBV in a human reference
	// genome.
	NcbiTaxonId int64 `json:"ncbiTaxonId,omitempty"`

	// ReferenceIds: The IDs of the reference objects that are part of this
	// set. Reference.md5checksum must be unique within this set.
	ReferenceIds []string `json:"referenceIds,omitempty"`

	// SourceAccessions: All known corresponding accession IDs in INSDC
	// (GenBank/ENA/DDBJ) ideally with a version number, for example
	// NC_000001.11.
	SourceAccessions []string `json:"sourceAccessions,omitempty"`

	// SourceURI: The URI from which the references were obtained.
	SourceURI string `json:"sourceURI,omitempty"`
}

type SearchAnnotationSetsRequest struct {
	// DatasetIds: The dataset IDs to search within. Caller must have READ
	// access to these datasets.
	DatasetIds []string `json:"datasetIds,omitempty"`

	// Name: Only return annotations sets for which a substring of the name
	// matches this string (case insensitive).
	Name string `json:"name,omitempty"`

	// PageSize: Specifies number of results to return in a single page. If
	// unspecified, it will default to 128. The maximum value is 1024.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: The continuation token, which is used to page through
	// large result sets. To get the next page of results, set this
	// parameter to the value of nextPageToken from the previous response.
	PageToken string `json:"pageToken,omitempty"`

	// ReferenceSetId: If specified, only annotation sets associated with
	// the given reference set are returned.
	ReferenceSetId string `json:"referenceSetId,omitempty"`

	// Types: If specified, only annotation sets that have any of these
	// types are returned.
	Types []string `json:"types,omitempty"`
}

type SearchAnnotationSetsResponse struct {
	// AnnotationSets: The matching annotation sets.
	AnnotationSets []*AnnotationSet `json:"annotationSets,omitempty"`

	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type SearchAnnotationsRequest struct {
	// AnnotationSetIds: The annotation sets to search within. The caller
	// must have READ access to these annotation sets. Required. All queried
	// annotation sets must have the same type.
	AnnotationSetIds []string `json:"annotationSetIds,omitempty"`

	// PageSize: Specifies number of results to return in a single page. If
	// unspecified, it will default to 256. The maximum value is 2048.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: The continuation token, which is used to page through
	// large result sets. To get the next page of results, set this
	// parameter to the value of nextPageToken from the previous response.
	PageToken string `json:"pageToken,omitempty"`

	// Range: If specified, this query matches only annotations that overlap
	// this range.
	Range *QueryRange `json:"range,omitempty"`
}

type SearchAnnotationsResponse struct {
	// Annotations: The matching annotations.
	Annotations []*Annotation `json:"annotations,omitempty"`

	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type SearchCallSetsRequest struct {
	// Name: Only return call sets for which a substring of the name matches
	// this string.
	Name string `json:"name,omitempty"`

	// PageSize: The maximum number of call sets to return.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: The continuation token, which is used to page through
	// large result sets. To get the next page of results, set this
	// parameter to the value of nextPageToken from the previous response.
	PageToken string `json:"pageToken,omitempty"`

	// VariantSetIds: Restrict the query to call sets within the given
	// variant sets. At least one ID must be provided.
	VariantSetIds []string `json:"variantSetIds,omitempty"`
}

type SearchCallSetsResponse struct {
	// CallSets: The list of matching call sets.
	CallSets []*CallSet `json:"callSets,omitempty"`

	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type SearchJobsRequest struct {
	// CreatedAfter: If specified, only jobs created on or after this date,
	// given in milliseconds since Unix epoch, will be returned.
	CreatedAfter int64 `json:"createdAfter,omitempty,string"`

	// CreatedBefore: If specified, only jobs created prior to this date,
	// given in milliseconds since Unix epoch, will be returned.
	CreatedBefore int64 `json:"createdBefore,omitempty,string"`

	// PageSize: Specifies the number of results to return in a single page.
	// Defaults to 128. The maximum value is 256.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: The continuation token which is used to page through large
	// result sets. To get the next page of results, set this parameter to
	// the value of the nextPageToken from the previous response.
	PageToken string `json:"pageToken,omitempty"`

	// ProjectNumber: Required. Only return jobs which belong to this Google
	// Developers Console project.
	ProjectNumber int64 `json:"projectNumber,omitempty,string"`

	// Status: Only return jobs which have a matching status.
	Status []string `json:"status,omitempty"`
}

type SearchJobsResponse struct {
	// Jobs: The list of jobs results, ordered newest to oldest.
	Jobs []*Job `json:"jobs,omitempty"`

	// NextPageToken: The continuation token which is used to page through
	// large result sets. Provide this value is a subsequent request to
	// return the next page of results. This field will be empty if there
	// are no more results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type SearchReadGroupSetsRequest struct {
	// DatasetIds: Restricts this query to read group sets within the given
	// datasets. At least one ID must be provided.
	DatasetIds []string `json:"datasetIds,omitempty"`

	// Name: Only return read group sets for which a substring of the name
	// matches this string.
	Name string `json:"name,omitempty"`

	// PageSize: Specifies number of results to return in a single page. If
	// unspecified, it will default to 256. The maximum value is 1024.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: The continuation token, which is used to page through
	// large result sets. To get the next page of results, set this
	// parameter to the value of nextPageToken from the previous response.
	PageToken string `json:"pageToken,omitempty"`
}

type SearchReadGroupSetsResponse struct {
	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ReadGroupSets: The list of matching read group sets.
	ReadGroupSets []*ReadGroupSet `json:"readGroupSets,omitempty"`
}

type SearchReadsRequest struct {
	// End: The end position of the range on the reference, 0-based
	// exclusive. If specified, referenceName must also be specified.
	End int64 `json:"end,omitempty,string"`

	// PageSize: Specifies number of results to return in a single page. If
	// unspecified, it will default to 256. The maximum value is 2048.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: The continuation token, which is used to page through
	// large result sets. To get the next page of results, set this
	// parameter to the value of nextPageToken from the previous response.
	PageToken string `json:"pageToken,omitempty"`

	// ReadGroupIds: The IDs of the read groups within which to search for
	// reads. All specified read groups must belong to the same read group
	// sets. Must specify one of readGroupSetIds or readGroupIds.
	ReadGroupIds []string `json:"readGroupIds,omitempty"`

	// ReadGroupSetIds: The IDs of the read groups sets within which to
	// search for reads. All specified read group sets must be aligned
	// against a common set of reference sequences; this defines the genomic
	// coordinates for the query. Must specify one of readGroupSetIds or
	// readGroupIds.
	ReadGroupSetIds []string `json:"readGroupSetIds,omitempty"`

	// ReferenceName: The reference sequence name, for example chr1, 1, or
	// chrX. If set to *, only unmapped reads are returned.
	ReferenceName string `json:"referenceName,omitempty"`

	// Start: The start position of the range on the reference, 0-based
	// inclusive. If specified, referenceName must also be specified.
	Start int64 `json:"start,omitempty,string"`
}

type SearchReadsResponse struct {
	// Alignments: The list of matching alignments sorted by mapped genomic
	// coordinate, if any, ascending in position within the same reference.
	// Unmapped reads, which have no position, are returned last and are
	// further sorted in ascending lexicographic order by fragment name.
	Alignments []*Read `json:"alignments,omitempty"`

	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type SearchReferenceSetsRequest struct {
	// Accessions: If present, return references for which the accession
	// matches any of these strings. Best to give a version number, for
	// example GCF_000001405.26. If only the main accession number is given
	// then all records with that main accession will be returned, whichever
	// version. Note that different versions will have different sequences.
	Accessions []string `json:"accessions,omitempty"`

	// AssemblyId: If present, return reference sets for which a substring
	// of their assemblyId matches this string (case insensitive).
	AssemblyId string `json:"assemblyId,omitempty"`

	// Md5checksums: If present, return references for which the md5checksum
	// matches. See ReferenceSet.md5checksum for details.
	Md5checksums []string `json:"md5checksums,omitempty"`

	// PageSize: Specifies the maximum number of results to return in a
	// single page.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: The continuation token, which is used to page through
	// large result sets. To get the next page of results, set this
	// parameter to the value of nextPageToken from the previous response.
	PageToken string `json:"pageToken,omitempty"`
}

type SearchReferenceSetsResponse struct {
	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ReferenceSets: The matching references sets.
	ReferenceSets []*ReferenceSet `json:"referenceSets,omitempty"`
}

type SearchReferencesRequest struct {
	// Accessions: If present, return references for which the accession
	// matches this string. Best to give a version number, for example
	// GCF_000001405.26. If only the main accession number is given then all
	// records with that main accession will be returned, whichever version.
	// Note that different versions will have different sequences.
	Accessions []string `json:"accessions,omitempty"`

	// Md5checksums: If present, return references for which the md5checksum
	// matches. See Reference.md5checksum for construction details.
	Md5checksums []string `json:"md5checksums,omitempty"`

	// PageSize: Specifies the maximum number of results to return in a
	// single page.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: The continuation token, which is used to page through
	// large result sets. To get the next page of results, set this
	// parameter to the value of nextPageToken from the previous response.
	PageToken string `json:"pageToken,omitempty"`

	// ReferenceSetId: If present, return only references which belong to
	// this reference set.
	ReferenceSetId string `json:"referenceSetId,omitempty"`
}

type SearchReferencesResponse struct {
	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// References: The matching references.
	References []*Reference `json:"references,omitempty"`
}

type SearchVariantSetsRequest struct {
	// DatasetIds: Exactly one dataset ID must be provided here. Only
	// variant sets which belong to this dataset will be returned.
	DatasetIds []string `json:"datasetIds,omitempty"`

	// PageSize: The maximum number of variant sets to return in a request.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: The continuation token, which is used to page through
	// large result sets. To get the next page of results, set this
	// parameter to the value of nextPageToken from the previous response.
	PageToken string `json:"pageToken,omitempty"`
}

type SearchVariantSetsResponse struct {
	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// VariantSets: The variant sets belonging to the requested dataset.
	VariantSets []*VariantSet `json:"variantSets,omitempty"`
}

type SearchVariantsRequest struct {
	// CallSetIds: Only return variant calls which belong to call sets with
	// these ids. Leaving this blank returns all variant calls. If a variant
	// has no calls belonging to any of these call sets, it won't be
	// returned at all. Currently, variants with no calls from any call set
	// will never be returned.
	CallSetIds []string `json:"callSetIds,omitempty"`

	// End: Required. The end of the window (0-based, exclusive) for which
	// overlapping variants should be returned.
	End int64 `json:"end,omitempty,string"`

	// MaxCalls: The maximum number of calls to return. However, at least
	// one variant will always be returned, even if it has more calls than
	// this limit.
	MaxCalls int64 `json:"maxCalls,omitempty"`

	// PageSize: The maximum number of variants to return.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: The continuation token, which is used to page through
	// large result sets. To get the next page of results, set this
	// parameter to the value of nextPageToken from the previous response.
	PageToken string `json:"pageToken,omitempty"`

	// ReferenceName: Required. Only return variants in this reference
	// sequence.
	ReferenceName string `json:"referenceName,omitempty"`

	// Start: Required. The beginning of the window (0-based, inclusive) for
	// which overlapping variants should be returned.
	Start int64 `json:"start,omitempty,string"`

	// VariantName: Only return variants which have exactly this name.
	VariantName string `json:"variantName,omitempty"`

	// VariantSetIds: Exactly one variant set ID must be provided. Only
	// variants from this variant set will be returned.
	VariantSetIds []string `json:"variantSetIds,omitempty"`
}

type SearchVariantsResponse struct {
	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Variants: The list of matching Variants.
	Variants []*Variant `json:"variants,omitempty"`
}

type StreamVariantsRequest struct {
	// CallSetIds: Only return variant calls which belong to call sets with
	// these ids. Leaving this blank returns all variant calls.
	CallSetIds []string `json:"callSetIds,omitempty"`

	// End: The end of the window (0-based, exclusive) for which overlapping
	// variants should be returned.
	End int64 `json:"end,omitempty,string"`

	// ReferenceName: Required. Only return variants in this reference
	// sequence.
	ReferenceName string `json:"referenceName,omitempty"`

	// Start: The beginning of the window (0-based, inclusive) for which
	// overlapping variants should be returned.
	Start int64 `json:"start,omitempty,string"`

	// VariantSetIds: Exactly one variant set ID must be provided. Only
	// variants from this variant set will be returned.
	VariantSetIds []string `json:"variantSetIds,omitempty"`
}

type Transcript struct {
	// CodingSequence: The range of the coding sequence for this transcript,
	// if any. To determine the exact ranges of coding sequence, intersect
	// this range with those of the exons, if any. If there are any exons,
	// the codingSequence must start and end within them.
	//
	// Note that in some
	// cases, the reference genome will not exactly match the observed mRNA
	// transcript e.g. due to variance in the source genome from reference.
	// In these cases, exon.frame will not necessarily match the expected
	// reference reading frame and coding exon reference bases cannot
	// necessarily be concatenated to produce the original transcript mRNA.
	CodingSequence *TranscriptCodingSequence `json:"codingSequence,omitempty"`

	// Exons: The exons that compose this transcript. This field should be
	// unset for genomes where transcript splicing does not occur, for
	// example prokaryotes.
	//
	//
	// Introns are regions of the transcript that are
	// not included in the spliced RNA product. Though not explicitly
	// modeled here, intron ranges can be deduced; all regions of this
	// transcript that are not exons are introns.
	//
	//
	// Exonic sequences do not
	// necessarily code for a translational product (amino acids). Only the
	// regions of exons bounded by the codingSequence correspond to coding
	// DNA sequence.
	//
	//
	// Exons are ordered by start position and may not
	// overlap.
	Exons []*TranscriptExon `json:"exons,omitempty"`

	// GeneId: The annotation ID of the gene from which this transcript is
	// transcribed.
	GeneId string `json:"geneId,omitempty"`
}

type TranscriptCodingSequence struct {
	// End: The end of the coding sequence on this annotation's reference
	// sequence, 0-based exclusive. Note that this position is relative to
	// the reference start, and not the containing annotation start.
	End int64 `json:"end,omitempty,string"`

	// Start: The start of the coding sequence on this annotation's
	// reference sequence, 0-based inclusive. Note that this position is
	// relative to the reference start, and not the containing annotation
	// start.
	Start int64 `json:"start,omitempty,string"`
}

type TranscriptExon struct {
	// End: The end position of the exon on this annotation's reference
	// sequence, 0-based exclusive. Note that this is relative to the
	// reference start, and not the containing annotation start.
	End int64 `json:"end,omitempty,string"`

	// Frame: The frame of this exon. Contains a value of 0, 1, or 2, which
	// indicates the offset of the first coding base of the exon within the
	// reading frame of the coding DNA sequence, if any. This field is
	// dependent on the strandedness of this annotation (see
	// Annotation.position.reverseStrand). For forward stranded annotations,
	// this offset is relative to the exon.start. For reverse strand
	// annotations, this offset is relative to the exon.end-1.
	//
	// Unset if
	// this exon does not intersect the coding sequence. Upon creation of a
	// transcript, the frame must be populated for all or none of the coding
	// exons.
	Frame *Int32Value `json:"frame,omitempty"`

	// Start: The start position of the exon on this annotation's reference
	// sequence, 0-based inclusive. Note that this is relative to the
	// reference start, and not the containing annotation start.
	Start int64 `json:"start,omitempty,string"`
}

type Variant struct {
	// AlternateBases: The bases that appear instead of the reference bases.
	AlternateBases []string `json:"alternateBases,omitempty"`

	// Calls: The variant calls for this particular variant. Each one
	// represents the determination of genotype with respect to this
	// variant.
	Calls []*Call `json:"calls,omitempty"`

	// Created: The date this variant was created, in milliseconds from the
	// epoch.
	Created int64 `json:"created,omitempty,string"`

	// End: The end position (0-based) of this variant. This corresponds to
	// the first base after the last base in the reference allele. So, the
	// length of the reference allele is (end - start). This is useful for
	// variants that don't explicitly give alternate bases, for example
	// large deletions.
	End int64 `json:"end,omitempty,string"`

	// Filter: A list of filters (normally quality filters) this variant has
	// failed. PASS indicates this variant has passed all filters.
	Filter []string `json:"filter,omitempty"`

	// Id: The Google generated ID of the variant, immutable.
	Id string `json:"id,omitempty"`

	// Info: A map of additional variant information.
	Info map[string][]string `json:"info,omitempty"`

	// Names: Names for the variant, for example a RefSNP ID.
	Names []string `json:"names,omitempty"`

	// Quality: A measure of how likely this variant is to be real. A higher
	// value is better.
	Quality float64 `json:"quality,omitempty"`

	// ReferenceBases: The reference bases for this variant. They start at
	// the given position.
	ReferenceBases string `json:"referenceBases,omitempty"`

	// ReferenceName: The reference on which this variant occurs. (such as
	// chr20 or X)
	ReferenceName string `json:"referenceName,omitempty"`

	// Start: The position at which this variant occurs (0-based). This
	// corresponds to the first base of the string of reference bases.
	Start int64 `json:"start,omitempty,string"`

	// VariantSetId: The ID of the variant set this variant belongs to.
	VariantSetId string `json:"variantSetId,omitempty"`
}

type VariantAnnotation struct {
	// AlternateBases: The alternate allele for this variant. If multiple
	// alternate alleles exist at this location, create a separate variant
	// for each one, as they may represent distinct conditions.
	AlternateBases string `json:"alternateBases,omitempty"`

	// ClinicalSignificance: Describes the clinical significance of a
	// variant. It is adapted from the ClinVar controlled vocabulary for
	// clinical significance described at:
	// http://www.ncbi.nlm.nih.gov/clinvar/docs/clinsig/
	ClinicalSignificance string `json:"clinicalSignificance,omitempty"`

	// Conditions: The set of conditions associated with this variant. A
	// condition describes the way a variant influences human health.
	Conditions []*VariantAnnotationCondition `json:"conditions,omitempty"`

	// Effect: Effect of the variant on the coding sequence.
	Effect string `json:"effect,omitempty"`

	// GeneId: Google annotation ID of the gene affected by this variant.
	// This should be provided when the variant is created.
	GeneId string `json:"geneId,omitempty"`

	// TranscriptIds: Google annotation IDs of the transcripts affected by
	// this variant. These should be provided when the variant is created.
	TranscriptIds []string `json:"transcriptIds,omitempty"`

	// Type: Type has been adapted from ClinVar's list of variant types.
	Type string `json:"type,omitempty"`
}

type VariantAnnotationCondition struct {
	// ConceptId: The MedGen concept id associated with this gene. Search
	// for these IDs at http://www.ncbi.nlm.nih.gov/medgen/
	ConceptId string `json:"conceptId,omitempty"`

	// ExternalIds: The set of external IDs for this condition.
	ExternalIds []*ExternalId `json:"externalIds,omitempty"`

	// Names: A set of names for the condition.
	Names []string `json:"names,omitempty"`

	// OmimId: The OMIM id for this condition. Search for these IDs at
	// http://omim.org/
	OmimId string `json:"omimId,omitempty"`
}

type VariantSet struct {
	// DatasetId: The dataset to which this variant set belongs. Immutable.
	DatasetId string `json:"datasetId,omitempty"`

	// Id: The Google-generated ID of the variant set. Immutable.
	Id string `json:"id,omitempty"`

	// Metadata: The metadata associated with this variant set.
	Metadata []*Metadata `json:"metadata,omitempty"`

	// ReferenceBounds: A list of all references used by the variants in a
	// variant set with associated coordinate upper bounds for each one.
	ReferenceBounds []*ReferenceBound `json:"referenceBounds,omitempty"`
}

// method id "genomics.annotationSets.create":

type AnnotationSetsCreateCall struct {
	s             *Service
	annotationset *AnnotationSet
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Create: Creates a new annotation set. Caller must have WRITE
// permission for the associated dataset.

func (r *AnnotationSetsService) Create(annotationset *AnnotationSet) *AnnotationSetsCreateCall {
	return &AnnotationSetsCreateCall{
		s:             r.s,
		annotationset: annotationset,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "annotationSets",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AnnotationSetsCreateCall) Fields(s ...googleapi.Field) *AnnotationSetsCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AnnotationSetsCreateCall) Do() (*AnnotationSet, error) {
	var returnValue *AnnotationSet
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.annotationset,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a new annotation set. Caller must have WRITE permission for the associated dataset.",
	//   "httpMethod": "POST",
	//   "id": "genomics.annotationSets.create",
	//   "path": "annotationSets",
	//   "request": {
	//     "$ref": "AnnotationSet"
	//   },
	//   "response": {
	//     "$ref": "AnnotationSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.annotationSets.delete":

type AnnotationSetsDeleteCall struct {
	s               *Service
	annotationSetId string
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Delete: Deletes an annotation set. Caller must have WRITE permission
// for the associated annotation set.

func (r *AnnotationSetsService) Delete(annotationSetId string) *AnnotationSetsDeleteCall {
	return &AnnotationSetsDeleteCall{
		s:               r.s,
		annotationSetId: annotationSetId,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "annotationSets/{annotationSetId}",
	}
}

func (c *AnnotationSetsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"annotationSetId": c.annotationSetId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes an annotation set. Caller must have WRITE permission for the associated annotation set.",
	//   "httpMethod": "DELETE",
	//   "id": "genomics.annotationSets.delete",
	//   "parameterOrder": [
	//     "annotationSetId"
	//   ],
	//   "parameters": {
	//     "annotationSetId": {
	//       "description": "The ID of the annotation set to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "annotationSets/{annotationSetId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.annotationSets.get":

type AnnotationSetsGetCall struct {
	s               *Service
	annotationSetId string
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Get: Gets an annotation set. Caller must have READ permission for the
// associated dataset.

func (r *AnnotationSetsService) Get(annotationSetId string) *AnnotationSetsGetCall {
	return &AnnotationSetsGetCall{
		s:               r.s,
		annotationSetId: annotationSetId,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "annotationSets/{annotationSetId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AnnotationSetsGetCall) Fields(s ...googleapi.Field) *AnnotationSetsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AnnotationSetsGetCall) Do() (*AnnotationSet, error) {
	var returnValue *AnnotationSet
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"annotationSetId": c.annotationSetId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets an annotation set. Caller must have READ permission for the associated dataset.",
	//   "httpMethod": "GET",
	//   "id": "genomics.annotationSets.get",
	//   "parameterOrder": [
	//     "annotationSetId"
	//   ],
	//   "parameters": {
	//     "annotationSetId": {
	//       "description": "The ID of the annotation set to be retrieved.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "annotationSets/{annotationSetId}",
	//   "response": {
	//     "$ref": "AnnotationSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.annotationSets.patch":

type AnnotationSetsPatchCall struct {
	s               *Service
	annotationSetId string
	annotationset   *AnnotationSet
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Patch: Updates an annotation set. The update must respect all
// mutability restrictions and other invariants described on the
// annotation set resource. Caller must have WRITE permission for the
// associated dataset. This method supports patch semantics.

func (r *AnnotationSetsService) Patch(annotationSetId string, annotationset *AnnotationSet) *AnnotationSetsPatchCall {
	return &AnnotationSetsPatchCall{
		s:               r.s,
		annotationSetId: annotationSetId,
		annotationset:   annotationset,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "annotationSets/{annotationSetId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AnnotationSetsPatchCall) Fields(s ...googleapi.Field) *AnnotationSetsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AnnotationSetsPatchCall) Do() (*AnnotationSet, error) {
	var returnValue *AnnotationSet
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"annotationSetId": c.annotationSetId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.annotationset,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an annotation set. The update must respect all mutability restrictions and other invariants described on the annotation set resource. Caller must have WRITE permission for the associated dataset. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "genomics.annotationSets.patch",
	//   "parameterOrder": [
	//     "annotationSetId"
	//   ],
	//   "parameters": {
	//     "annotationSetId": {
	//       "description": "The ID of the annotation set to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "annotationSets/{annotationSetId}",
	//   "request": {
	//     "$ref": "AnnotationSet"
	//   },
	//   "response": {
	//     "$ref": "AnnotationSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.annotationSets.search":

type AnnotationSetsSearchCall struct {
	s                           *Service
	searchannotationsetsrequest *SearchAnnotationSetsRequest
	caller_                     googleapi.Caller
	params_                     url.Values
	pathTemplate_               string
}

// Search: Searches for annotation sets that match the given criteria.
// Results are returned in a deterministic order. Caller must have READ
// permission for the queried datasets.

func (r *AnnotationSetsService) Search(searchannotationsetsrequest *SearchAnnotationSetsRequest) *AnnotationSetsSearchCall {
	return &AnnotationSetsSearchCall{
		s: r.s,
		searchannotationsetsrequest: searchannotationsetsrequest,
		caller_:                     googleapi.JSONCall{},
		params_:                     make(map[string][]string),
		pathTemplate_:               "annotationSets/search",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AnnotationSetsSearchCall) Fields(s ...googleapi.Field) *AnnotationSetsSearchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AnnotationSetsSearchCall) Do() (*SearchAnnotationSetsResponse, error) {
	var returnValue *SearchAnnotationSetsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.searchannotationsetsrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Searches for annotation sets that match the given criteria. Results are returned in a deterministic order. Caller must have READ permission for the queried datasets.",
	//   "httpMethod": "POST",
	//   "id": "genomics.annotationSets.search",
	//   "path": "annotationSets/search",
	//   "request": {
	//     "$ref": "SearchAnnotationSetsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchAnnotationSetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.annotationSets.update":

type AnnotationSetsUpdateCall struct {
	s               *Service
	annotationSetId string
	annotationset   *AnnotationSet
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Update: Updates an annotation set. The update must respect all
// mutability restrictions and other invariants described on the
// annotation set resource. Caller must have WRITE permission for the
// associated dataset.

func (r *AnnotationSetsService) Update(annotationSetId string, annotationset *AnnotationSet) *AnnotationSetsUpdateCall {
	return &AnnotationSetsUpdateCall{
		s:               r.s,
		annotationSetId: annotationSetId,
		annotationset:   annotationset,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "annotationSets/{annotationSetId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AnnotationSetsUpdateCall) Fields(s ...googleapi.Field) *AnnotationSetsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AnnotationSetsUpdateCall) Do() (*AnnotationSet, error) {
	var returnValue *AnnotationSet
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"annotationSetId": c.annotationSetId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.annotationset,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an annotation set. The update must respect all mutability restrictions and other invariants described on the annotation set resource. Caller must have WRITE permission for the associated dataset.",
	//   "httpMethod": "PUT",
	//   "id": "genomics.annotationSets.update",
	//   "parameterOrder": [
	//     "annotationSetId"
	//   ],
	//   "parameters": {
	//     "annotationSetId": {
	//       "description": "The ID of the annotation set to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "annotationSets/{annotationSetId}",
	//   "request": {
	//     "$ref": "AnnotationSet"
	//   },
	//   "response": {
	//     "$ref": "AnnotationSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.annotations.batchCreate":

type AnnotationsBatchCreateCall struct {
	s                             *Service
	batchcreateannotationsrequest *BatchCreateAnnotationsRequest
	caller_                       googleapi.Caller
	params_                       url.Values
	pathTemplate_                 string
}

// BatchCreate: Creates one or more new annotations atomically. All
// annotations must belong to the same annotation set. Caller must have
// WRITE permission for this annotation set. For optimal performance,
// batch positionally adjacent annotations together.
//
//
// If the request
// has a systemic issue, such as an attempt to write to an inaccessible
// annotation set, the entire RPC will fail accordingly. For lesser data
// issues, when possible an error will be isolated to the corresponding
// batch entry in the response; the remaining well formed annotations
// will be created normally.

func (r *AnnotationsService) BatchCreate(batchcreateannotationsrequest *BatchCreateAnnotationsRequest) *AnnotationsBatchCreateCall {
	return &AnnotationsBatchCreateCall{
		s: r.s,
		batchcreateannotationsrequest: batchcreateannotationsrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "annotations:batchCreate",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AnnotationsBatchCreateCall) Fields(s ...googleapi.Field) *AnnotationsBatchCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AnnotationsBatchCreateCall) Do() (*BatchAnnotationsResponse, error) {
	var returnValue *BatchAnnotationsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.batchcreateannotationsrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates one or more new annotations atomically. All annotations must belong to the same annotation set. Caller must have WRITE permission for this annotation set. For optimal performance, batch positionally adjacent annotations together.\n\n\nIf the request has a systemic issue, such as an attempt to write to an inaccessible annotation set, the entire RPC will fail accordingly. For lesser data issues, when possible an error will be isolated to the corresponding batch entry in the response; the remaining well formed annotations will be created normally.",
	//   "httpMethod": "POST",
	//   "id": "genomics.annotations.batchCreate",
	//   "path": "annotations:batchCreate",
	//   "request": {
	//     "$ref": "BatchCreateAnnotationsRequest"
	//   },
	//   "response": {
	//     "$ref": "BatchAnnotationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.annotations.create":

type AnnotationsCreateCall struct {
	s             *Service
	annotation    *Annotation
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Create: Creates a new annotation. Caller must have WRITE permission
// for the associated annotation set.

func (r *AnnotationsService) Create(annotation *Annotation) *AnnotationsCreateCall {
	return &AnnotationsCreateCall{
		s:             r.s,
		annotation:    annotation,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "annotations",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AnnotationsCreateCall) Fields(s ...googleapi.Field) *AnnotationsCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AnnotationsCreateCall) Do() (*Annotation, error) {
	var returnValue *Annotation
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.annotation,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a new annotation. Caller must have WRITE permission for the associated annotation set.",
	//   "httpMethod": "POST",
	//   "id": "genomics.annotations.create",
	//   "path": "annotations",
	//   "request": {
	//     "$ref": "Annotation"
	//   },
	//   "response": {
	//     "$ref": "Annotation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.annotations.delete":

type AnnotationsDeleteCall struct {
	s             *Service
	annotationId  string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes an annotation. Caller must have WRITE permission for
// the associated annotation set.

func (r *AnnotationsService) Delete(annotationId string) *AnnotationsDeleteCall {
	return &AnnotationsDeleteCall{
		s:             r.s,
		annotationId:  annotationId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "annotations/{annotationId}",
	}
}

func (c *AnnotationsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"annotationId": c.annotationId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes an annotation. Caller must have WRITE permission for the associated annotation set.",
	//   "httpMethod": "DELETE",
	//   "id": "genomics.annotations.delete",
	//   "parameterOrder": [
	//     "annotationId"
	//   ],
	//   "parameters": {
	//     "annotationId": {
	//       "description": "The ID of the annotation set to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "annotations/{annotationId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.annotations.get":

type AnnotationsGetCall struct {
	s             *Service
	annotationId  string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets an annotation. Caller must have READ permission for the
// associated annotation set.

func (r *AnnotationsService) Get(annotationId string) *AnnotationsGetCall {
	return &AnnotationsGetCall{
		s:             r.s,
		annotationId:  annotationId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "annotations/{annotationId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AnnotationsGetCall) Fields(s ...googleapi.Field) *AnnotationsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AnnotationsGetCall) Do() (*Annotation, error) {
	var returnValue *Annotation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"annotationId": c.annotationId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets an annotation. Caller must have READ permission for the associated annotation set.",
	//   "httpMethod": "GET",
	//   "id": "genomics.annotations.get",
	//   "parameterOrder": [
	//     "annotationId"
	//   ],
	//   "parameters": {
	//     "annotationId": {
	//       "description": "The ID of the annotation set to be retrieved.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "annotations/{annotationId}",
	//   "response": {
	//     "$ref": "Annotation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.annotations.patch":

type AnnotationsPatchCall struct {
	s             *Service
	annotationId  string
	annotation    *Annotation
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates an annotation. The update must respect all mutability
// restrictions and other invariants described on the annotation
// resource. Caller must have WRITE permission for the associated
// dataset. This method supports patch semantics.

func (r *AnnotationsService) Patch(annotationId string, annotation *Annotation) *AnnotationsPatchCall {
	return &AnnotationsPatchCall{
		s:             r.s,
		annotationId:  annotationId,
		annotation:    annotation,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "annotations/{annotationId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AnnotationsPatchCall) Fields(s ...googleapi.Field) *AnnotationsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AnnotationsPatchCall) Do() (*Annotation, error) {
	var returnValue *Annotation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"annotationId": c.annotationId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.annotation,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an annotation. The update must respect all mutability restrictions and other invariants described on the annotation resource. Caller must have WRITE permission for the associated dataset. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "genomics.annotations.patch",
	//   "parameterOrder": [
	//     "annotationId"
	//   ],
	//   "parameters": {
	//     "annotationId": {
	//       "description": "The ID of the annotation set to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "annotations/{annotationId}",
	//   "request": {
	//     "$ref": "Annotation"
	//   },
	//   "response": {
	//     "$ref": "Annotation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.annotations.search":

type AnnotationsSearchCall struct {
	s                        *Service
	searchannotationsrequest *SearchAnnotationsRequest
	caller_                  googleapi.Caller
	params_                  url.Values
	pathTemplate_            string
}

// Search: Searches for annotations that match the given criteria.
// Results are returned ordered by start position. Annotations that have
// matching start positions are ordered deterministically. Caller must
// have READ permission for the queried annotation sets.

func (r *AnnotationsService) Search(searchannotationsrequest *SearchAnnotationsRequest) *AnnotationsSearchCall {
	return &AnnotationsSearchCall{
		s: r.s,
		searchannotationsrequest: searchannotationsrequest,
		caller_:                  googleapi.JSONCall{},
		params_:                  make(map[string][]string),
		pathTemplate_:            "annotations/search",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AnnotationsSearchCall) Fields(s ...googleapi.Field) *AnnotationsSearchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AnnotationsSearchCall) Do() (*SearchAnnotationsResponse, error) {
	var returnValue *SearchAnnotationsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.searchannotationsrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Searches for annotations that match the given criteria. Results are returned ordered by start position. Annotations that have matching start positions are ordered deterministically. Caller must have READ permission for the queried annotation sets.",
	//   "httpMethod": "POST",
	//   "id": "genomics.annotations.search",
	//   "path": "annotations/search",
	//   "request": {
	//     "$ref": "SearchAnnotationsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchAnnotationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.annotations.update":

type AnnotationsUpdateCall struct {
	s             *Service
	annotationId  string
	annotation    *Annotation
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates an annotation. The update must respect all mutability
// restrictions and other invariants described on the annotation
// resource. Caller must have WRITE permission for the associated
// dataset.

func (r *AnnotationsService) Update(annotationId string, annotation *Annotation) *AnnotationsUpdateCall {
	return &AnnotationsUpdateCall{
		s:             r.s,
		annotationId:  annotationId,
		annotation:    annotation,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "annotations/{annotationId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AnnotationsUpdateCall) Fields(s ...googleapi.Field) *AnnotationsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AnnotationsUpdateCall) Do() (*Annotation, error) {
	var returnValue *Annotation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"annotationId": c.annotationId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.annotation,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an annotation. The update must respect all mutability restrictions and other invariants described on the annotation resource. Caller must have WRITE permission for the associated dataset.",
	//   "httpMethod": "PUT",
	//   "id": "genomics.annotations.update",
	//   "parameterOrder": [
	//     "annotationId"
	//   ],
	//   "parameters": {
	//     "annotationId": {
	//       "description": "The ID of the annotation set to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "annotations/{annotationId}",
	//   "request": {
	//     "$ref": "Annotation"
	//   },
	//   "response": {
	//     "$ref": "Annotation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.callsets.create":

type CallsetsCreateCall struct {
	s             *Service
	callset       *CallSet
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Create: Creates a new call set.

func (r *CallsetsService) Create(callset *CallSet) *CallsetsCreateCall {
	return &CallsetsCreateCall{
		s:             r.s,
		callset:       callset,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "callsets",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CallsetsCreateCall) Fields(s ...googleapi.Field) *CallsetsCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CallsetsCreateCall) Do() (*CallSet, error) {
	var returnValue *CallSet
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.callset,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a new call set.",
	//   "httpMethod": "POST",
	//   "id": "genomics.callsets.create",
	//   "path": "callsets",
	//   "request": {
	//     "$ref": "CallSet"
	//   },
	//   "response": {
	//     "$ref": "CallSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.callsets.delete":

type CallsetsDeleteCall struct {
	s             *Service
	callSetId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a call set.

func (r *CallsetsService) Delete(callSetId string) *CallsetsDeleteCall {
	return &CallsetsDeleteCall{
		s:             r.s,
		callSetId:     callSetId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "callsets/{callSetId}",
	}
}

func (c *CallsetsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"callSetId": c.callSetId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a call set.",
	//   "httpMethod": "DELETE",
	//   "id": "genomics.callsets.delete",
	//   "parameterOrder": [
	//     "callSetId"
	//   ],
	//   "parameters": {
	//     "callSetId": {
	//       "description": "The ID of the call set to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "callsets/{callSetId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.callsets.get":

type CallsetsGetCall struct {
	s             *Service
	callSetId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a call set by ID.

func (r *CallsetsService) Get(callSetId string) *CallsetsGetCall {
	return &CallsetsGetCall{
		s:             r.s,
		callSetId:     callSetId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "callsets/{callSetId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CallsetsGetCall) Fields(s ...googleapi.Field) *CallsetsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CallsetsGetCall) Do() (*CallSet, error) {
	var returnValue *CallSet
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"callSetId": c.callSetId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a call set by ID.",
	//   "httpMethod": "GET",
	//   "id": "genomics.callsets.get",
	//   "parameterOrder": [
	//     "callSetId"
	//   ],
	//   "parameters": {
	//     "callSetId": {
	//       "description": "The ID of the call set.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "callsets/{callSetId}",
	//   "response": {
	//     "$ref": "CallSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.callsets.patch":

type CallsetsPatchCall struct {
	s             *Service
	callSetId     string
	callset       *CallSet
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates a call set. This method supports patch semantics.

func (r *CallsetsService) Patch(callSetId string, callset *CallSet) *CallsetsPatchCall {
	return &CallsetsPatchCall{
		s:             r.s,
		callSetId:     callSetId,
		callset:       callset,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "callsets/{callSetId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CallsetsPatchCall) Fields(s ...googleapi.Field) *CallsetsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CallsetsPatchCall) Do() (*CallSet, error) {
	var returnValue *CallSet
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"callSetId": c.callSetId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.callset,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a call set. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "genomics.callsets.patch",
	//   "parameterOrder": [
	//     "callSetId"
	//   ],
	//   "parameters": {
	//     "callSetId": {
	//       "description": "The ID of the call set to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "callsets/{callSetId}",
	//   "request": {
	//     "$ref": "CallSet"
	//   },
	//   "response": {
	//     "$ref": "CallSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.callsets.search":

type CallsetsSearchCall struct {
	s                     *Service
	searchcallsetsrequest *SearchCallSetsRequest
	caller_               googleapi.Caller
	params_               url.Values
	pathTemplate_         string
}

// Search: Gets a list of call sets matching the criteria.
//
// Implements
// GlobalAllianceApi.searchCallSets.

func (r *CallsetsService) Search(searchcallsetsrequest *SearchCallSetsRequest) *CallsetsSearchCall {
	return &CallsetsSearchCall{
		s: r.s,
		searchcallsetsrequest: searchcallsetsrequest,
		caller_:               googleapi.JSONCall{},
		params_:               make(map[string][]string),
		pathTemplate_:         "callsets/search",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CallsetsSearchCall) Fields(s ...googleapi.Field) *CallsetsSearchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CallsetsSearchCall) Do() (*SearchCallSetsResponse, error) {
	var returnValue *SearchCallSetsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.searchcallsetsrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a list of call sets matching the criteria.\n\nImplements GlobalAllianceApi.searchCallSets.",
	//   "httpMethod": "POST",
	//   "id": "genomics.callsets.search",
	//   "path": "callsets/search",
	//   "request": {
	//     "$ref": "SearchCallSetsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchCallSetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.callsets.update":

type CallsetsUpdateCall struct {
	s             *Service
	callSetId     string
	callset       *CallSet
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates a call set.

func (r *CallsetsService) Update(callSetId string, callset *CallSet) *CallsetsUpdateCall {
	return &CallsetsUpdateCall{
		s:             r.s,
		callSetId:     callSetId,
		callset:       callset,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "callsets/{callSetId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CallsetsUpdateCall) Fields(s ...googleapi.Field) *CallsetsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CallsetsUpdateCall) Do() (*CallSet, error) {
	var returnValue *CallSet
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"callSetId": c.callSetId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.callset,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a call set.",
	//   "httpMethod": "PUT",
	//   "id": "genomics.callsets.update",
	//   "parameterOrder": [
	//     "callSetId"
	//   ],
	//   "parameters": {
	//     "callSetId": {
	//       "description": "The ID of the call set to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "callsets/{callSetId}",
	//   "request": {
	//     "$ref": "CallSet"
	//   },
	//   "response": {
	//     "$ref": "CallSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.datasets.create":

type DatasetsCreateCall struct {
	s             *Service
	dataset       *Dataset
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Create: Creates a new dataset.

func (r *DatasetsService) Create(dataset *Dataset) *DatasetsCreateCall {
	return &DatasetsCreateCall{
		s:             r.s,
		dataset:       dataset,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "datasets",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsCreateCall) Fields(s ...googleapi.Field) *DatasetsCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatasetsCreateCall) Do() (*Dataset, error) {
	var returnValue *Dataset
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.dataset,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a new dataset.",
	//   "httpMethod": "POST",
	//   "id": "genomics.datasets.create",
	//   "path": "datasets",
	//   "request": {
	//     "$ref": "Dataset"
	//   },
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.datasets.delete":

type DatasetsDeleteCall struct {
	s             *Service
	datasetId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a dataset.

func (r *DatasetsService) Delete(datasetId string) *DatasetsDeleteCall {
	return &DatasetsDeleteCall{
		s:             r.s,
		datasetId:     datasetId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "datasets/{datasetId}",
	}
}

func (c *DatasetsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"datasetId": c.datasetId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a dataset.",
	//   "httpMethod": "DELETE",
	//   "id": "genomics.datasets.delete",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "The ID of the dataset to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "datasets/{datasetId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.datasets.get":

type DatasetsGetCall struct {
	s             *Service
	datasetId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a dataset by ID.

func (r *DatasetsService) Get(datasetId string) *DatasetsGetCall {
	return &DatasetsGetCall{
		s:             r.s,
		datasetId:     datasetId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "datasets/{datasetId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsGetCall) Fields(s ...googleapi.Field) *DatasetsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatasetsGetCall) Do() (*Dataset, error) {
	var returnValue *Dataset
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"datasetId": c.datasetId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a dataset by ID.",
	//   "httpMethod": "GET",
	//   "id": "genomics.datasets.get",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "The ID of the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "datasets/{datasetId}",
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.datasets.list":

type DatasetsListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists datasets within a project.

func (r *DatasetsService) List() *DatasetsListCall {
	return &DatasetsListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "datasets",
	}
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results returned by this request. If unspecified, defaults to 50.
func (c *DatasetsListCall) PageSize(pageSize int64) *DatasetsListCall {
	c.params_.Set("pageSize", fmt.Sprintf("%v", pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, which is used to page through large result sets. To get the
// next page of results, set this parameter to the value of
// nextPageToken from the previous response.
func (c *DatasetsListCall) PageToken(pageToken string) *DatasetsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// ProjectNumber sets the optional parameter "projectNumber": The
// project to list datasets for.
func (c *DatasetsListCall) ProjectNumber(projectNumber int64) *DatasetsListCall {
	c.params_.Set("projectNumber", fmt.Sprintf("%v", projectNumber))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsListCall) Fields(s ...googleapi.Field) *DatasetsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatasetsListCall) Do() (*ListDatasetsResponse, error) {
	var returnValue *ListDatasetsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists datasets within a project.",
	//   "httpMethod": "GET",
	//   "id": "genomics.datasets.list",
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of results returned by this request. If unspecified, defaults to 50.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of nextPageToken from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectNumber": {
	//       "description": "The project to list datasets for.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "datasets",
	//   "response": {
	//     "$ref": "ListDatasetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.datasets.patch":

type DatasetsPatchCall struct {
	s             *Service
	datasetId     string
	dataset       *Dataset
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates a dataset. This method supports patch semantics.

func (r *DatasetsService) Patch(datasetId string, dataset *Dataset) *DatasetsPatchCall {
	return &DatasetsPatchCall{
		s:             r.s,
		datasetId:     datasetId,
		dataset:       dataset,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "datasets/{datasetId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsPatchCall) Fields(s ...googleapi.Field) *DatasetsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatasetsPatchCall) Do() (*Dataset, error) {
	var returnValue *Dataset
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"datasetId": c.datasetId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.dataset,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a dataset. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "genomics.datasets.patch",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "The ID of the dataset to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "datasets/{datasetId}",
	//   "request": {
	//     "$ref": "Dataset"
	//   },
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.datasets.undelete":

type DatasetsUndeleteCall struct {
	s             *Service
	datasetId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Undelete: Undeletes a dataset by restoring a dataset which was
// deleted via this API. This operation is only possible for a week
// after the deletion occurred.

func (r *DatasetsService) Undelete(datasetId string) *DatasetsUndeleteCall {
	return &DatasetsUndeleteCall{
		s:             r.s,
		datasetId:     datasetId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "datasets/{datasetId}/undelete",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsUndeleteCall) Fields(s ...googleapi.Field) *DatasetsUndeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatasetsUndeleteCall) Do() (*Dataset, error) {
	var returnValue *Dataset
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"datasetId": c.datasetId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Undeletes a dataset by restoring a dataset which was deleted via this API. This operation is only possible for a week after the deletion occurred.",
	//   "httpMethod": "POST",
	//   "id": "genomics.datasets.undelete",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "The ID of the dataset to be undeleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "datasets/{datasetId}/undelete",
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.datasets.update":

type DatasetsUpdateCall struct {
	s             *Service
	datasetId     string
	dataset       *Dataset
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates a dataset.

func (r *DatasetsService) Update(datasetId string, dataset *Dataset) *DatasetsUpdateCall {
	return &DatasetsUpdateCall{
		s:             r.s,
		datasetId:     datasetId,
		dataset:       dataset,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "datasets/{datasetId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsUpdateCall) Fields(s ...googleapi.Field) *DatasetsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatasetsUpdateCall) Do() (*Dataset, error) {
	var returnValue *Dataset
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"datasetId": c.datasetId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.dataset,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a dataset.",
	//   "httpMethod": "PUT",
	//   "id": "genomics.datasets.update",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "The ID of the dataset to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "datasets/{datasetId}",
	//   "request": {
	//     "$ref": "Dataset"
	//   },
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.experimental.jobs.create":

type ExperimentalJobsCreateCall struct {
	s                            *Service
	experimentalcreatejobrequest *ExperimentalCreateJobRequest
	caller_                      googleapi.Caller
	params_                      url.Values
	pathTemplate_                string
}

// Create: Creates and asynchronously runs an ad-hoc job. This is an
// experimental call and may be removed or changed at any time.

func (r *ExperimentalJobsService) Create(experimentalcreatejobrequest *ExperimentalCreateJobRequest) *ExperimentalJobsCreateCall {
	return &ExperimentalJobsCreateCall{
		s: r.s,
		experimentalcreatejobrequest: experimentalcreatejobrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "experimental/jobs/create",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ExperimentalJobsCreateCall) Fields(s ...googleapi.Field) *ExperimentalJobsCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ExperimentalJobsCreateCall) Do() (*ExperimentalCreateJobResponse, error) {
	var returnValue *ExperimentalCreateJobResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.experimentalcreatejobrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates and asynchronously runs an ad-hoc job. This is an experimental call and may be removed or changed at any time.",
	//   "httpMethod": "POST",
	//   "id": "genomics.experimental.jobs.create",
	//   "path": "experimental/jobs/create",
	//   "request": {
	//     "$ref": "ExperimentalCreateJobRequest"
	//   },
	//   "response": {
	//     "$ref": "ExperimentalCreateJobResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.read_write",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.jobs.cancel":

type JobsCancelCall struct {
	s             *Service
	jobId         string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Cancel: Cancels a job by ID. Note that it is possible for partial
// results to be generated and stored for cancelled jobs.

func (r *JobsService) Cancel(jobId string) *JobsCancelCall {
	return &JobsCancelCall{
		s:             r.s,
		jobId:         jobId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "jobs/{jobId}/cancel",
	}
}

func (c *JobsCancelCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"jobId": c.jobId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Cancels a job by ID. Note that it is possible for partial results to be generated and stored for cancelled jobs.",
	//   "httpMethod": "POST",
	//   "id": "genomics.jobs.cancel",
	//   "parameterOrder": [
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "Required. The ID of the job.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "jobs/{jobId}/cancel",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.jobs.get":

type JobsGetCall struct {
	s             *Service
	jobId         string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a job by ID.

func (r *JobsService) Get(jobId string) *JobsGetCall {
	return &JobsGetCall{
		s:             r.s,
		jobId:         jobId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "jobs/{jobId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsGetCall) Fields(s ...googleapi.Field) *JobsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *JobsGetCall) Do() (*Job, error) {
	var returnValue *Job
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"jobId": c.jobId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a job by ID.",
	//   "httpMethod": "GET",
	//   "id": "genomics.jobs.get",
	//   "parameterOrder": [
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "Required. The ID of the job.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "jobs/{jobId}",
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.jobs.search":

type JobsSearchCall struct {
	s                 *Service
	searchjobsrequest *SearchJobsRequest
	caller_           googleapi.Caller
	params_           url.Values
	pathTemplate_     string
}

// Search: Gets a list of jobs matching the criteria.

func (r *JobsService) Search(searchjobsrequest *SearchJobsRequest) *JobsSearchCall {
	return &JobsSearchCall{
		s:                 r.s,
		searchjobsrequest: searchjobsrequest,
		caller_:           googleapi.JSONCall{},
		params_:           make(map[string][]string),
		pathTemplate_:     "jobs/search",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsSearchCall) Fields(s ...googleapi.Field) *JobsSearchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *JobsSearchCall) Do() (*SearchJobsResponse, error) {
	var returnValue *SearchJobsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.searchjobsrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a list of jobs matching the criteria.",
	//   "httpMethod": "POST",
	//   "id": "genomics.jobs.search",
	//   "path": "jobs/search",
	//   "request": {
	//     "$ref": "SearchJobsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchJobsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.readgroupsets.align":

type ReadgroupsetsAlignCall struct {
	s                         *Service
	alignreadgroupsetsrequest *AlignReadGroupSetsRequest
	caller_                   googleapi.Caller
	params_                   url.Values
	pathTemplate_             string
}

// Align: Aligns read data from existing read group sets or files from
// Google Cloud Storage. See the  alignment and variant calling
// documentation for more details.

func (r *ReadgroupsetsService) Align(alignreadgroupsetsrequest *AlignReadGroupSetsRequest) *ReadgroupsetsAlignCall {
	return &ReadgroupsetsAlignCall{
		s: r.s,
		alignreadgroupsetsrequest: alignreadgroupsetsrequest,
		caller_:                   googleapi.JSONCall{},
		params_:                   make(map[string][]string),
		pathTemplate_:             "readgroupsets/align",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReadgroupsetsAlignCall) Fields(s ...googleapi.Field) *ReadgroupsetsAlignCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReadgroupsetsAlignCall) Do() (*AlignReadGroupSetsResponse, error) {
	var returnValue *AlignReadGroupSetsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.alignreadgroupsetsrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Aligns read data from existing read group sets or files from Google Cloud Storage. See the  alignment and variant calling documentation for more details.",
	//   "httpMethod": "POST",
	//   "id": "genomics.readgroupsets.align",
	//   "path": "readgroupsets/align",
	//   "request": {
	//     "$ref": "AlignReadGroupSetsRequest"
	//   },
	//   "response": {
	//     "$ref": "AlignReadGroupSetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.read_write",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.readgroupsets.call":

type ReadgroupsetsCallCall struct {
	s                        *Service
	callreadgroupsetsrequest *CallReadGroupSetsRequest
	caller_                  googleapi.Caller
	params_                  url.Values
	pathTemplate_            string
}

// Call: Calls variants on read data from existing read group sets or
// files from Google Cloud Storage. See the  alignment and variant
// calling documentation for more details.

func (r *ReadgroupsetsService) Call(callreadgroupsetsrequest *CallReadGroupSetsRequest) *ReadgroupsetsCallCall {
	return &ReadgroupsetsCallCall{
		s: r.s,
		callreadgroupsetsrequest: callreadgroupsetsrequest,
		caller_:                  googleapi.JSONCall{},
		params_:                  make(map[string][]string),
		pathTemplate_:            "readgroupsets/call",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReadgroupsetsCallCall) Fields(s ...googleapi.Field) *ReadgroupsetsCallCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReadgroupsetsCallCall) Do() (*CallReadGroupSetsResponse, error) {
	var returnValue *CallReadGroupSetsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.callreadgroupsetsrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Calls variants on read data from existing read group sets or files from Google Cloud Storage. See the  alignment and variant calling documentation for more details.",
	//   "httpMethod": "POST",
	//   "id": "genomics.readgroupsets.call",
	//   "path": "readgroupsets/call",
	//   "request": {
	//     "$ref": "CallReadGroupSetsRequest"
	//   },
	//   "response": {
	//     "$ref": "CallReadGroupSetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.read_write",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.readgroupsets.delete":

type ReadgroupsetsDeleteCall struct {
	s              *Service
	readGroupSetId string
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Delete: Deletes a read group set.

func (r *ReadgroupsetsService) Delete(readGroupSetId string) *ReadgroupsetsDeleteCall {
	return &ReadgroupsetsDeleteCall{
		s:              r.s,
		readGroupSetId: readGroupSetId,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "readgroupsets/{readGroupSetId}",
	}
}

func (c *ReadgroupsetsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"readGroupSetId": c.readGroupSetId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a read group set.",
	//   "httpMethod": "DELETE",
	//   "id": "genomics.readgroupsets.delete",
	//   "parameterOrder": [
	//     "readGroupSetId"
	//   ],
	//   "parameters": {
	//     "readGroupSetId": {
	//       "description": "The ID of the read group set to be deleted. The caller must have WRITE permissions to the dataset associated with this read group set.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "readgroupsets/{readGroupSetId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.readgroupsets.export":

type ReadgroupsetsExportCall struct {
	s                          *Service
	exportreadgroupsetsrequest *ExportReadGroupSetsRequest
	caller_                    googleapi.Caller
	params_                    url.Values
	pathTemplate_              string
}

// Export: Exports read group sets to a BAM file in Google Cloud
// Storage.
//
// Note that currently there may be some differences between
// exported BAM files and the original BAM file at the time of import.
// In particular, comments in the input file header will not be
// preserved, some custom tags will be converted to strings, and
// original reference sequence order is not necessarily preserved.

func (r *ReadgroupsetsService) Export(exportreadgroupsetsrequest *ExportReadGroupSetsRequest) *ReadgroupsetsExportCall {
	return &ReadgroupsetsExportCall{
		s: r.s,
		exportreadgroupsetsrequest: exportreadgroupsetsrequest,
		caller_:                    googleapi.JSONCall{},
		params_:                    make(map[string][]string),
		pathTemplate_:              "readgroupsets/export",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReadgroupsetsExportCall) Fields(s ...googleapi.Field) *ReadgroupsetsExportCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReadgroupsetsExportCall) Do() (*ExportReadGroupSetsResponse, error) {
	var returnValue *ExportReadGroupSetsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.exportreadgroupsetsrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Exports read group sets to a BAM file in Google Cloud Storage.\n\nNote that currently there may be some differences between exported BAM files and the original BAM file at the time of import. In particular, comments in the input file header will not be preserved, some custom tags will be converted to strings, and original reference sequence order is not necessarily preserved.",
	//   "httpMethod": "POST",
	//   "id": "genomics.readgroupsets.export",
	//   "path": "readgroupsets/export",
	//   "request": {
	//     "$ref": "ExportReadGroupSetsRequest"
	//   },
	//   "response": {
	//     "$ref": "ExportReadGroupSetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.read_write",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.readgroupsets.get":

type ReadgroupsetsGetCall struct {
	s              *Service
	readGroupSetId string
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Get: Gets a read group set by ID.

func (r *ReadgroupsetsService) Get(readGroupSetId string) *ReadgroupsetsGetCall {
	return &ReadgroupsetsGetCall{
		s:              r.s,
		readGroupSetId: readGroupSetId,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "readgroupsets/{readGroupSetId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReadgroupsetsGetCall) Fields(s ...googleapi.Field) *ReadgroupsetsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReadgroupsetsGetCall) Do() (*ReadGroupSet, error) {
	var returnValue *ReadGroupSet
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"readGroupSetId": c.readGroupSetId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a read group set by ID.",
	//   "httpMethod": "GET",
	//   "id": "genomics.readgroupsets.get",
	//   "parameterOrder": [
	//     "readGroupSetId"
	//   ],
	//   "parameters": {
	//     "readGroupSetId": {
	//       "description": "The ID of the read group set.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "readgroupsets/{readGroupSetId}",
	//   "response": {
	//     "$ref": "ReadGroupSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.readgroupsets.import":

type ReadgroupsetsImportCall struct {
	s                          *Service
	importreadgroupsetsrequest *ImportReadGroupSetsRequest
	caller_                    googleapi.Caller
	params_                    url.Values
	pathTemplate_              string
}

// Import: Creates read group sets by asynchronously importing the
// provided information.
//
// Note that currently comments in the input file
// header are not imported and some custom tags will be converted to
// strings, rather than preserving tag types. The caller must have WRITE
// permissions to the dataset.

func (r *ReadgroupsetsService) Import(importreadgroupsetsrequest *ImportReadGroupSetsRequest) *ReadgroupsetsImportCall {
	return &ReadgroupsetsImportCall{
		s: r.s,
		importreadgroupsetsrequest: importreadgroupsetsrequest,
		caller_:                    googleapi.JSONCall{},
		params_:                    make(map[string][]string),
		pathTemplate_:              "readgroupsets/import",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReadgroupsetsImportCall) Fields(s ...googleapi.Field) *ReadgroupsetsImportCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReadgroupsetsImportCall) Do() (*ImportReadGroupSetsResponse, error) {
	var returnValue *ImportReadGroupSetsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.importreadgroupsetsrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates read group sets by asynchronously importing the provided information.\n\nNote that currently comments in the input file header are not imported and some custom tags will be converted to strings, rather than preserving tag types. The caller must have WRITE permissions to the dataset.",
	//   "httpMethod": "POST",
	//   "id": "genomics.readgroupsets.import",
	//   "path": "readgroupsets/import",
	//   "request": {
	//     "$ref": "ImportReadGroupSetsRequest"
	//   },
	//   "response": {
	//     "$ref": "ImportReadGroupSetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.read_write",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.readgroupsets.patch":

type ReadgroupsetsPatchCall struct {
	s              *Service
	readGroupSetId string
	readgroupset   *ReadGroupSet
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Patch: Updates a read group set. This method supports patch
// semantics.

func (r *ReadgroupsetsService) Patch(readGroupSetId string, readgroupset *ReadGroupSet) *ReadgroupsetsPatchCall {
	return &ReadgroupsetsPatchCall{
		s:              r.s,
		readGroupSetId: readGroupSetId,
		readgroupset:   readgroupset,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "readgroupsets/{readGroupSetId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReadgroupsetsPatchCall) Fields(s ...googleapi.Field) *ReadgroupsetsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReadgroupsetsPatchCall) Do() (*ReadGroupSet, error) {
	var returnValue *ReadGroupSet
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"readGroupSetId": c.readGroupSetId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.readgroupset,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a read group set. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "genomics.readgroupsets.patch",
	//   "parameterOrder": [
	//     "readGroupSetId"
	//   ],
	//   "parameters": {
	//     "readGroupSetId": {
	//       "description": "The ID of the read group set to be updated. The caller must have WRITE permissions to the dataset associated with this read group set.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "readgroupsets/{readGroupSetId}",
	//   "request": {
	//     "$ref": "ReadGroupSet"
	//   },
	//   "response": {
	//     "$ref": "ReadGroupSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.readgroupsets.search":

type ReadgroupsetsSearchCall struct {
	s                          *Service
	searchreadgroupsetsrequest *SearchReadGroupSetsRequest
	caller_                    googleapi.Caller
	params_                    url.Values
	pathTemplate_              string
}

// Search: Searches for read group sets matching the
// criteria.
//
// Implements GlobalAllianceApi.searchReadGroupSets.

func (r *ReadgroupsetsService) Search(searchreadgroupsetsrequest *SearchReadGroupSetsRequest) *ReadgroupsetsSearchCall {
	return &ReadgroupsetsSearchCall{
		s: r.s,
		searchreadgroupsetsrequest: searchreadgroupsetsrequest,
		caller_:                    googleapi.JSONCall{},
		params_:                    make(map[string][]string),
		pathTemplate_:              "readgroupsets/search",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReadgroupsetsSearchCall) Fields(s ...googleapi.Field) *ReadgroupsetsSearchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReadgroupsetsSearchCall) Do() (*SearchReadGroupSetsResponse, error) {
	var returnValue *SearchReadGroupSetsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.searchreadgroupsetsrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Searches for read group sets matching the criteria.\n\nImplements GlobalAllianceApi.searchReadGroupSets.",
	//   "httpMethod": "POST",
	//   "id": "genomics.readgroupsets.search",
	//   "path": "readgroupsets/search",
	//   "request": {
	//     "$ref": "SearchReadGroupSetsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchReadGroupSetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.readgroupsets.update":

type ReadgroupsetsUpdateCall struct {
	s              *Service
	readGroupSetId string
	readgroupset   *ReadGroupSet
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Update: Updates a read group set.

func (r *ReadgroupsetsService) Update(readGroupSetId string, readgroupset *ReadGroupSet) *ReadgroupsetsUpdateCall {
	return &ReadgroupsetsUpdateCall{
		s:              r.s,
		readGroupSetId: readGroupSetId,
		readgroupset:   readgroupset,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "readgroupsets/{readGroupSetId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReadgroupsetsUpdateCall) Fields(s ...googleapi.Field) *ReadgroupsetsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReadgroupsetsUpdateCall) Do() (*ReadGroupSet, error) {
	var returnValue *ReadGroupSet
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"readGroupSetId": c.readGroupSetId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.readgroupset,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a read group set.",
	//   "httpMethod": "PUT",
	//   "id": "genomics.readgroupsets.update",
	//   "parameterOrder": [
	//     "readGroupSetId"
	//   ],
	//   "parameters": {
	//     "readGroupSetId": {
	//       "description": "The ID of the read group set to be updated. The caller must have WRITE permissions to the dataset associated with this read group set.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "readgroupsets/{readGroupSetId}",
	//   "request": {
	//     "$ref": "ReadGroupSet"
	//   },
	//   "response": {
	//     "$ref": "ReadGroupSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.readgroupsets.coveragebuckets.list":

type ReadgroupsetsCoveragebucketsListCall struct {
	s              *Service
	readGroupSetId string
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// List: Lists fixed width coverage buckets for a read group set, each
// of which correspond to a range of a reference sequence. Each bucket
// summarizes coverage information across its corresponding genomic
// range.
//
// Coverage is defined as the number of reads which are aligned
// to a given base in the reference sequence. Coverage buckets are
// available at several precomputed bucket widths, enabling retrieval of
// various coverage 'zoom levels'. The caller must have READ permissions
// for the target read group set.

func (r *ReadgroupsetsCoveragebucketsService) List(readGroupSetId string) *ReadgroupsetsCoveragebucketsListCall {
	return &ReadgroupsetsCoveragebucketsListCall{
		s:              r.s,
		readGroupSetId: readGroupSetId,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "readgroupsets/{readGroupSetId}/coveragebuckets",
	}
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return in a single page. If unspecified, defaults to
// 1024. The maximum value is 2048.
func (c *ReadgroupsetsCoveragebucketsListCall) PageSize(pageSize int64) *ReadgroupsetsCoveragebucketsListCall {
	c.params_.Set("pageSize", fmt.Sprintf("%v", pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, which is used to page through large result sets. To get the
// next page of results, set this parameter to the value of
// nextPageToken from the previous response.
func (c *ReadgroupsetsCoveragebucketsListCall) PageToken(pageToken string) *ReadgroupsetsCoveragebucketsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// RangeEnd sets the optional parameter "range.end": The end position of
// the range on the reference, 0-based exclusive. If specified,
// referenceName must also be specified.
func (c *ReadgroupsetsCoveragebucketsListCall) RangeEnd(rangeEnd int64) *ReadgroupsetsCoveragebucketsListCall {
	c.params_.Set("range.end", fmt.Sprintf("%v", rangeEnd))
	return c
}

// RangeReferenceName sets the optional parameter "range.referenceName":
// The reference sequence name, for example chr1, 1, or chrX.
func (c *ReadgroupsetsCoveragebucketsListCall) RangeReferenceName(rangeReferenceName string) *ReadgroupsetsCoveragebucketsListCall {
	c.params_.Set("range.referenceName", fmt.Sprintf("%v", rangeReferenceName))
	return c
}

// RangeStart sets the optional parameter "range.start": The start
// position of the range on the reference, 0-based inclusive. If
// specified, referenceName must also be specified.
func (c *ReadgroupsetsCoveragebucketsListCall) RangeStart(rangeStart int64) *ReadgroupsetsCoveragebucketsListCall {
	c.params_.Set("range.start", fmt.Sprintf("%v", rangeStart))
	return c
}

// TargetBucketWidth sets the optional parameter "targetBucketWidth":
// The desired width of each reported coverage bucket in base pairs.
// This will be rounded down to the nearest precomputed bucket width;
// the value of which is returned as bucketWidth in the response.
// Defaults to infinity (each bucket spans an entire reference sequence)
// or the length of the target range, if specified. The smallest
// precomputed bucketWidth is currently 2048 base pairs; this is subject
// to change.
func (c *ReadgroupsetsCoveragebucketsListCall) TargetBucketWidth(targetBucketWidth int64) *ReadgroupsetsCoveragebucketsListCall {
	c.params_.Set("targetBucketWidth", fmt.Sprintf("%v", targetBucketWidth))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReadgroupsetsCoveragebucketsListCall) Fields(s ...googleapi.Field) *ReadgroupsetsCoveragebucketsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReadgroupsetsCoveragebucketsListCall) Do() (*ListCoverageBucketsResponse, error) {
	var returnValue *ListCoverageBucketsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"readGroupSetId": c.readGroupSetId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists fixed width coverage buckets for a read group set, each of which correspond to a range of a reference sequence. Each bucket summarizes coverage information across its corresponding genomic range.\n\nCoverage is defined as the number of reads which are aligned to a given base in the reference sequence. Coverage buckets are available at several precomputed bucket widths, enabling retrieval of various coverage 'zoom levels'. The caller must have READ permissions for the target read group set.",
	//   "httpMethod": "GET",
	//   "id": "genomics.readgroupsets.coveragebuckets.list",
	//   "parameterOrder": [
	//     "readGroupSetId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of results to return in a single page. If unspecified, defaults to 1024. The maximum value is 2048.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of nextPageToken from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "range.end": {
	//       "description": "The end position of the range on the reference, 0-based exclusive. If specified, referenceName must also be specified.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "range.referenceName": {
	//       "description": "The reference sequence name, for example chr1, 1, or chrX.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "range.start": {
	//       "description": "The start position of the range on the reference, 0-based inclusive. If specified, referenceName must also be specified.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "readGroupSetId": {
	//       "description": "Required. The ID of the read group set over which coverage is requested.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "targetBucketWidth": {
	//       "description": "The desired width of each reported coverage bucket in base pairs. This will be rounded down to the nearest precomputed bucket width; the value of which is returned as bucketWidth in the response. Defaults to infinity (each bucket spans an entire reference sequence) or the length of the target range, if specified. The smallest precomputed bucketWidth is currently 2048 base pairs; this is subject to change.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "readgroupsets/{readGroupSetId}/coveragebuckets",
	//   "response": {
	//     "$ref": "ListCoverageBucketsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.reads.search":

type ReadsSearchCall struct {
	s                  *Service
	searchreadsrequest *SearchReadsRequest
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
}

// Search: Gets a list of reads for one or more read group sets. Reads
// search operates over a genomic coordinate space of reference sequence
// & position defined over the reference sequences to which the
// requested read group sets are aligned.
//
// If a target positional range
// is specified, search returns all reads whose alignment to the
// reference genome overlap the range. A query which specifies only read
// group set IDs yields all reads in those read group sets, including
// unmapped reads.
//
// All reads returned (including reads on subsequent
// pages) are ordered by genomic coordinate (reference sequence &
// position). Reads with equivalent genomic coordinates are returned in
// a deterministic order.
//
// Implements GlobalAllianceApi.searchReads.

func (r *ReadsService) Search(searchreadsrequest *SearchReadsRequest) *ReadsSearchCall {
	return &ReadsSearchCall{
		s:                  r.s,
		searchreadsrequest: searchreadsrequest,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "reads/search",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReadsSearchCall) Fields(s ...googleapi.Field) *ReadsSearchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReadsSearchCall) Do() (*SearchReadsResponse, error) {
	var returnValue *SearchReadsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.searchreadsrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a list of reads for one or more read group sets. Reads search operates over a genomic coordinate space of reference sequence \u0026 position defined over the reference sequences to which the requested read group sets are aligned.\n\nIf a target positional range is specified, search returns all reads whose alignment to the reference genome overlap the range. A query which specifies only read group set IDs yields all reads in those read group sets, including unmapped reads.\n\nAll reads returned (including reads on subsequent pages) are ordered by genomic coordinate (reference sequence \u0026 position). Reads with equivalent genomic coordinates are returned in a deterministic order.\n\nImplements GlobalAllianceApi.searchReads.",
	//   "httpMethod": "POST",
	//   "id": "genomics.reads.search",
	//   "path": "reads/search",
	//   "request": {
	//     "$ref": "SearchReadsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchReadsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.references.get":

type ReferencesGetCall struct {
	s             *Service
	referenceId   string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a reference.
//
// Implements GlobalAllianceApi.getReference.

func (r *ReferencesService) Get(referenceId string) *ReferencesGetCall {
	return &ReferencesGetCall{
		s:             r.s,
		referenceId:   referenceId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "references/{referenceId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReferencesGetCall) Fields(s ...googleapi.Field) *ReferencesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReferencesGetCall) Do() (*Reference, error) {
	var returnValue *Reference
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"referenceId": c.referenceId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a reference.\n\nImplements GlobalAllianceApi.getReference.",
	//   "httpMethod": "GET",
	//   "id": "genomics.references.get",
	//   "parameterOrder": [
	//     "referenceId"
	//   ],
	//   "parameters": {
	//     "referenceId": {
	//       "description": "The ID of the reference.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "references/{referenceId}",
	//   "response": {
	//     "$ref": "Reference"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.references.search":

type ReferencesSearchCall struct {
	s                       *Service
	searchreferencesrequest *SearchReferencesRequest
	caller_                 googleapi.Caller
	params_                 url.Values
	pathTemplate_           string
}

// Search: Searches for references which match the given
// criteria.
//
// Implements GlobalAllianceApi.searchReferences.

func (r *ReferencesService) Search(searchreferencesrequest *SearchReferencesRequest) *ReferencesSearchCall {
	return &ReferencesSearchCall{
		s: r.s,
		searchreferencesrequest: searchreferencesrequest,
		caller_:                 googleapi.JSONCall{},
		params_:                 make(map[string][]string),
		pathTemplate_:           "references/search",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReferencesSearchCall) Fields(s ...googleapi.Field) *ReferencesSearchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReferencesSearchCall) Do() (*SearchReferencesResponse, error) {
	var returnValue *SearchReferencesResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.searchreferencesrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Searches for references which match the given criteria.\n\nImplements GlobalAllianceApi.searchReferences.",
	//   "httpMethod": "POST",
	//   "id": "genomics.references.search",
	//   "path": "references/search",
	//   "request": {
	//     "$ref": "SearchReferencesRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchReferencesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.references.bases.list":

type ReferencesBasesListCall struct {
	s             *Service
	referenceId   string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists the bases in a reference, optionally restricted to a
// range.
//
// Implements GlobalAllianceApi.getReferenceBases.

func (r *ReferencesBasesService) List(referenceId string) *ReferencesBasesListCall {
	return &ReferencesBasesListCall{
		s:             r.s,
		referenceId:   referenceId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "references/{referenceId}/bases",
	}
}

// End sets the optional parameter "end": The end position (0-based,
// exclusive) of this query. Defaults to the length of this reference.
func (c *ReferencesBasesListCall) End(end int64) *ReferencesBasesListCall {
	c.params_.Set("end", fmt.Sprintf("%v", end))
	return c
}

// PageSize sets the optional parameter "pageSize": Specifies the
// maximum number of bases to return in a single page.
func (c *ReferencesBasesListCall) PageSize(pageSize int64) *ReferencesBasesListCall {
	c.params_.Set("pageSize", fmt.Sprintf("%v", pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, which is used to page through large result sets. To get the
// next page of results, set this parameter to the value of
// nextPageToken from the previous response.
func (c *ReferencesBasesListCall) PageToken(pageToken string) *ReferencesBasesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Start sets the optional parameter "start": The start position
// (0-based) of this query. Defaults to 0.
func (c *ReferencesBasesListCall) Start(start int64) *ReferencesBasesListCall {
	c.params_.Set("start", fmt.Sprintf("%v", start))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReferencesBasesListCall) Fields(s ...googleapi.Field) *ReferencesBasesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReferencesBasesListCall) Do() (*ListBasesResponse, error) {
	var returnValue *ListBasesResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"referenceId": c.referenceId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists the bases in a reference, optionally restricted to a range.\n\nImplements GlobalAllianceApi.getReferenceBases.",
	//   "httpMethod": "GET",
	//   "id": "genomics.references.bases.list",
	//   "parameterOrder": [
	//     "referenceId"
	//   ],
	//   "parameters": {
	//     "end": {
	//       "description": "The end position (0-based, exclusive) of this query. Defaults to the length of this reference.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Specifies the maximum number of bases to return in a single page.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of nextPageToken from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "referenceId": {
	//       "description": "The ID of the reference.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "start": {
	//       "description": "The start position (0-based) of this query. Defaults to 0.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "references/{referenceId}/bases",
	//   "response": {
	//     "$ref": "ListBasesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.referencesets.get":

type ReferencesetsGetCall struct {
	s              *Service
	referenceSetId string
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Get: Gets a reference set.
//
// Implements
// GlobalAllianceApi.getReferenceSet.

func (r *ReferencesetsService) Get(referenceSetId string) *ReferencesetsGetCall {
	return &ReferencesetsGetCall{
		s:              r.s,
		referenceSetId: referenceSetId,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "referencesets/{referenceSetId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReferencesetsGetCall) Fields(s ...googleapi.Field) *ReferencesetsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReferencesetsGetCall) Do() (*ReferenceSet, error) {
	var returnValue *ReferenceSet
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"referenceSetId": c.referenceSetId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a reference set.\n\nImplements GlobalAllianceApi.getReferenceSet.",
	//   "httpMethod": "GET",
	//   "id": "genomics.referencesets.get",
	//   "parameterOrder": [
	//     "referenceSetId"
	//   ],
	//   "parameters": {
	//     "referenceSetId": {
	//       "description": "The ID of the reference set.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "referencesets/{referenceSetId}",
	//   "response": {
	//     "$ref": "ReferenceSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.referencesets.search":

type ReferencesetsSearchCall struct {
	s                          *Service
	searchreferencesetsrequest *SearchReferenceSetsRequest
	caller_                    googleapi.Caller
	params_                    url.Values
	pathTemplate_              string
}

// Search: Searches for reference sets which match the given
// criteria.
//
// Implements GlobalAllianceApi.searchReferenceSets.

func (r *ReferencesetsService) Search(searchreferencesetsrequest *SearchReferenceSetsRequest) *ReferencesetsSearchCall {
	return &ReferencesetsSearchCall{
		s: r.s,
		searchreferencesetsrequest: searchreferencesetsrequest,
		caller_:                    googleapi.JSONCall{},
		params_:                    make(map[string][]string),
		pathTemplate_:              "referencesets/search",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReferencesetsSearchCall) Fields(s ...googleapi.Field) *ReferencesetsSearchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReferencesetsSearchCall) Do() (*SearchReferenceSetsResponse, error) {
	var returnValue *SearchReferenceSetsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.searchreferencesetsrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Searches for reference sets which match the given criteria.\n\nImplements GlobalAllianceApi.searchReferenceSets.",
	//   "httpMethod": "POST",
	//   "id": "genomics.referencesets.search",
	//   "path": "referencesets/search",
	//   "request": {
	//     "$ref": "SearchReferenceSetsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchReferenceSetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.streamingVariantStore.streamvariants":

type StreamingVariantStoreStreamvariantsCall struct {
	s                     *Service
	streamvariantsrequest *StreamVariantsRequest
	caller_               googleapi.Caller
	params_               url.Values
	pathTemplate_         string
}

// Streamvariants: Returns a stream of all the variants matching the
// search request, ordered by reference name, position, and ID.

func (r *StreamingVariantStoreService) Streamvariants(streamvariantsrequest *StreamVariantsRequest) *StreamingVariantStoreStreamvariantsCall {
	return &StreamingVariantStoreStreamvariantsCall{
		s: r.s,
		streamvariantsrequest: streamvariantsrequest,
		caller_:               googleapi.JSONCall{},
		params_:               make(map[string][]string),
		pathTemplate_:         "streamingVariantStore/streamvariants",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StreamingVariantStoreStreamvariantsCall) Fields(s ...googleapi.Field) *StreamingVariantStoreStreamvariantsCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *StreamingVariantStoreStreamvariantsCall) Do() (*Variant, error) {
	var returnValue *Variant
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.streamvariantsrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Returns a stream of all the variants matching the search request, ordered by reference name, position, and ID.",
	//   "httpMethod": "POST",
	//   "id": "genomics.streamingVariantStore.streamvariants",
	//   "path": "streamingVariantStore/streamvariants",
	//   "request": {
	//     "$ref": "StreamVariantsRequest"
	//   },
	//   "response": {
	//     "$ref": "Variant"
	//   }
	// }

}

// method id "genomics.variants.create":

type VariantsCreateCall struct {
	s             *Service
	variant       *Variant
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Create: Creates a new variant.

func (r *VariantsService) Create(variant *Variant) *VariantsCreateCall {
	return &VariantsCreateCall{
		s:             r.s,
		variant:       variant,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "variants",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsCreateCall) Fields(s ...googleapi.Field) *VariantsCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *VariantsCreateCall) Do() (*Variant, error) {
	var returnValue *Variant
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.variant,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a new variant.",
	//   "httpMethod": "POST",
	//   "id": "genomics.variants.create",
	//   "path": "variants",
	//   "request": {
	//     "$ref": "Variant"
	//   },
	//   "response": {
	//     "$ref": "Variant"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variants.delete":

type VariantsDeleteCall struct {
	s             *Service
	variantId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a variant.

func (r *VariantsService) Delete(variantId string) *VariantsDeleteCall {
	return &VariantsDeleteCall{
		s:             r.s,
		variantId:     variantId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "variants/{variantId}",
	}
}

func (c *VariantsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"variantId": c.variantId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a variant.",
	//   "httpMethod": "DELETE",
	//   "id": "genomics.variants.delete",
	//   "parameterOrder": [
	//     "variantId"
	//   ],
	//   "parameters": {
	//     "variantId": {
	//       "description": "The ID of the variant to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "variants/{variantId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variants.get":

type VariantsGetCall struct {
	s             *Service
	variantId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a variant by ID.

func (r *VariantsService) Get(variantId string) *VariantsGetCall {
	return &VariantsGetCall{
		s:             r.s,
		variantId:     variantId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "variants/{variantId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsGetCall) Fields(s ...googleapi.Field) *VariantsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *VariantsGetCall) Do() (*Variant, error) {
	var returnValue *Variant
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"variantId": c.variantId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a variant by ID.",
	//   "httpMethod": "GET",
	//   "id": "genomics.variants.get",
	//   "parameterOrder": [
	//     "variantId"
	//   ],
	//   "parameters": {
	//     "variantId": {
	//       "description": "The ID of the variant.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "variants/{variantId}",
	//   "response": {
	//     "$ref": "Variant"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.variants.search":

type VariantsSearchCall struct {
	s                     *Service
	searchvariantsrequest *SearchVariantsRequest
	caller_               googleapi.Caller
	params_               url.Values
	pathTemplate_         string
}

// Search: Gets a list of variants matching the criteria.
//
// Implements
// GlobalAllianceApi.searchVariants.

func (r *VariantsService) Search(searchvariantsrequest *SearchVariantsRequest) *VariantsSearchCall {
	return &VariantsSearchCall{
		s: r.s,
		searchvariantsrequest: searchvariantsrequest,
		caller_:               googleapi.JSONCall{},
		params_:               make(map[string][]string),
		pathTemplate_:         "variants/search",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsSearchCall) Fields(s ...googleapi.Field) *VariantsSearchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *VariantsSearchCall) Do() (*SearchVariantsResponse, error) {
	var returnValue *SearchVariantsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.searchvariantsrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a list of variants matching the criteria.\n\nImplements GlobalAllianceApi.searchVariants.",
	//   "httpMethod": "POST",
	//   "id": "genomics.variants.search",
	//   "path": "variants/search",
	//   "request": {
	//     "$ref": "SearchVariantsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchVariantsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.variants.update":

type VariantsUpdateCall struct {
	s             *Service
	variantId     string
	variant       *Variant
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates a variant's names and info fields. All other
// modifications are silently ignored. Returns the modified variant
// without its calls.

func (r *VariantsService) Update(variantId string, variant *Variant) *VariantsUpdateCall {
	return &VariantsUpdateCall{
		s:             r.s,
		variantId:     variantId,
		variant:       variant,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "variants/{variantId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsUpdateCall) Fields(s ...googleapi.Field) *VariantsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *VariantsUpdateCall) Do() (*Variant, error) {
	var returnValue *Variant
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"variantId": c.variantId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.variant,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a variant's names and info fields. All other modifications are silently ignored. Returns the modified variant without its calls.",
	//   "httpMethod": "PUT",
	//   "id": "genomics.variants.update",
	//   "parameterOrder": [
	//     "variantId"
	//   ],
	//   "parameters": {
	//     "variantId": {
	//       "description": "The ID of the variant to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "variants/{variantId}",
	//   "request": {
	//     "$ref": "Variant"
	//   },
	//   "response": {
	//     "$ref": "Variant"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variantsets.delete":

type VariantsetsDeleteCall struct {
	s             *Service
	variantSetId  string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes the contents of a variant set. The variant set object
// is not deleted.

func (r *VariantsetsService) Delete(variantSetId string) *VariantsetsDeleteCall {
	return &VariantsetsDeleteCall{
		s:             r.s,
		variantSetId:  variantSetId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "variantsets/{variantSetId}",
	}
}

func (c *VariantsetsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"variantSetId": c.variantSetId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes the contents of a variant set. The variant set object is not deleted.",
	//   "httpMethod": "DELETE",
	//   "id": "genomics.variantsets.delete",
	//   "parameterOrder": [
	//     "variantSetId"
	//   ],
	//   "parameters": {
	//     "variantSetId": {
	//       "description": "The ID of the variant set to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "variantsets/{variantSetId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variantsets.export":

type VariantsetsExportCall struct {
	s                       *Service
	variantSetId            string
	exportvariantsetrequest *ExportVariantSetRequest
	caller_                 googleapi.Caller
	params_                 url.Values
	pathTemplate_           string
}

// Export: Exports variant set data to an external destination.

func (r *VariantsetsService) Export(variantSetId string, exportvariantsetrequest *ExportVariantSetRequest) *VariantsetsExportCall {
	return &VariantsetsExportCall{
		s:                       r.s,
		variantSetId:            variantSetId,
		exportvariantsetrequest: exportvariantsetrequest,
		caller_:                 googleapi.JSONCall{},
		params_:                 make(map[string][]string),
		pathTemplate_:           "variantsets/{variantSetId}/export",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsetsExportCall) Fields(s ...googleapi.Field) *VariantsetsExportCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *VariantsetsExportCall) Do() (*ExportVariantSetResponse, error) {
	var returnValue *ExportVariantSetResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"variantSetId": c.variantSetId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.exportvariantsetrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Exports variant set data to an external destination.",
	//   "httpMethod": "POST",
	//   "id": "genomics.variantsets.export",
	//   "parameterOrder": [
	//     "variantSetId"
	//   ],
	//   "parameters": {
	//     "variantSetId": {
	//       "description": "Required. The ID of the variant set that contains variant data which should be exported. The caller must have READ access to this variant set.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "variantsets/{variantSetId}/export",
	//   "request": {
	//     "$ref": "ExportVariantSetRequest"
	//   },
	//   "response": {
	//     "$ref": "ExportVariantSetResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variantsets.get":

type VariantsetsGetCall struct {
	s             *Service
	variantSetId  string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a variant set by ID.

func (r *VariantsetsService) Get(variantSetId string) *VariantsetsGetCall {
	return &VariantsetsGetCall{
		s:             r.s,
		variantSetId:  variantSetId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "variantsets/{variantSetId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsetsGetCall) Fields(s ...googleapi.Field) *VariantsetsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *VariantsetsGetCall) Do() (*VariantSet, error) {
	var returnValue *VariantSet
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"variantSetId": c.variantSetId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a variant set by ID.",
	//   "httpMethod": "GET",
	//   "id": "genomics.variantsets.get",
	//   "parameterOrder": [
	//     "variantSetId"
	//   ],
	//   "parameters": {
	//     "variantSetId": {
	//       "description": "Required. The ID of the variant set.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "variantsets/{variantSetId}",
	//   "response": {
	//     "$ref": "VariantSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.variantsets.importVariants":

type VariantsetsImportVariantsCall struct {
	s                     *Service
	variantSetId          string
	importvariantsrequest *ImportVariantsRequest
	caller_               googleapi.Caller
	params_               url.Values
	pathTemplate_         string
}

// ImportVariants: Creates variant data by asynchronously importing the
// provided information.
//
// The variants for import will be merged with
// any existing data and each other according to the behavior of
// mergeVariants. In particular, this means for merged VCF variants that
// have conflicting INFO fields, some data will be arbitrarily
// discarded. As a special case, for single-sample VCF files, QUAL and
// FILTER fields will be moved to the call level; these are sometimes
// interpreted in a call-specific context. Imported VCF headers are
// appended to the metadata already in a variant set.

func (r *VariantsetsService) ImportVariants(variantSetId string, importvariantsrequest *ImportVariantsRequest) *VariantsetsImportVariantsCall {
	return &VariantsetsImportVariantsCall{
		s:                     r.s,
		variantSetId:          variantSetId,
		importvariantsrequest: importvariantsrequest,
		caller_:               googleapi.JSONCall{},
		params_:               make(map[string][]string),
		pathTemplate_:         "variantsets/{variantSetId}/importVariants",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsetsImportVariantsCall) Fields(s ...googleapi.Field) *VariantsetsImportVariantsCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *VariantsetsImportVariantsCall) Do() (*ImportVariantsResponse, error) {
	var returnValue *ImportVariantsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"variantSetId": c.variantSetId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.importvariantsrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates variant data by asynchronously importing the provided information.\n\nThe variants for import will be merged with any existing data and each other according to the behavior of mergeVariants. In particular, this means for merged VCF variants that have conflicting INFO fields, some data will be arbitrarily discarded. As a special case, for single-sample VCF files, QUAL and FILTER fields will be moved to the call level; these are sometimes interpreted in a call-specific context. Imported VCF headers are appended to the metadata already in a variant set.",
	//   "httpMethod": "POST",
	//   "id": "genomics.variantsets.importVariants",
	//   "parameterOrder": [
	//     "variantSetId"
	//   ],
	//   "parameters": {
	//     "variantSetId": {
	//       "description": "Required. The variant set to which variant data should be imported.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "variantsets/{variantSetId}/importVariants",
	//   "request": {
	//     "$ref": "ImportVariantsRequest"
	//   },
	//   "response": {
	//     "$ref": "ImportVariantsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.read_write",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variantsets.mergeVariants":

type VariantsetsMergeVariantsCall struct {
	s                    *Service
	variantSetId         string
	mergevariantsrequest *MergeVariantsRequest
	caller_              googleapi.Caller
	params_              url.Values
	pathTemplate_        string
}

// MergeVariants: Merges the given variants with existing variants. Each
// variant will be merged with an existing variant that matches its
// reference sequence, start, end, reference bases, and alternative
// bases. If no such variant exists, a new one will be created.
//
// When
// variants are merged, the call information from the new variant is
// added to the existing variant, and other fields (such as key/value
// pairs) are discarded.

func (r *VariantsetsService) MergeVariants(variantSetId string, mergevariantsrequest *MergeVariantsRequest) *VariantsetsMergeVariantsCall {
	return &VariantsetsMergeVariantsCall{
		s:                    r.s,
		variantSetId:         variantSetId,
		mergevariantsrequest: mergevariantsrequest,
		caller_:              googleapi.JSONCall{},
		params_:              make(map[string][]string),
		pathTemplate_:        "variantsets/{variantSetId}/mergeVariants",
	}
}

func (c *VariantsetsMergeVariantsCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"variantSetId": c.variantSetId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.mergevariantsrequest,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Merges the given variants with existing variants. Each variant will be merged with an existing variant that matches its reference sequence, start, end, reference bases, and alternative bases. If no such variant exists, a new one will be created.\n\nWhen variants are merged, the call information from the new variant is added to the existing variant, and other fields (such as key/value pairs) are discarded.",
	//   "httpMethod": "POST",
	//   "id": "genomics.variantsets.mergeVariants",
	//   "parameterOrder": [
	//     "variantSetId"
	//   ],
	//   "parameters": {
	//     "variantSetId": {
	//       "description": "The destination variant set.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "variantsets/{variantSetId}/mergeVariants",
	//   "request": {
	//     "$ref": "MergeVariantsRequest"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variantsets.patch":

type VariantsetsPatchCall struct {
	s             *Service
	variantSetId  string
	variantset    *VariantSet
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates a variant set's metadata. All other modifications are
// silently ignored. This method supports patch semantics.

func (r *VariantsetsService) Patch(variantSetId string, variantset *VariantSet) *VariantsetsPatchCall {
	return &VariantsetsPatchCall{
		s:             r.s,
		variantSetId:  variantSetId,
		variantset:    variantset,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "variantsets/{variantSetId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsetsPatchCall) Fields(s ...googleapi.Field) *VariantsetsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *VariantsetsPatchCall) Do() (*VariantSet, error) {
	var returnValue *VariantSet
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"variantSetId": c.variantSetId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.variantset,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a variant set's metadata. All other modifications are silently ignored. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "genomics.variantsets.patch",
	//   "parameterOrder": [
	//     "variantSetId"
	//   ],
	//   "parameters": {
	//     "variantSetId": {
	//       "description": "The ID of the variant to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "variantsets/{variantSetId}",
	//   "request": {
	//     "$ref": "VariantSet"
	//   },
	//   "response": {
	//     "$ref": "VariantSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variantsets.search":

type VariantsetsSearchCall struct {
	s                        *Service
	searchvariantsetsrequest *SearchVariantSetsRequest
	caller_                  googleapi.Caller
	params_                  url.Values
	pathTemplate_            string
}

// Search: Returns a list of all variant sets matching search
// criteria.
//
// Implements GlobalAllianceApi.searchVariantSets.

func (r *VariantsetsService) Search(searchvariantsetsrequest *SearchVariantSetsRequest) *VariantsetsSearchCall {
	return &VariantsetsSearchCall{
		s: r.s,
		searchvariantsetsrequest: searchvariantsetsrequest,
		caller_:                  googleapi.JSONCall{},
		params_:                  make(map[string][]string),
		pathTemplate_:            "variantsets/search",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsetsSearchCall) Fields(s ...googleapi.Field) *VariantsetsSearchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *VariantsetsSearchCall) Do() (*SearchVariantSetsResponse, error) {
	var returnValue *SearchVariantSetsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.searchvariantsetsrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Returns a list of all variant sets matching search criteria.\n\nImplements GlobalAllianceApi.searchVariantSets.",
	//   "httpMethod": "POST",
	//   "id": "genomics.variantsets.search",
	//   "path": "variantsets/search",
	//   "request": {
	//     "$ref": "SearchVariantSetsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchVariantSetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.variantsets.update":

type VariantsetsUpdateCall struct {
	s             *Service
	variantSetId  string
	variantset    *VariantSet
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates a variant set's metadata. All other modifications are
// silently ignored.

func (r *VariantsetsService) Update(variantSetId string, variantset *VariantSet) *VariantsetsUpdateCall {
	return &VariantsetsUpdateCall{
		s:             r.s,
		variantSetId:  variantSetId,
		variantset:    variantset,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "variantsets/{variantSetId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsetsUpdateCall) Fields(s ...googleapi.Field) *VariantsetsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *VariantsetsUpdateCall) Do() (*VariantSet, error) {
	var returnValue *VariantSet
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"variantSetId": c.variantSetId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.variantset,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a variant set's metadata. All other modifications are silently ignored.",
	//   "httpMethod": "PUT",
	//   "id": "genomics.variantsets.update",
	//   "parameterOrder": [
	//     "variantSetId"
	//   ],
	//   "parameters": {
	//     "variantSetId": {
	//       "description": "The ID of the variant to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "variantsets/{variantSetId}",
	//   "request": {
	//     "$ref": "VariantSet"
	//   },
	//   "response": {
	//     "$ref": "VariantSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}
