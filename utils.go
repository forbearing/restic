package restic

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// print is used to output stdout and stderr in real time.
func print(stdout, stderr io.ReadCloser, done chan struct{}) {
	stopCh := make(chan struct{}, 2)
	defer stdout.Close()
	defer stderr.Close()
	//errCh := make(chan error, 2)

	// A goroutine that outputs stdout in real time.
	go func() {
		defer func() {
			stopCh <- struct{}{}
		}()
		scanner := bufio.NewScanner(stdout)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			fmt.Printf("%s", scanner.Text())
		}
		err := scanner.Err()
		// if stdout already closed, stop the goroutine.
		if errors.Is(err, os.ErrClosed) {
			return
		}
		if err != nil {
			fmt.Println("scanner output stdout error:", err)
			return
		}
	}()

	// A goroutine that outputs stderr in real time.
	go func() {
		defer func() {
			stopCh <- struct{}{}
		}()
		scanner := bufio.NewScanner(stderr)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			fmt.Printf("%s", scanner.Text())
		}
		err := scanner.Err()
		// if stderr already closed, stop the goroutine.
		if errors.Is(err, os.ErrClosed) {
			return
		}
		if err != nil {
			fmt.Println("scanner output stderr error:", err)
			return
		}
	}()

	<-stopCh
	<-stopCh
	done <- struct{}{}
}

// concat restic command all flags
func concat(f interface{}) string {
	var s string
	t := reflect.TypeOf(f).Elem()
	v := reflect.ValueOf(f).Elem()

	for i := 0; i < v.NumField(); i++ {
		// PkgPath is the package path that qualified a lows case (unexported)
		// field name. It is empty for upper case (exported) field names.
		// Skip unexported fields
		if len(t.Field(i).PkgPath) != 0 {
			continue
		}
		knd := v.Field(i).Kind()
		typ := v.Field(i).Type().String()
		val := v.Field(i).Interface()
		nam := t.Field(i).Name
		//nam := v.Type().Field(i).Name
		tag := t.Field(i).Tag.Get("json")

		if !v.IsValid() {
			continue
		}
		if v.IsZero() {
			continue
		}
		// If json tag have multiple values, the frist value separated by "," is
		// restic flag name.
		tag = strings.Split(tag, ",")[0]

		_ = knd
		_ = typ
		_ = val
		_ = nam
		_ = tag

		switch typ {
		case "string":
			l, ok := val.(string)
			if ok {
				if len(l) == 0 {
					continue
				}
				s = s + " " + tag + "=" + l
				s = strings.TrimSpace(s)
			}
		case "[]string":
			l, ok := val.([]string)
			if ok {
				if l == nil {
					continue
				}
				s = s + " " + tag + "=" + strings.Join(l, ",")
				s = strings.TrimSpace(s)
			}
		case "int":
			l, ok := val.(int)
			if ok {
				s = s + " " + tag + "=" + strconv.Itoa(l)
				s = strings.TrimSpace(s)
			}
		case "int64":
			l, ok := val.(int64)
			if ok {
				s = s + " " + tag + "=" + strconv.FormatInt(l, 10)
				s = strings.TrimSpace(s)
			}
		case "bool":
			l, ok := val.(bool)
			if ok {
				if l == false {
					continue
				}
				s = s + " " + tag
				s = strings.TrimSpace(s)
			}
		case "map[string]string":
			l, ok := val.(map[string]string)
			if ok {
				var ts string
				for key, val := range l {
					ts = ts + key + "=" + val + ","
				}
				ts = strings.TrimSuffix(ts, ",")
				s = s + " " + tag + "=" + ts
				s = strings.TrimSpace(s)
			}
		}
	}

	return s
}

// concatFlags will concat restic and it's sub-commands all flags
func concatAll(cl ...Command) string {
	var s string
	for _, c := range cl {
		s = s + c.Flags()
	}
	return s
}
