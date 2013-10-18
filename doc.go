// Copyright (c) 2013 Black Square Media Ltd. All rights reserved.
// (The MIT License)

// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// 'Software'), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:

// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
// CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
// SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

/*
This package implements a parser (with optional validation) and a generator for
OpenRTB 2.1 requests and responses.

Spec

http://www.iab.net/media/file/OpenRTB-API-Specification-Version-2-1-FINAL.pdf

Usage

  // Handle a HTTP request
  http.HandleFunc("/bid", func(w http.ResponseWriter, r *http.Request) {
    defer r.Body().Close()

    req, err := openrtb.ParseRequest(r.Body(), true)
    if err != nil {
      log.Println("ERROR %s", err.Error())
    } else {
      log.Println("INFO  Received bid request %s", *req.Id)
    }

    w.WriteHeader(204) // respond with 'no bid'
  })

Implemented

- Request parsing
- Most objects
- PMP additions

Todo

- Response parsing
- Generating request and responses
- Content objects
*/
package openrtb
