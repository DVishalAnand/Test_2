/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

//go:generate bash -c "protoc -I /usr/include --proto_path=$MAGMA_ROOT --go_out=plugins=grpc:$MAGMA_ROOT/.. $MAGMA_ROOT/cwf/protos/mconfig/*.proto"
package protos
## This is change we are doing for test-1