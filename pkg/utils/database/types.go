/*
 * Copyright 2021 kloeckner.i GmbH
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package database

// Credentials contains credentials to connect database
type Credentials struct {
	Name             string
	Username         string
	Password         string
	ConnectionString string
}

// DatabaseAddress contains host and port of a database instance
type DatabaseAddress struct {
	Host string
	Port uint16
}

// AdminCredentials contains admin username and password of database server
type AdminCredentials struct {
	Username string `yaml:"user"`
	Password string `yaml:"password"`
}

// Database is interface for CRUD operate of different types of databases
type Database interface {
	createDatabase(admin AdminCredentials) error
	createUser(admin AdminCredentials) error
	deleteDatabase(admin AdminCredentials) error
	deleteUser(admin AdminCredentials) error
	CheckStatus() error
	GetCredentials() Credentials
	ParseAdminCredentials(data map[string][]byte) (AdminCredentials, error)
	GetDatabaseAddress() DatabaseAddress
}
