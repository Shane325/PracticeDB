package main

import (
    "encoding/binary"
    "encoding/json"
    "bytes"
    "fmt"
    "io"
)

type byteReader struct {
    io.Reader
    byteBuf []byte
}

func newByteReader(r io.Reader) *byteReader {
    return &byteReader{
        Reader: r,
        byteBuf: make([]byte, 1),
    }
}

func (b *byteReader) ReadByte() (byte, error) {
    n, err := b.Reader.Read(b.byteBuf)
    if err != nil {
        return 0, fmt.Errorf("ReadByte: error reading byte: %v", err)
    }

    if n != 1 {
        return 0, fmt.Errorf("ReadByte: expected to read one byte")
    }
    return b.byteBuf[0], nil
}

func main() {
    // // Expressions
    // notEquals := newNotEquals("genres", "Comedy")

    // // Iterator nodes
    // scanner := newScanner("ml-20m/movies.csv")
    // selection := newSelection(notEquals, scanner)
    // projection := newProjection("title", selection)
    // limit := newLimit(15, projection)


    // for limit.next() {
    //     fmt.Println(limit.execute())
    // }
    newValueA := Value{name: "firstName", value: "Shane"}
    newValueB := Value{name: "firstName", value: "John"}
    newTupleA := Tuple{values: []Value{newValueA}}
    newTupleB := Tuple{values: []Value{newValueB}}
    tuples := []Tuple{newTupleA, newTupleB}
    columns := []string{"firstName"}

    buf := bytes.NewBuffer(nil)
    writer := newFileWriter(len(tuples), columns, buf)
    err := writer.Append(newTupleA)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(buf.Len())
    fmt.Println(buf)
    err = writer.Append(newTupleB)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(buf.Len())
    fmt.Println(buf)
    // Read logic
    r := newByteReader(buf)
    headerLength, err := binary.ReadUvarint(r)
    headerBytes := make([]byte, headerLength)
    _, err = io.ReadFull(r, headerBytes)
    fmt.Println(headerBytes)
    if err != nil {
        fmt.Println(err)
    }
    header := &Header{}
    err = json.Unmarshal(headerBytes, header)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(header.numRows)
    fmt.Println(header.columnNames)
}
