/*
 * Copyright (c) "Neo4j"
 * Neo4j Sweden AB [https://neo4j.com]
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package neo4j

import "github.com/SeBBBe/neo4j-go-driver/v5/neo4j/internal/auth"

// AuthToken contains credentials to be sent over to the neo4j server.
type AuthToken = auth.Token

const keyScheme = "scheme"
const schemeNone = "none"
const schemeBasic = "basic"
const schemeKerberos = "kerberos"
const schemeBearer = "bearer"
const keyPrincipal = "principal"
const keyCredentials = "credentials"
const keyRealm = "realm"

// NoAuth generates an empty authentication token
func NoAuth() AuthToken {
	return AuthToken{Tokens: map[string]any{
		keyScheme: schemeNone,
	}}
}

// BasicAuth generates a basic authentication token with provided username, password and realm
func BasicAuth(username string, password string, realm string) AuthToken {
	tokens := map[string]any{
		keyScheme:      schemeBasic,
		keyPrincipal:   username,
		keyCredentials: password,
	}

	if realm != "" {
		tokens[keyRealm] = realm
	}

	return AuthToken{Tokens: tokens}
}

// KerberosAuth generates a kerberos authentication token with provided base-64 encoded kerberos ticket
func KerberosAuth(ticket string) AuthToken {
	token := AuthToken{
		Tokens: map[string]any{
			keyScheme: schemeKerberos,
			// Backwards compatibility: Neo4j servers pre 4.4 require the presence of the principal.
			keyPrincipal:   "",
			keyCredentials: ticket,
		},
	}

	return token
}

// BearerAuth generates an authentication token with the provided base-64 value generated by a Single Sign-On provider
func BearerAuth(token string) AuthToken {
	result := AuthToken{
		Tokens: map[string]any{
			keyScheme:      schemeBearer,
			keyCredentials: token,
		},
	}

	return result
}

// CustomAuth generates a custom authentication token with provided parameters
func CustomAuth(scheme string, username string, password string, realm string, parameters map[string]any) AuthToken {
	tokens := map[string]any{
		keyScheme:    scheme,
		keyPrincipal: username,
	}

	if password != "" {
		tokens[keyCredentials] = password
	}

	if realm != "" {
		tokens[keyRealm] = realm
	}

	if len(parameters) > 0 {
		tokens["parameters"] = parameters
	}

	return AuthToken{Tokens: tokens}
}
