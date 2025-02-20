// Copyright 2022 IBM Corp.
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
// limitations under the License.package datasource

package contract

import (
	"github.com/ibm-hyper-protect/terraform-provider-hpcr/encrypt"
	E "github.com/ibm-hyper-protect/terraform-provider-hpcr/fp/either"
)

// CryptoEncryptAndSignContract returns the crypto implementation of the encryption and signer
func CryptoEncryptAndSignContract(cert []byte) func([]byte) func(RawMap) E.Either[error, RawMap] {
	return EncryptAndSignContract(encrypt.CryptoEncryptBasic(cert), encrypt.CryptoSignDigest, encrypt.CryptoPublicKey)
}
