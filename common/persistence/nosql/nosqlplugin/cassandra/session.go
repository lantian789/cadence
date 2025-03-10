// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cassandra

import (
	"time"

	"github.com/uber/cadence/common/config"
	"github.com/uber/cadence/common/persistence/nosql/nosqlplugin/cassandra/gocql"
)

const (
	cassandraProtoVersion = 4
	defaultSessionTimeout = 10 * time.Second
)

// CreateSession creates a new session
// TODO this will be converted to private later, after all cassandra code moved to plugin pkg
func CreateSession(cfg config.Cassandra) (gocql.Session, error) {
	return cfg.CQLClient.CreateSession(gocql.ClusterConfig{
		Hosts:             cfg.Hosts,
		Port:              cfg.Port,
		User:              cfg.User,
		Password:          cfg.Password,
		Keyspace:          cfg.Keyspace,
		Region:            cfg.Region,
		Datacenter:        cfg.Datacenter,
		MaxConns:          cfg.MaxConns,
		TLS:               cfg.TLS,
		ProtoVersion:      cassandraProtoVersion,
		Consistency:       gocql.LocalQuorum,
		SerialConsistency: gocql.LocalSerial,
		Timeout:           defaultSessionTimeout,
	})
}
