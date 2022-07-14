package datasource

import (
	"strconv"
	"time"

	"crypto/sha256"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-provider-hpcr/common"
	"github.com/terraform-provider-hpcr/data"
	"github.com/terraform-provider-hpcr/fp"
	B "github.com/terraform-provider-hpcr/fp/bytes"
	E "github.com/terraform-provider-hpcr/fp/either"
	"github.com/terraform-provider-hpcr/validation"

	F "github.com/terraform-provider-hpcr/fp/function"
)

func setUniqueID(d *schema.ResourceData) *schema.ResourceData {
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return d
}

var (
	seqResourceData = E.SequenceArray[error, *schema.ResourceData]()
	setRendered     = fp.ResourceDataSet[string](common.KeyRendered)
	setText         = fp.ResourceDataSet[string](common.KeyText)
	setSha256       = fp.ResourceDataSet[string](common.KeySha256)
	getJson         = fp.ResourceDataGet[any](common.KeyJson)
	getText         = fp.ResourceDataGet[string](common.KeyText)
	getPubKey       = fp.ResourceDataGet[string](common.KeyCert)

	// encode as sha256
	computeSha256 = F.Flow3(
		E.Map[error](sha256.Sum256),
		E.Map[error](func(hash [sha256.Size]byte) string { return fmt.Sprintf("%x", hash) }),
		E.Map[error](setSha256),
	)

	// encode as text
	computeText = F.Flow2(
		E.Map[error](B.ToString),
		E.Map[error](setText),
	)

	schemaCertIn = schema.Schema{
		Type:             schema.TypeString,
		Description:      "Certificate used to encrypt the JSON document in PEM format",
		Optional:         true,
		Default:          data.DefaultCertificate,
		ValidateDiagFunc: validation.DiagCertificate,
	}

	schemaTextOut = schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}

	schemaRenderedOut = schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}

	schemaSha256Out = schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}

	schemaJsonIn = schema.Schema{
		Type:        schema.TypeMap,
		Required:    true,
		Description: "JSON Document to archive",
	}

	schemaTextIn = schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "Text to archive",
	}
)