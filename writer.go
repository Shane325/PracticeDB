package main

import (
    "encoding/binary"
    "encoding/json"
    "fmt"
    "io"
)

type FileWriter struct {
    // Args
    numRows int
    columnNames []string
    w io.Writer

    // State
    numWritten int
    uvarintBuffer []byte
}

type Header struct {
    numRows int
    columnNames []string
}

func newFileWriter(numRows int, columnNames []string, w io.Writer) *FileWriter {
    return &FileWriter{
        numRows: numRows,
        columnNames: columnNames,
        w: w,
        uvarintBuffer: make([]byte, binary.MaxVarintLen64),
    }
}

func (w *FileWriter) Append(t Tuple) error {
    if w.numWritten == 0 {
        err := w.WriteHeader()
        if err != nil {
            return fmt.Errorf(
                "Writer: Append: error writing the header: %v", err,
            )
        }
    }

    if len(t.values) != len(w.columnNames) {
        return fmt.Errorf(
            "Writer: Append: tried to write tuple %v with %d values, but writer expects %d columns",
            t, len(t.values), len(w.columnNames),
        )
    }

    for _, v := range t.values {
        if err := w.WriteUVarInt(uint64(len(v.value))); err != nil {
            return fmt.Errorf(
                "Writer: Append: error writing uvarint: %v", err,
            )
        }
        if _, err := w.w.Write([]byte(v.value)); err != nil {
            return fmt.Errorf(
                "Writer: Append: error writing string %s: error %v",
                v.value, err,
            )
        }
    }

    w.numWritten++

    return nil
}

func (w *FileWriter) WriteHeader() error {
    header := Header{
        numRows: w.numRows,
        columnNames: w.columnNames,
    }
    headerBytes, err := json.Marshal(&header)
    if err != nil {
        return fmt.Errorf(
            "Writer: WriteHeader: error marshaling header: %v err: %v",
            header, err,
        )
    }

    if err := w.WriteUVarInt(uint64(len(headerBytes))); err != nil {
        return fmt.Errorf(
            "Writer: WriteHeader: error writing header bytes uvarint length: %v", err,
        )
    }

    if _, err := w.w.Write(headerBytes); err != nil {
        return fmt.Errorf(
            "Writer: WriteHeader: error writing header bytes: %v", err,
        )
    }

    return nil
}

func (w *FileWriter) WriteUVarInt(x uint64) error {
    varintLen := binary.PutUvarint(w.uvarintBuffer, x)
    _, err := w.w.Write(w.uvarintBuffer[:varintLen])
    if err != nil {
        return fmt.Errorf(
            "Writer: WriteUVarInt: error writing uvarint: %v", err,
        )
    }
    return nil
}
