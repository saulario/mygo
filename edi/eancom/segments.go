// Package eancom ...
/**
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
package eancom

// BgmType Beginnig of message
type BgmType struct {
}

// ComType Communication contact
type ComType struct {
}

// CtaType Contact information
type CtaType struct {
}

// DtmType Date/time/period
type DtmType struct {
}

// NadType Name and address
type NadType struct {
}

// RffType Reference
type RffType struct {
}

// UnaType Service string advice
type UnaType struct {
	Una1 UNA1Type
	Una2 UNA2Type
	Una3 UNA3Type
	Una4 UNA4Type
	Una5 UNA5Type
	Una6 UNA6Type
}

// UnbType Interchange header
type UnbType struct {
	S0001  S0001Type
	S0002  S0002Type
	S0003  S0003Type
	S0004  S0004Type
	De0020 De0020Type
	S0005  S0005Type
	De0026 De0026Type
	De0029 De0029Type
	De0031 De0031Type
	De0032 De0032Type
	De0035 De0035Type
}

// UnhType Message header
type UnhType struct {
}

// DefaultUna defines default values for the segment
func DefaultUna() UnaType {
	return UnaType{
		Una1: ":",
		Una2: "+",
		Una3: ".",
		Una4: "?",
		Una5: "*",
		Una6: "'",
	}
}
