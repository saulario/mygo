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

// S0001Type SYNTAX IDENTIFIER
type S0001Type struct {
	De0001 De0001Type
	De0002 De0002Type
	De0080 De0080Type
	De0133 De0133Type
}

// S0002Type INTERCHANGE SENDER
type S0002Type struct {
	De0004 De0004Type
	De0007 De0007Type
	De0008 De0008Type
	De0042 De0042Type
}

// S0003Type INTERCHANGE RECIPIENT
type S0003Type struct {
	De0010 De0010Type
	De0007 De0007Type
	De0014 De0014Type
	De0046 De0046Type
}

// S0004Type DATE AND TIME OF PREPARATION
type S0004Type struct {
	De0017 De0017Type
	De0019 De0019Type
}

// S0005Type RECIPIENT REFERENCE/PASSWORD DETAILS
type S0005Type struct {
	De0022 De0022Type
	De0025 De0025Type
}

// Serialize ...
func (s S0001Type) Serialize(una UnaType) string {
	ds := string(una.Una1)
	retval := string(s.De0001) + ds +
		string(s.De0002) + ds +
		string(s.De0080) + ds +
		string(s.De0133)
	return cleanTrailingSeparator(retval, ds)
}

// Serialize ...
func (s S0002Type) Serialize(una UnaType) string {
	ds := string(una.Una1)
	retval := string(s.De0004) + ds +
		string(s.De0007) + ds +
		string(s.De0008) + ds +
		string(s.De0042)
	return cleanTrailingSeparator(retval, ds)
}

// Serialize ...
func (s S0003Type) Serialize(una UnaType) string {
	ds := string(una.Una1)
	retval := string(s.De0010) + ds +
		string(s.De0007) + ds +
		string(s.De0014) + ds +
		string(s.De0046)
	return cleanTrailingSeparator(retval, ds)
}

// Serialize ...
func (s S0004Type) Serialize(una UnaType) string {
	ds := string(una.Una2)
	retval := string(s.De0017) + ds + string(s.De0019)
	return cleanTrailingSeparator(retval, ds)
}

// Serialize ...
func (s S0005Type) Serialize(una UnaType) string {
	ds := string(una.Una2)
	retval := string(s.De0022) + ds + string(s.De0025)
	return cleanTrailingSeparator(retval, ds)
}
