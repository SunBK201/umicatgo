package log

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type uctFormatter struct {
}

func (formatter *uctFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	formatTime := entry.Time.Format("2006-01-02 15:04:05")
	b.WriteString(fmt.Sprintf("[%s][%s]: %s\n", formatTime, strings.ToUpper(entry.Level.String()), entry.Message))
	return b.Bytes(), nil
}

func SetLogConf() {
	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile("umicatgo.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file umicatgo.log failed: %v", err)
	}
	logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))
	formatter := &uctFormatter{}
	logrus.SetFormatter(formatter)
	logrus.SetLevel(logrus.TraceLevel)
}
