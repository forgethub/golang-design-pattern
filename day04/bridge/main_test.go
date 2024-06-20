package main

import (
    "testing"

    /* "github.com/stretchr/testify/assert" */
)

func TestErrorNotification_Notify(t *testing.T) {
    mac := &Mac{printer: &HP{}}
    mac.Print()
    linux := &Linux{printer: &Epson{}}
    linux.Print()
    /* assert.Nil(t, err) */
}
