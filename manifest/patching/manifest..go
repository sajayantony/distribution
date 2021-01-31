package patching

const (
	// HeaderPatchManifest header checked for requesting patched manifest
	HeaderPatchManifest = "X-Meta-Content-Patches"

	// MediaTypePatchManifest specifies the mediaType for the current version.
	MediaTypePatchManifest = "application/vnd.oci.artifact.patch.v1+json"
)
