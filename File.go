// File.go.

//============================================================================//
//
// Copyright © 2018..2020 by McArcher.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
//============================================================================//
//
// Web Site:		'https://github.com/neverwinter-nights'.
// Author:			McArcher.
// Web Site Address is an Address in the global Computer Internet Network.
//
//============================================================================//

// The 'File' Class.

package bencode

import (
	"bufio"
	"io/ioutil"
	"os"
	"time"
)

// A File.
type File struct {
	path   string
	osFile *os.File
}

// File's Constructor.
func NewFile(
	filePath string,
) (result *File) {
	result = new(File)

	result.path = filePath

	return
}

// Closes a File.
func (f *File) close() (err error) {
	err = f.osFile.Close()
	return
}

// Reads the Contents of an opened File.
func (f File) getContents() (fileContents []byte, err error) {

	// Fool Check.
	if f.osFile == nil {
		err = ErrFileNotInitialized
		return
	}

	_, err = f.osFile.Seek(0, 0)
	if err != nil {
		return
	}

	fileContents, err = ioutil.ReadFile(f.path)
	if err != nil {
		return
	}

	return
}

// Opens a File.
func (f *File) open() (err error) {
	f.osFile, err = os.Open(f.path)
	return
}

// Parses an input File into an Interface.
// Also stores some additional Data, all packed into an Object.
func (f File) Parse() (result *DecodedObject, err error) {

	// Open the File and prepare a Stream Reader.
	err = f.open()
	if err != nil {
		return
	}
	defer func() {
		// Close the File.
		var derr error
		derr = f.close()
		if derr != nil {
			err = combineErrors(err, derr)
		}
	}()
	var bufioReader *bufio.Reader = bufio.NewReader(f.osFile)

	// Parse the File encoded with 'bencode' Encoding into an Object.
	var decoder *Decoder = NewDecoder(bufioReader)
	var ifc interface{}
	ifc, err = decoder.readBencodedValue()
	if err != nil {
		return
	}

	// Get the File Contents.
	var fileContents []byte
	fileContents, err = f.getContents()
	if err != nil {
		return
	}

	// Prepare Result.
	var object *DecodedObject
	object = &DecodedObject{
		FilePath:        f.path,
		SourceData:      fileContents,
		DecodedObject:   ifc,
		DecodeTimestamp: time.Now().Unix(),
	}

	// Perform a Self-Check.
	var ok bool
	ok = object.MakeSelfCheck()
	if !ok {
		err = ErrSelfCheck
		return
	}

	// Calculate BTIH.
	err = object.CalculateBtih()
	if err != nil {
		return
	}

	result = object
	return
}
