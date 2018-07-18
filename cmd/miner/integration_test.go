/*
 * Copyright 2018 The ThunderDB Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"database/sql"
	"path/filepath"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/thunderdb/ThunderDB/client"
	"gitlab.com/thunderdb/ThunderDB/utils"
	"gitlab.com/thunderdb/ThunderDB/utils/log"
)

var (
	baseDir        = utils.GetProjectSrcDir()
	testWorkingDir = FJ(baseDir, "./test/")
	logDir         = FJ(testWorkingDir, "./log/")
)

var FJ = filepath.Join

func _TestBuild(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	utils.Build()
}

func startNodes() {
	// start 3bps
	go utils.RunCommand(
		FJ(baseDir, "./bin/thunderdbd"),
		[]string{"-config", FJ(testWorkingDir, "./node_0/config.yaml")},
		"leader", testWorkingDir, logDir, false,
	)
	go utils.RunCommand(
		FJ(baseDir, "./bin/thunderdbd"),
		[]string{"-config", FJ(testWorkingDir, "./node_1/config.yaml")},
		"follower1", testWorkingDir, logDir, false,
	)
	go utils.RunCommand(
		FJ(baseDir, "./bin/thunderdbd"),
		[]string{"-config", FJ(testWorkingDir, "./node_2/config.yaml")},
		"follower2", testWorkingDir, logDir, false,
	)

	time.Sleep(time.Second * 3)

	// start 3miners
	go utils.RunCommand(
		FJ(baseDir, "./bin/thunderminerd"),
		[]string{"-config", FJ(testWorkingDir, "./node_miner_0/config.yaml")},
		"miner1", testWorkingDir, logDir, false,
	)
	go utils.RunCommand(
		FJ(baseDir, "./bin/thunderminerd"),
		[]string{"-config", FJ(testWorkingDir, "./node_miner_1/config.yaml")},
		"miner2", testWorkingDir, logDir, false,
	)
	go utils.RunCommand(
		FJ(baseDir, "./bin/thunderminerd"),
		[]string{"-config", FJ(testWorkingDir, "./node_miner_2/config.yaml")},
		"miner3", testWorkingDir, logDir, false,
	)

}

func _TestFullProcess(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	Convey("test full process", t, func() {
		startNodes()
		time.Sleep(5 * time.Second)

		var err error
		err = client.Init(FJ(testWorkingDir, "./node_c/config.yaml"), []byte(""))
		So(err, ShouldBeNil)

		// create
		dsn, err := client.Create(client.ResourceMeta{Node: 1})
		So(err, ShouldBeNil)

		log.Infof("the created database dsn is %v", dsn)

		db, err := sql.Open("thunderdb", dsn)
		So(err, ShouldBeNil)

		_, err = db.Exec("CREATE TABLE test (test int)")
		So(err, ShouldBeNil)

		_, err = db.Exec("INSERT INTO test VALUES(?)", 4)
		So(err, ShouldBeNil)

		row := db.QueryRow("SELECT * FROM test LIMIT 1")

		var result int
		err = row.Scan(&result)
		So(err, ShouldBeNil)
		So(result, ShouldEqual, 4)

		err = db.Close()
		So(err, ShouldBeNil)

		err = client.Drop(dsn)
		So(err, ShouldBeNil)
	})
}
